package entity

type CoursesDiscipline struct {
	CourseID     int32  `gorm:"column:course_id;primaryKey" json:"course_id"`
	DisciplineID string `gorm:"column:discipline_id;primaryKey" json:"discipline_id"`
}
