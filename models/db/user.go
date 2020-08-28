package db

import "time"

type User struct {
	ID                     int       `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Avatar                 string    `gorm:"type:varchar(300);NOT NULL"`
	Name                   string    `gorm:"type:varchar(50);NOT NULL"`
	Password               string    `gorm:"type:varchar(70);unique_index;NOT NULL"`
	Email                  string    `gorm:"type:varchar(60);unique_index;NOT NULL"`
	Verified               bool      `gorm:"NOT NULL;defualt:default:0"`
	JoinAt                 time.Time `gorm:"NOT NULL"`
	VerifyCode             string    `gorm:"type:varchar(10)"`
	VerifyCodePassword     string    `gorm:"type:varchar(10)"`
	VerifyCodePasswordTime *time.Time
}
