package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type OpenAPI struct {
	OpenapiVersion string `yaml:"openapi"`
	Info           struct {
		Title     string `yaml:"title"`
		ServiceID string `yaml:"x-service"`
		Version   string `yaml:"version"`
	} `yaml:"info"`
}

// ReadOpenAPI reads an openapi file in yaml format and return as OpenAPI struct
func ReadOpenAPI(path string) (OpenAPI, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return OpenAPI{}, err
	}

	var data OpenAPI

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return OpenAPI{}, err
	}

	return data, nil
}
