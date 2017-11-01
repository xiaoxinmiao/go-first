package main

import (
	"first/data"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Name3)
	// fmt.Println("hello world")
	// time.Sleep(5e9)
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "hello echo api")
	// })

	// e.GET("/pay", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "wx pay")
	// })
	// e.Use(middleware.CORS())
	// e.Start(":1111")
	var name string = "liubo"
	name2 := "yumeng"
	fmt.Println(name)
	fmt.Println(name2)

	age := 18
	fmt.Println(age)

	var age2 int32 = 19
	var age3 int64 = int64(age2)

	fmt.Println(age2)
	fmt.Println(age3)

	age4 := strconv.FormatInt(age3, 10)
	fmt.Println(age4)

	height := "170"

	height1, _ := strconv.ParseInt(height, 10, 32)
	height2 := int32(height1)
	fmt.Println(height2)
	//fmt.Println(err1)

	prices := []int{1, 2, 3, 56, 7}
	slice := prices[0:3]
	fmt.Println(slice)

	for k, v := range prices {
		fmt.Println(k, v)
	}

	mm := make(map[string]string, 0)
	mm["yangying"] = "white"
	mm["weiwei"] = "black"
	for k, v := range mm {
		if v != "weiwei" {
			fmt.Println(k, v)
		}
	}

	if height8, err1 := strconv.ParseInt(height, 10, 32); err1 != nil {
		fmt.Println(height8)
	}

	i := 12

	switch i {
	case 1, 2, 3:
		fmt.Println("come into 1,2,3")
	case 12:
		fmt.Println("come into 12")
		fallthrough
	default:
		fmt.Println("come into default")
	}

	fmt.Println(run(70))
	fmt.Println(run2(70))

	fmt.Println(run3("a"))

	fmt.Println(data.Walk())
}

func run(speed int) int {
	return speed * 10
}

func run2(speed int) (bonus int) {
	bonus = speed * 9
	return
}

func run3(speed string) (bonus int64, err error) {
	var speedInt int64
	speedInt, err = strconv.ParseInt(speed, 10, 32)
	if err != nil {
		return
	}
	bonus = speedInt * 9
	return
}
