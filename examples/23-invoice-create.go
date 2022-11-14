package main

import (
	"fmt"
	"github.com/mk-j/go-netsuite-rest"
	"strings"
)

func main() {

	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	address := netsuite.NSAddress{
		Attention: "Harrison Fjord",
		Addressee: "LLC Corp Inc.",
		AddrPhone: "867-5309",
		Addr1:     "1600 Amphitheatre Parkway",
		Addr2:     "",
		City:      "Mountain View",
		State:     "CA",
		Zip:       "94043",
		Country:   "US",
	}

	sorder := netsuite.NSSalesOrder{
		ExternalId:      "SO1234567" + netsuite.RandString(4),
		OtherRefNum:     "12345",
		Email:           "jim.johnson@example.com",
		BillingAddress:  address,
		ShippingAddress: address,
		TranDate:        "2022-10-13",
	}
	sorder.SetCurrency("USD")
	sorder.SetEntity("CZ1234565")
	sorder.SetSubsidiary("WUSA")
	sorder.SetTerms("NET 30")
	sorder.AddItem(netsuite.NSSalesOrderItem{
		Quantity:    1.0,
		Amount:      3000,
		Description: "Golden Widget",
		Item:        map[string]string{"externalId": "PZ2"},
		Line:        1,
	})
	sorder.AddItem(netsuite.NSSalesOrderItem{
		Quantity:    1.0,
		Amount:      2000,
		Description: "Silver Widget",
		Item:        map[string]string{"externalId": "PZ4"},
		Line:        2,
	})

	response := conn.At("/salesOrder").POST(customer)
	fmt.Println("POST /salesOrder")
	fmt.Printf("STATUS: %s\n", conn.LastStatus)
	fmt.Println(response)
	lastInsert := conn.LastInsertId

	if lastInsert > 0 {
		resp := conn.At("/salesOrder/%d", lastInsert).GET()
		fmt.Println("GET /salesOrder")
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Println(resp)

		conn.DELETE("/salesOrder/%d", lastInsert)
		fmt.Printf("DELETE /salesOrder/%d\n", lastInsert)
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
	}

}
