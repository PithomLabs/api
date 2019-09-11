package lambdas

import (
	"log"
	"net/http"

	mail "github.com/komfy/api/pkg/email"
	err "github.com/komfy/api/pkg/error"
)

// VerifyHandler is the endpoint for /verify which is used for mail verification
func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	verificationCode := req.URL.Query().Get("verify_code")

	verr := mail.VerifyUser(verificationCode)
	if verr != nil {
		err.HandleErrorInHTTP(resp, verr)
		log.Print(verr)
	}
}
