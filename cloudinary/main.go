//Package cloudinary provides an easy way of connection between go and cloudinary
package cloudinary

import (
	"errors"
	"fmt"
	"net/url"
	"os"
)

//Service represents cloudinary service
type Service struct {
	cloudName string
	apiKey    string
	apiSecret string
	uploadURL *url.URL
	adminURL  *url.URL
	verbose   bool
	simulate  bool
	resType   int
}

const (
	baseUploadURL string = "https://api.cloudinary.com/v1_1"
	imageType     int    = 0
)

//Dial configurates cloudinary service
// cloudinary://api_key:api_secret@cloud_name
func Dial(uri string) (*Service, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "cloudinary" {
		return nil, errors.New("URL scheme is not cloudinary")
	}
	secret, exists := u.User.Password()
	if !exists {
		return nil, errors.New("There is no api secret provided in URL")
	}
	s := &Service{
		cloudName: u.Hostname(),
		apiKey:    u.User.Username(),
		apiSecret: secret,
		resType:   imageType,
		simulate:  false,
		verbose:   false,
	}
	up, err := url.Parse(fmt.Sprintf("%s/%s/image/upload/", baseUploadURL, s.cloudName))
	if err != nil {
		return nil, err
	}
	s.uploadURL = up
	admURL, err := url.Parse(fmt.Sprintf("%s/%s", baseUploadURL, s.cloudName))
	if err != nil {
		return nil, err
	}
	admURL.User = url.UserPassword(s.apiKey, s.apiSecret)
	s.adminURL = admURL

	return s, nil
}

//UploadFile receives file, most probably from Multipart Form, uploads to cloud and returns
//a link to the resource
func (s *Service) UploadFile(file *os.File, randomPublicID bool) (string, error) {
	var publicID string
	fi, err := file.Stat()
	if err != nil {
		return "", err
	}
	if fi.Size() == 0 {
		return "", fmt.Errorf("Not allowed to upload empty files: ", fi.Name())
	}
	if !randomPublicID {
		publicID = fi.Name()
	} else {
		publicID = ""
	}

	return "", nil
}
