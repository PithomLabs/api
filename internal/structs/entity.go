package structs

// Entity struct represent the Posts and Comments
type Entity struct {
	ID     uint `gorm:"primary_key"`
	UserID uint `gorm:"primary_key"`
	// This Type field correspond to the ENTITY_TYPE
	Type      string
	Inside    Content `gorm:"EMBEDDED"`
	Likes     uint
	CreatedAt uint64 `gorm:"column:created_at" json:"created_at"`
	EditedAt  uint64 `gorm:"column:edited_at" json:"edited_at"`
	AnswerOf  uint   `gorm:"column:answer_of"`
}

// Content represent the inside of the post
type Content struct {
	Type        string `gorm:"column:content_type"`
	Description string `gorm:"column:description"`
	Source      string `gorm:"column:source"`
	NSFW        bool   `gorm:"column:NSFW"`
}
