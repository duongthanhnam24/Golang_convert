package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Nhập số: ")
	fmt.Scan(&input)

	// number, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("Lỗi: Không thể chuyển đổi chuỗi thành số")
	// 	return
	// }
	// numberinit := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	number := input

	digits := []int{}
	if number == 0 {
		digits = append([]int{0}, digits...)
	}
	for number > 0 {
		digit := number % 10
		digits = append([]int{digit}, digits...)
		number /= 10
	}
	// var last string
	text := []string{}
	fmt.Println("Chữ số là :")
	for _, digit := range digits {
		switch digit {
		case 0:
			text = append([]string{"Không"}, text...)
		case 1:
			fmt.Println("Một")
			text = append([]string{"Không"}, text...)

		case 2:
			text = append([]string{"Hai"}, text...)

		case 3:
			text = append([]string{"Ba"}, text...)

		case 4:
			text = append([]string{"Bốn"}, text...)

		case 5:
			text = append([]string{"Năm"}, text...)

		case 6:
			text = append([]string{"Sáu"}, text...)

		case 7:
			text = append([]string{"Bảy"}, text...)

		case 8:
			text = append([]string{"Tám"}, text...)

		case 9:
			text = append([]string{"Chín"}, text...)

		default:
			fmt.Println("Số ngày không hợp lệ")
		}
		// fmt.Println(digit)
	}
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]

		switch len(text) {
		case 2:
			fmt.Printf("%v Mươi %v", text[i], text[j])

		default:
			fmt.Printf("%v", text[i])
		}

		// return fmt.Printf("%v %v", text[i], text[j])
	}
}
