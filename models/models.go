package models

var USER_ACTIONS = []string{"view", "add_to_product", "purchase"}

type UserActivity struct {
	UserID    uint   `json:"user_id" faker:"oneof: 1, 2, 3, 4, 5"`
	Timestamp string `json:"timestamp" faker:"time"`
	Action    string `json:"action" faker:"oneof: view, add_to_product, purchase"`
	ProductID uint   `json:"product_id" faker:"oneof: 1, 2, 3, 4, 5"`
}
