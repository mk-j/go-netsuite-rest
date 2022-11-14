package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
)

func main() {
	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	//get customer by internalId
	custInternalId := 1234567
	jsonResponse := conn.GET("/customer/%v", custInternalId)
	customer:= netsuite.UnmarshalJSONToCustomer(jsonResponse)
	fmt.Printf("GET /customer/%v\n", custInternalId)
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println("--------------------------")
	fmt.Println("FULL RESPONSE:")
	fmt.Println(jsonResponse)
	fmt.Println("--------------------------")
	fmt.Println("JUST THE FIELDS I WANT:")
	fmt.Println(netsuite.JsonEncode(customer))

	//get customer by externalId
	custExternalId := "CZ1234567"
	jsonResponse2 := conn.GET("/customer/eid:%v", custExternalId)
	customer2:= netsuite.UnmarshalJSONToCustomer(jsonResponse2)
	fmt.Printf("GET /customer/eid: %v\n", custExternalId)
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(netsuite.JsonEncode(customer2))
}
