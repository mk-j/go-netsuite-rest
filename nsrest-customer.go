package netsuite

import (
    "log"
    "encoding/json"
)

//not a complete list of struct fields... 
//just the ones we feel like using in this app
//for more fields, see: https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html
type NSCustomer struct {
    InternalId          string `json:"id,omitempty"`
    ExternalId          string `json:"externalId,omitempty"`
    CompanyName         string `json:"companyName,omitempty"`
    Email               string `json:"email,omitempty"`
    EntityId            string `json:"entityId,omitempty"`    
    EntityStatus        map[string]interface{} `json:"entityStatus,omitempty"`
    CustomForm          map[string]interface{} `json:"customForm,omitempty"` 
    Terms               map[string]interface{} `json:"terms,omitempty"`
    Subsidiary          map[string]interface{} `json:"subsidiary,omitempty"`
    AccountCreateDate   string `json:"custentity_account_create_date,omitempty"`
}

func (this *NSCustomer) SetSubsidiary(SubsidiaryExternalId string) {
    this.Subsidiary = map[string]interface{} {"externalId":SubsidiaryExternalId}
}
func (this *NSCustomer) SetTerms(TermsExternalId string) {
    this.Terms = map[string]interface{} {"externalId":TermsExternalId}
}
func (this *NSCustomer) SetTermsId(TermsInternalId string) {
    this.Terms = map[string]interface{} {"id":TermsInternalId}
}
func (this *NSCustomer) SetCustomFormId(CustomFormId string) {
    this.CustomForm = map[string]interface{} {"id":CustomFormId}
}


func UnmarshalJSONToCustomer(json_input string) *NSCustomer {
    ncustomer := new(NSCustomer)
    if err := json.Unmarshal([]byte(json_input), ncustomer); err != nil {
        log.Printf("Error unmarshalling customer: %v\n", err)
    }
    return ncustomer
}
