package models

import "gorm.io/gorm"

type TimeSlot struct {
    gorm.Model
    Label     string `json:"label"`
    StartHour int    `json:"start_hour"`
    EndHour   int    `json:"end_hour"`
    Emoji     string `json:"emoji"`
}

