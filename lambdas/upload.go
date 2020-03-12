package lambdas

import (
	"encoding/json"
	"net/http"

	"github.com/komfy/cloudinary"
)

//UploadHandler handles upload requests for cloudinary.
type UploadHandler struct {
	service *cloudinary.Service
}

func (uh *UploadHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	if uh.service == nil {
		http.Error(res, "cloudinary service is not initiated, it's empty", http.StatusBadRequest)
	}
	file, fh, err := req.FormFile("image")
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	fname := fh.Filename
	upResp, err := uh.service.Upload(fname, file, false)
	res.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(upResp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.WriteHeader(200)
	res.Write(json)
}
