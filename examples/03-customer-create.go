package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
)

func main() {

	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	customer:= netsuite.NSCustomer{
		CompanyName  :"Elephant and Pebbles Waffle Company",
		Email        :"kettle-cooked@example.com",
		ExternalId   :"CZ1234567" + netsuite.RandString(4),
	}
	fmt.Println(netsuite.JsonEncode(customer))

	response:= conn.At("/customer").POST(customer)
	fmt.Println("POST /customer")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(response)
	lastInsert := conn.LastInsertId

	if lastInsert>0 {
		resp := conn.At("/customer/%d", lastInsert).GET()
		fmt.Println("GET /customer")
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Println(resp)

		conn.DELETE("/customer/%d", lastInsert)
		fmt.Printf("DELETE /customer/%d\n", lastInsert)
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
	}
}
