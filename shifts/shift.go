package shifts

// An Interface representing a shift
type IShift interface {
	StartTime() int
	EndTime() int
	Users() []int
}

// A struct representing a shift
type Shift struct {
	startTime int
	endTime   int
	userIds   []int
}

// A regular function returning a pointer to a shift
func NewShift(start int, end int, users []int) *Shift {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Shift{
		startTime: start,
		endTime:   end,
		userIds:   users,
	}
}

func (s *Shift) StartTime() int {
	return s.startTime
}

func (s *Shift) EndTime() int {
	return s.endTime
}

func (s *Shift) Users() []int {
	return s.userIds
}

func (s *Shift) AddUser(userId int) {
	s.userIds = append(s.userIds, userId)
}
