package models

type GarbageDay struct {
	BaseModel
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`

	DayOfWeek string `gorm:"size:255" json:"day_of_week,omitempty"`

	Garbage string `gorm:"size:255" json:"garbage,omitempty"`
}

type BaseGarbageDayResponse struct {
	BaseModel
	Garbage string `gorm:"size:255" json:"garbage,omitempty"`
}
