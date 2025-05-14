package generator

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"github.com/katalinas/openapi-schema-mapper/internal/parser"
)

type TypeMapper struct {
	typeMap map[string]string
}

func NewTypeMapper() *TypeMapper {
	return &TypeMapper{
		typeMap: map[string]string{
			"integer": "int",
			"number":  "float64",
			"string":  "string",
			"boolean": "bool",
		},
	}
}

func (m *TypeMapper) Convert(openAPIType string) string {
	if goType, ok := m.typeMap[openAPIType]; ok {
		return goType
	}
	return "interface{}"
}

func GenerateSchemaMap(schemaName string, schema parser.Schema, mapper *TypeMapper) string {
	var result string
	result += fmt.Sprintf("var %s = map[string]string{\n", schemaName)

	if schema.RawProperties != nil && schema.RawProperties.Kind == yaml.MappingNode {
		for i := 0; i < len(schema.RawProperties.Content); i += 2 {
			key := schema.RawProperties.Content[i]
			value := schema.RawProperties.Content[i+1]

			propName := key.Value
			var propType string
			for j := 0; j < len(value.Content); j += 2 {
				k := value.Content[j]
				v := value.Content[j+1]
				if k.Value == "type" {
					propType = v.Value
					break
				}
			}

			goType := mapper.Convert(propType)
			result += fmt.Sprintf("\t\"%s\": \"%s\",\n", propName, goType)
		}
	}
	result += "}\n\n"
	return result
}
