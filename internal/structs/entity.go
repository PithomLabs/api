package structs

import "github.com/graph-gophers/graphql-go"

// Entity struct represent the Posts and Comments
type Entity struct {
	ID     graphql.ID `gorm:"column:entity_id;primary_key"`
	UserID uint
	// This Type field correspond to the ENTITY_TYPE
	Type      string  `gorm:"default:'post'"`
	Inside    Content `gorm:"EMBEDDED"`
	Likes     uint
	CreatedAt uint64     `gorm:"column:created_at" json:"created_at"`
	EditedAt  uint64     `gorm:"column:edited_at" json:"edited_at"`
	AnswerOf  graphql.ID `gorm:"column:answer_of" json:"answer_of"`
	// This is only used for recursivity and filtering
	Depth uint `gorm:"-" json:"-"`
}

// Content represent the inside of the post
type Content struct {
	Type   string  `gorm:"column:content_type;default:'text'"`
	Text   string  `gorm:"column:text"`
	NSFW   bool    `gorm:"column:NSFW"`
	Source []Asset `gorm:"-"`
}

type Asset struct {
	ID           graphql.ID `gorm:"column:asset_id;primary_key"`
	Width        uint
	Height       uint
	ResourceType string `gorm:"column:resource_type;default:'image'" json:"resource_type"`
	URL          string
	SecureURL    string `gorm:"column:secure_url" json:"secure_url"`
	CreatedAt    uint64 `gorm:"column:created_at" json:"created_at"`
}
