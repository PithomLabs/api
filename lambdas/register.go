package lambdas

import (
	"log"
	"net/http"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign/register"
)

func RegisterHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	resp.Header().Set("Access-Control-Allow-Headers", "X-Captcha")

	if req.Method == http.MethodOptions {
		return
	}

	// validationErrors is a string here
	validationErrors, rErr := register.NewUser(req)
	if len(validationErrors) > 0 {
		// Show and log rErr error first so there is no
		// WriteHeader duplicate
		err.ShowOnBrowser(resp, rErr)
		log.Println(rErr)

		sjErr := err.SendJSON(resp, validationErrors)
		if sjErr != nil {
			log.Println(sjErr)
			return
		}
		return
	}

	if rErr != nil {
		err.ShowOnBrowser(resp, rErr)
		log.Println(rErr)
		return
	}

	// resp.Write will mark the request as 200
	resp.Write([]byte("Successful request"))

}
