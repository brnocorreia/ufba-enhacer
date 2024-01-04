package entity

type Course struct {
	Code        int32   `gorm:"column:code;primaryKey" json:"code"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	Shift       string  `gorm:"column:shift;not null" json:"shift"`
	MinDuration float64 `gorm:"column:min_duration;not null" json:"min_duration"`
	MaxDuration float64 `gorm:"column:max_duration;not null" json:"max_duration"`
	LegalBase   string  `gorm:"column:legal_base;not null" json:"legal_base"`
	Description string  `gorm:"column:description;not null" json:"description"`
	ObWorkload  int32   `gorm:"column:ob_workload;not null" json:"ob_workload"`
	OpWorkload  int32   `gorm:"column:op_workload;not null" json:"op_workload"`
	AcWorkload  int32   `gorm:"column:ac_workload;not null" json:"ac_workload"`
}
