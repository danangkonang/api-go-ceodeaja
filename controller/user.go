package controller

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/danangkonang/rest-api/config"
	"github.com/danangkonang/rest-api/helper"
)

type Users struct {
	Id         int
	Email      string
	Username   sql.NullString
	Avatar     sql.NullString
	Active     bool
	Created_at time.Time
	Updated_at time.Time
}

type UpdateProfile struct {
	Id         int
	Email      string
	Username   sql.NullString
	Avatar     string
	Active     bool
	Created_at time.Time
	Updated_at time.Time
}

func GetProfil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	db := config.Connect()
	sqlStatement := `SELECT id, email, username, avatar, active, created_at, updated_at FROM users WHERE id = $1`
	post := Users{}
	row := db.QueryRow(sqlStatement, cek["Id"])
	err := row.Scan(&post.Id, &post.Email, &post.Username, &post.Avatar, &post.Active, &post.Created_at, &post.Updated_at)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"error": map[string]interface{}{},
		}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		json.NewEncoder(w).Encode(post)
	default:
		panic(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	sqlStatement := "SELECT id,username,email,active,created_at,updated_at from users"
	result, err := db.Query(sqlStatement)
	checkErr(err)
	var posts []interface{}
	for result.Next() {
		var u Users
		err := result.Scan(&u.Id, &u.Username, &u.Email, &u.Active, &u.Created_at, &u.Updated_at)
		checkErr(err)
		res := map[string]interface{}{
			"id":         u.Id,
			"name":       u.Username,
			"email":      u.Email,
			"active":     u.Active,
			"created_at": u.Created_at,
			"updates_at": u.Updated_at,
		}
		posts = append(posts, res)
	}
	defer result.Close()
	defer db.Close()
	json.NewEncoder(w).Encode(posts)
}

func UpdateProfilAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decode := json.NewDecoder(r.Body)
	var ava UpdateProfile
	err := decode.Decode(&ava)
	checkErr(err)
	dec, err := base64.StdEncoding.DecodeString(ava.Avatar)
	checkErr(err)
	namaFile := helper.RandomString(32) + `.png`
	f, err := os.Create("./files/profiles/" + namaFile)
	checkErr(err)
	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	db := config.Connect()
	var res UpdateProfile
	var avatarUpdate Users
	getAvatar := "SELECT avatar FROM users WHERE id = $1"
	errG := db.QueryRow(getAvatar, cek["Id"]).Scan(&avatarUpdate.Avatar)
	checkErr(errG)
	avatar := avatarUpdate.Avatar
	if avatar.Valid {
		removeAvatarName := "./files/profiles/" + avatar.String
		sqlStatement := "UPDATE users SET avatar = $1, updated_at = $2 WHERE id = $3 RETURNING id, email, username, avatar, active, created_at, updated_at"
		err = db.QueryRow(sqlStatement, namaFile, timeNow(), cek["Id"]).Scan(&res.Id, &res.Email, &res.Username, &res.Avatar, &res.Active, &res.Created_at, &res.Updated_at)
		checkErr(err)
		resJes := map[string]interface{}{
			"id":         res.Id,
			"email":      res.Email,
			"username":   res.Username,
			"avatar":     res.Avatar,
			"active":     res.Active,
			"created_at": res.Created_at.Format("2006-01-02 15:04:05"),
			"update_at":  res.Updated_at.Format("2006-01-02 15:04:05"),
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
		errRm := os.Remove(removeAvatarName)
		checkErr(errRm)
	} else {
		sqlStatement := "UPDATE users SET avatar = $1, updated_at = $2 WHERE id = $3 RETURNING id, email, username, avatar, active, created_at, updated_at"
		err = db.QueryRow(sqlStatement, namaFile, timeNow(), cek["Id"]).Scan(&res.Id, &res.Email, &res.Username, &res.Avatar, &res.Active, &res.Created_at, &res.Updated_at)
		checkErr(err)
		resJes := map[string]interface{}{
			"id":         res.Id,
			"email":      res.Email,
			"username":   res.Username,
			"avatar":     res.Avatar,
			"active":     res.Active,
			"created_at": res.Created_at.Format("2006-01-02 15:04:05"),
			"update_at":  res.Updated_at.Format("2006-01-02 15:04:05"),
		}
		re, _ := json.Marshal(resJes)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	}
}

func UpdateProfilData(w http.ResponseWriter, r *http.Request) {

}

func timeNow() string {
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	return waktu
}
