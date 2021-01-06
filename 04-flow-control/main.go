package main

import (
	"fmt"
)

func main() {

	fmt.Println("for loop counter: initialization; condition; increment")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	items := make(map[int32]string)
	items[1] = "wheel"
	items[2] = "glass"
	items[3] = "door"
	items[4] = "windshield"

	fmt.Println("Getting values from a map")
	for _, value := range items {
		fmt.Println(value)
	}

	fmt.Println("Getting keys and values from a map")
	for id, value := range items {
		fmt.Printf("%d ==== %s \n", id, value)
	}

	fmt.Println("Switch condition")
	var dayOfWeek = 2
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		{
			fmt.Println("Saturday")
			fmt.Println("Weekend.")
		}
	case 7:
		{
			fmt.Println("Sunday")
			fmt.Println("Weekend.")
		}
	default:
		fmt.Println("Invalid day")
	}

	fmt.Println("Grouped switch condition")
	switch dayOfWeek := 5; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	fmt.Println("No expression switch condition")
	var BMI = 21.0
	switch {
	case BMI < 18.5:
		fmt.Println("You're underweight")
	case BMI >= 18.5 && BMI < 25.0:
		fmt.Println("Your weight is normal")
	case BMI >= 25.0 && BMI < 30.0:
		fmt.Println("You're overweight")
	default:
		fmt.Println("You're obese")
	}

	fmt.Println("End")
}
