package netsuite

import (
)

type NSCreditMemoApply struct {
    Amount                   float64 `json:"amount,omitempty"` 
    Doc                      string `json:"line,omitempty"` 
    Apply                    bool `json:"apply,omitempty"` 
}

type NSCreditMemoItem struct {
    Quantity                 float64 `json:"quantity,omitempty"` 
    Amount                   float64 `json:"amount,omitempty"` 
    Description              string `json:"description,omitempty"`     //ProductCache::productName($item_line['product_id']);
    Item                     map[string]string `json:"item,omitempty"` //for externalId of ProductId
    Line                     int64 `json:"line,omitempty"` 
    OrderLine                int64 `json:"orderLine,omitempty"` 
}

//not a complete list of creditmemo fields... 
//just the ones we feel like using in this app
//for more fields, see: https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html
type NSCreditMemo struct {
    ExternalId          string `json:"externalId,omitempty"` 
    Currency            map[string]interface{} `json:"currency,omitempty"` 
    TranDate            string `json:"tranDate,omitempty"` 
    Entity              map[string]interface{} `json:"entity,omitempty"`
    Item                map[string][]NSCreditMemoItem `json:"item,omitempty"`
    CreatedFrom         map[string]interface{} `json:"createdFrom,omitempty"`
    ApplyList           map[string][]NSCreditMemoApply `json:"applyList,omitempty"`
}



