package models

type GPSDevice struct {
	ID          uint      `gorm:"primaryKey" json:"id_device"`
	DeviceName  string    `gorm:"type:varchar(50);not null" json:"device_name"`
	Description string    `gorm:"type:text" json:"description"`
}
