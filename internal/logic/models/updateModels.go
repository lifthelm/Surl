package models

type UpdateUser struct {
	UserName *string
	Email    *string
	Password *string
	UserRole *UserRole
}

type UpdateProject struct {
	Name        *string
	Description *string
}

type UpdateLink struct {
	URLShort   *string
	DefLongURL *string
}

type UpdateRoute struct {
	LongURL  *string
	Platform *PlatformType
	AZone    *GeoType
}
