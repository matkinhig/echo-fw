package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/leekchan/accounting"
	"github.com/matkinhig/echo-fw/config"
	"github.com/matkinhig/echo-fw/route"
	service "github.com/matkinhig/echo-fw/services"

	"net/http"

	"github.com/labstack/echo"

	"database/sql"

	_ "github.com/mattn/go-oci8"
)

type User struct {
	Name    string `json:name`
	Address string `json:address`
}

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"course_name"`
	AcademyName string `json:"acedemy_name"`
}

func main() {
	fmt.Println("start golang")

	route.Public()

	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))

	fmt.Println("The Config value:", config.Config)
	// testServer()
	// connectORACLE()
	crudService()
}

func crudService() {
	e := echo.New()

	//Route
	e.POST("/student", service.CreateStudent)
	e.GET("/student/:id", service.GetStudent)
	e.PUT("/student/:id", service.UpdateStudent)
	e.DELETE("/student/:id", service.DeleteUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func testServer() {
	e := echo.New()
	e.GET("/", allocServer)
	e.GET("/info/:id", getInfomation)
	e.POST("info", addUser)
	e.POST("/student", addStudent)
	e.POST("/people", addPeople)
	e.Start(":8000")
}

func addPeople(c echo.Context) error {
	aPeople := Student{}
	defer c.Request().Body.Close()
	err := c.Bind(&aPeople) //cach 3 bind Data request input
	if err != nil {
		log.Printf("can not decode the request : %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is a people %#v : ", aPeople)
	return c.String(http.StatusOK, "we are receive the people %s : ")
}

func addStudent(c echo.Context) error {
	aStudent := Student{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&aStudent) //cach 2 NewDecoder
	if err != nil {
		log.Printf("can not decode the request : %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is a student %#v : ", aStudent)
	return c.String(http.StatusOK, "we are receive the student %s : ")
}

func allocServer(c echo.Context) error {
	return c.String(http.StatusOK, "my first website by golang")
}

func getInfomation(c echo.Context) error {
	name := c.QueryParam("name")
	address := c.QueryParam("address")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your name : %s \n your address : %s \n", name, address))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    name,
			"address": address,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "cant parse to json",
	})
}

func addUser(c echo.Context) error {
	teo := User{}

	b, err := ioutil.ReadAll(c.Request().Body) //cach 1 post data

	if err != nil {
		log.Println("cant read data request %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &teo)
	if err != nil {
		log.Println("cant unmarshal json %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	defer c.Request().Body.Close()
	fmt.Printf("info cua teo la : %+v", teo)

	return c.String(http.StatusOK, "add user")
}

func connectORACLE() {
	println("start")
	db, err := sql.Open("oci8", "USER_NHS/nhsVB123456@192.168.254.34:1521/sid?tcbsweb")
	if err != nil {
		log.Fatal(err)
	}
	println("Connection succcess!!")
	rows, err := db.Query("SELECT sysdate  FROM dual")
	if err != nil {
		log.Fatalln("err:", err.Error)
	}
	var (
		sysdate string
	)
	for rows.Next() {
		if err = rows.Scan(&sysdate); err != nil {
			log.Fatalln("error fetching", err)
		}
		log.Println(sysdate)
	}
}
