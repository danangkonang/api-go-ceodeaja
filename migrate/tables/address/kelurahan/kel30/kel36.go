package kel30

import "github.com/danangkonang/ceodeaja-go/config"

func Kel36() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	`)
}
