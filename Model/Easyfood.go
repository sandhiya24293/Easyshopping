package Model

type Fooddelivery struct {
	Hotelname    string
	Food         []Foodcount
	Address      string
	Mobilenumber string
	Status       string
	Loginid      string
}

type Foodcount struct {
	Foodname string
	Quantity int
}

type Foodtrack struct {
	Foodid       int
	Hotelname    string
	Address      string
	Mobilenumber string
	Status       string
	Loginid      string
}

//type Updateinstant struct {
//	Easyid int
//	Status string
//}
type Foodtrackreesponse struct {
	Details     Foodtrack
	Fooddetails []Foodcount
}
