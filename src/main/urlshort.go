package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort"
	)

func main() {
	yamlPath := flag.String("mappingYaml", "/home/akshay/repos/gophercises/src/urlshort/mapping.yaml",
		"Path to yaml containing mappings of url to redirection")
	flag.Parse()

	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlData, err := ioutil.ReadFile(*yamlPath)
	if err != nil {
		fmt.Println("Error file opening yaml file with path ", *yamlPath)
	}
	yamlHandler, _ := urlshort.YAMLHandler(yamlData, mapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}


