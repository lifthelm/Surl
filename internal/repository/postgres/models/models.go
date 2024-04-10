package models

import (
	"time"
)

const (
	UserRegular = "user"
	UserAdmin   = "admin"

	ProjectOwner  = "owner"
	ProjectMember = "member"

	LinkOwner    = "owner"
	LinkRedactor = "redactor"
)

type User struct {
	UserID           uint   `gorm:"primaryKey"`
	UserName         string `gorm:"column:username"`
	Email            string
	Password         string
	RegistrationDate time.Time
	UserRole         string
}

type Project struct {
	ProjectID   uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type UserProjectRelation struct {
	RelationID uint `gorm:"primaryKey"`
	UserID     uint
	ProjectID  uint
	UserRole   string
}

type Link struct {
	LinkID     uint `gorm:"primaryKey"`
	DefLongURL string
	ShortURL   string
	CreateDate time.Time
}

type UserLinkRelation struct {
	RelationID uint `gorm:"primaryKey"`
	UserID     uint
	LinkID     uint
	UserRole   string
}

type ProjectLinkRelation struct {
	RelationID uint `gorm:"primaryKey"`
	ProjectID  uint
	LinkID     uint
}

type Platform struct {
	PlatformID  uint `gorm:"primaryKey"`
	Description string
}

type AvailabilityZone struct {
	AZoneID     uint `gorm:"primaryKey"`
	Description string
}

type LinkRoute struct {
	RouteID    uint `gorm:"primaryKey"`
	LinkID     uint
	PlatformID uint
	AZoneID    uint
	LongURL    string
}

type ClickStat struct {
	RecordID      uint `gorm:"primaryKey"`
	RouteID       uint
	ClickDateTime time.Time
	UserIPAddress string
	UserAgent     string
	UserCountry   string
}
