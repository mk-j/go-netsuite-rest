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
	salesOrderId := conn.LastInsertId

	if salesOrderId > 0 {
		resp2 := conn.At("/salesOrder/%v", salesOrderId).GET()
		fmt.Println("GET /salesOrder")
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Println(resp2)

		resp3 := conn.POST("/salesOrder/%v/!transform/itemFulfillment", salesOrderId)
		fmt.Printf("POST /salesOrder/%v/!transform/itemFulfillment\n", salesOrderId)
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Printf(resp3)

		fmt.Println(xutils.JsonEncode(invoiceRequest))
		resp4 := conn.At("/salesOrder/%v/!transform/invoice", salesOrderId).POST(invoiceRequest)
		fmt.Printf("POST /salesOrder/%v/!transform/invoice\n", salesOrderId)
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Printf(resp4)
		invoiceId := conn.LastInsertId

		if invoiceId > 0 {
			resp5 := conn.GET("/invoice/%v", invoiceId)
			fmt.Printf("/invoice/%v\n", invoiceId)
			fmt.Printf("STATUS: %s\n", conn.LastStatus)
			fmt.Println(resp5)

			fmt.Printf("DELETE-ing invoice...\n")
			conn.DELETE("/invoice/%v", invoiceId)
		}
		fmt.Printf("DELETE-ing salesOrder...\n")
		conn.DELETE("/salesOrder/%v", salesOrderId)
	}
}
