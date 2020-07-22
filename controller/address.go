package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/danangkonang/rest-api/config"
	"github.com/gorilla/mux"
)

type Province struct {
	Id      int64
	Provinc string
}

func ShowProvinces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	result, err := db.Query("SELECT id, provinsi FROM provinsi")
	checkErr(err)
	var prov []interface{}
	for result.Next() {
		var p Province
		err := result.Scan(&p.Id, &p.Provinc)
		checkErr(err)
		res := map[string]interface{}{
			"id":   p.Id,
			"name": p.Provinc,
		}
		prov = append(prov, res)
	}
	defer db.Close()
	q, _ := json.Marshal(prov)
	w.Write(q)
}

func ShowProvince(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	sqlStatement := `SELECT id, provinsi FROM provinsi WHERE id = $1`
	var post Province
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Provinc)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"data": "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"id":      post.Id,
			"user_id": post.Provinc,
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

func ShowCitys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	result, err := db.Query("SELECT id,kota,provinsi_id FROM kota WHERE provinsi_id = $1", params["id"])
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
	q, _ := json.Marshal(citi)
	w.Write(q)
}

func ShowCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id,kota,provinsi_id FROM kota WHERE id = $1`
	var post City
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Province_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"data": "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"id":          post.Id,
			"name":        post.Name,
			"province_id": post.Province_id,
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

func ShowKecamatans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	result, err := db.Query("SELECT id,kecamatan,kota_id,provinsi_id FROM kecamatan WHERE kota_id = $1", params["id"])
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
	q, _ := json.Marshal(kec)
	w.Write(q)
}

func ShowKecamatan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id,kecamatan,kota_id,provinsi_id FROM kecamatan WHERE id = $1`
	var post Kecamatan
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Kota_id, &post.Provinsi_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"data": "null",
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		resJes := map[string]interface{}{
			"id":          post.Id,
			"name":        post.Name,
			"kota_id":     post.Kota_id,
			"province_id": post.Provinsi_id,
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

func ShowKelurahans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	sqlStatement := "SELECT id,kelurahan,kecamatan_id,kota_id,provinsi_id FROM kelurahan WHERE kecamatan_id = $1"
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
	q, _ := json.Marshal(kel)
	w.Write(q)
}

func ShowKelurahan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sqlStatement := `SELECT id,kelurahan,kecamatan_id,kota_id,provinsi_id FROM kelurahan WHERE id = $1`
	var post Kelurahan
	db := config.Connect()
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.Name, &post.Kecamatan_id, &post.Kota_id, &post.Provinsi_id)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"data": "null",
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
