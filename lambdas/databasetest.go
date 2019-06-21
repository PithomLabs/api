package lambdas

import (
	"fmt"
	"net/http"

	dtb "github.com/komfy/api/pkg/database"
)

func DatabaseTest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	db := dtb.OpenDatabase()
	defer db.CloseDB()

	user := db.AskUserByID(query.Get("user_id"))
	fmt.Fprintln(w, user)

}
