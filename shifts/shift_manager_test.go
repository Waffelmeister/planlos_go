package shifts

// An Test for Shift manager
import "testing"

func TestShiftManager(t *testing.T) {
	shiftManager := NewShiftManager()

	shiftManager.AddShift(*NewShift(0, 10, []int{1, 2, 3}))
	shiftManager.AddShift(*NewShift(10, 20, []int{1, 2, 3}))
	shiftManager.AddShift(*NewShift(20, 30, []int{1, 2, 3}))

	if len(shiftManager.Shifts()) != 3 {
		t.Errorf("Expected 3 shifts, got %d", len(shiftManager.Shifts()))
	}

	if len(shiftManager.ShiftsForUser(1)) != 3 {
		t.Errorf("Expected 3 shifts for user, got %d", len(shiftManager.ShiftsForUser(1)))
	}

	if len(shiftManager.ShiftsForUserAtTime(1, 5)) != 1 {
		t.Errorf("Expected 1 shifts for user at time, got %d", len(shiftManager.ShiftsForUserAtTime(1, 5)))
	}

	if len(shiftManager.ShiftsForUserAtTimeWithUser(1, 5, 2)) != 2 {
		t.Errorf("Expected 2 shifts for user at time with user, got %d", len(shiftManager.ShiftsForUserAtTimeWithUser(1, 5, 2)))
	}

	if len(shiftManager.ShiftsForUserAtTimeWithUserAndRole(1, 5, 2, "role")) != 2 {
		t.Errorf("Expected 2 shifts for user at time with user and role, got %d", len(shiftManager.ShiftsForUserAtTimeWithUserAndRole(1, 5, 2, "role")))
	}

}
