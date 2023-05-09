package main

import (
	"planlos/shifts"
)

func main() {
	shiftManager := shifts.NewShiftManager()
	shiftManager.AddShift(*shifts.NewShift(0, 10, []int{1, 2, 3}))

}
