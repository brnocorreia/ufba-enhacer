package entity

type Discipline struct {
	Code         string `gorm:"column:code;primaryKey" json:"code"`
	Name         string `gorm:"column:name;not null" json:"name"`
	Workload     int32  `gorm:"column:workload;not null" json:"workload"`
	Department   string `gorm:"column:department;not null" json:"department"`
	Program      string `gorm:"column:program;not null" json:"program"`
	Objective    string `gorm:"column:objective;not null" json:"objective"`
	Content      string `gorm:"column:content;not null" json:"content"`
	Bibliography string `gorm:"column:bibliography;not null" json:"bibliography"`
}
