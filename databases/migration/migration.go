package main

import (
	"fmt"
	"log"

	"github.com/onizukazaza/tar-ecom-api/config"
	"github.com/onizukazaza/tar-ecom-api/databases"
	"github.com/jmoiron/sqlx"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx, err := db.Connect().Beginx()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	if err := runMigrations(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Migration failed: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("All migrations applied successfully!")
}

func runMigrations(tx *sqlx.Tx) error {
	migrations := []struct {
		query     string
		tableName string
	}{
		{query: createUserTable, tableName: "users"},
		{query: createAddressTable, tableName: "addresses"},
		{query: createProductTable, tableName: "products"},
		{query: createProductImageTable, tableName: "product_image"},
		{query: createColorTable, tableName: "color"},
		{query: createSizeTable, tableName: "size"},
		{query: createProductVariationTable, tableName: "product_variation"},
		{query: createOrderTable, tableName: "order"},
		{query: createOrderItemTable, tableName: "order_item"},
	}

	for _, migration := range migrations {
		if _, err := tx.Exec(migration.query); err != nil {
			return fmt.Errorf("failed to migrate table %s: %v", migration.tableName, err)
		}
		fmt.Printf("Table %s migrated successfully\n", migration.tableName)
	}

	return nil
}

const createUserTable = `
CREATE TYPE IF NOT EXISTS user_role AS ENUM ('buyer', 'seller', 'admin');
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	username VARCHAR(60),
	email VARCHAR(100),
	password VARCHAR(255),
	role user_role DEFAULT 'buyer',
	profile_image VARCHAR(255),
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
)`

const createAddressTable = `
CREATE TABLE IF NOT EXISTS addresses (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	recipient_name VARCHAR(60) NOT NULL,
	province VARCHAR(50),
	district VARCHAR(100),
	subdistrict VARCHAR(100),
	postal VARCHAR(10),
	address_line VARCHAR(128),
	contact VARCHAR(15),
	favorite BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createProductTable = `
CREATE TABLE IF NOT EXISTS products (
	id UUID PRIMARY KEY,
	name VARCHAR(100),
	description VARCHAR(255),
	seller_id UUID NOT NULL REFERENCES users(id),
	gender VARCHAR(10),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createProductImageTable = `
CREATE TABLE IF NOT EXISTS product_image (
	id UUID PRIMARY KEY,
	image_url VARCHAR(255),
	product_id UUID NOT NULL REFERENCES products(id),
	is_primary BOOLEAN,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createColorTable = `
CREATE TABLE IF NOT EXISTS color (
	id UUID PRIMARY KEY,
	type VARCHAR(50),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createSizeTable = `
CREATE TABLE IF NOT EXISTS size (
	id UUID PRIMARY KEY,
	type VARCHAR(10),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createProductVariationTable = `
CREATE TABLE IF NOT EXISTS product_variation (
	id UUID PRIMARY KEY,
	color_id UUID REFERENCES color(id),
	size_id UUID NOT NULL REFERENCES size(id),
	product_id UUID NOT NULL REFERENCES products(id),
	variation_price DECIMAL(10,2),
	image_variation VARCHAR(255),
	quantity INT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createOrderTable = `
CREATE TYPE IF NOT EXISTS order_status AS ENUM ('pending', 'delivered', 'cancelled');
CREATE TABLE IF NOT EXISTS order (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL REFERENCES users(id),
	status order_status DEFAULT 'pending',
	total_price DECIMAL(10,2),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`

const createOrderItemTable = `
CREATE TABLE IF NOT EXISTS order_item (
	id UUID PRIMARY KEY,
	order_id UUID NOT NULL REFERENCES orders(id),
	product_variation_id UUID NOT NULL REFERENCES product_variations(id),
	quantity INT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
)`
