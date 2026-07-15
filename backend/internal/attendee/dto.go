package attendee

// Using avi because this data is the result of joined table.
type AttendeeViewItem struct {
	Name                  string `json:"name"`
	PublicId              string `json:"public_id"`
	SchoolName            string `json:"school"`
	Grade                 string `json:"grade"`
	Subgrade              string `json:"subgrade"`
	ContactNumber         string `json:"telp"`
	GuardianContactNumber string `json:"guardian_telp"`
}

// http DTO
type AddAttendeeForm struct {
	Name                  string `json:"name"`
	Gender                string `json:"gender"`
	SchoolId              int    `json:"school_id"`
	GradeId               int    `json:"grade_id"`
	SubgradeId            int    `json:"subgrade_id"`
	ContactNumber         string `json:"telp"`
	GuardianContactNumber string `json:"guardian_telp"`
}
