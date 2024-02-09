package models

type Department struct {
	ID             int    `json:"id"`
	DepartmentName string `json:"departmentName"`
	StaffQuantity  int    `json:"staffQuantity"`
}
