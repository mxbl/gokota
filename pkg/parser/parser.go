package parser

import (
	"fmt"
	"gokota/pkg/config"
	"io"
	"io/fs"
	"os"
)

type Parser interface {
	Open(filename string) (bool, error)
	Pagecount(filename string) int64
	DocumentType() string
}

// Tries to autodetect the document format.
// Returns the correct handler function or nil if the format is unknown.
func GetDocumentParser(
	filename string,
) (Parser, error) {
	fd, info, err := open(filename)
	if err != nil {
		return nil, err
	}

	// Read the first and the last block of the file
	first := make([]byte, config.FIRST_BLOCK_SIZE)
	_, err = fd.Read(first)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("Error reading first block: %q", err)
		}
	}
	last := make([]byte, config.LAST_BLOCK_SIZE)
	_, err = fd.ReadAt(last, info.Size()-int64(config.LAST_BLOCK_SIZE))
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("Error reading last block: %q", err)
		}
	}

	// Check all listet document type parsers, each parser checks if this
	// file is of his type, we return the first parser that matches.
	// TODO: does the order matter than?
	parser := []Parser{
		new(PDFParser),
	}

	for _, p := range parser {
		// TODO: Maybe handle error herr?
		ok, err := p.Open(filename)
		if err != nil {
			fmt.Printf("Error opening parser: %s\n", filename)
			fmt.Printf("Error: %q\n", err)
		}
		if ok {
			return p, nil
		}
	}
	return nil, fmt.Errorf("Unknown file type")
}

func open(filename string) (f *os.File, info fs.FileInfo, err error) {
	f, err = os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("Error opening file: %q", err)
	}

	info, err = f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("Error getting file info: %q", err)
	}

	return f, info, nil
}
