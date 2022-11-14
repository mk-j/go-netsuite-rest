package netsuite

import (
    //"fmt"
    "crypto/hmac"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/rand"
    "encoding/base64"
    "encoding/binary"
)

func HashHmacSha1(message string, secret string) string {
    h := hmac.New(sha1.New, []byte(secret))
    h.Write([]byte(message))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HashHmacSha256(message string, secret string) string {
    h := hmac.New(sha256.New, []byte(secret))
    h.Write([]byte(message))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }
    return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func RandString(s int) string {
    b, _ := GenerateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b)[0:s]
}

func RandInt(min int64, max int64) int64 {
    if min > max {
        swapme:=min
        min=max
        max=swapme
    }
    b, _ := GenerateRandomBytes(8)   
    randomInt := int64(binary.LittleEndian.Uint64(b)>>1)
    return int64(min) + randomInt%(max-min)
}

