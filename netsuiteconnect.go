package netsuite

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"reflect"
	"strconv"
)

type NetSuiteConnector struct {
	Auth struct {
		Account        string `yaml:"account"`
		ConsumerKey    string `yaml:"consumerKey"`
		ConsumerSecret string `yaml:"consumerSecret"`
		Token          string `yaml:"token"`
		TokenSecret    string `yaml:"tokenSecret"`
		APIURL         string `yaml:"apiURL"`
	} `yaml:"netsuite"`
	ConfigFile     string
	LastInsertId   int
	LastStatusCode int
	LastStatus     string
}

type NetSuiteRequest struct {
	Connection *NetSuiteConnector
	RequestURL string
	QueryParams map[string]string
	RequestBody interface{} //can pass in json encoded string, or a struct or map here
}

// https://docs.oracle.com/en/cloud/saas/netsuite/ns-online-help/chapter_1540391670.html
/* REST web services always load record instances in "Edit" mode, including GET requests. 
 * When a user without the Administrator role tries to run a GET request for an employee 
 * record of a user with the Administrator role, this error message will appear: 
 * "For security reasons, only an administrator is allowed to edit an administrator record". 
 * This prevents users without the Administrator role from editing the employee record of a 
 * user with the Administrator role. */

//----------------------------------------------------------------------
func (r *NetSuiteConnector) ReadConfig(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, r)
	if err != nil {
		panic(err)
	}
	if r.Auth.Account == "" {
		panic("Missing from config.yml file: " + "auth: Account")
	}
	if r.Auth.APIURL == "" {
		panic("Missing from config.yml file: " + "auth: APIURL")
	}
	if r.Auth.Token == "" {
		panic("Missing from config.yml file: " + "auth: Token")
	}
	if r.Auth.TokenSecret == "" {
		panic("Missing from config.yml file: " + "auth: TokenSecret")
	}
	if r.Auth.ConsumerKey == "" {
		panic("Missing from config.yml file: " + "auth: ConsumerKey")
	}
	if r.Auth.ConsumerSecret == "" {
		panic("Missing from config.yml file: " + "auth: ConsumerSecret")
	}
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) do(method string, path string, post_json interface{}) string {
	if r.Auth.Account=="" {
		if len(r.ConfigFile)>0 {
			r.ReadConfig(r.ConfigFile)
		} else {
			panic("Config.yml file not specified")
		}
	}

	post_json_string:=""
	switch str:=reflect.TypeOf(post_json).Kind(); str {
		case reflect.String:
			post_json_string = post_json.(string)
		case reflect.Struct:
			post_json_string = JsonEncode(post_json)
		default:
			post_json_string = JsonEncode(post_json)
	}

	form := map[string]string{}
	apiurl := strings.TrimSuffix(r.Auth.APIURL, "/")
	apipath := strings.TrimPrefix(path, "/")
	restapiurl := fmt.Sprintf("%s/%s", apiurl, apipath)

	var post_body io.Reader
	if method != "GET" {
		post_body = strings.NewReader(post_json_string)
	}
	req, err := http.NewRequest(method, restapiurl, post_body)

	authorization := r.authorization(method, restapiurl, form)
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {authorization},
		"Cookie":        {"NS_ROUTING_VERSION=LAGGING"},
	}
	//log.Printf("[%s]\n", authorization)

	client := http.Client{}
	resp, err := client.Do(req)

	r.LastStatus = resp.Status
	r.LastStatusCode, _ = strconv.Atoi(resp.Status[0:3])

	//log.Printf("STATUS: [%s]\n", resp.Status)
	//if location, ok := resp.Header["Location"]; ok {
	//	log.Printf("Header: [%s]\n", location)
	//}
	
	//for k,v := range(resp.Header) {
	//	log.Printf("Header: [%v] %v\n", k, v)
	//}

	r.LastInsertId = extractLastInsertId(resp.Header)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return JsonPrettyPrint(string(body))
}

