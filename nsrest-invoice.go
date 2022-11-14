package netsuite

import (
)

//not a complete list of struct fields... 
//just the ones we feel like using in this app
//for more fields, see: https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html
type NSInvoice struct {
    BillingAddress      NSAddress `json:"billingAddress,omitempty"` 
    ShippingAddress     NSAddress `json:"shippingAddress,omitempty"` 
    ExternalId          string `json:"externalId,omitempty"` 
    InternalId          string `json:"internalId,omitempty"` 
    TranDate            string `json:"tranDate,omitempty"` 
    Email               string `json:"email,omitempty"` 
    OtherRefNum         string `json:"otherRefNum,omitempty"`    
    Account             map[string]interface{} `json:"account,omitempty"`
    CreatedFrom         map[string]interface{} `json:"createdFrom,omitempty"`
    CustomForm          map[string]interface{} `json:"customForm,omitempty"` 
    Terms               map[string]interface{} `json:"terms,omitempty,omitempty"`
    Currency            map[string]interface{} `json:"currency,omitempty"` 
    Entity              map[string]interface{} `json:"entity,omitempty"`
    SalesRep            map[string]interface{} `json:"salesRep,omitempty"`
    Subsidiary          map[string]interface{} `json:"subsidiary,omitempty"`
    Item                map[string][]NSInvoiceItem `json:"item,omitempty"`
    InvoiceMethod       string `json:"custbody_invoice_method,omitempty"`  //example custom field
}

type NSInvoiceItem struct {
    Quantity                 float64 `json:"quantity,omitempty"` 
    Amount                   float64 `json:"amount,omitempty"` 
    Description              string `json:"description,omitempty"`     
    Item                     map[string]string `json:"item,omitempty"`
    Line                     int64 `json:"line,omitempty"` 
    OrderLine                int64 `json:"orderLine,omitempty"` 
}

func (this *NSInvoice) AddItem(itemToAdd NSInvoiceItem) {
    if len(this.Item)==0 {
        this.Item = map[string][]NSInvoiceItem{"items": []NSInvoiceItem{},}
    }
    this.Item["items"] = append(this.Item["items"], itemToAdd)
}

