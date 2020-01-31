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

	rErr, missingPasswordCriteria := register.NewUser(req)
	if len(missingPasswordCriteria) > 0 {
		// Whether the SendJSON functions succeeded or not
		// We need to return from the function so we avoid the redirection
		// Plus we need to log.Println the previous error
		log.Println(rErr)
		jErr := err.SendJSON(resp, missingPasswordCriteria)
		if jErr != nil {
			err.ShowOnBrowser(resp, jErr)
			log.Println(jErr)
		}
		return
	}
	if rErr != nil {
		err.ShowOnBrowser(resp, rErr)
		log.Println(rErr)
		return
	}

	http.Redirect(resp, req, redirectRegURL, http.StatusSeeOther)
}
