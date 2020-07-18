package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	 var handler http.HandlerFunc =  func(writer http.ResponseWriter, request *http.Request) {
	 	if mappedUrl, found := pathsToUrls[request.URL.Path]; ! found {
	 		fmt.Println("No mapping found for ", request.URL.String())
	 		fallback.ServeHTTP(writer, request)
		} else {
			 http.Redirect(writer, request, mappedUrl, http.StatusFound)
		 }
	 }
	 return handler
}


// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathAndUrls := [] PathAndUrls{}
	err := yaml.Unmarshal(yamlData, &pathAndUrls)
	if err != nil {
		fmt.Println("Error parsing yaml", err)
	}
	pathToUrl := make(map[string]string)
	for _, pathAndUrl := range pathAndUrls {
		pathToUrl[pathAndUrl.Path] = pathAndUrl.Url
	}

	return MapHandler(pathToUrl, fallback), err
}


type PathAndUrls struct {
	Path string `yaml:"path"`
	Url string	`yaml:"url"`
}