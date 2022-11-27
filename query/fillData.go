package query

import (
	"gorm.io/gorm"
)

func FillDB(db *gorm.DB) error {
	if db.Migrator().HasTable("customer_companies") {
		db.Exec("COPY customer_companies(company_id, company_name) FROM '/tmp/test_data/customer_companies.csv' DELIMITER ',' CSV HEADER;")
	}
	if db.Migrator().HasTable("customers") {
		db.Exec("COPY customers(user_id, login, password, name, company_id, credit_cards) FROM '/tmp/test_data/customers.csv' DELIMITER ',' CSV HEADER;")
	}
	if db.Migrator().HasTable("order_items") {
		db.Exec("if not exists(select 1 from table)	begin COPY order_items(id, order_id,price_per_unit, quantity, product) FROM '/tmp/test_data/order_items.csv' DELIMITER ',' CSV HEADER end;")
	}
	if db.Migrator().HasTable("orders") {
		db.Exec("COPY orders(id,created_at, order_name, customer_id) FROM '/tmp/test_data/orders.csv' DELIMITER ',' CSV HEADER;")
	}
	if db.Migrator().HasTable("deliveries") {
		db.Exec("COPY deliveries(id, order_item_id, delivered_quantity) FROM '/tmp/test_data/deliveries.csv' DELIMITER ',' CSV HEADER;")
	}
	
	return nil
}
