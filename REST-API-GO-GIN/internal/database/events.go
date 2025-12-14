package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId" binding:"required"` // 所屬用戶的 ID
	Name        string `json:"name" binding:"required" min:"1" max:"100"`
	Discription string `json:"discription" binding:"required" min:"1" max:"500"`
	Date        string `json:"date" binding:"required" example:"2023-12-31" format:"date"`
	location    string `json:"location" binding:"required" min:"1" max:"200"`
}
