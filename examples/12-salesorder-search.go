package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
	"strings"
)

func main() {

	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	var response string
	fmt.Println("Grab a few sample salesOrders? (Y/N)")
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		orderlist := conn.GET("/salesOrder/?limit=5&offset=0")
		fmt.Println("GET /salesOrder/?limit=5&offset=0")
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Println(orderlist)

		fmt.Println("Press the Enter Key to continue")
		fmt.Scanln()
		fmt.Println("")
	}

}
