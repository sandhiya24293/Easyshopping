package Model

type Fooddelivery struct {
	Hotelname    string
	Food         []Foodcount
	Address      string
	Mobilenumber string
	Status       string
	Loginid      string
}
type Ownfood struct {
	Dishname   string
	Rate       int
	Platecount int
}
type Responseownfood struct {
	ID         int
	Dishname   string
	Rate       int
	Platecount int
	Status     string
}

type Ownorderplaced struct {
	Loginid     string
	Food        []Ownfood
	TotalAmount int
	Noofitems   int
	Date        string
}

type ResOwnorderplaced struct {
	Id          int
	Loginid     string
	TotalAmount int
	Noofitems   int
	Date        string
	Status      string
}

type ResOwnorderplacedAll struct {
	Res  ResOwnorderplaced
	Food []Ownfood
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
