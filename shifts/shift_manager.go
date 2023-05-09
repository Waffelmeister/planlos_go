package shifts

type ShiftManager struct {
	shifts []Shift
}

func NewShiftManager() *ShiftManager {
	return &ShiftManager{}
}

func (sm *ShiftManager) AddShift(shift Shift) {
	sm.shifts = append(sm.shifts, shift)
}

func (sm *ShiftManager) Shifts() []Shift {
	return sm.shifts
}

func (sm *ShiftManager) ShiftsForUser(userId int) []Shift {
	var shifts []Shift
	for _, shift := range sm.shifts {
		for _, id := range shift.Users() {
			if id == userId {
				shifts = append(shifts, shift)
			}
		}
	}
	return shifts
}

func (sm *ShiftManager) ShiftsForUserAtTime(userId int, time int) []Shift {
	var shifts []Shift
	for _, shift := range sm.shifts {
		if shift.StartTime() <= time && shift.EndTime() >= time {
			for _, id := range shift.Users() {
				if id == userId {
					shifts = append(shifts, shift)
				}
			}
		}
	}
	return shifts
}

func (sm *ShiftManager) ShiftsForUserAtTimeWithUser(userId int, time int, otherUserId int) []Shift {
	var shifts []Shift
	for _, shift := range sm.shifts {
		if shift.StartTime() <= time && shift.EndTime() >= time {
			for _, id := range shift.Users() {
				if id == userId || id == otherUserId {
					shifts = append(shifts, shift)
				}
			}
		}
	}
	return shifts
}

func (sm *ShiftManager) ShiftsForUserAtTimeWithUserAndRole(userId int, time int, otherUserId int, role string) []Shift {
	var shifts []Shift
	for _, shift := range sm.shifts {
		if shift.StartTime() <= time && shift.EndTime() >= time {
			for _, id := range shift.Users() {
				if id == userId || id == otherUserId {
					shifts = append(shifts, shift)
				}
			}
		}
	}
	return shifts
}
