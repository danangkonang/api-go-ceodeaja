package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/danangkonang/ceodeaja-go/config"
	"github.com/gorilla/mux"
)

type Province struct {
	IdProvinsi int64
	Provinc    string
}

// fix
func ShowProvinces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	result, err := db.Query("SELECT id_provinsi, provinsi FROM provinsi")
	checkErr(err)
	var prov []interface{}
	for result.Next() {
		var p Province
		err := result.Scan(&p.IdProvinsi, &p.Provinc)
		checkErr(err)
		res := map[string]interface{}{
			"id":   p.IdProvinsi,
			"name": p.Provinc,
		}
		prov = append(prov, res)
	}
	defer db.Close()
	resProv := map[string]interface{}{
		"success": true,
		"data":    prov,
	}
	q, _ := json.Marshal(resProv)
	w.Write(q)
}

// fix
func ShowProvince(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	sqlStatement := `SELECT id_provinsi, provinsi FROM provinsi WHERE id_provinsi = $1`
	var post Province
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.IdProvinsi, &post.Provinc)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"data":    "null",
			"success": true,
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"id":      post.IdProvinsi,
				"user_id": post.Provinc,
			},
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	default:
		panic(err)
	}
}

type City struct {
	Id          int64
	Name        string
	Province_id int64
}

// fix
func ShowCitys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	result, err := db.Query("SELECT id_kabupaten,kabupaten,provinsi_id FROM kabupaten WHERE provinsi_id = $1", params["id"])
	checkErr(err)
	var citi []interface{}
	for result.Next() {
		var p City
		err := result.Scan(&p.Id, &p.Name, &p.Province_id)
		checkErr(err)
		res := map[string]interface{}{
			"id":          p.Id,
			"name":        p.Name,
			"province_id": p.Province_id,
		}
		citi = append(citi, res)
	}
	defer db.Close()
	if len(citi) > 0 {
		resKab := map[string]interface{}{
			"success": true,
			"data":    citi,
		}
		q, _ := json.Marshal(resKab)
		w.Write(q)
	} else {
		f2 := make([]string, 0)
		res := map[string]interface{}{
			"success": true,
			"data":    f2,
		}
		q, _ := json.Marshal(res)
		w.Write(q)
	}
}

// fix
func ShowCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id_kabupaten,kabupaten,provinsi_id FROM kabupaten WHERE id_kabupaten = $1`
	var post City
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Province_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"success": true,
			"data":    "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"id":          post.Id,
				"name":        post.Name,
				"province_id": post.Province_id,
			},
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	default:
		panic(err)
	}
}

type Kecamatan struct {
	Id          int32
	Name        string
	Kota_id     int32
	Provinsi_id int32
}

// fix
func ShowKecamatans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	result, err := db.Query("SELECT id_kecamatan,kecamatan,kabupaten_id,provinsi_id FROM kecamatan WHERE kabupaten_id = $1", params["id"])
	checkErr(err)
	var kec []interface{}
	for result.Next() {
		var p Kecamatan
		err := result.Scan(&p.Id, &p.Name, &p.Kota_id, &p.Provinsi_id)
		checkErr(err)
		res := map[string]interface{}{
			"id":          p.Id,
			"name":        p.Name,
			"kota_id":     p.Kota_id,
			"province_id": p.Provinsi_id,
		}
		kec = append(kec, res)
	}
	defer db.Close()
	resKec := map[string]interface{}{
		"success": true,
		"data":    kec,
	}
	q, _ := json.Marshal(resKec)
	w.Write(q)
}

func ShowKecamatan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id_kecamatan,kecamatan,kabupaten_id,provinsi_id FROM kecamatan WHERE id_kecamatan = $1`
	var post Kecamatan
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Kota_id, &post.Provinsi_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"success": true,
			"data":    "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"id":          post.Id,
				"name":        post.Name,
				"kota_id":     post.Kota_id,
				"province_id": post.Provinsi_id,
			},
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	default:
		panic(err)
	}
}

type Kelurahan struct {
	Id           int32
	Name         string
	Kecamatan_id int32
	Kota_id      int32
	Provinsi_id  int32
}

// fix
func ShowKelurahans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	sqlStatement := "SELECT id_kelurahan,kelurahan,kecamatan_id,kabupaten_id,provinsi_id FROM kelurahan WHERE kecamatan_id = $1"
	result, err := db.Query(sqlStatement, params["id"])
	checkErr(err)
	var kel []interface{}
	for result.Next() {
		var p Kelurahan
		err := result.Scan(&p.Id, &p.Name, &p.Kecamatan_id, &p.Kota_id, &p.Provinsi_id)
		checkErr(err)
		res := map[string]interface{}{
			"id":           p.Id,
			"name":         p.Name,
			"kecamatan_id": p.Kecamatan_id,
			"kota_id":      p.Kota_id,
			"province_id":  p.Provinsi_id,
		}
		kel = append(kel, res)
	}
	defer db.Close()
	if len(kel) > 0 {
		res := map[string]interface{}{
			"success": true,
			"data":    kel,
		}
		q, _ := json.Marshal(res)
		w.Write(q)
	} else {
		f2 := make([]string, 0)
		res := map[string]interface{}{
			"success": true,
			"data":    f2,
		}
		q, _ := json.Marshal(res)
		w.Write(q)
	}
}

// fix
func ShowKelurahan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id_kelurahan,kelurahan,kecamatan_id,kabupaten_id,provinsi_id FROM kelurahan WHERE id_kelurahan = $1`
	var post Kelurahan
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Kecamatan_id, &post.Kota_id, &post.Provinsi_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"success": true,
			"data":    "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"id":           post.Id,
			"province_id":  post.Provinsi_id,
			"kota_id":      post.Kota_id,
			"kecamatan_id": post.Kecamatan_id,
			"name":         post.Name,
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	default:
		panic(err)
	}
}
