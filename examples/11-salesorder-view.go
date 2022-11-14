package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
	//"strings"
)

func main() {

	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	//fetch by externalId
	salesOrderJSON := conn.GET("/salesOrder/eid:GOTEST_NSTRAN79019")
	fmt.Println(salesOrderJSON)

	//fetch by internalId
	orderInfo := conn.GET("/salesOrder/15816500")
	fmt.Println("GET /salesOrder/15816500")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(orderInfo)

	shipAddress := conn.GET("/salesOrder/3366853/shippingAddress")
	fmt.Println("GET /salesOrder/3366853/shippingAddress")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(shipAddress)

	billAddress := conn.GET("/salesOrder/3366853/billingAddress")
	fmt.Println("GET /salesOrder/3366853/billingAddress")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(billAddress)

	orderItemList := conn.GET("/salesOrder/3366853/item")
	fmt.Println("GET /salesOrder/3366853/item")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(orderItemList)

	orderItems := conn.GET("/salesOrder/3366853/item/1")
	fmt.Println("GET /salesOrder/3366853/item/1")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(orderItems)
}
