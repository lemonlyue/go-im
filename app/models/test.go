package models

type Test struct {
	BaseModel
	Name string `gorm:"column:name;not null;type:varchar(50)" json:"name,imitempty"`
	CommonTimestampsField
}

func (Test) TableName() string {
	return "test"
}

func NewTestModel() *Test {
	return &Test{}
}
