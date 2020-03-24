package lambdas

import (
	"fmt"
	"log"
	"net/http"
	"os"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign/authenticate"
)

var redirectAuthURL = os.Getenv("redir_auth_url") + "/set_cookie?token=%s"

// AuthenticationHandler handle the /auth endpoint
func AuthenticationHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if req.Method == http.MethodOptions {
		return
	}

	result := authenticate.User(req)
	if result.Error != nil {
		err.ShowOnBrowser(resp, result.Error)
		log.Println(result.Error)
		return
	}

	url := fmt.Sprintf(redirectAuthURL, result.JWT)

	http.Redirect(resp, req, url, http.StatusSeeOther)
}
