package repository

import (
	"database/sql"
	"time"
)

type OmzetRepository interface {
	GetMerchantOmzet(merchantId int, limit int, offset int) (map[string]interface{}, error)
	GetOutletOmzet(outletId int, limit int, offset int) (map[string]interface{}, error)
	GetMerchantName(merchantId int) (string, error)
	GetOutletName(outletId int) (string, string, error)
}

type mysqlOmzetRepository struct {
	db *sql.DB
}

func NewMySQLOmzetRepository(db *sql.DB) OmzetRepository {
	return &mysqlOmzetRepository{
		db: db,
	}
}

func (m *mysqlOmzetRepository) GetMerchantOmzet(merchantId int, limit int, offset int) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT SUM(Transactions.bill_total), Transactions.created_at FROM Merchants JOIN Transactions ON Merchants.id=Transactions.merchant_id WHERE Transactions.merchant_id=? GROUP BY DATE(Transactions.`created_at`) LIMIT ? OFFSET ?", merchantId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]interface{})

	for rows.Next() {
		var date time.Time
		var omzet int
		err = rows.Scan(
			&omzet,
			&date,
		)
		if err != nil {
			return nil, err
		}
		r := make(map[string]interface{})
		r["omzet"] = omzet
		dateString := date.Format("2006-01-02")
		if err != nil {
			return nil, err
		}
		result[dateString] = r
	}

	return result, nil

}

func (m *mysqlOmzetRepository) GetOutletOmzet(outletId int, limit int, offset int) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT SUM(Transactions.bill_total), Transactions.created_at FROM Merchants JOIN Outlets ON Merchants.id=Outlets.merchant_id JOIN Transactions ON Outlets.id=Transactions.outlet_id WHERE Transactions.outlet_id=? GROUP BY DATE(Transactions.`created_at`), Outlets.id LIMIT ? OFFSET ?", outletId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]interface{})

	for rows.Next() {
		var date time.Time
		var omzet int
		err = rows.Scan(
			&omzet,
			&date,
		)
		if err != nil {
			return nil, err
		}
		r := make(map[string]interface{})
		r["omzet"] = omzet
		dateString := date.Format("2006-01-02")
		if err != nil {
			return nil, err
		}
		result[dateString] = r
	}

	return result, nil

}

func (m *mysqlOmzetRepository) GetMerchantName(merchantId int) (string, error) {
	row := m.db.QueryRow("SELECT Merchants.merchant_name FROM Merchants WHERE id=?", merchantId)
	var merchantName string
	err := row.Scan(&merchantName)
	if err != nil {
		return "", err
	}

	return merchantName, nil

}

func (m *mysqlOmzetRepository) GetOutletName(outletId int) (string, string, error) {
	row := m.db.QueryRow("SELECT Merchants.merchant_name, Outlets.outlet_name FROM Outlets JOIN Merchants ON Merchants.id=Outlets.merchant_id WHERE Outlets.id=?", outletId)
	var merchantName string
	var outletName string
	err := row.Scan(&merchantName, &outletName)
	if err != nil {
		return "", "", err
	}

	return merchantName, outletName, nil

}
