package library

import (
	"fmt"
)

func ValidateProduct(name, address, description, weight_unit string, weight, prince, category, stock int64, active, condition bool, image []string) {
	// switch {
	// case len(name) == 0 && len(prince) == 0 && len(address) == 0 && len(description) == 0:
	// 	fmt.Println("product_name,prince,address,description=required")
	// default:
	// 	fmt.Println("oke")
	// }
}

func Validate(key string, data []string) {
	for _, va := range data {
		switch va {
		case "required":
			fmt.Println("required")
		case "email":
			fmt.Println("email")
		default:
			fmt.Println("tong")
		}
	}
}
