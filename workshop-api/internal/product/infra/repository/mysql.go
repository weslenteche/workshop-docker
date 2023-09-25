package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weslenteche/workshop-api/config"
	"github.com/weslenteche/workshop-api/internal/product/entity"
	"log"
)

type ProductMysqlRepository struct {
	db *sql.DB
}

func NewProductMysqlRepository() entity.ProductRepository {
	envs := config.LoadEnvs()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		envs.DBUser, envs.DBPassword, envs.DBHost, envs.DBPort, envs.DBDatabase,
	)
	db, err := sql.Open(
		"mysql",
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &ProductMysqlRepository{
		db: db,
	}
}

func (r *ProductMysqlRepository) FindAll(ctx context.Context) []entity.Product {
	rows, err := r.db.Query("SELECT external_id, name, description, price FROM products")
	if err != nil {
		log.Print(err)
		return make([]entity.Product, 0)
	}
	defer rows.Close()
	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		_ = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		products = append(products, product)
	}
	return products
}

func (r *ProductMysqlRepository) Save(ctx context.Context, product *entity.Product) {
	stmt, err := r.db.Prepare("INSERT INTO products (external_id, name, description, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}

	_, err = stmt.Exec(
		product.ID,
		product.Name,
		product.Description,
		product.Price,
	)
	if err != nil {
		log.Print(err)
	}
}
