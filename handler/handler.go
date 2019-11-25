package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/matkinhig/echo-fw/db"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "done check health")
}

func GetAllStudent(c echo.Context) error {
	// aStudent := types.Student{
	// 	ID: 1, Firstname: "Luc", Lastname: "Nguyen", Age: 80, Classname: "nc golang", Email: "matkinhig@gmail.cpm",
	// }
	// bStudent := types.Student{
	// 	ID: 2, Firstname: "Teo", Lastname: "Tran", Age: 100, Classname: "nc golang", Email: "matkinhig@gmail.cpm",
	// }
	// cStudent := types.Student{
	// 	ID: 3, Firstname: "Ty", Lastname: "Le", Age: 50, Classname: "nc golang", Email: "matkinhig@gmail.cpm",
	// }
	// dStudent := types.Student{
	// 	ID: 4, Firstname: "Dan", Lastname: "Cao", Age: 20, Classname: "nc golang", Email: "matkinhig@gmail.cpm",
	// }

	// result := []types.Student{aStudent, bStudent, cStudent, dStudent}
	result, err := db.GetAllStudent()

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, result)
}
