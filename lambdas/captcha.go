package lambdas

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/komfy/api/internal/captcha"
	err "github.com/komfy/api/pkg/error"
)

// GetCaptchaHandler is the handler function for captcha generation
// And image representation of it
func GetCaptchaHandler(resp http.ResponseWriter, req *http.Request) {
	/*id, digits := captcha.CreateCaptchaAndShow(resp)

	log.Printf("Captcha ID: %s", id)
	log.Printf("  > Digits: %v", digits)*/

}

// VerifyCaptchaHandler is the handler function for captcha verification
func VerifyCaptchaHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	resp.Header().Set("Access-Control-Allow-Headers", "X-Captcha-ID")

	fmt.Println(req.Method)

	if req.Method == "OPTIONS" {
		return

	}

	id := req.Header.Get("X-Captcha-ID")
	digits := req.URL.Query().Get("digits")

	if id != "" && digits != "" {
		/*if captcha.VerifyCaptcha(id, digits, false) {
			resp.Write([]byte("Captcha has been solved"))

		} else {
			err.HandleErrorInHTTP(resp, err.ErrCaptchaInvalid)
			log.Print(err.ErrCaptchaInvalid.Error())

		}*/

	} else if id == "" {
		err.HandleErrorInHTTP(resp, err.ErrCaptchaHeaderMissing)
		log.Print(err.ErrCaptchaHeaderMissing.Error())

	} else if digits == "" {
		err.HandleErrorInHTTP(resp, err.ErrDigitsMissing)
		log.Print(err.ErrDigitsMissing.Error())

	}

}
