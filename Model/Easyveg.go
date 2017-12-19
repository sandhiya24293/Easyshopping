package Model

type Vegetables struct {
	Vegid     int
	Vegetable string
	Type      string
	Rate      int
	Weight    string
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

type Orderplaced struct {
	Noofproducts int
	Loginid      string
	Products     []Product
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
