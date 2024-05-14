package models

type Employee struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	EmployeeCode string `json:"employeeCode,omitempty" bson:"employeeCode,omitempty"`
	StaffType    string `json:"staffType,omitempty" bson:"staffType,omitempty"`
}
