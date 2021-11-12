package request

type CourseAdd struct {
	Name       string `json:"name"`
	CourseHour int    `json:"courseHour"`
	CourseImg  string `json:"courseImg"`
	CourseDesc string `json:"courseDesc"`
	SchoolId   int    `json:"schoolId"`
}

type CourseUpdate struct {
	ID         uint   `json:"ID"`
	Name       string `json:"name"`
	CourseHour int    `json:"courseHour"`
	CourseImg  string `json:"courseImg"`
	CourseDesc string `json:"courseDesc"`
	SchoolId   int    `json:"schoolId"`
}
