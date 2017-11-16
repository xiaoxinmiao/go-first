package hw

import (
	"net/http"

	"github.com/labstack/echo"
)

/*
1.写一套水果的api，要求包括增删改查
其中，数据库，自己用全局变量的集合模拟
*/
func Hw2() {
	Data()
	e := echo.New()
	e.GET("/fruits", GetAllFruit)
	e.GET("/fruits/:id", GetFruit)
	e.POST("/fruits", PostFruit)
	e.PUT("/fruits/:id", PutFruit)
	e.DELETE("/fruits/:id", DeleteFruit)
	e.Start(":5000")

}

var fruits []fruit

type fruit struct {
	Id    int64
	Name  string
	Price float32
	Color string
}

func Data() {
	fruits = []fruit{
		fruit{Id: 1, Name: "apple", Price: 10, Color: "red"},
		fruit{Id: 2, Name: "pear", Price: 15, Color: "green"},
	}
}

func GetAllFruit(c echo.Context) error {
	return c.JSON(http.StatusOK, fruits)
}
func GetFruit(c echo.Context) error {
	return nil
}

func PostFruit(c echo.Context) error {
	return nil
}

func PutFruit(c echo.Context) error {
	return nil
}

func DeleteFruit(c echo.Context) error {
	return nil
}
