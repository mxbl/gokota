package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type PDFParser struct {
	f    io.ReaderAt
	size int64
}

func (pdf *PDFParser) Open(filename string) (bool, error) {
	f, info, err := open(filename)
	if err != nil {
		f.Close()
		return false, err
	}

	if ok, err := validate(f, info.Size()); !ok {
		return false, err
	}

	pdf.f = f
	pdf.size = info.Size()
	return true, nil
}

func validate(f io.ReaderAt, size int64) (bool, error) {
	buf := make([]byte, 25)
	f.ReadAt(buf, 0)
	if !bytes.HasPrefix(buf, []byte("%PDF-1.")) || buf[7] < '0' ||
		buf[7] > '7' ||
		buf[8] != '\r' && buf[8] != '\n' {
		return false, fmt.Errorf("not a PDF file: invalid header")
	}

	end := size
	const endChunk = 100
	buf = make([]byte, endChunk)
	f.ReadAt(buf, end-endChunk)
	for len(buf) > 0 && buf[len(buf)-1] == '\n' || buf[len(buf)-1] == '\r' {
		buf = buf[:len(buf)-1]
	}
	buf = bytes.TrimRight(buf, "\r\n\t ")
	if !bytes.HasSuffix(buf, []byte("%%EOF")) {
		return false, fmt.Errorf("not a PDF file: missing %%%%EOF")
	}

	i := findLastLine(buf, "startxref")
	if i < 0 {
		return false, fmt.Errorf("malformed PDF file: missing final startxref")
	}

	return true, nil
}

func (pdf *PDFParser) Pagecount(filename string) int64 {
	return 0
}

// Return an identifier string for this document type
func (pdf *PDFParser) DocumentType() string {
	return "pdf"
}

func parseDocument(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip comments
		if line[:2] == "% " {
			continue
		}
	}

	return nil
}
