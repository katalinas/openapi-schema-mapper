// pkg/mapper/mapper.go
package mapper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/katalinas/openapi-schema-mapper/internal/generator"
	"github.com/katalinas/openapi-schema-mapper/internal/parser"
	"github.com/katalinas/openapi-schema-mapper/internal/writer"
)

type Config struct {
	SpecsDir  string // OpenAPI规范文件目录
	OutputPkg string // 生成的Go文件包名
}

func GenerateAll(config Config) error {
	mapper := generator.NewTypeMapper()
	header := fmt.Sprintf("package %s\n\n", config.OutputPkg)

	return filepath.Walk(config.SpecsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.Name() != "openapi.yaml" {
			return nil
		}

		return processSingleFile(path, mapper, header)
	})
}

func processSingleFile(path string, mapper *generator.TypeMapper, header string) error {
	def, err := parser.ParseOpenAPIDefinition(path)
	if err != nil {
		return fmt.Errorf("解析文件 %s 失败: %w", path, err)
	}

	dir := filepath.Dir(path)
	base := filepath.Base(dir)
	outputFile := fmt.Sprintf("%s_ApiSchemasMap.go", base)

	var content string
	for name, schema := range def.Components.Schemas {
		content += generator.GenerateSchemaMap(name, schema, mapper)
	}

	return writer.WriteFormattedFile(dir, outputFile, header+content)
}
