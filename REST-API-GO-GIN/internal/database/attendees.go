package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"` // 所屬用戶的 ID
	EventId int `json:"eventId"`
}
