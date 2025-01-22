package models

type GPSData struct {
	ID             uint    `gorm:"primaryKey" json:"id_data"`
	DeviceID       uint    `gorm:"not null" json:"id_device"`
	Date           string  `gorm:"type:date;not null" json:"date"`
	Time           string  `gorm:"type:time;not null" json:"time"`
	Latitude       float64 `gorm:"type:numeric(10,7);not null" json:"latitude"`
	Longitude      float64 `gorm:"type:numeric(10,7);not null" json:"longitude"`
	Sattelite      int     `gorm:"type:int" json:"sattelite"`
	Altitude       float64 `gorm:"type:numeric(8,2)" json:"altitude"`
	Speed          float64 `gorm:"type:numeric(6,2)" json:"speed"`
	BatteryVoltage float64 `gorm:"type:numeric(5,2)" json:"battery_voltage"`
	BatteryCapacity float64 `gorm:"type:numeric(5,2)" json:"battery_capacity"`
}
