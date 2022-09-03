package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

// SignHMAC signs msg with key by using HMAC-SHA256 function.
func SignHMAC(msg, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return mac.Sum(nil)
}

// VerifyHMAC verifies hash signed by SignHMAC msg with key by using HMAC-SHA256
// function.
func VerifyHMAC(msg, key, hash []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(hash, mac.Sum(nil))
}
