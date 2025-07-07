package models

type Song struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Artist      string
	YouTubeLink string

	VibeID     uint
	Vibe       Vibe `gorm:"foreignKey:VibeID"`

	TimeSlotID uint
	TimeSlot   TimeSlot `gorm:"foreignKey:TimeSlotID"`
}

