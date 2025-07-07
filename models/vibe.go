package models

import "gorm.io/gorm"

type Vibe struct {
    gorm.Model
    Label string `json:"label"`
    Emoji string `json:"emoji"`
}
