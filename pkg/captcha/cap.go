package captcha

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
)

var (
	// IsInitialize enable/disable MemoryStorage creation
	IsInitialize = false
	store        captcha.Store
)

// InitializeMemoryStorage create a new MemoryStorage object
// uses it as the default one and change the IsInitialize value
func InitializeMemoryStorage() {
	IsInitialize = !IsInitialize

	memory := captcha.NewMemoryStore(1000, time.Hour*1)
	captcha.SetCustomStore(memory)
	store = memory

}

// CreateCaptchaAndShow create a captcha and load it to the /captcha endpoint
func CreateCaptchaAndShow(resp http.ResponseWriter) {
	// Create a new captcha and store it
	id := captcha.New()
	// Get digits based on the captcha id
	digits := store.Get(id, false)
	// Create an image based on the id and digits of captcha
	image := captcha.NewImage(id, digits, captcha.StdWidth, captcha.StdHeight)

	resp.Header().Set("X-Captcha-ID", id)
	image.WriteTo(resp)

}

// VerifyCaptcha verify if a captcha is valid
func VerifyCaptcha(id string, digits string) bool {
	storedDigits := store.Get(id, true)

	storedString := fromByteToString(storedDigits)

	return storedString == digits

}

func fromByteToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	var str string
	for _, v := range b {
		str += fmt.Sprintf("%v,", v)

	}

	return str[:len(str)-1]
}
