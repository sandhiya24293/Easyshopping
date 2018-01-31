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
	Image      string
}
type Responseownfood struct {
	ID         int
	Dishname   string
	Rate       int
	Platecount int
	Image      string
	Status     string
}

type Ownorderplaced struct {
	Loginid     string
	Food        []Ownfood
	TotalAmount int
	Noofitems   int
	Date        string
}
type SendOwnfood struct {
	Data []Responseownfood
}
type ResOwnorderplaced struct {
	Id          int
	Loginid     string
	TotalAmount int
	Noofitems   int
	Date        string
	Status      string
}

type Ownfood1 struct {
	Id         int
	Dishname   string
	Rate       int
	Platecount int
	Image      string
}
type ResOwnorderplacedAll struct {
	Res         ResOwnorderplaced
	Emailid     string
	Phonenumber string
	Food        []Ownfood
}
type ResOwnorderplacedAll1 struct {
	Res         ResOwnorderplaced
	Emailid     string
	Phonenumber string
	Address     string

	Food []Ownfood1
}
type Foodcount struct {
	Foodname string
	Quantity int
}
type Foodcount1 struct {
	Foodid   int
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

type Foodtrackreesponse1 struct {
	Details     Foodtrack
	Fooddetails []Foodcount1
}
