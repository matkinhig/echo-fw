package route

import (
	"github.com/labstack/echo"
	"github.com/matkinhig/echo-fw/handler"
)

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/public")
	g.GET("/health", handler.CheckHealth)
	g.GET("/student", handler.GetAllStudent)
}

func Staff() {

}

func Private() {

}
