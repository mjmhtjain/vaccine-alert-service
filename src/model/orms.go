package model

type AppointmentSessionORM struct {
	Id                     string `gorm:"primaryKey;autoIncrement:false;column:id"`
	CenterIDFK             int    `gorm:"column:center_idfk"`
	Date                   string `gorm:"column:date"`
	AvailableCapacity      int    `gorm:"column:available_capacity"`
	MinAgeLimit            int    `gorm:"column:min_age_limit"`
	VaccineIDKF            string `gorm:"column:vaccine_idfk"`
	AvailableCapacityDose1 int    `gorm:"column:available_capacity_dose1"`
	AvailableCapacityDose2 int    `gorm:"column:available_capacity_dose2"`
}

type VaccineORM struct {
	Id   string `gorm:"primaryKey;autoIncrement:false;column:id"`
	Name string `gorm:"column:name"`
}

type CenterORM struct {
	Id           int    `gorm:"primaryKey;autoIncrement:false;column:id"`
	Name         string `gorm:"column:name"`
	Address      string `gorm:"column:address"`
	StateName    string `gorm:"column:state_name"`
	DistrictName string `gorm:"column:district_name"`
	Pincode      int    `gorm:"column:pincode"`
}

// TableNames
func (CenterORM) TableName() string {
	return "center_info"
}

func (VaccineORM) TableName() string {
	return "vaccine"
}

func (AppointmentSessionORM) TableName() string {
	return "appointment_session"
}
