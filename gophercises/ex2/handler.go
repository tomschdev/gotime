package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	// http.HandlerFunc is a type that is a func with following params
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := urls[r.URL.Path]; ok { // determines whether map contains entry with corresponding key
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml
	urlYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	// convert yaml to map
	urlMap := buildMap(urlYaml)
	// return a map handler using the map
	return MapHandler(urlMap, fallback), nil
}
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlJson, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}
	urlMap := buildMap(urlJson)
	return MapHandler(urlMap, fallback), nil
}

type pathToUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func parseYAML(yml []byte) ([]pathToUrl, error) {
	var pathsToUrls []pathToUrl
	err := yaml.Unmarshal(yml, &pathsToUrls)
	if err != nil {
		return nil, err
	}
	return pathsToUrls, nil

}
func parseJSON(jsn []byte) ([]pathToUrl, error) {
	var pathsToUrls []pathToUrl
	err := json.Unmarshal(jsn, &pathsToUrls)
	if err != nil {
		return nil, err
	}
	return pathsToUrls, nil
}
func buildMap(ptu []pathToUrl) map[string]string {
	urlMap := make(map[string]string)
	for _, pu := range ptu {
		urlMap[pu.Path] = pu.URL
	}
	return urlMap
}
