package accounter

import (
	"gokota/pkg/parser"
)

type Analyzer struct {
	options  map[string]string
	filename string
}

func NewAnalyzer(options map[string]string, filename string) *Analyzer {
	return &Analyzer{options, filename}
}

func (a *Analyzer) Pagecount() (int64, error) {
	parser, err := parser.GetDocumentParser(a.filename)
	if err != nil {
		return 0, err
	}

	return parser.Pagecount(a.filename), nil
}
