package main

import (
	"fmt"
	"log"

	"github.com/onizukazaza/tar-ecom-api/config"
	"github.com/onizukazaza/tar-ecom-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)


	userMigrate(db)
	addressMigrate(db)
	productMigrate(db)
	productImageMigrate(db)
	productVariationMigrate(db)
	productVariationOptionMigrate(db)
	stockMigrate(db)
	cartMigrate(db)
	cartItemMigrate(db)
}

func userMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(128) NOT NULL,
		email VARCHAR(128) UNIQUE NOT NULL,
		role VARCHAR(20) NOT NULL DEFAULT 'buyer',
		profile_image VARCHAR(256) DEFAULT 'placeholder',
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "users")
}

func addressMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS addresses (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		recipient_name VARCHAR(128) NOT NULL,
		postal_address VARCHAR(256) NOT NULL,
		house_number VARCHAR(128) NOT NULL,
		contact_number VARCHAR(128) NOT NULL,
		favorite BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "addresses")
}

func productMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		seller_id INT NOT NULL REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE,
		product_name VARCHAR(128) UNIQUE NOT NULL,
		description VARCHAR(128) NOT NULL,
		is_archive BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "products")
}

func productImageMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS product_images (
		id SERIAL PRIMARY KEY,
		product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
		image_url VARCHAR(256) NOT NULL,
		is_primary BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "product_images")
}

func productVariationMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS product_variations (
		id SERIAL PRIMARY KEY,
		product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
		variation_value VARCHAR(128) NOT NULL,
		image_variation VARCHAR(128) UNIQUE NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "product_variations")
}

func productVariationOptionMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS product_variation_options (
		id SERIAL PRIMARY KEY,
		product_variation_id INT NOT NULL REFERENCES product_variations(id) ON DELETE CASCADE ON UPDATE CASCADE,
		option_value VARCHAR(128) NOT NULL
	)`
	executeMigration(db, query, "product_variation_options")
}

func stockMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS stock (
		id SERIAL PRIMARY KEY,
		product_variation_option_id INT NOT NULL REFERENCES product_variation_options(id) ON DELETE CASCADE ON UPDATE CASCADE,
		quantity INT NOT NULL
	)`
	executeMigration(db, query, "stock")
}

func cartMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS carts (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		total_price NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`
	executeMigration(db, query, "carts")
}

func cartItemMigrate(db databases.Database) {
	query := `
	CREATE TABLE IF NOT EXISTS cart_items (
		id SERIAL PRIMARY KEY,
		cart_id INT NOT NULL REFERENCES carts(id) ON DELETE CASCADE ON UPDATE CASCADE,
		product_variation_option_id INT NOT NULL REFERENCES product_variation_options(id) ON DELETE CASCADE ON UPDATE CASCADE,
		total_price NUMERIC(10, 2) NOT NULL
	)`
	executeMigration(db, query, "cart_items")
}


func executeMigration(db databases.Database, query, tableName string) {
	_, err := db.Connect().Exec(query)
	if err != nil {
		log.Fatalf("Failed to migrate table %s: %v", tableName, err)
	} else {
		fmt.Printf("Table %s migrated successfully\n", tableName)
	}
}
