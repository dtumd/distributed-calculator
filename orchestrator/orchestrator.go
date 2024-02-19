package orchestrator

import (
	calc "distr-calc/calculator"
	db "distr-calc/db"
	mdl "distr-calc/model"
	"fmt"
	"time"
)

var resources = map[string][]mdl.ComputingResource{
	"Resources": {
		{Status: "Work", Name: "computing server", LastPing: time.Now()},
		{Status: "Reconnect", Name: "computing server", LastPing: time.Now().Add(time.Minute * (-5))},
		{Status: "Lost", Name: "computing server", LastPing: time.UnixMicro(0)},
	},
}

func GetResources() map[string][]mdl.ComputingResource {
	return resources
}

func AddExpression(expr string) mdl.Expression {
	expression := db.SaveExpressions(expr, "Calculating", "?")

	calculate(expression)

	return expression
}

func Init() {
	go func() {
		fmt.Println("Initialize orchestrator")
	}()
}

func calculate(expression mdl.Expression) {
	go func() {
		fmt.Println("calculate")

		time.Sleep(10 * time.Second)

		res, err := calc.Calculate(expression.Value)
		if err != nil {
			expression.Status = "Error"
			expression.Result = "Error"
		} else {
			expression.Status = "Ready"
			expression.Result = fmt.Sprint(res)
		}
		db.UpdateExpression(expression)
	}()
}
