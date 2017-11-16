package lesson

import (
	"fmt"
	"strconv"
)

/*
lesson one: base
1.variable,slice,map
2.for if switch
3.string to int, int to string
4.method,folder
*/

func First() {
	fmt.Println("hello world")
	var name string = "chen.yafei"
	fmt.Println(name)
	age := 18
	fmt.Printf("\n%T,%v", age, age)

	var age2 int64 = 18
	fmt.Printf("\n%T,%v", age2, age2)

	prices := []int{1, 2, 3, 5}
	price := prices[0:3]
	fmt.Println(price)

	for k, v := range prices {
		fmt.Println(k, v)
	}

	mm := make(map[string]string, 0)
	mm["chen"] = "chen.yafei"
	mm["xiao"] = "xiao.xinmiao"
	// for k, v := range mm {
	// 	fmt.Println(k, v)
	// }
	for k, v := range mm {
		if k == "chen" {
			fmt.Println(k, v)
		}
	}

	color := "black"
	switch color {
	case "red", "yellow":
		fmt.Println("I like.")
	case "black":
		fmt.Println("I don't like")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	//lesson two
	price_red := SaleApple("red")
	price_black := SaleApple("black")
	fmt.Printf("\nred apple's price is %v", price_red)
	fmt.Printf("\nblack apple's price is %v", price_black)

	price2, discount2 := SaleApple2("red")
	fmt.Printf("\nprice:%v,discount:%v", price2, discount2)

	var dollar int64 = 1000
	fmt.Printf("%T,%T", dollar, int(dollar))

	var dollar2 int = 10000
	fmt.Printf("%v", "$"+strconv.FormatInt(int64(dollar2), 10))

	dollar3 := "12"
	dollar4, err := strconv.ParseInt(dollar3, 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dollar4)

}

func SaleApple(color string) int {
	price := 0
	switch color {
	case "red", "yellow":
		price = 5
	case "black":
		price = 10
	default:
		price = 0
	}
	return price
}

func SaleApple2(color string) (price int, discount int) {
	switch color {
	case "red", "yellow":
		price = 5
		discount = price / 2
	case "black":
		price = 10
		discount = price / 2
	default:
		price = 0
	}
	return
}
