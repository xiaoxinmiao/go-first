package lesson

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type people struct {
	Id     string
	Hair   string
	Ear    string
	Weight int
}

func (people) run(name string) (speed float32) {
	switch name {
	case "chen.yafei":
		speed = 20
	case "bo.erte":
		speed = 9.8
	default:
		speed = 100
	}
	return
}

//Second good
func Second() {
	p := people{
		Hair: "black hair",
	}
	fmt.Printf("%+v", p)

	fmt.Println()
	speed := p.run("chen.yafei")
	fmt.Printf("speed is %v", speed)

	//1.
	p1 := people{}
	p1.Hair = "111"
	//2.
	p2 := new(people)
	p2.Hair = "222"
	fmt.Println(p1, p2)

	api()

}

func api() {

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/persons", GetPerson)
	e.GET("/persons/:id", GetPerson)
	e.Start(":5000")

}

func GetPersons(c echo.Context) error {
	p := people{
		Hair: "black",
	}
	return c.JSON(http.StatusOK, p)
}

func GetPerson(c echo.Context) error {
	id := c.Param("id")
	var p people
	if id == "1" {
		p = people{
			Hair: "yellow",
		}
	}
	return c.JSON(http.StatusOK, p)
}
