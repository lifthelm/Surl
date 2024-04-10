package models

import (
	"math/rand"
	"time"
)

type UUID uint

func NewUUID() UUID {
	return UUID(rand.Int())
} // TODO implement proper UUID generation

type UserRole byte

const (
	UserRegular UserRole = iota
	UserAdmin
)

type UserLinkRole byte

const (
	LinkOwner UserLinkRole = iota
	LinkRedactor
)

type UserProjectRole byte

const (
	ProjectOwner UserProjectRole = iota
	ProjectMember
)

type User struct {
	CreatedAt time.Time
	UserName  string
	Email     string
	Password  string
	ID        UUID
	UserRole  UserRole
}

type UserLink struct {
	ID       UUID
	UserID   UUID
	LinkID   UUID
	UserRole UserLinkRole
}

type UserProject struct {
	ID        UUID
	UserID    UUID
	ProjectID UUID
	UserRole  UserProjectRole
}

type Project struct {
	Name        string
	Description string
	ID          UUID
}

type ProjectLink struct {
	ID        UUID
	ProjectID UUID
	LinkID    UUID
}

type Link struct {
	CreatedAt  time.Time
	URLShort   string
	DefLongURL string
	ID         UUID
}

type Platform struct {
	Description string
	ID          UUID
}

type AvailabilityZone struct {
	Description string
	ID          UUID
}

type Route struct {
	LongURL  string
	LinkID   UUID
	ID       UUID
	Platform PlatformType
	AZone    GeoType
}

type Stat struct {
	CreatedAt    time.Time
	UserAgent    string
	UserIP       string
	RouteID      UUID
	ID           UUID
	UserPlatform PlatformType
	UserAZone    GeoType
}
