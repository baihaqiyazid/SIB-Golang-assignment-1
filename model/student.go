package model

type Student struct{
	ID string `json:id`
	Name string `json:"name"`
	Code string `json:"code"`
	Address string `json:"address"`
	Occupation string `json:"occupation"`
	JoiningReason string `json:"joining_reason"`
}

type StudentList struct{
	StudentList []Student `json:"students"`
}