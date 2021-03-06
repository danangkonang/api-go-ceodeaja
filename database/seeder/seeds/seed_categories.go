package seeds

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Categories() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO categories (category,created_at,updated_at)VALUES
		('handphone','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('komputer & laptop','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('elektronik & listrik','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('otomotif','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('fashion pria','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('fashion wanita','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('kesehatan & kecantikan','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('ibu & anak','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('jasa & pekerjaan','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('olahraga & hobbi','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('properti & bangunan','2006-01-02 15:04:05','2006-01-02 15:04:05'),
		('rumah tangga','2006-01-02 15:04:05','2006-01-02 15:04:05')
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert categories")
}
