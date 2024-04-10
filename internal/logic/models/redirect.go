package models

type RequestData struct {
	UserAgent string
	IP        string
}

type GeoType byte

const (
	GeoGeneral GeoType = iota
	GeoAmerica
	GeoAsia
	GeoEurope
	GeoAfrica
)

type PlatformType byte

const (
	PlatformGeneral PlatformType = iota
	PlatformAndroid
	PlatformIOS
)

type RouteFilter struct {
	Platform PlatformType
	AZone    GeoType
}
