package sign

import (
	"github.com/komfy/api/internal/structs"
)

// A simple struct in order to transport information
// from functions inside the authentication and registration
type Transport struct {
	User  *structs.User
	Error error
	Bool  bool
}



func CreateErrorTransport(transportError error) Transport {
	return Transport{
		User:  nil,
		Error: transportError,
	}
}

func CreateUserTransport(transportUser *structs.User) Transport {
	return Transport{
		User:  transportUser,
		Error: nil,
	}
}

func CreateBoolTransport(transportBool bool) Transport {
	return Transport{
		Bool:  transportBool,
		Error: nil,
	}
}
