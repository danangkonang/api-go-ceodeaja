package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/danangkonang/ceodeaja-go/controller"
	"github.com/danangkonang/ceodeaja-go/middleware"

	// "github.com/danangkonang/ceodeaja-go/migrate"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	STATIC_DIR = "/files/"
)

func main() {
	// migrate.Migrate()
	handleRequest()
}

func handleRequest() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	PORT := os.Getenv("PORT")
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	r.HandleFunc("/tes", controller.Tes).Methods("GET")

	// auth
	r.HandleFunc("/v1/login", controller.Login).Methods("POST")
	r.HandleFunc("/v1/registrasi", controller.Registrasi).Methods("POST")
	r.HandleFunc("/v1/request-reset-password", middleware.Auth(controller.RequestResetPassword)).Methods("POST")
	r.HandleFunc("/v1/refres-token", middleware.Auth(controller.RefresToken)).Methods("GET")
	// product
	r.HandleFunc("/v1/post-product", middleware.Auth(controller.PostProduct)).Methods("POST")
	r.HandleFunc("/v1/get-product/{id}", controller.ShowProduct).Methods("GET")
	r.HandleFunc("/v1/get-products", controller.ShowProducts).Methods("GET")
	r.HandleFunc("/v1/get-products/{id}", controller.ShowProductsByCategory).Methods("GET")
	r.HandleFunc("/v1/delete-products/{id}", middleware.Auth(controller.DestroyProduct)).Methods("DELETE")
	// address
	r.HandleFunc("/v1/provinsi", controller.ShowProvinces).Methods("GET")
	r.HandleFunc("/v1/provinsi/{id}", controller.ShowProvince).Methods("GET")
	r.HandleFunc("/v1/kabupaten-in-provinsi/{id}", controller.ShowCitys).Methods("GET")
	r.HandleFunc("/v1/kabupaten/{id}", controller.ShowCity).Methods("GET")
	r.HandleFunc("/v1/kecamatan-in-kabupaten/{id}", controller.ShowKecamatans).Methods("GET")
	r.HandleFunc("/v1/kecamatan/{id}", controller.ShowKecamatan).Methods("GET")
	r.HandleFunc("/v1/kelurahan-in-kecamatan/{id}", controller.ShowKelurahans).Methods("GET")
	r.HandleFunc("/v1/kelurahan/{id}", controller.ShowKelurahan).Methods("GET")
	// user
	r.HandleFunc("/v1/users", middleware.Auth(controller.GetUsers)).Methods("GET")
	r.HandleFunc("/v1/me", middleware.Auth(controller.GetProfil)).Methods("GET")
	r.HandleFunc("/v1/image-profile", middleware.Auth(controller.UpdateProfilAvatar)).Methods("PUT")
	r.HandleFunc("/v1/data-profile", middleware.Auth(controller.UpdateProfilData)).Methods("PUT")
	// category
	r.HandleFunc("/v1/category", controller.ShowCategory).Methods("GET")
	r.HandleFunc("/v1/categories", controller.ShowCategories).Methods("GET")
	r.HandleFunc("/v1/sub-categories/{id}", controller.ShowSubCategories).Methods("GET")
	// email
	r.HandleFunc("/v1/email", controller.SendEmail).Methods("POST")
	r.HandleFunc("/v1/verify-accunt/?email={email}&token={token}", controller.ActivasiAccunt).Methods("POST")

	fmt.Println("local server started at http://localhost:" + PORT)
	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Authorization",
	}
	method := []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	origin := []string{"*"}
	http.ListenAndServe(":"+PORT, handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))
}
