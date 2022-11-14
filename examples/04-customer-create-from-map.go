package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
)

func main() {

	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	//We can also use maps... to create the JSON for the rest API
	request1 := map[string]string{
		"companyName":  "Elephant and Bubbles Waffle Company",
		"email":        "kettle-cooked@example.com",
		"externalId":   "CZ1234567" + netsuite.RandString(4),
		"terms":        "2",
		"subsidiary":   "1",
		"customForm":   "1", //directly specifying an id, implies "id"/"internalId"
	}

	//terms, subsidiary and other fields 
	request2 := map[string]interface{}{
		"companyName": "Elephant and Jiggles Waffle Company",
		"email":       "kettle-cooked@example.com",
		"externalId":  "CZ1234567" + netsuite.RandString(4),
		"terms": map[string]string{
			"id": "2", //you can also specify "id"/"internalId" in a map
		},
		"subsidiary": map[string]string{
			"externalId": "WUSA", //map[string]string] allows to to specify an externalId
		},
		"customForm": map[string]string{
			"id": "1",
		},
	}
	
	//--1
	conn.At("/customer").POST(request1)
	fmt.Println("POST /customer")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	customerId := conn.LastInsertId

	conn.DELETE("/customer/%d", customerId)
	fmt.Printf("DELETE /customer/%d\n", customerId)
	fmt.Printf("STATUS: %s\n", conn.LastStatus)

	//--2
	conn.At("/customer").POST(request2)
	fmt.Println("POST /customer")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	customerId2 := conn.LastInsertId

	conn.At("/customer/%d", customerId2).DELETE()
	fmt.Printf("DELETE /customer/%d\n", customerId2)
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
}
