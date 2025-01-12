package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateCoupon(*Coupon) error
	DeleteCoupon(int) error
	UpdateCoupon(*Coupon) error
	GetCoupons() ([]*Coupon, error)
	GetCouponByID(int) (*Coupon, error)
	CreateBxGyItem(*BxGy) error
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {
	connStr := "host=localhost port=5432 user=postgres password=0000 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Printf("Database connected...\n")

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() error {
	if err := s.CreateCouponTable(); err != nil {
		return err
	}

	if err := s.CreateCouponUsageTable(); err != nil {
		return err
	}

	if err := s.CreateBXGYItemTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgressStore) CreateCouponTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS coupons (
			coupon_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			code VARCHAR(255) UNIQUE NOT NULL,
			discount_type INT NOT NULL,
			discount_value DECIMAL(10, 2),
			start_date TIMESTAMP,
			end_date TIMESTAMP,
			details JSON
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateCouponUsageTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS coupon_usage (
			usage_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			coupon_id INT,
			user_id INT,
			order_id INT,
			usage_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (coupon_id) REFERENCES coupons(coupon_id)
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateBXGYItemTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS bxgy_items (
			bxgy_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			coupon_id INT UNIQUE,
			bx_item_list JSON,
			gy_item_list JSON,
			FOREIGN KEY (coupon_id) REFERENCES coupons(coupon_id)
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateCoupon(c *Coupon) error {
	detailsJSON, err := json.Marshal(c.Details)
	if err != nil {
		return err
	}
	_, err = s.db.Query(
		`
			INSERT INTO coupons (
				code,
				discount_type,
				discount_value,
				start_date,
				end_date,
				details
			)
			VALUES(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
			)	
		`,
		c.Code, c.DiscoutType, c.DiscountValue, c.StartDate, c.EndDate, detailsJSON,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) UpdateCoupon(c *Coupon) error {
	detailsJSON, err := json.Marshal(c.Details)
	if err != nil {
		return err
	}
	_, err = s.db.Query(
		`
			UPDATE coupons SET (
				code,
				discount_type,
				discount_value,
				start_date,
				end_date,
				details
			)
			VALUES(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
		) WHERE coupon_id = $6
		`,
		c.Code, c.DiscoutType, c.DiscountValue, c.StartDate, c.EndDate, c.ID, string(detailsJSON),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) DeleteCoupon(id int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var exists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM bxgy_items WHERE coupon_id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = tx.Exec("DELETE FROM bxgy_items WHERE coupon_id = $1", id)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec("DELETE FROM coupons WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
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
		&coupon.Code,
		&coupon.DiscoutType,
		&coupon.DiscountValue,
		&coupon.StartDate,
		&coupon.EndDate,
		&coupon.Details,
	)
	return coupon, err
}

func (s *PostgressStore) CreateBxGyItem(i *BxGy) error {
	bxItemListJSON, err := json.Marshal(i.BxItemList)
	if err != nil {
		return err
	}

	gyItemListJSON, err := json.Marshal(i.GyItemList)
	if err != nil {
		return err
	}
	_, err = s.db.Query(
		`
			INSERT INTO bxgy_items (
				coupon_id,
				bx_item_list,
				gy_item_list
			)
			VALUES(
				$1,
				$2,
				$3
			)	
		`,
		i.CouponID, string(bxItemListJSON), string(gyItemListJSON),
	)
	return err
}
