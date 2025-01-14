package entities

type Order struct {
	Id           string      `json:"id" gorm:"primaryKey"`
	UserId       string      `json:"userId"`
	Item         []OrderItem `json:"item"`
	TotalPrice   float64     `json:"totalPrice"`
	Status       string      `json:"status"`
	CreationDate string      `json:"creationDate"`
	LastUpdate   string      `json:"lastUpdate"`
}

type OrderItem struct {
	ProductID int64
	Quantity  int
	Price     float64
}

func (o *Order) Update(order *Order) {
	o.UserId = order.UserId
	o.Status = order.Status
	o.LastUpdate = order.LastUpdate
}

func (o *Order) InValidOrder(order *Order) bool {
	return order.UserId == "" || order.Status == ""
}
