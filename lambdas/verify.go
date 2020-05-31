package lambdas

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/mail"
)

// VerifyHandler is the endpoint for /verify which is used for mail verification
func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	verificationCode := req.URL.Query().Get("verify_code")

	vErr := mail.Verify(graphql.ID(verificationCode))
	if vErr != nil {
		err.ShowOnBrowser(resp, vErr)
		log.Println(vErr)
		return
	}

	resp.Write([]byte("Account email as been verified"))
}
