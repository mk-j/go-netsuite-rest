package netsuite

import (
)

type NSSalesOrderItem struct {
    Quantity                 float64 `json:"quantity,omitempty"` 
    Amount                   float64 `json:"amount,omitempty"` 
    Description              string `json:"description,omitempty"`     
    Item                     map[string]string `json:"item,omitempty"` 
    Line                     int64 `json:"line,omitempty"` 
}

//not a complete list of salesorder fields... 
//just the ones we feel like using in this app
//for more fields, see: https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html
type NSSalesOrder struct {
    InternalId          string `json:"internalId,omitempty"` 
    ExternalId          string `json:"externalId,omitempty"` 
    CustomForm          map[string]interface{} `json:"customForm,omitempty"`
    Currency            map[string]interface{} `json:"currency,omitempty"` 
    Memo                string `json:"memo,omitempty"`
    OtherRefNum         string `json:"otherRefNum,omitempty"`
    TranDate            string `json:"tranDate,omitempty"` 
    Entity              map[string]interface{} `json:"entity,omitempty"`
    Item                map[string][]NSSalesOrderItem `json:"item,omitempty"`
    Subsidiary          map[string]interface{} `json:"subsidiary,omitempty"`
    Terms               map[string]interface{} `json:"terms,omitempty"`
    SalesRep            map[string]interface{} `json:"salesRep,omitempty"`
    Message             string `json:"message,omitempty"` 
    LineOfBusiness      string `json:"class,omitempty"`
    Email               string `json:"email,omitempty"`
    BillingAddress      NSAddress `json:"billingAddress,omitempty"`
    ShippingAddress     NSAddress `json:"shippingAddress,omitempty"`
    AltLocalName        string `json:"custbody_local_name,omitempty"` //example custom field
}
//----------------------------------------------------------------------
func (this *NSSalesOrder) AddItem(itemToAdd NSSalesOrderItem) {
    if len(this.Item)==0 {
        this.Item = map[string][]NSSalesOrderItem{"items": []NSSalesOrderItem{},}
    }
    this.Item["items"] = append(this.Item["items"], itemToAdd)
}

func (this *NSSalesOrder) SetCurrency(CurrencyExternalId string) {
    this.Currency = map[string]interface{} {"externalId":CurrencyExternalId}
}
func (this *NSSalesOrder) SetEntity(EntityExternalId string) {
    this.Entity = map[string]interface{} {"externalId":EntityExternalId}
}
func (this *NSSalesOrder) SetSubsidiary(SubsidiaryExternalId string) {
    this.Subsidiary = map[string]interface{} {"externalId":SubsidiaryExternalId}
}
func (this *NSSalesOrder) SetTerms(TermsExternalId string) {
    this.Terms = map[string]interface{} {"externalId":TermsExternalId}
}
func (this *NSSalesOrder) SetTermsId(TermsInternalId string) {
    this.Terms = map[string]interface{} {"id":TermsInternalId}
}
func (this *NSSalesOrder) SetSalesRepId(SalesRepInternalId string) {
    this.SalesRep = map[string]interface{} {"id":SalesRepInternalId}
}
func (this *NSSalesOrder) SetCustomFormId(CustomFormId string) {
    this.CustomForm = map[string]interface{} {"id":CustomFormId}
}
