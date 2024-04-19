package orchestrator

import (
	"context"
	"fmt"
	"log"
	"time"
	db "yc/distr-calc/db"
	mdl "yc/distr-calc/model"
	calcp "yc/distr-calc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func AddExpression(expr string, login string) mdl.Expression {
	fmt.Println("AddExpression, user login: " + login)

	expression := db.SaveExpression(expr, "Calculating", "?", login) // prepare expression

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

		host := "localhost"
		port := "8333"

		addr := fmt.Sprintf("%s:%s", host, port)

		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Println("could not connect to grpc server: ", err)
		}

		defer conn.Close()
		grpcClient := calcp.NewCalculateServiceClient(conn)

		res, err := grpcClient.Calculate(context.TODO(), &calcp.CalculateRequest{Expression: expression.Value})

		if err != nil {
			expression.Status = "Error"
			expression.Result = "Error"
		} else {
			expression.Status = "Ready"
			expression.Result = fmt.Sprint(res.Result)
		}
		db.UpdateExpression(expression)
	}()

}
