package Mysql

import (
	Model "Easyshopping/Model"
	_ "database/sql"

	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func Getvegetables_DB(Data string) (Vegarray []Model.Vegetables) {

	var Vegetabledata Model.Vegetables
	rows, err := OpenConnection["Rentmatics"].Query("select * from easyvegetables where Type=?", Data)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Vegetabledata.Vegid,
			&Vegetabledata.Vegetable,
			&Vegetabledata.Type,
			&Vegetabledata.Rate1kg,
			&Vegetabledata.Rate500gm,
			&Vegetabledata.Rate250gm,
			&Vegetabledata.Display,
			&Vegetabledata.Pictureurl,
		)
		Vegarray = append(Vegarray, Vegetabledata)

	}

	return
}

func GetSingleprod_DB(Productid int) (Responsedata Model.Vegetables) {

	rows, err := OpenConnection["Rentmatics"].Query("select * from easyvegetables where easyid=?", Productid)
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for rows.Next() {

		rows.Scan(

			&Responsedata.Vegid,
			&Responsedata.Vegetable,
			&Responsedata.Type,
			&Responsedata.Rate1kg,
			&Responsedata.Rate500gm,
			&Responsedata.Rate250gm,
			&Responsedata.Display,
			&Responsedata.Pictureurl,
		)

	}

	return
}

func Getveg() (Vegresponse []Model.Vegetables) {
	Vegresponse = Getvegetables_DB("Veg")
	return
}
func GetLeaves() (Vegresponse []Model.Vegetables) {
	Vegresponse = Getvegetables_DB("Leaves")
	return
}

func GetNonveg() (Vegresponse []Model.Vegetables) {
	Vegresponse = Getvegetables_DB("NonVeg")
	return
}

func GetFruits() (Vegresponse []Model.Vegetables) {
	Vegresponse = Getvegetables_DB("Fruits")
	return
}

func GetSeafood_DB() (Vegresponse []Model.Vegetables) {
	Vegresponse = Getvegetables_DB("Seafood")
	return
}
func Search_DB(searchstring string) (Vegarray []Model.Vegetables) {

	var Array []string
	var Product string
	var Productall []string
	var Vegetabledata Model.Vegetables
	row, err := OpenConnection["Rentmatics"].Query("select Easyproduct from easyvegetables")
	if err != nil {
		log.Println("Error -DB: Get User", err)
	}
	for row.Next() {

		row.Scan(

			&Product,
		)
		Productall = append(Productall, Product)

	}

	for _, v := range Productall {
		searchvar := "(?i)" + searchstring
		re1, _ := regexp.Compile(searchvar)
		result := re1.MatchString(v)

		if result == true {
			matchstring := v
			Array = append(Array, matchstring)

		}

	}

	for _, v := range Array {
		rows, err := OpenConnection["Rentmatics"].Query("select * from easyvegetables where Easyproduct=?", v)
		if err != nil {
			log.Println("Error -DB: Get User", err)
		}

		for rows.Next() {

			rows.Scan(

				&Vegetabledata.Vegid,
				&Vegetabledata.Vegetable,
				&Vegetabledata.Type,
				&Vegetabledata.Rate1kg,
				&Vegetabledata.Rate500gm,
				&Vegetabledata.Rate250gm,
				&Vegetabledata.Display,
				&Vegetabledata.Pictureurl,
			)
			Vegarray = append(Vegarray, Vegetabledata)

		}

	}

	return Vegarray
}
