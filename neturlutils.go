package netsuite

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func extractLastInsertId(h http.Header) int {
	//https://1234567-sb2.suitetalk.api.netsuite.com/services/rest/record/v1/customer/987654
	if location, ok := h["Location"]; ok {
		urllist := filterStartsWith(location, "https://")
		lastIdStr := arrayLast(strings.Split(arrayFirst(urllist), "/"))
		i, _ := strconv.Atoi(lastIdStr)
		return i
	}
	return 0
}

func arrayFirst(array []string) string {
	if len(array) > 0 {
		return array[0]
	}
	return ""
}

func arrayLast(array []string) string {
	if len(array) > 0 {
		return array[len(array)-1]
	}
	return ""
}

func filterStartsWith(input []string, prefix string) []string {
	output := []string{}
	for _, v := range input {
		if v[0:len(prefix)] == prefix {
			output = append(output, v)
		}
	}
	return output
}

func urlWithoutQueryString(neturl string) string {
	noQueryUrl, _ := url.Parse(neturl)
	noQueryUrl.Scheme = strings.ToLower(noQueryUrl.Scheme)
	noQueryUrl.Host = strings.ToLower(noQueryUrl.Host)
	noQueryUrl.RawQuery = ""
	return noQueryUrl.String()
}

func getQueryParams(neturl string) url.Values {
	u, _ := url.Parse(neturl)
	queryParams := u.Query()
	return queryParams
}

func PercentEncode(input string) string {
	var buf bytes.Buffer
	for _, b := range []byte(input) {
		// if in unreserved set
		if shouldEscape(b) {
			buf.Write([]byte(fmt.Sprintf("%%%02X", b)))
		} else {
			// do not escape, write byte as-is
			buf.WriteByte(b)
		}
	}
	return buf.String()
}

func shouldEscape(c byte) bool {
	// RFC3986 2.3 unreserved characters
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}
	switch c {
	case '-', '.', '_', '~':
		return false
	}
	// all other bytes must be escaped
	return true
}


