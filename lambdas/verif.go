package lambdas

import (
	"log"
	"net/http"

	mail "github.com/komfy/api/pkg/email"
)

// VerifyHandler is the endpoint for /verify which is used for mail verification
func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	verificationCode := req.URL.Query().Get("token")

	err := mail.VerifyUser(verificationCode)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
}
