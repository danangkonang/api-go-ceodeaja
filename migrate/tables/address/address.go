package address

import (
	"github.com/danangkonang/rest-api/config"
	"github.com/danangkonang/rest-api/migrate/tables/address/kabupaten"
	"github.com/danangkonang/rest-api/migrate/tables/address/kecamatan"
	"github.com/danangkonang/rest-api/migrate/tables/address/kelurahan"
	"github.com/danangkonang/rest-api/migrate/tables/address/provinsi"
)

func Address() {
	db := config.Connect()
	db.Exec(`DROP TABLE provinsi`)
	db.Exec(`DROP TABLE kota`)
	db.Exec(`DROP TABLE kecamatan`)
	db.Exec(`DROP TABLE kelurahan`)
	db.Exec(`CREATE TABLE provinsi(
		id serial PRIMARY KEY,
		provinsi VARCHAR (225) NOT NULL
	)`)
	db.Exec(`CREATE TABLE kota(
		id serial PRIMARY KEY,
		kota VARCHAR (225) NOT NULL,
		provinsi_id INTEGER NOT NULL
	)`)
	db.Exec(`CREATE TABLE kecamatan(
		id serial PRIMARY KEY,
		kecamatan VARCHAR (225) NOT NULL,
		kota_id INTEGER NOT NULL,
		provinsi_id INTEGER NOT NULL
	)`)
	db.Exec(`CREATE TABLE kelurahan(
		id serial PRIMARY KEY,
		kelurahan VARCHAR (225) NOT NULL,
		kecamatan_id INTEGER NOT NULL,
		kota_id INTEGER NOT NULL,
		provinsi_id INTEGER NOT NULL
	)`)
	provinsi.CreateProvinsi()
	kabupaten.CreateKota()
	kecamatan.CreateKecamatan()
	kelurahan.Kelurahan()
}
