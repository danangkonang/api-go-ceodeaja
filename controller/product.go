package controller

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/ceodeaja-go/config"
	"github.com/danangkonang/ceodeaja-go/helper"

	"github.com/gorilla/mux"
)

type ResBarang struct {
	Id            string
	User_id       string
	Product_name  string
	Images        string
	Prince        string
	Description   string
	Brand         sql.NullString
	Address       string
	Condition     bool
	Stock         int64
	Active        bool
	CategoryId    int64
	SubCategoryId int64
	Weight        int64
	Weight_unit   string
	Created_at    time.Time
	Updated_at    time.Time
}

type ReqBarang struct {
	Id           string
	User_id      string
	Prince       int64
	Stock        int64
	CategoryId   int64
	Weight       int64
	Product_name string
	Description  string
	Address      string
	Weight_unit  string
	Condition    bool
	Active       bool
	Images       []string
	Created_at   time.Time
	Updated_at   time.Time
}

type Categori struct {
	Category    string
	SubCategory string
	Item        string
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var req ReqBarang
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	fmt.Println("erorr bos")
	// 	fmt.Println(err)
	// }
	// upload, handle, err := r.FormFile("gambar")
	// if err != nil {
	// 	panic(err)
	// }
	// defer upload.Close()
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// fileName := handle.Filename
	// if r.FormValue("gambar") != "" {
	// 	fileName = fmt.Sprintf("%s%s", r.FormValue("gambar"))
	// }

	// fmt.Println("handle", handle.Filename)
}

func PostProductBackUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ReqBarang
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}
	newFile := []string{}
	for _, val := range req.Images {
		randName := helper.UnixRandomString(32)
		coI := strings.Index(string(val), ",")
		rawImage := string(val)[coI+1:]
		buff := bytes.Buffer{}
		switch strings.TrimSuffix(val[5:coI], ";base64") {
		case "image/png":
			reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(rawImage))
			_, format, err := image.DecodeConfig(reader)
			if err != nil {
				log.Fatal(err)
			}
			fileName := randName + "." + format
			ioutil.WriteFile("./files/products/"+fileName, buff.Bytes(), 0644)
			newFile = append(newFile, fileName)
		case "image/jpeg":
			dec, err := base64.StdEncoding.DecodeString(rawImage)
			checkErr(err)
			format := "jpeg"
			fileName := randName + "." + format
			f, err := os.Create("./files/products/" + fileName)
			checkErr(err)
			defer f.Close()
			if _, err := f.Write(dec); err != nil {
				panic(err)
			}
			if err := f.Sync(); err != nil {
				panic(err)
			}
			newFile = append(newFile, fileName)
		}
	}
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	gambar := strings.Join(newFile, ",")
	userId := cek["Id"]
	db := config.Connect()
	sqlStatement := `
		INSERT INTO products(id_product, user_id, product_name, images, prince, description, address, condition, stock, category_id, active, weight, weight_unit, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING *`
	var resp ResBarang
	err := db.QueryRow(sqlStatement, helper.UnixRandomString(32), userId, req.Product_name, gambar, req.Prince, req.Description, req.Address, req.Condition, req.Stock, req.CategoryId, req.Active, req.Weight, req.Weight_unit, helper.GetTime(), helper.GetTime()).Scan(
		&resp.Id,
		&resp.User_id,
		&resp.Product_name,
		&resp.Images,
		&resp.Prince,
		&resp.Description,
		&resp.Brand,
		&resp.Address,
		&resp.Condition,
		&resp.Stock,
		&resp.CategoryId,
		&resp.Weight,
		&resp.Weight_unit,
		&resp.Active,
		&resp.Created_at,
		&resp.Updated_at,
	)
	checkErr(err)
	defer db.Close()
	avatars := strings.Split(resp.Images, ",")
	resJes := map[string]interface{}{
		"id":           resp.Id,
		"user_id":      resp.User_id,
		"product_name": resp.Product_name,
		"images":       avatars,
		"prince":       resp.Prince,
		"description":  resp.Description,
		"brand":        resp.Brand.String,
		"address":      resp.Address,
		"new":          resp.Condition,
		"category": map[string]interface{}{
			"category_id": resp.CategoryId,
		},
		"active": resp.Active,
		"create": resp.Created_at.Format("2006-01-02 15:04:05"),
		"update": resp.Updated_at.Format("2006-01-02 15:04:05"),
	}
	re, _ := json.Marshal(resJes)
	w.Write(re)
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	sqlStatement := `SELECT id, user_id, product_name, images, prince, description, address, new, active, created_at, updated_at FROM products WHERE id = $1`
	var post ResBarang
	row := db.QueryRow(sqlStatement, params["id"])
	err := row.Scan(&post.Id, &post.User_id, &post.Product_name, &post.Images, &post.Prince, &post.Description, &post.Address, &post.Condition, &post.Active, &post.Created_at, &post.Updated_at)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{}
		r, _ := json.Marshal(res)
		w.Write(r)
	case nil:
		avatars := strings.Split(post.Images, ",")
		resJes := map[string]interface{}{
			"id":           post.Id,
			"user_id":      post.User_id,
			"product_name": post.Product_name,
			"images":       avatars,
			"prince":       post.Prince,
			"ddescription": post.Description,
			"address":      post.Address,
			"new":          post.Condition,
			"active":       post.Active,
			"create":       post.Created_at.Format("2006-01-02 15:04:05"),
			"update":       post.Updated_at.Format("2006-01-02 15:04:05"),
		}
		res := map[string]interface{}{
			"success": true,
			"data":    resJes,
		}
		re, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(re)
	default:
		panic(err)
	}
}

