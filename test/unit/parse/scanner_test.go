package parse

import (
	"strings"
	"testing"
	prs "yc/distr-calc/parse"
)

func TestRead(t *testing.T) {
	s := "1 + 2"
	r := strings.NewReader(s)
	sc := prs.NewScanner(r)
	sc.Read()
}
