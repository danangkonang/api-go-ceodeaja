package controller

import (
	"fmt"
	"net/http"

	"github.com/danangkonang/ceodeaja-go/helper"
)

func Tes(w http.ResponseWriter, r *http.Request) {
	fmt.Println(helper.UnixRandomString(32))
}
