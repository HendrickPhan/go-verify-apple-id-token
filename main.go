package main

import (
	"fmt"

	"github.com/HendrickPhan/go-verify-apple-id-token/validator"
)

var idToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6ImJkOWVjZDNiODNhOTNkMGNjMTczODBmZGQzZTRiZTVmIn0.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLmV4YW1wbGUuYXBwIiwiZXhwIjoxNzI3MjU1MjIzLCJpYXQiOjE3MjcxNjg4MjMsInN1YiI6InJhbmRvbV9hcHBsZV91c2VyX2lkIiwiY19oYXNoIjoicmFuZG9tX2NfaGFzX3ZhbHVlIiwiZW1haWwiOiJyYW5kb20tdmFsdWVAcHJpdmF0ZXJlbGF5LmFwcGxlaWQuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImlzX3ByaXZhdGVfZW1haWwiOnRydWUsImF1dGhfdGltZSI6MTcyNzE2ODgyMywibm9uY2Vfc3VwcG9ydGVkIjp0cnVlfQ.VqR1igwpz1uaNQznmjFicjDySSQJD-nj5ToiS1kwSyqtNjDUpgiGW3hvSjjdsQkpcKmKfayNjE-30O_ewB2Nhg"

func main() {
	client := validator.NewClient()

	jwtClaims, err := client.VerifyIdToken("com.example.app", idToken)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(jwtClaims.Email)
	}
}
