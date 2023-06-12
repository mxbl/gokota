package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func getPdfsFrom(root string) []string {
	paths := []string{}
	_ = filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".pdf" {
				paths = append(paths, path)
			}
			return nil
		},
	)
	return paths
}

func TestGetDocumentParserForPdfs(t *testing.T) {
	for _, filename := range getPdfsFrom("/home/mx/Documents") {
		parser, err := GetDocumentParser(filename)
		if err != nil {
			t.Logf("Error: %q for %s", err, filename)
			continue
		}

		if parser.DocumentType() != "pdf" {
			t.Errorf(
				"Expected PDF, got %s for file %s",
				parser.DocumentType(),
				filename,
			)
		}
	}
}
