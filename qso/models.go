package qso

import (
	"github.com/jinzhu/gorm"
)

type Qso struct {
	gorm.Model
	Dateon   string `json:"dateon"`
	Timeon   string `json:"timeon"`
	Callsign string    `gorm:"Not Null" json:"callsign"`
	Band     int       `gorm:"Not Null" json:"band"`
	Mode     string    `gorm:"Not Null" json:"mode"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	County   string    `json:"county"`
	Country  string    `json:"country"`
	Name     string    `json:"name"`
	Qslr     bool      `json:"qslr"`
	Qsls     bool      `json:"qsls"`
	Rstr     int       `json:"rstr"`
	Rsts     int       `json:"rsts"`
	Power    int       `json:"power"`
	Remarks  string    `json:"remarks"`
}
