package main

import (
	"fmt"
	"planlos/shifts"
)

func main() {
	shift := shifts.NewShift(0, 10, []int{1, 2, 3})
	shift.AddUser(4)
	fmt.Println(shift.StartTime())
	fmt.Println(shift.EndTime())
	fmt.Println(shift.Users())
}
