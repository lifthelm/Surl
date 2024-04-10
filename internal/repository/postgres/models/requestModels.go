package models

type RequestProjectOwner struct {
	ProjectID   uint `gorm:"primaryKey"`
	Name        string
	Description string
	OwnerID     uint
}

type RequestProjectRole struct {
	ProjectID   uint `gorm:"primaryKey"`
	Name        string
	Description string
	UserRole    string
}

type RequestRouteFilter struct {
	LinkID     uint
	PlatformID uint
	AZoneID    uint
}
