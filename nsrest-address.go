package netsuite

import (
)

//not a complete list of struct fields... 
//just the ones we feel like using in this app
//for more fields, see: https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html
type NSAddress struct {
    Attention    string `json:"attention,omitempty"`
    Addressee    string `json:"addressee,omitempty"`
    AddrPhone    string `json:"addrPhone,omitempty"`
    Addr1        string `json:"addr1,omitempty"`
    Addr2        string `json:"addr2,omitempty"`
    Addr3        string `json:"addr3,omitempty"`
    City         string `json:"city,omitempty"`
    State        string `json:"state,omitempty"`
    Zip          string `json:"zip,omitempty"`
    Country      string `json:"country,omitempty"`
}


