package hw

import (
	"fmt"
)

// 1、打印输出 5 个"Hello world!"
// 2、打印输出 1-100 中的奇数
// 3、打印菱形
// 4.打印空心菱形

func Hw1() {
	FiveHello()
	PrintOdd()
	Diamond(22)
	DiamondEmpty(22)
}

func FiveHello() {
	for index := 0; index < 5; index++ {
		fmt.Println("hello world")
	}
}

func PrintOdd() {
	for index := 0; index < 100; index++ {
		if index%2 == 1 {
			fmt.Println(index)
		}
	}
}

func Diamond(count int) {
	count = BecomeOdd(count)
	midCount := count / 2
	//up
	for index := 0; index < midCount; index++ {
		//1.print blank
		for j := 0; j < midCount-index; j++ {
			fmt.Printf("%v", " ")
		}

		for j := 0; j < 2*index+1; j++ {
			fmt.Printf("%v", "*")
		}
		fmt.Println()

	}
	//down
	for index := 0; index < midCount+1; index++ {
		//1.print blank
		for j := 0; j < index; j++ {
			fmt.Printf("%v", " ")
		}

		for j := 0; j < count-2*index; j++ {
			fmt.Printf("%v", "*")
		}
		fmt.Println()
	}

}

func DiamondEmpty(count int) {
	count = BecomeOdd(count)
	midCount := count / 2
	//up
	for index := 0; index < midCount; index++ {
		//1.print blank
		for j := 0; j < midCount-index; j++ {
			fmt.Printf("%v", " ")
		}

		fmt.Printf("%v", "*")
		for j := 0; j < 2*index-1; j++ {
			fmt.Printf("%v", " ")
		}
		if index != 0 {
			fmt.Printf("%v", "*")
		}

		fmt.Println()

	}
	//down
	for index := 0; index < midCount+1; index++ {
		//1.print blank
		for j := 0; j < index; j++ {
			fmt.Printf("%v", " ")
		}

		fmt.Printf("%v", "*")

		for j := 0; j < count-2-2*index; j++ {
			fmt.Printf("%v", " ")
		}
		if index != midCount {
			fmt.Printf("%v", "*")

		}
		fmt.Println()
	}
}

func BecomeOdd(count int) (result int) {
	result = count
	if count%2 == 0 {
		result = count + 1
	}
	return
}
