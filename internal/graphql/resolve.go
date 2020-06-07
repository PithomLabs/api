package graphql

import (
	"fmt"
	"strconv"

	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/structs"
)

//RootResolver represents a root query for graphql
type RootResolver struct{}

//User resolves user type
func (r *RootResolver) User(args struct{ UserID int32 }) (*UserResolver, error) {
	user, err := database.GetUserByID(int64(args.UserID))
	if err != nil {
		return &UserResolver{&structs.User{}}, err
	}
	return &UserResolver{user}, err
}

//Posts resolves Posts field
func (r *RootResolver) Posts(args struct{ UserID int32 }) (*[]*PostResolver, error) {
	pResolvers := &[]*PostResolver{}
	posts, err := database.GetAllEntitiesFromUser(int64(args.UserID), "post")
	if err != nil {
		return pResolvers, err
	}
	for _, post := range *posts {
		*pResolvers = append(*pResolvers, &PostResolver{post})
	}
	return pResolvers, err
}

//Post resolves single post
func (r *RootResolver) Post(args struct{ PostID int32 }) (*PostResolver, error) {
	post, err := database.GetEntityByID(int64(args.PostID), "post")
	if err != nil {
		return &PostResolver{post}, err
	}
	return &PostResolver{post}, nil
}

//Comment is a resolver for comment
//It returns PostResolver, because comment and post is the same entity.
//So, the way to resolve them is same.
func (r *RootResolver) Comment(args struct{ CommentID int32 }) (*PostResolver, error) {
	return r.Post(struct{ PostID int32 }{args.CommentID})
}

//Comments is root resolver for comments by userID or PostID
func (r *RootResolver) Comments(args struct {
	PostID int32
	UserID int32
}) (*[]*PostResolver, error) {
	if args.PostID != 0 {
		comments, err := database.GetEntitiesByAnswerOf(int64(args.PostID))
		if err != nil {
			return nil, err
		}
		prs := make([]*PostResolver, 0)
		for _, comment := range comments {
			prs = append(prs, &PostResolver{comment})
		}
		return &prs, nil
	}
	if args.UserID != 0 {
		return r.Posts(struct{ UserID int32 }{int32(args.UserID)})
	}
	return nil, fmt.Errorf("expected any of given args to be in query: PostID, UserID")
}

//UserInfo resolves info about author of post
func (r *RootResolver) UserInfo(args struct{ UserID int32 }) (*UserResolver, error) {
	return r.User(struct{ UserID int32 }{args.UserID})
}

//UserResolver ...
type UserResolver struct{ *structs.User }

//ID ...
func (ur *UserResolver) ID() int32 {
	return int32(ur.User.ID)
}

//Username ...
func (ur *UserResolver) Username() string {
	return ur.User.Username
}

//Fullname ...
func (ur *UserResolver) Fullname() *string {
	return &ur.User.Fullname
}

//Bio ...
func (ur *UserResolver) Bio() *string {
	return &ur.User.Bio
}

//CreatedAt ...
func (ur *UserResolver) CreatedAt() string {
	return strconv.Itoa(int(ur.User.CreatedAt))
}
func (ur *UserResolver) AvatarUrl() string {
	return ur.AvatarURL
}

//Posts ...
func (ur *UserResolver) Posts() (*[]*PostResolver, error) {
	root := &RootResolver{}
	return root.Posts(struct{ UserID int32 }{int32(ur.User.ID)})
}

//Comments is comment resolver for post
func (ur *UserResolver) Comments() (*[]*PostResolver, error) {
	root := &RootResolver{}
	return root.Comments(struct {
		PostID int32
		UserID int32
	}{UserID: int32(ur.User.ID)})
}

func (ur *UserResolver) Nsfw() bool {
	return ur.User.Settings.ShowNSFW
}

func (ur *UserResolver) Settings() *UserSettingsResolver {
	usr := &UserSettingsResolver{ur.User.Settings}
	return usr
}

type UserSettingsResolver struct{ structs.Settings }

func (usr *UserSettingsResolver) ShowNsfw() bool {
	return usr.ShowNSFW
}
func (usr *UserSettingsResolver) ShowLikes() bool {
	return usr.Settings.ShowLikes
}
func (usr *UserSettingsResolver) NsfwPage() bool {
	return usr.Settings.NSFWPage
}

//PostResolver ...
type PostResolver struct{ *structs.Entity }

//ID ...
func (pr *PostResolver) ID() int32 {
	return int32(pr.Entity.ID)
}

//Likes ...
func (pr *PostResolver) Likes() int32 {
	return int32(pr.Entity.Likes)
}

//Liked ...
// func (pr *PostResolver) Liked() *bool {
// 	return pr.L
// }

//Inside ...
func (pr *PostResolver) Inside() (*ContentResolver, error) {
	return &ContentResolver{&pr.Entity.Inside}, nil
}

//Comments is comment resolver for post
func (pr *PostResolver) Comments() (*[]*PostResolver, error) {
	root := &RootResolver{}
	return root.Comments(struct {
		PostID int32
		UserID int32
	}{PostID: int32(pr.Entity.ID)})
}

//CreatedAt ...
func (pr *PostResolver) CreatedAt() string {
	ca := strconv.Itoa(int(pr.Entity.CreatedAt))
	return ca
}

//EditedAt ...
func (pr *PostResolver) EditedAt() *string {
	ea := strconv.Itoa(int(pr.Entity.EditedAt))
	return &ea
}

//ContentResolver resolves content of post
type ContentResolver struct{ *structs.Content }

//Type ...
func (cr *ContentResolver) Type() string {
	return cr.Content.Type
}

//Text ...
func (cr *ContentResolver) Text() string {
	return cr.Content.Text
}

//Nsfw ...
func (cr *ContentResolver) Nsfw() bool {
	return cr.Content.NSFW
}

//Author resolves author field for post
func (pr *PostResolver) Author() (*UserResolver, error) {
	root := &RootResolver{}
	return root.UserInfo(struct{ UserID int32 }{int32(pr.UserID)})
}

//Source resolves
func (cr *ContentResolver) Source() *[]*AssetResolver {
	assetResolvers := make([]*AssetResolver, 0)
	for _, source := range cr.Content.Source {
		assetResolvers = append(assetResolvers, &AssetResolver{&source})
	}
	return &assetResolvers
}

//AssetResolver ...
type AssetResolver struct{ *structs.Asset }

//ID ...
func (ar *AssetResolver) ID() int32 {
	return int32(ar.Asset.ID)
}

//Width ...
func (ar *AssetResolver) Width() int32 {
	return int32(ar.Asset.Width)
}

//Height ...
func (ar *AssetResolver) Height() int32 {
	return int32(ar.Asset.Height)
}

//ResourceType ...
func (ar *AssetResolver) ResourceType() *string {
	return &ar.Asset.ResourceType
}

//Url ...
func (ar *AssetResolver) Url() *string {
	return &ar.Asset.URL
}

//SecureUrl ...
func (ar *AssetResolver) SecureUrl() *string {
	return &ar.Asset.SecureURL
}

//CreatedAt ...
func (ar *AssetResolver) CreatedAt() string {
	return strconv.Itoa(int(ar.Asset.CreatedAt))
}
