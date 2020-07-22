package kel30

import "github.com/danangkonang/rest-api/config"

func Kel32() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	`)
}
