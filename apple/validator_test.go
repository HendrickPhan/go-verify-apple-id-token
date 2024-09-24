package apple

import (
	"fmt"
	"testing"
)

var idToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6ImJkOWVjZDNiODNhOTNkMGNjMTczODBmZGQzZTRiZTVmIn0.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLmV4YW1wbGUuYXBwIiwiZXhwIjoxNzI3MjU1MjIzLCJpYXQiOjE3MjcxNjg4MjMsInN1YiI6InJhbmRvbV9hcHBsZV91c2VyX2lkIiwiY19oYXNoIjoicmFuZG9tX2NfaGFzX3ZhbHVlIiwiZW1haWwiOiJyYW5kb20tdmFsdWVAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzX3ByaXZhdGVfZW1haWwiOnRydWUsImF1dGhfdGltZSI6MTcyNzE2ODgyMywibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.VqR1igwpz1uaNQznmjFicjDySSQJD-nj5ToiS1kwSyqtNjDUpgiGW3hvSjjdsQkpcKmKfayNjE-30O_ewB2Nhg"
var wrongToken = ""

func TestVerifyIdToken(t *testing.T) {
	client := New()

	jwtClaims, err := client.VerifyIdToken("com.example.app", idToken)
	if err != nil {
		t.Errorf("%s", err)
	} else {
		fmt.Println(jwtClaims.Email)
		fmt.Println(jwtClaims.AuthTime)
		fmt.Println(jwtClaims.Aud)
		fmt.Println(jwtClaims.IsPrivateEmail)
		fmt.Println(jwtClaims.EmailVerified)
	}
}

func TestVerifyWrongToken(t *testing.T) {
	client := New()

	jwtClaims, err := client.VerifyIdToken("com.ios.sample", wrongToken)
	if err != nil {
		t.Errorf("%s", err)
	} else {
		fmt.Println(jwtClaims.Email)
	}
}
