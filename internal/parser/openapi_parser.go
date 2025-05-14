package parser

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Schema struct {
	RawProperties *yaml.Node `yaml:"properties"`
}

type Components struct {
	Schemas map[string]Schema `yaml:"schemas"`
}

type OpenAPIDefinition struct {
	Components Components `yaml:"components"`
}

func ParseOpenAPIDefinition(path string) (*OpenAPIDefinition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var def OpenAPIDefinition
	if err := yaml.Unmarshal(data, &def); err != nil {
		return nil, err
	}

	return &def, nil
}
