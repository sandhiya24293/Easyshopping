package Model

type Senddata struct {
	Data []Vegetables
}
type Vegetables struct {
	Vegid      int
	Vegetable  string
	Type       string
	Rate1kg    int
	Rate500gm  int
	Rate250gm  int
	Display    string
	Pictureurl string
}
type ProductResponse struct {
	Vegetables []Vegetables
	Nonveg     []Vegetables
	Fruits     []Vegetables
	Fish       []Vegetables
}
type Search struct {
	Searchproduct string
}

type Product struct {
	Product string
	Rate    int
	Weight  string
}

type Product1 struct {
	Id      int
	Product string
	Rate    int
	Weight  string
}
type Orderplaced struct {
	Noofproducts int
	Loginid      string
	Products     []Product
	TotalAmount  int
	Date         string
}

type ResponseOrderplaced struct {
	Billid      string
	Address     string
	Phonenumber string

	Totalproduct string
	Products     []Product1
	TotalAmount  int
	Date         string
}

type Orderdata struct {
	Userid      string
	Uniqueid    string
	Emailid     string
	Phonenumber string
}

type Placeorderresponse struct {
	Orderid int
	Billid  string
}

type Ordertrack struct {
	Loginid string
}
type Ordertrackid struct {
	Orderid int
}

type Ordertrackresponse struct {
	Ordertrackingid string
	Orderid         string
	Orderplacedid   string
	Loginid         string
	Status          string
}
type Trackingresponse struct {
	Ordertrackingid int
	Orderid         int
	Orderplacedid   string
	Loginid         string
	Status          string
}
type Trackingresponses struct {
	Ordertrackingid int
	Orderid         int
	Orderplacedid   string
	Loginid         string
	Status          string
	Address         string
	Phonenumber     string
}

type OrderPlacedhistory struct {
	Orderid      string
	Billid       string
	Totalproduct int
	Totalamount  int
	Date         string
}

type Orderhistory struct {
	Orderdetails   OrderPlacedhistory
	Productdetails []Product
}
type Catergorylist struct {
	CategoryName string
	Url          string
}
type Tracking struct {
	Address string
	Track   []Trackingresponse
}

type Datares struct {
	Data []Catergorylist
}

type Ordercancel struct {
	Loginid string
	Orderid int
}
