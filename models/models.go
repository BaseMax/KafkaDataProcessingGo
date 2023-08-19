package models

var USER_ACTIONS = []string{"view", "add_to_product", "purchase"}

type User struct {
	ID       uint   `json:"id" faker:"-"`
	Username string `json:"user" faker:"username"`
	Password string `json:"pass" faker:"password"`
	Email    string `json:"email" faker:"email"`
}

type UserActivity struct {
	UserID    uint   `json:"user_id" faker:"oneof: 1, 2, 3, 4, 5"`
	Timestamp string `json:"timestamp" faker:"time"`
	Action    string `json:"action" faker:"oneof: view, add_to_product, purchase"`
	ProductID uint   `json:"product_id" faker:"oneof: 1, 2, 3, 4, 5"`
}

type Product struct {
	ID          uint    `json:"id" faker:"-"`
	Name        string  `json:"name" faker:"word"`
	Description string  `json:"description" faker:"paragraph"`
	Price       float64 `faker:"oneof: 4.95, 9.99, 31997.97"`
}
