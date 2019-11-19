// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package accounts

// _Entity represents all types with @key
type _Entity interface {
	Is_Entity()
}

type Entity struct {
	FindReviewByID   *Review  `json:"findReviewByID"`
	FindUserByID     *User    `json:"findUserByID"`
	FindProductByUpc *Product `json:"findProductByUpc"`
}

type Product struct {
	Upc     string    `json:"upc"`
	Reviews []*Review `json:"reviews"`
}

func (Product) Is_Entity() {}

type User struct {
	ID       string    `json:"id"`
	Username *string   `json:"username"`
	Reviews  []*Review `json:"reviews"`
}

func (User) Is_Entity() {}