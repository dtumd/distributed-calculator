package parse

import (
	"strings"
	"testing"
	prs "yc/distr-calc/parse"
)

func TestScan(t *testing.T) {
	s := "1 + 2"
	p := prs.NewParser(strings.NewReader(s))
	p.Scan()
}

func TestScanIgnoreWhitespace(t *testing.T) {
	s := "1 + 2"
	p := prs.NewParser(strings.NewReader(s))
	p.ScanIgnoreWhitespace()
}

func TestUnscan(t *testing.T) {
	s := "1 + 2"
	p := prs.NewParser(strings.NewReader(s))
	p.Unscan()
}

func TestParse(t *testing.T) {
	s := "1 + 2"
	p := prs.NewParser(strings.NewReader(s))
	p.ScanIgnoreWhitespace()
}
