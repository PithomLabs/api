package graphql

import (
	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/structs"
)

//InputProfile is an input for user modification
type inputProfile struct {
	ID        int32
	Fullname  *string
	Bio       *string
	AvatarURL *string
	Username  *string
	Settings  *inputSettings
}
type inputSettings struct {
	ShowLikes *bool
	ShowNSFW  *bool
	NSFWPage  *bool
}

//UpdateProfile is a graphql a resolver of gql mutation updateProfile
func (r *RootResolver) UpdateProfile(args struct {
	Profile inputProfile
}) (*UserResolver, error) {
	u := &structs.User{
		ID:        int64(args.Profile.ID),
		Username:  *args.Profile.Username,
		Fullname:  *args.Profile.Fullname,
		Bio:       *args.Profile.Bio,
		AvatarURL: *args.Profile.AvatarURL,
		Settings: structs.Settings{
			ShowLikes: *args.Profile.Settings.ShowLikes,
			NSFWPage:  *args.Profile.Settings.NSFWPage,
			ShowNSFW:  *args.Profile.Settings.ShowNSFW,
		},
	}
	u, err := database.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	ur := &UserResolver{u}
	return ur, nil

}
