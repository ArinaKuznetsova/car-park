package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/DonRIn/carpark/controllers"
)

func Run() {
	e := echo.New()
	e.Use(middleware.CORS())


 	e.GET("/drivers/route/:idRoute", controllers.GetDrivers)
	e.GET("/buses/:idRoute", controllers.GetBuses)
	e.GET("/routes/:idRoute/length", controllers.GetLengthRoutes)
	e.GET("/route/:idRoute/:time", controllers.GetTimeRoutes)
	e.GET("/drivers/bus/:idBus", controllers.GetDriverBus)
	e.GET("/info", controllers.GetInfo)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	}))
	g.POST("/driver", controllers.AddDriver)
	g.PUT("/bus/length", controllers.PutLength)
	g.DELETE("/bus/del/:id", controllers.DeleteBus)


	e.Start(":1323")
}

