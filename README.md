# go-netsuite
Go/Golang  library for connecting to netsuite rest API

This requires a config/config.yml file in the following format:
```
netsuite:
    apiURL: https://1234567-sb1.suitetalk.api.netsuite.com/services/rest/record/v1
    account: 1234567_SB1
    token: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    tokenSecret: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    consumerKey: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
    consumerSecret: 0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff
```


```go
package main

import (
    "fmt"
    "github.com/mk-j/go-netsuite-rest"
)

func main() {
	conn := netsuite.NetSuiteConnector{}
	conn.ReadConfig("./config/config.yml")

	jsonResponse:= conn.GET("/customer/1234567")
	restCustomer:= restdata.UnmarshalJSONToCustomer(jsonResponse)
	fmt.Println(JsonEncode(restCustomer))
}

```

Documention on the NetSuite REST API can be found at:
- https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/chapter_1540391670.html
- https://system.netsuite.com/help/helpcenter/en_US/APIs/REST_API_Browser/record/v1/2022.1/index.html#tag-salesOrder

On Date Formats;
- https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/section_157305090645.html

On Filtering Data
- https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/section_1545222128.html

| Field Type | Allowed Operators |
|----|----|
| None | EMPTY, EMPTY_NOT |
| Boolean | IS, IS_NOT |
| Double, Integer, Float, Number, Duration | ANY_OF, ANY_OF_NOT, BETWEEN, BETWEEN_NOT, EQUAL, EQUAL_NOT, GREATER, GREATER_NOT, GREATER_OR_EQUAL, GREATER_OR_EQUAL_NOT, LESS, LESS_NOT, LESS_OR_EQUAL, LESS_OR_EQUAL_NOT, WITHIN, WITHIN_NOT |
| String | CONTAIN, CONTAIN_NOT, IS, IS_NOT, START_WITH, START_WITH_NOT, END_WITH, END_WITH_NOT |
| Date / Time | AFTER, AFTER_NOT, BEFORE, BEFORE_NOT, ON, ON_NOT, ON_OR_AFTER, ON_OR_AFTER_NOT, ON_OR_BEFORE, ON_OR_BEFORE_NOT |    



| Operators with varying  numbers of values | Example
|----|----|
| Unary operators: The EMPTY and EMPTY_NOT operators do not accept any values. | ?q=companyName EMPTY | 
| Ternary operators: The BETWEEN, BETWEEN_NOT, WITHIN, and WITHIN_NOT operators accept two values. |  ?q=id BETWEEN_NOT [1, 42] | 
| N-ary operators: The ANY_OF and ANY_OF_NOT operators do accept one or any higher number of values. |  ?q=id ANY_OF [1, 2, 3, 4, 5] | 


