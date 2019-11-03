package lambdas

import (
	"log"
	"net/http"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/mail"
)

// VerifyHandler is the endpoint for /verify which is used for mail verification
func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	verificationCode := req.URL.Query().Get("verify_code")

	verr := mail.Verify(verificationCode)
	if verr != nil {
		err.ShowOnBrowser(resp, verr)
		log.Print(verr)
	}

	resp.Write([]byte("Account email as been verified"))
}
