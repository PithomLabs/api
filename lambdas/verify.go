package lambdas

import (
	"log"
	"net/http"
	"strconv"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/mail"
)

// VerifyHandler is the endpoint for /verify which is used for mail verification
func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	verificationCode := req.URL.Query().Get("verify_code")

	code, cErr := strconv.ParseUint(verificationCode, 10, 64)
	if cErr != nil {
		err.ShowOnBrowser(resp, cErr)
		log.Println(cErr)
		return
	}

	vErr := mail.Verify(uint(code))
	if vErr != nil {
		err.ShowOnBrowser(resp, vErr)
		log.Println(vErr)
		return
	}

	resp.Write([]byte("Account email as been verified"))
}
