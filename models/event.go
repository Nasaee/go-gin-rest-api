package models

import (
	"time"

	"github.com/Nasaee/go-gin-rest-api/db"
)

// Gin จะ ตรวจสอบด้วย ว่าค่าในแต่ละ field ที่มี binding:"required" ถูกส่งมาหรือเปล่า ถ้าไม่ส่งมาจะ return error
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"        binding:"required,min=1,max=200"`
	Description string    `json:"description" binding:"required,min=1,max=2000"`
	Location    string    `json:"location"    binding:"required"`
	DateTime    time.Time `json:"dateTime"    binding:"required"`
	UserID      int       `json:"userId"`
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query) // stmt คือ statement
	if err != nil {
		return err
	}

	// defer = “สั่งให้โค้ดนี้ไปรันทีหลังสุด ก่อนที่ฟังก์ชันจะจบการทำงาน” พูดง่าย ๆ คือ “ตอนนี้ยังไม่ต้องทำ เดี๋ยวทำตอนออกจากฟังก์ชัน”
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	/*
		.Next() เมธอดนี้ใช้ เลื่อน cursor ไปยังแถวถัดไป ในผลลัพธ์
		- ถ้ามีแถวถัดไป → Next() จะคืนค่า true
		- ถ้าไม่มีแล้ว → คืนค่า false แล้วออกจาก loop
	*/
	for rows.Next() {
		var event Event
		// .Scan() เพื่อ อ่านค่าของแถวนี้ แล้ว ใส่ลงในตัวแปร ที่เตรียมไว้
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

/*
Note:
- .Exec() ใช้สําหรับการสร้าง แก้ไข ลบ ข้อมูล
- .Query() ใช้สําหรับการดึงข้อมูล
*/
