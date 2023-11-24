package qso

import (
	"github.com/jinzhu/gorm"
)

type Qso struct {
	gorm.Model
	Dateon   string `json:"dateon" example:"2023-10-28"`
	Timeon   string `json:"timeon" example:"16:34:58.947Z"`
	Callsign string `gorm:"Not Null" json:"callsign" example:"W1AW"`
	Band     int    `gorm:"Not Null" json:"band" example:"20"`
	Mode     string `gorm:"Not Null" json:"mode" example:"cw"`
	City     string `json:"city" example:"New Haven"`
	State    string `json:"state" example:"CT"`
	County   string `json:"county" example:"New Haven"`
	Country  string `json:"country" example:"USA"`
	Name     string `json:"name" example:"John Smith"`
	Qslr     bool   `json:"qslr" example:"false"`
	Qsls     bool   `json:"qsls" example:"false"`
	Rstr     int    `json:"rstr" example:"599"`
	Rsts     int    `json:"rsts" example:"599"`
	Power    int    `json:"power" example:"100"`
	Remarks  string `json:"remarks" example:"This is a test contact with W1AW"`
}
