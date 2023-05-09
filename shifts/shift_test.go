package shifts

import "testing"

func TestShift(t *testing.T) {
	shift := NewShift(0, 10, []int{1, 2, 3})

	shift.AddUser(4)

	if shift.StartTime() != 0 {
		t.Errorf("Expected 0, got %d", shift.StartTime())
	}
	if shift.EndTime() != 10 {
		t.Errorf("Expected 10, got %d", shift.EndTime())
	}
	if len(shift.Users()) != 4 {
		t.Errorf("Expected 4 users, got %d", len(shift.Users()))
	}

}
