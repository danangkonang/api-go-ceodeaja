package seeds

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func CreateProvinces() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO provinces (id,province) VALUES
		(1, 'BANTEN'),
		(2, 'DKI JAKARTA'),
		(3, 'JAWA BARAT'),
		(4, 'JAWA TENGAH'),
		(5, 'DI YOGYAKARTA'),
		(6, 'JAWA TIMUR'),
		(7, 'BALI'),
		(8, 'NANGGROE ACEH DARUSSALAM (NAD)'),
		(9, 'SUMATERA UTARA'),
		(10, 'SUMATERA BARAT'),
		(11, 'RIAU'),
		(12, 'KEPULAUAN RIAU'),
		(13, 'JAMBI'),
		(14, 'BENGKULU'),
		(15, 'SUMATERA SELATAN'),
		(16, 'BANGKA BELITUNG'),
		(17, 'LAMPUNG'),
		(18, 'KALIMANTAN BARAT'),
		(19, 'KALIMANTAN TENGAH'),
		(20, 'KALIMANTAN SELATAN'),
		(21, 'KALIMANTAN TIMUR'),
		(22, 'KALIMANTAN UTARA'),
		(23, 'SULAWESI BARAT'),
		(24, 'SULAWESI SELATAN'),
		(25, 'SULAWESI TENGGARA'),
		(26, 'SULAWESI TENGAH'),
		(27, 'GORONTALO'),
		(28, 'SULAWESI UTARA'),
		(29, 'MALUKU'),
		(30, 'MALUKU UTARA'),
		(31, 'NUSA TENGGARA BARAT (NTB)'),
		(32, 'NUSA TENGGARA TIMUR (NTT)'),
		(33, 'PAPUA BARAT'),
		(34, 'PAPUA')
	`)
	if err != nil {
		fmt.Println("error insert province")
		panic(err)
	}
	fmt.Println("insert province")
}
