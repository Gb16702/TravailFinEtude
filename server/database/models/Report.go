package models;

type Report struct {
	ID 				uint 	`gorm:"primaryKey" json:"id"`
	UserID 			uint 	`gorm:"not null" json:"user_id"`
	TargetID 		uint 	`gorm:"not null" json:"target_id"`
	Reason 	    	string 	`gorm:"not null"`
	CreatedAt 		int64
}
