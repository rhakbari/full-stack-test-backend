package query

import (
	"gorm.io/gorm"
)

func FillDB(db *gorm.DB) error {
	if db.Migrator().HasTable("customer_companies") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM order_items limit 1) THEN
		ELSE
		COPY customer_companies(company_id, company_name) FROM '/tmp/test_data/customer_companies.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("customers") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM order_items limit 1) THEN
		ELSE
		COPY customers(user_id, login, password, name, company_id, credit_cards) FROM '/tmp/test_data/customers.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("order_items") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM order_items limit 1) THEN
		ELSE
		COPY order_items(id, order_id, price_per_unit, quantity, product)
		FROM '/tmp/test_data/order_items.csv' 
		DELIMITER ','
		CSV header;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("orders") {

		db.Exec(`DO $$
		begin IF EXISTS (select * FROM order_items limit 1) THEN
		ELSE
		COPY orders(id,created_at, order_name, customer_id) FROM '/tmp/test_data/orders.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("deliveries") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM order_items limit 1) THEN
		ELSE
		COPY deliveries(id, order_item_id, delivered_quantity) FROM '/tmp/test_data/deliveries.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}

	return nil
}
