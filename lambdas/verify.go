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

	vErr := mail.Verify(verificationCode)
	if vErr != nil {
		err.ShowOnBrowser(resp, vErr)
		log.Print(vErr)
		return
	}

	resp.Write([]byte("Account email as been verified"))
}
