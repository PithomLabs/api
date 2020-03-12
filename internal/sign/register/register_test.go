package register

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/komfy/api/internal/sign"
	"github.com/komfy/api/internal/structs"
)

func TestExtractUserFromJson(t *testing.T) {
	t.Run("Passing test with jsons", func(t *testing.T) {
		expect := structs.User{1, "Koshqua", "Kind$OfPass123", "myemail@gmail.com", "", "", "", 1233, true, structs.Settings{}}
		reqBody, err := json.Marshal(expect)
		if err != nil {
			t.Errorf("Impossible to marshal json")
		}
		req, err := http.NewRequest("POST", "https://localhost:3000/reg", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Errorf("impossible to send req")
		}
		req.Header.Set("Content-type", jSON)
		userChan := make(chan (sign.Transport))
		go extractUser(req, userChan)
		info := <-userChan
		close(userChan)
		if !reflect.DeepEqual(expect, *info.User) {
			t.Errorf("\nExpected to get %+v, \ngot %+v", expect, *info.User)
		}
	})
	t.Run("Passing tests with UrlEncoded", func(t *testing.T) {
		expect := structs.User{1, "Koshqua", "Kind$OfPass123", "myemail@gmail.com", "", "", "", 1233, true, structs.Settings{}}
		reqStr := "username=Koshqua&password=Kind$OfPass123&email=myemail@gmail.com"
		req, err := http.NewRequest("POST", "https://localhost:3000/reg", bytes.NewBufferString(reqStr))
		if err != nil {
			t.Errorf("impossible to send req")
		}
		req.Header.Set("Content-type", urlencoded)
		userChan := make(chan (sign.Transport))
		go extractUser(req, userChan)
		info := <-userChan
		close(userChan)
		if !reflect.DeepEqual(expect, *info.User) {
			t.Errorf("\nExpected to get %+v, \ngot %+v", expect, *info.User)
		}
	})
}
