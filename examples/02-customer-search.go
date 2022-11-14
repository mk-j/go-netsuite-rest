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
	fmt.Println("Grab a few sample customers? (Y/N) [this could take a while]")
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		//list of customers
		//date formats are discussed in https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/section_157305090645.html
		q := `dateCreated ON_OR_AFTER "1/31/2022" AND dateCreated BEFORE "12/31/2022"`
		
		urlparams:=map[string]string{"limit":"5","offset":"0","q":q}
		resp := conn.At("/customer/").Query(urlparams).GET()
		fmt.Println("GET /customer/?limit=5&offset=0&q=" + netsuite.PercentEncode(q))
		fmt.Printf("STATUS: %s\n", conn.LastStatus)
		fmt.Println(resp)

		fmt.Println("Press the Enter Key to continue")
		fmt.Scanln()
		fmt.Println("")
	}
}
