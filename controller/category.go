package controller

import (
	"encoding/json"
	"net/http"

	"github.com/danangkonang/rest-api/config"
	"github.com/gorilla/mux"
)

type Category struct {
	Id       int64
	Category string
}

type Sub struct {
	Id          int64
	CategoryId  string
	SubCategory string
}

type Items struct {
	Id            int64
	CategoryId    string
	SubCategoryId string
	Items         string
}

func ShowCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sqlStatement := `SELECT id,category FROM categories`
	db := config.Connect()
	result, err := db.Query(sqlStatement)
	checkErr(err)
	var resCat []interface{}
	for result.Next() {
		var res Category
		err := result.Scan(&res.Id, &res.Category)
		checkErr(err)
		query := `SELECT id,category_id,sub_category FROM sub_categories WHERE category_id=$1`
		rslt, err := db.Query(query, res.Id)
		checkErr(err)
		var resSubCat []interface{}
		for rslt.Next() {
			var resSub Sub
			err := rslt.Scan(&resSub.Id, &resSub.CategoryId, &resSub.SubCategory)
			checkErr(err)
			querys := `SELECT id,category_id,sub_category_id,sub_item FROM sub_categori_items WHERE sub_category_id=$1`
			resItm, err := db.Query(querys, resSub.Id)
			checkErr(err)
			var resItem []interface{}
			for resItm.Next() {
				var resItems Items
				err := resItm.Scan(&resItems.Id, &resItems.CategoryId, &resItems.SubCategoryId, &resItems.Items)
				checkErr(err)
				g := map[string]interface{}{
					"id":              resItems.Id,
					"category_id":     resItems.CategoryId,
					"sub_category_id": resItems.SubCategoryId,
					"name":            resItems.Items,
				}
				resItem = append(resItem, g)
			}
			rj := map[string]interface{}{
				"id_sub_categori": resSub.Id,
				"categori_id":     resSub.CategoryId,
				"name":            resSub.SubCategory,
				"items":           resItem,
			}
			resSubCat = append(resSubCat, rj)
		}
		rjs := map[string]interface{}{
			"id_category":  res.Id,
			"name":         res.Category,
			"sub_category": resSubCat,
		}
		resCat = append(resCat, rjs)
	}
	defer db.Close()
	res := map[string]interface{}{
		"success": true,
		"data":    resCat,
	}
	o, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}

func ShowCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sqlStatement := `SELECT id,category FROM categories`
	db := config.Connect()
	result, err := db.Query(sqlStatement)
	checkErr(err)
	var resCat []interface{}
	for result.Next() {
		var res Category
		err := result.Scan(&res.Id, &res.Category)
		checkErr(err)
		rjs := map[string]interface{}{
			"id_category": res.Id,
			"name":        res.Category,
		}
		resCat = append(resCat, rjs)
	}
	defer db.Close()
	q, _ := json.Marshal(resCat)
	w.Write(q)
}

func ShowSubCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := config.Connect()
	query := `SELECT id,category_id,sub_category FROM sub_categories WHERE category_id=$1`
	rslt, err := db.Query(query, params["id"])
	checkErr(err)
	var resSubCat []interface{}
	for rslt.Next() {
		var resSub Sub
		err := rslt.Scan(&resSub.Id, &resSub.CategoryId, &resSub.SubCategory)
		checkErr(err)
		rj := map[string]interface{}{
			"id_sub_categori": resSub.Id,
			"id_categori":     resSub.CategoryId,
			"name":            resSub.SubCategory,
		}
		resSubCat = append(resSubCat, rj)
	}
	defer db.Close()
	q, _ := json.Marshal(resSubCat)
	w.Write(q)
}