// ----------------------------------------------------------------------
func (r *NetSuiteConnector) isDev() bool {
	if strings.Contains(strings.ToLower(r.Auth.Account), "sb") {
		return true
	}
	return false
}
// ----------------------------------------------------------------------
func (r *NetSuiteConnector) authorization(method string, neturl string, form map[string]string) string {
	if method == "DELETE" && !r.isDev() {
		log.Printf("DELETE REST call attempted.... terminated prematurely, DELETE in PROD disallowed")
		return ""
	}
	ubase := urlWithoutQueryString(neturl)
	queryParams := getQueryParams(neturl)

	nonce := RandString(11)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	realm := r.Auth.Account
	consumer_key := r.Auth.ConsumerKey
	consumer_secret := r.Auth.ConsumerSecret
	token := r.Auth.Token
	token_secret := r.Auth.TokenSecret
	oauthParams := map[string]string{
		"oauth_consumer_key":     consumer_key,
		"oauth_nonce":            nonce,
		"oauth_signature_method": "HMAC-SHA256",
		"oauth_timestamp":        timestamp,
		"oauth_token":            token,
		"oauth_version":          "1.0",
	}

	params := []string{}
	for k, v := range form {
		params = append(params, fmt.Sprintf("%s=%s", k, PercentEncode(v)))
	}
	for k, v := range queryParams {
		for _, vv := range v {
			params = append(params, fmt.Sprintf("%s=%s", k, PercentEncode(vv)))
		}
	}
	for k, v := range oauthParams {
		params = append(params, fmt.Sprintf("%s=%s", k, PercentEncode(v)))
	}
	sort.Strings(params)

	paramsJoined := strings.Join(params, "&")
	baseString := strings.Join([]string{method, PercentEncode(ubase), PercentEncode(paramsJoined)}, "&")
	//fmt.Printf("[base] [%s]\n", baseString)
	secretString := strings.Join([]string{consumer_secret, token_secret}, "&")
	signature := url.QueryEscape(HashHmacSha256(baseString, secretString))

	authstr := fmt.Sprintf("OAuth realm=%q,oauth_consumer_key=%q,oauth_token=%q,oauth_signature_method=%q,"+
		"oauth_timestamp=%q,oauth_nonce=%q,oauth_version=%q,oauth_signature=%q",
		realm, consumer_key, token, "HMAC-SHA256", timestamp, nonce, "1.0", signature)
	//fmt.Printf("[auth] [%s]\n", authstr)
	return authstr
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) GET(patharg string, args ...interface{}) string {
	path := fmt.Sprintf(patharg, args...)
	json_post_body:=""
	return r.do("GET", path, json_post_body)
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) POST(patharg string, args ...interface{}) string {
	path := fmt.Sprintf(patharg, args...)
	json_post_body:=""
	return r.do("POST", path, json_post_body)
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) DELETE(patharg string, args ...interface{}) string {
	path := fmt.Sprintf(patharg, args...)
	json_post_body := ""
	return r.do("DELETE", path, json_post_body)
}
//----------------------------------------------------------------------
func (r *NetSuiteConnector) At(patharg string, args ...interface{}) NetSuiteRequest {
	path := fmt.Sprintf(patharg, args...)
	return NetSuiteRequest{ RequestURL:path, Connection:r }
}
//----------------------------------------------------------------------
//----------------------------------------------------------------------
func (r NetSuiteRequest) Query(QueryParams map[string]string) NetSuiteRequest {
	r.QueryParams=QueryParams
	return r
}
func (r NetSuiteRequest) GET() string {
	if len(r.QueryParams)>0 {
		pieces := []string{}
		for k,v := range(r.QueryParams) {
			pieces = append(pieces, fmt.Sprintf("%s=%s", PercentEncode(k), PercentEncode(v)))
		}
		return r.Connection.do("GET", r.RequestURL+"?"+strings.Join(pieces, "&"), "")
	}
	json_post_body:=""
	return r.Connection.do("GET", r.RequestURL, json_post_body)
}
//----------------------------------------------------------------------
func (r NetSuiteRequest) POST(post_json interface{}) string {
	return r.Connection.do("POST", r.RequestURL, post_json)
}
func (r NetSuiteRequest) INSERT(post_json interface{}) string {
	return r.Connection.do("POST", r.RequestURL, post_json)
}
//----------------------------------------------------------------------
func (r NetSuiteRequest) PATCH(post_json interface{}) string {
	return r.Connection.do("PATCH", r.RequestURL, post_json)
}
func (r NetSuiteRequest) UPDATE(post_json interface{}) string {
	return r.Connection.do("PATCH", r.RequestURL, post_json)
}
//----------------------------------------------------------------------
func (r NetSuiteRequest) PUT(post_json interface{}) string {
	return r.Connection.do("PUT", r.RequestURL, post_json)
}
func (r NetSuiteRequest) UPSERT(post_json interface{}) string {
	return r.Connection.do("PUT", r.RequestURL, post_json)
}
//----------------------------------------------------------------------
func (r NetSuiteRequest) DELETE() string {
	json_post_body := ""
	return r.Connection.do("DELETE", r.RequestURL, json_post_body)
}
//----------------------------------------------------------------------