func PostProduct2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ReqBarang
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}
	newFile := []string{}
	for _, val := range req.Images {
		idx := strings.Index(val, ";base64,")
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(val[idx+8:]))
		buff := bytes.Buffer{}
		_, err := buff.ReadFrom(reader)
		checkErr(err)
		_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
		checkErr(err)
		nama := helper.RandomString(32)
		fileName := nama + "." + fm
		ioutil.WriteFile("./files/products/"+fileName, buff.Bytes(), 0644)
		newFile = append(newFile, fileName)
	}
	gmbr := strings.Join(newFile, ",")
	mesege, k := requiredProduct(req.Product_name, req.Address, req.Description)

	// fmt.Println(k)
	if k != true {
		res := string(mesege)
		var responRequired map[string]interface{}
		json.Unmarshal([]byte(res), &responRequired)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mesege)
		return
	}
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	// fmt.Println(gmbr)
	// fmt.Println(mesege)
	// fmt.Println(cek["Id"])
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	db := config.Connect()
	defer db.Close()
	sqlStatement := `
		INSERT INTO products(user_id, product_name, images, prince, description, address, new, category_id,sub_category_id,active, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING *`
	var resp ResBarang
	err := db.QueryRow(sqlStatement, cek["Id"], req.Product_name, gmbr, req.Prince, req.Description, req.Address, req.Condition, 2, 1, false, waktu, waktu).Scan(
		&resp.Id,
		&resp.User_id,
		&resp.Product_name,
		&resp.Images,
		&resp.Prince,
		&resp.Description,
		&resp.Brand,
		&resp.Address,
		&resp.Condition,
		&resp.CategoryId,
		&resp.SubCategoryId,
		&resp.Active,
		&resp.Created_at,
		&resp.Updated_at,
	)
	checkErr(err)
	avatars := strings.Split(resp.Images, ",")
	resJes := map[string]interface{}{
		"id":           resp.Id,
		"user_id":      resp.User_id,
		"product_name": resp.Product_name,
		"images":       avatars,
		"prince":       resp.Prince,
		"description":  resp.Description,
		"brand":        resp.Brand.String,
		"address":      resp.Address,
		"new":          resp.Condition,
		"category": map[string]interface{}{
			"category_id":     resp.CategoryId,
			"sub_category_id": resp.SubCategoryId,
		},
		"active": resp.Active,
		"create": resp.Created_at.Format("2006-01-02 15:04:05"),
		"update": resp.Updated_at.Format("2006-01-02 15:04:05"),
	}
	re, _ := json.Marshal(resJes)
	w.Write(re)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func requiredProduct(name, address, description string) ([]byte, bool) {
	switch {
	case len(name) == 0 && len(address) == 0 && len(description) == 0:
		res := map[string]interface{}{
			"error": map[string]interface{}{
				"product_name": "Required",
				"prince":       "Required",
				"address":      "Required",
				"description":  "Required",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	default:
		return nil, true
	}
}

type Person struct {
	Friends []string
}

func ShowProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod []interface{}
	db := config.Connect()
	query := `
	SELECT 
	p.id_product, 
	p.user_id, 
	p.product_name, 
	p.images, 
	p.prince, 
	p.description, 
	p.address, 
	p.condition, 
	p.active, 
	p.created_at, 
	p.updated_at,
	i.category_id
	FROM products p 
	LEFT JOIN sub_items i ON p.category_id = i.id_sub_item
	`
	result, err := db.Query(query)
	checkErr(err)
	for result.Next() {
		var post ResBarang
		var cat Categori
		err := result.Scan(&post.Id, &post.User_id, &post.Product_name, &post.Images, &post.Prince, &post.Description, &post.Address, &post.Condition, &post.Active, &post.Created_at, &post.Updated_at, &cat.Item)
		checkErr(err)
		avatars := strings.Split(post.Images, ",")
		resJes := map[string]interface{}{
			"id":           post.Id,
			"category":     cat.Item,
			"user_id":      post.User_id,
			"product_name": post.Product_name,
			"images":       avatars,
			"prince":       post.Prince,
			"description":  post.Description,
			"address":      post.Address,
			"new":          post.Condition,
			"active":       post.Active,
			"created_at":   post.Created_at.Format("2006-01-02 15:04:05"),
			"updated_at":   post.Updated_at.Format("2006-01-02 15:04:05"),
		}
		prod = append(prod, resJes)
	}
	defer db.Close()
	if len(prod) > 0 {
		res := map[string]interface{}{
			"success": true,
			"data":    prod,
		}
		q, _ := json.Marshal(res)
		w.Write(q)
	} else {
		f2 := make([]string, 0)
		json2, _ := json.Marshal(f2)
		w.Write(json2)
	}

	// if string(q) == "null" {
	// } else {
	// }
}

func ShowProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var prod []interface{}
	db := config.Connect()
	query := `
	SELECT 
	p.id, 
	p.user_id, 
	p.product_name, 
	p.images, 
	p.prince, 
	p.description, 
	p.address, 
	p.new, 
	p.active, 
	p.created_at, 
	p.updated_at,
	c.category,
	s.sub_category
	FROM products p
	INNER JOIN categories c ON p.category_id = c.id
	INNER JOIN subCategories s ON p.sub_category_id = s.id
	WHERE c.id=$1
	`
	result, err := db.Query(query, params["id"])
	checkErr(err)
	for result.Next() {
		var post ResBarang
		var cat Categori
		err := result.Scan(&post.Id, &post.User_id, &post.Product_name, &post.Images, &post.Prince, &post.Description, &post.Address, &post.Condition, &post.Active, &post.Created_at, &post.Updated_at, &cat.Category, &cat.SubCategory)
		checkErr(err)
		avatars := strings.Split(post.Images, ",")
		resJes := map[string]interface{}{
			"id":           post.Id,
			"category":     cat.Category,
			"sub_category": cat.SubCategory,
			"user_id":      post.User_id,
			"product_name": post.Product_name,
			"images":       avatars,
			"prince":       post.Prince,
			"description":  post.Description,
			"address":      post.Address,
			"new":          post.Condition,
			"active":       post.Active,
			"created_at":   post.Created_at.Format("2006-01-02 15:04:05"),
			"updated_at":   post.Updated_at.Format("2006-01-02 15:04:05"),
		}
		prod = append(prod, resJes)
	}
	defer db.Close()
	q, _ := json.Marshal(prod)
	w.Write(q)
}

func DestroyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1")
	checkErr(err)
	_, err = stmt.Exec(params["id"])
	checkErr(err)
	res := map[string]interface{}{
		"message": "id " + params["id"] + " deleted",
	}
	userJson, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ReqBarang
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	db := config.Connect()
	defer db.Close()
	sqlStatement := `
		UPDATE products SET product_name=$1, prince=$2, description=$3, address=$4, new=$5, category_id=$6,sub_category_id=$7, active=$8, updated_at=$9 WHERE id=$10 RETURNING *`
	var resp ResBarang
	err := db.QueryRow(sqlStatement, req.Product_name, req.Prince, req.Description, req.Address, req.Condition, 2, 1, false, waktu, cek["Id"]).Scan(
		&resp.Id,
		&resp.User_id,
		&resp.Product_name,
		&resp.Images,
		&resp.Prince,
		&resp.Description,
		&resp.Brand,
		&resp.Address,
		&resp.Condition,
		&resp.CategoryId,
		&resp.SubCategoryId,
		&resp.Active,
		&resp.Created_at,
		&resp.Updated_at,
	)
	checkErr(err)
	avatars := strings.Split(resp.Images, ",")
	resJes := map[string]interface{}{
		"id":           resp.Id,
		"user_id":      resp.User_id,
		"product_name": resp.Product_name,
		"images":       avatars,
		"prince":       resp.Prince,
		"description":  resp.Description,
		"brand":        resp.Brand.String,
		"address":      resp.Address,
		"new":          resp.Condition,
		"category": map[string]interface{}{
			"category_id":     resp.CategoryId,
			"sub_category_id": resp.SubCategoryId,
		},
		"active": resp.Active,
		"create": resp.Created_at.Format("2006-01-02 15:04:05"),
		"update": resp.Updated_at.Format("2006-01-02 15:04:05"),
	}
	re, _ := json.Marshal(resJes)
	w.Write(re)
}

func UpdateImageProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ReqBarang
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
	}
	newFile := []string{}
	for _, val := range req.Images {
		idx := strings.Index(val, ";base64,")
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(val[idx+8:]))
		buff := bytes.Buffer{}
		_, err := buff.ReadFrom(reader)
		checkErr(err)
		_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
		checkErr(err)
		nama := helper.RandomString(32)
		fileName := nama + "." + fm
		ioutil.WriteFile("./files/products/"+fileName, buff.Bytes(), 0644)
		newFile = append(newFile, fileName)
	}
	gmbr := strings.Join(newFile, ",")
	mesege, k := requiredProduct(req.Product_name, req.Address, req.Description)
	if k != true {
		res := string(mesege)
		var responRequired map[string]interface{}
		json.Unmarshal([]byte(res), &responRequired)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mesege)
		return
	}
	authorizationHeader := r.Header.Get("Authorization")
	cek := helper.DataToken(authorizationHeader)
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	db := config.Connect()
	defer db.Close()
	sqlStatement := `
		UPDATE products SET images=$1, updated_at=$2 WHERE id=$3 RETURNING *`
	var resp ResBarang
	err := db.QueryRow(sqlStatement, gmbr, waktu, cek["Id"]).Scan(
		&resp.Id,
		&resp.User_id,
		&resp.Product_name,
		&resp.Images,
		&resp.Prince,
		&resp.Description,
		&resp.Brand,
		&resp.Address,
		&resp.Condition,
		&resp.CategoryId,
		&resp.SubCategoryId,
		&resp.Active,
		&resp.Created_at,
		&resp.Updated_at,
	)
	checkErr(err)
	avatars := strings.Split(resp.Images, ",")
	resJes := map[string]interface{}{
		"id":           resp.Id,
		"user_id":      resp.User_id,
		"product_name": resp.Product_name,
		"images":       avatars,
		"prince":       resp.Prince,
		"description":  resp.Description,
		"brand":        resp.Brand.String,
		"address":      resp.Address,
		"new":          resp.Condition,
		"category": map[string]interface{}{
			"category_id":     resp.CategoryId,
			"sub_category_id": resp.SubCategoryId,
		},
		"active": resp.Active,
		"create": resp.Created_at.Format("2006-01-02 15:04:05"),
		"update": resp.Updated_at.Format("2006-01-02 15:04:05"),
	}
	re, _ := json.Marshal(resJes)
	w.Write(re)
}
