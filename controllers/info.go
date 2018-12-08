package controllers

import (
	"database/sql"
	_"github.com/lib/pq"
	"encoding/json"
	"github.com/DonRIn/carpark/db"
	"github.com/DonRIn/carpark/schemas"
	"github.com/labstack/echo"
	"net/http"
)


func GetDrivers(c echo.Context) error {
	idRoute := c.Param("idRoute")
	rows, err := db.Conn.Query(`SELECT fullname FROM drivers JOIN specifications ON drivers.id_bus=specifications.id_bus JOIN routes ON routes.id_route = specifications.id_route WHERE routes.id_route=$1`,idRoute)
	if err != nil { return err }

	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.Drivers{Name: names})
}

func GetBuses(c echo.Context) error {
	idRoute := c.Param("idRoute")
	rows, err := db.Conn.Query(`SELECT id_bus FROM specifications JOIN routes ON routes.id_route = specifications.id_route WHERE routes.id_route=$1`,idRoute)
	if err != nil { return err }

	defer rows.Close()
	nums := make([]int, 0)
	for rows.Next() {
		var num int
		if err := rows.Scan(&num); err != nil {
			return err
		}
		nums = append(nums, num)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.Buses{Nums: nums})
}

func GetTimeRoutes(c echo.Context) error {
	idRoute := c.Param("idRoute")
	time := c.Param("time")
	var rows *sql.Rows
	if idRoute != "all" {
		if time == "start" {
			rows, _ = db.Conn.Query(`SELECT id_route, mov_start FROM routes WHERE id_route=$1`,idRoute)
		} else {
			rows, _ = db.Conn.Query(`SELECT id_route, mov_end FROM routes WHERE id_route=$1`,idRoute)
		}

	} else {
		if time == "start" {
			rows, _ = db.Conn.Query(`SELECT id_route, mov_start FROM routes`)
		} else {
			rows, _ = db.Conn.Query(`SELECT id_route, mov_end FROM routes`)
		}
	}

	defer rows.Close()
	routes := make([]schemas.RoutesMove, 0)
	for rows.Next() {
		var route schemas.RoutesMove
		if err := rows.Scan(&route.IdRoute, &route.Time); err != nil {
			return err
		}
		routes = append(routes, route)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.RoutesM{Data: routes})
}

func GetLengthRoutes(c echo.Context) error {
	idRoute := c.Param("idRoute")
	var rows *sql.Rows
	if idRoute != "all" {
		rows, _ = db.Conn.Query(`SELECT id_route, length FROM routes WHERE id_route=$1`,idRoute)
	} else {
		rows, _ = db.Conn.Query(`SELECT id_route, length FROM routes`)
	}

	defer rows.Close()
	routes := make([]schemas.RoutesLength, 0)
	for rows.Next() {
		var route schemas.RoutesLength
		if err := rows.Scan(&route.IdRoute, &route.Length); err != nil {
			return err
		}
		routes = append(routes, route)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.RoutesL{Data: routes})
}

func GetDriverBus(c echo.Context) error {
	idBus := c.Param("idBus")
	rows, _ := db.Conn.Query(`SELECT fullname FROM drivers JOIN specifications ON drivers.id_bus=specifications.id_bus WHERE specifications.id_bus=$1`,idBus)

	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.Drivers{Name: names})

}

func AddDriver(c echo.Context) error {
	dr := new(schemas.Driver)
	if err := c.Bind(dr); err != nil {
		return err
	}
	_, err := db.Conn.Query(`INSERT INTO drivers VALUES ($1, $2, $3, $4)`,dr.Fullname,dr.Class,dr.Exp,dr.IdBus)
	return err
}

func PutLength(c echo.Context) error {
	rl := new(schemas.RoutesLength)
	if err := c.Bind(rl); err != nil {
		return err
	}
	_, err := db.Conn.Query(`UPDATE routes SET length=$1 WHERE id_route=$2`, rl.Length, rl.IdRoute)
	return err
}

func DeleteBus(c echo.Context) error {
	bus := c.Param("id")
	_, err := db.Conn.Query(`DELETE FROM specifications WHERE id_bus=$1`,bus)
	return err
}

func GetInfo(c echo.Context) error {
	rows, _ := db.Conn.Query(`SELECT count(id_bus) FROM specifications`)
	var num int
	for rows.Next() {
		if err := rows.Scan(&num); err != nil {
			return err
		}
	}
	rows, _ = db.Conn.Query(`SELECT DISTINCT type FROM specifications`)
	types := make([]string, 0)
	for rows.Next() {
		var typ string
		if err := rows.Scan(&typ); err != nil {
			return err
		}
		types = append(types, typ)
	}
	rows, _ = db.Conn.Query(`SELECT id_route, mov_start, bus_interval FROM routes`)
	routes := make([]schemas.RouteInfo, 0)
	for rows.Next() {
		var route schemas.RouteInfo
		if err := rows.Scan(&route.IdRoute, &route.MovStart, &route.BusInterval); err != nil {
			return err
		}
		routes = append(routes, route)
	}
	rows, _ = db.Conn.Query(`SELECT fullname, class FROM drivers`)
	drivers := make([]schemas.DriverInfo, 0)
	for rows.Next() {
		var driver schemas.DriverInfo
		if err := rows.Scan(&driver.Fullname, &driver.Class); err != nil {
			return err
		}
		drivers = append(drivers, driver)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&schemas.Info{ CountBus:num, BusTypes:types, Routes:routes, Drivers:drivers})

}