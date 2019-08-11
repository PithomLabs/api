package lambdas

import (
	"log"
	"net/http"

	"github.com/komfy/api/pkg/captcha"
	err "github.com/komfy/api/pkg/error"
)

// GetCaptchaHandler is the handler function for captcha generation
// And image representation of it
func GetCaptchaHandler(resp http.ResponseWriter, req *http.Request) {
	captcha.CreateCaptchaAndShow(resp)

}

// VerifyCaptchaHandler is the handler function for captcha verification
func VerifyCaptchaHandler(resp http.ResponseWriter, req *http.Request) {
	id := req.Header.Get("X-Captcha-ID")
	digits := req.URL.Query().Get("digits")

	if id != "" && digits != "" {
		if captcha.VerifyCaptcha(id, digits) {
			resp.Write([]byte("Captcha has been solved"))

		} else {
			resp.Write([]byte("Captcha has not been solved, try again with another one"))

		}

	} else if id == "" {
		err.HandleErrorInHTTP(resp, err.ErrCaptchaHeaderMissing)
		log.Print(err.ErrCaptchaHeaderMissing.Error())

	} else if digits == "" {
		err.HandleErrorInHTTP(resp, err.ErrDigitsMissing)
		log.Print(err.ErrDigitsMissing.Error())

	}

}
