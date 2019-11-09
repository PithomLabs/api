package sign

import (
	"github.com/komfy/api/internal/structs"
)

type Transport struct {
	User  *structs.User
	JWT   string
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

func CreateJWTTransport(transportJWT string) Transport {
	return Transport{
		JWT:   transportJWT,
		Error: nil,
	}
}
