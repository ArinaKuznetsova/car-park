package schemas

type Route struct {
	ID int `json:"route"`
}
type Driver struct {
	Fullname string `json:"fullname"`
	Class int `json:"class"`
	Exp float64 `json:"experience"`
	IdBus int `json:"idBus"`
}

type Drivers struct {
	Name []string `json:"name"`
}
type Buses struct {
	Nums []int `json:"numbers"`
}
type RoutesMove struct {
	IdRoute int `json:"route"`
	Time string	`json:"moveTime"`
}
type RoutesM struct {
	Data []RoutesMove `json:"data"`
}

type RoutesLength struct {
	IdRoute int `json:"route"`
	Length float64 `json:"length"`
}
type RoutesL struct {
	Data []RoutesLength `json:"data"`
}
type RouteInfo struct {
	IdRoute int `json:"route"`
	MovStart string `json:"moveStart"`
	BusInterval int `json:"busInterval"`
}
type DriverInfo struct {
	Fullname string `json:"fullname"`
	Class int `json:"class"`
}
type Info struct {
	CountBus int `json:"countBus"`
	BusTypes []string `json:"busTypes"`
	Routes []RouteInfo `json:"routes"`
	Drivers []DriverInfo `json:"drivers"`
}
