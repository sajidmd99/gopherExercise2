package handlers

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		redirectPath, ok := pathsToUrls[r.URL.Path]

		if ok {
			http.Redirect(w, r, redirectPath, http.StatusFound)
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
// invalid YAML data.parsedYAML := parseYAML(yml)

//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(inputYAML []byte) ([]path2Url, error) {
	var path2Urls []path2Url
	err := yaml.Unmarshal(inputYAML, &path2Urls)
	if err != nil {
		return nil, err
	}

	return path2Urls, nil
}

func buildMap(path2Urls []path2Url) map[string]string {
	pathUrls := make(map[string]string)
	for _, pu := range path2Urls {
		pathUrls[pu.Path] = pu.Url
	}

	return pathUrls
}

type path2Url struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
