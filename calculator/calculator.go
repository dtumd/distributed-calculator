package calculator

import (
	"distr-calc/parse"
	"fmt"
	"strings"
)

func Calculate(s string) (float64, error) {
	p := parse.NewParser(strings.NewReader(s))
	// fmt.Printf("%+v\n", p)
	stack, err := p.Parse()
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
		return 0, err
	}
	// fmt.Printf("%+v\n", stack)
	stack = parse.ShuntingYard(stack)
	// fmt.Printf("%+v\n", stack)
	answer := parse.SolvePostfix(stack)
	return answer, nil
}
