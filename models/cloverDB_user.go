package models

type Student struct {
	Sno            string `json:"sno"`
	Completed      bool   `json:"completed"`
	HealthCodePath string `json:"healthCodePath"`
	PathCodePath   string `json:"pathCodePath"`
	CloseCodePath  string `json:"closeCodePath"`
}

func (Student) TableStudent() string {
	return "student"
}
