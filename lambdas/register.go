package lambdas

import (
	"log"
	"net/http"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign/register"
)

const redirectRegURL = "https://komfy.now.sh/verify_email"

func RegisterHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	resp.Header().Set("Access-Control-Allow-Headers", "X-Captcha")

	if req.Method == http.MethodOptions {
		return
	}

	rErr := register.NewUser(req)
	if rErr != nil {
		err.ShowOnBrowser(resp, rErr)
		log.Println(rErr)
		return
	}

	http.Redirect(resp, req, redirectRegURL, http.StatusSeeOther)

}
