package test

type CourseItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Selected int    `json:"selected"`
}

type Course struct {
	Code int        `json:"code"`
	Data CourseItem `json:"data"`
	Msg  string     `json:"msg"`
}

type AllCourse struct {
	Code int          `json:"code"`
	Data []CourseItem `json:"data"`
	Msg  string       `json:"msg"`
}

type StudentItem struct {
	Stuid string `json:"stuid"`
	Name  string `json:"name"`
}

type StudentCourseItem struct {
	Stuid  string `json:"stuid"`
	Name   string `json:"name"`
	Course []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"course"`
}

type StudentCourse struct {
	Code int               `json:"code"`
	Data StudentCourseItem `json:"data"`
	Msg  string            `json:"msg"`
}

type CourseResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
