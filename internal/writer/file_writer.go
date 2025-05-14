package writer

import (
	"go/format"
	"os"
	"path/filepath"
)

func WriteFormattedFile(dir, filename string, content string) error {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, filename), formatted, 0644)
}
