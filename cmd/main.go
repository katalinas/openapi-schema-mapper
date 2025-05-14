// cmd/main.go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/katalinas/openapi-schema-mapper/internal/generator"
	"github.com/katalinas/openapi-schema-mapper/internal/parser"
	"github.com/katalinas/openapi-schema-mapper/internal/writer"
)

func main() {
	mapper := generator.NewTypeMapper()
	header := "package specs\n\n"

	root, _ := os.Getwd()
	specsDir := filepath.Join(root, "..", "specs")

	err := filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.Name() != "openapi.yaml" {
			return nil
		}

		def, err := parser.ParseOpenAPIDefinition(path)
		if err != nil {
			return err
		}

		dir := filepath.Dir(path)
		base := filepath.Base(dir)
		outputFile := fmt.Sprintf("%s_ApiSchemasMap.go", base)

		var content string
		for name, schema := range def.Components.Schemas {
			content += generator.GenerateSchemaMap(name, schema, mapper)
		}

		full := header + content
		return writer.WriteFormattedFile(dir, outputFile, full)
	})

	if err != nil {
		log.Fatalf("处理失败: %v", err)
	}
	log.Println("所有 ApiSchemasMap.go 文件生成完成")
}