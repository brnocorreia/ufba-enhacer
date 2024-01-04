package entity

type Prerequisite struct {
	CourseID     int32  `gorm:"column:course_id;primaryKey" json:"course_id"`
	DisciplineID string `gorm:"column:discipline_id;primaryKey" json:"discipline_id"`
	DependsOn    string `gorm:"column:depends_on;primaryKey" json:"depends_on"`
}
