// cmd/main.go
package main

import (
	"log"

	"github.com/katalinas/openapi-schema-mapper/internal/config"
	"github.com/katalinas/openapi-schema-mapper/pkg/mapper"
)

func main() {
	config := config.Config{
		SpecsDir:  "../specs",
		OutputPkg: "specs",
	}

	if err := mapper.GenerateAll(config); err != nil {
		log.Fatalf("生成失败: %v", err)
	}
	log.Println("所有 ApiSchemasMap.go 文件生成完成")
}
