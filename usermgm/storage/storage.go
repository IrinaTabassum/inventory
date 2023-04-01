package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

type User struct {
	ID        int          `json:"id" form:"-" db:"id"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Email     string       `json:"email" db:"email"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	Role      string       `json:"role" db:"role"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total     int          `json:"total" db:"total"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.When(u.ID == 0).Error("The username field is required."),
		),
		validation.Field(&u.Email,
			validation.Required.When(u.ID == 0).Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.When(u.ID == 0).Error("The password field is required."),
		),
	)
}

type Login struct {
	Username string
	Password string
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&l.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

type Supplier struct {
	ID        int          `form:"-" db:"id"`
	Name      string       `json:"name" db:"name"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total     int          `json:"total" db:"total"`
}

type SupplierFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

func (s Supplier) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name,
			validation.Required.Error("The name field is required."),
		),
	)
}

type Customer struct {
	ID        int          `form:"-" db:"id"`
	Name      string       `json:"name" db:"name"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total     int          `json:"total" db:"total"`
}

type CustomerFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("The name field is required."),
		),
	)
}

type Category struct {
	ID        int          `form:"-" db:"id"`
	Name      string       `json:"name" db:"name"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total     int          `json:"total" db:"total"`
}

type CategoryFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

func (c Category) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("The name field is required."),
		),
	)
}

type Product struct {
	ID           int          `form:"-" db:"id"`
	CategoryId   int          `json:"category_id" db:"category_id"`
	Name         string       `json:"name" db:"name"`
	CategoryName string       `json:"category_name" db:"category_name"`
	Description  string       `json:"description" db:"description"`
	Price        float32      `json:"price" db:"price"`
	Quantity     int          `json:"quantity" db:"quantity"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total        int          `json:"total" db:"total"`
}

type ProductFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name,
			validation.Required.Error("The name field is required."),
		),
		validation.Field(&p.CategoryId,
			validation.Required.Error("category id is required."),
		),
		validation.Field(&p.Price,
			validation.Required.Error("Price is required."),
		),
	)
}

type Purchase struct {
	ID           int          `form:"-" db:"id"`
	SupplierId   int          `json:"supplier_id" db:"supplier_id"`
	SupplierName string       `json:"name" db:"name"`
	ProductId    int          `json:"product_id" db:"product_id"`
	Quantity     int          `json:"quantity" db:"quantity"`
	UnitPrice    float32      `json:"unit_price" db:"unit_price"`
	TotalPrice   float32      `json:"total_price" db:"total_price"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total        int          `json:"total" db:"total"`
}

type PurchaseFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

func (p Purchase) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.SupplierId,
			validation.Required.Error("Supplier Id is required."),
		),
		validation.Field(&p.ProductId,
			validation.Required.Error("Product id is required."),
		),
		validation.Field(&p.Quantity,
			validation.Required.Error("Quantity is required."),
		),
	)
}

type Stock struct {
	ID          int       `form:"-" db:"id"`
	ProductId   int       `json:"product_id" db:"product_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	ProductNane string    `json:"name" db:"name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Total       int       `json:"total" db:"total"`
}

type StockFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

type Sell struct {
	ID           int       `form:"-" db:"id"`
	CustomerId   int       `json:"customer_id" db:"customer_id"`
	CustomerName string    `json:"name" db:"name"`
	TotalPrice   float32   `json:"total_price" db:"total_price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Total        int       `json:"total" db:"total"`
	SelPro       map[int32]int32
}

func (s Sell) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.CustomerId,
			validation.Required.Error("Customer Id is required."),
		),
		validation.Field(&s.SelPro,
			validation.Required.Error("SelPro is required."),
		),
	)
}

type SellFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

type SellProduct struct {
	ID             int       `form:"-" db:"id"`
	SellId         int       `json:"sell_id" db:"sell_id"`
	ProductId      int       `json:"product_id" db:"product_id"`
	ProductName    string    `json:"name" db:"name"`
	Quantity       int       `json:"quantity" db:"quantity"`
	UnitPrice      float32   `json:"unit_price" db:"unit_price"`
	TotalUnitPrice float32   `json:"total_unit_price" db:"total_unit_price"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Total          int       `json:"total" db:"total"`
}

type SellProductmap struct {
	SellId int
	SelPro map[int32]int32
}

type ProductStock struct {
	ID            int     `form:"-" db:"id"`
	ProductId     int     `json:"product_id" db:"product_id"`
	ProductName   string  `json:"name" db:"name"`
	UnitPrice     float32 `json:"price" db:"price"`
	StockQuantity int     `json:"quantity" db:"quantity"`
}

type SoldProduct struct {
	ProductId      int
	ProductName    string
	UnitPrice      float32
	Quantity       int
	TotalUnitPrice float32
}

type SaleReport struct {
	ProductId        int          `json:"product_id" db:"product_id"`
	ProductName      string       `json:"name" db:"name"`
	PurchaseQuantity int          `json:"purchase_quantity" db:"purchase_quantity"`
	SellQuantity     int          `json:"sell_quantity" db:"sell_quantity"`
	StockQuantity    int          `json:"stock_quantity" db:"stock_quantity"`
	DeletedAt        sql.NullTime `json:"deleted_at" db:"deleted_at"`
	Total            int          `json:"total" db:"total"`
}
type SaleReportFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}
