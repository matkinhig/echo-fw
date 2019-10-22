package services

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Student struct {
	ID             int    `json:"id"`
	MaHocSinh      string `json:"mahocsinh"`
	HoTen          string `json:"hoten"`
	SoTien         int    `json:"sotien"`
	DienThoai      string `json:"dienthoai"`
	ThoiGianHocPhi string `json:"thoigianhocphi"`
	GhiChu         string `json:"ghichu"`
	TrangThai      string `json:"trangthai"`
	NgayThanhToan  string `json:"ngaythanhtoan"`
	NgayTao        string `json:"ngaytao"`
	HanThanhToan   string `json:"hanthanhtoan"`
}

var (
	tmpStudent = map[int]*Student{}
	seq        = 1
)

//----------
// Handlers
//---------

func CreateStudent(c echo.Context) error {
	u := &Student{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	tmpStudent[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, tmpStudent[id])
}

func UpdateStudent(c echo.Context) error {
	st := new(Student)
	if err := c.Bind(st); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	tmpStudent[id].ID = st.ID
	return c.JSON(http.StatusOK, tmpStudent[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(tmpStudent, id)
	return c.NoContent(http.StatusNoContent)
}
