package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateCoupon(*Coupon) error
	DeleteCoupon(int) error
	UpdateCoupon(*Coupon) error
	GetCoupons() ([]*Coupon, error)
	GetCouponByID(int) (*Coupon, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=root sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	if err := s.CreateCouponTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) CreateCouponTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS Coupons (
			id SERIAL PRIMARY KEY,
			code VARCHAR(50),
			discount_type VARCHAR(50),
			discount_value DECIMAL(10, 2),
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP NOT NULL,
			min_cart_value DECIMAL(10, 2),
			applies_to_item JSONB
		)
	`
	// applies_to_ites contains array of sku according to coupon type
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateCoupon(c *Coupon) error {
	_, err := s.db.Query(
		`
			INSERT INTO coupons (
				code,
				discount_type,
				discount_value,
				start_date,
				end_date,
				min_cart_value,
				applies_to_item
			)
			VALUES(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7
			)	
		`,
		c.DiscoutType, c.DiscountValue, c.StartDate, c.EndDate, c.MinCartValue, c.AppliesToItem,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) UpdateCoupon(*Coupon) error {
	return nil
}

func (s *PostgressStore) DeleteCoupon(id int) error {
	_, err := s.db.Query("DELETE FROM coupons WHERE id = $1", id)
	return err
}

func (s *PostgressStore) GetCoupons() ([]*Coupon, error) {
	rows, err := s.db.Query("SELECT * FROM coupons")
	if err != nil {
		return nil, err
	}

	coupons := []*Coupon{}
	for rows.Next() {
		coupon, err := scanIntoCoupons(rows)

		if err != nil {
			return nil, err
		}

		coupons = append(coupons, coupon)
	}

	return coupons, nil
}

func (s *PostgressStore) GetCouponByID(id int) (*Coupon, error) {
	rows, err := s.db.Query("SELECT * FROM coupons WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoCoupons(rows)
	}
	return nil, fmt.Errorf("coupon %d not found", id)
}

func scanIntoCoupons(rows *sql.Rows) (*Coupon, error) {
	coupon := &Coupon{}

	err := rows.Scan(
		&coupon.ID,
		&coupon.DiscoutType,
		&coupon.DiscountValue,
		&coupon.StartDate,
		&coupon.EndDate,
		&coupon.MinCartValue,
		&coupon.AppliesToItem,
	)
	return coupon, err
}
