package query

import (
	"gorm.io/gorm"
)

func FillDB(db *gorm.DB) error {
	if db.Migrator().HasTable("customer_companies") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM customer_companies limit 1) THEN
		ELSE
		COPY customer_companies(company_id, company_name) FROM '/tmp/test_data/customer_companies.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("customers") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM customers limit 1) THEN
		ELSE
		COPY customers(user_id, login, password, name, company_id, credit_cards) FROM '/tmp/test_data/customers.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("order_items") {
		db.Exec(`CREATE TABLE IF NOT EXISTS public.order_items_temp
		(
			id text,
			order_id text,
			price_per_unit numeric,
			quantity int4,
			product text
		);
		
		CREATE OR REPLACE PROCEDURE public.transfer(
			)
		LANGUAGE 'plpgsql'
		AS $BODY$
		begin
		
		copy order_items_temp(id, order_id, price_per_unit, quantity, product)
		FROM '/tmp/test_data/order_items.csv'
		DELIMITER ','
		CSV header;
		
		INSERT INTO order_items(id, order_id, price_per_unit, quantity, product)
			SELECT id, order_id, price_per_unit, quantity, product
			FROM order_items_temp
			WHERE not EXISTS (
				SELECT id, order_id, price_per_unit, quantity, product
				FROM order_items
				WHERE order_items_temp.id = order_items.id);
				commit;
		
		truncate table order_items_temp;
		end;
		$BODY$;
		ALTER PROCEDURE public.transfer()
			OWNER TO postgres;
		
		CALL public.transfer();
		
		select * from order_items;
		
		select * from order_items_temp;`)
	}
	if db.Migrator().HasTable("orders") {

		db.Exec(`DO $$
		begin IF EXISTS (select * FROM orders limit 1) THEN
		ELSE
		COPY orders(id,created_at, order_name, customer_id) FROM '/tmp/test_data/orders.csv'
		DELIMITER ',' 
		CSV HEADER;
		END if;
		END $$;`)
	}
	if db.Migrator().HasTable("deliveries") {
		db.Exec(`DO $$
		begin IF EXISTS (select * FROM deliveries limit 1) THEN
		ELSE
		COPY deliveries(id, order_item_id, delivered_quantity) FROM '/tmp/test_data/deliveries.csv' DELIMITER ',' CSV HEADER;
		END if;
		END $$;`)
	}

	return nil
}
