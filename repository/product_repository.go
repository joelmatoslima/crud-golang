package repository

import (
	"database/sql"
	"fmt"
	"lsport/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}

}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	const query = "SELECT * FROM product"
	var rows, err = pr.connection.Query(query)

	if err != nil {
		fmt.Println("Erro ao tentar obter todos os productos => ", err)

	}

	var productList []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			fmt.Println("Erro ao tentar PARSSEAR O PRODUTO => ", err)
			return []model.Product{}, err
		}

		productList = append(productList, product)

	}

	rows.Close()
	return productList, nil

}

func (pq *ProductRepository) CreateProduct(product model.Product) (uint, error) {
	var id uint
	query, err := pq.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println("Ao no pq.connection.Prepare")
		return 0, err
	}

	query.QueryRow(product.Name, product.Price).Scan(&id)

	query.Close()
	return id, nil

}

func (pq *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pq.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println("Ao no pq.connection.Prepare =>>>", err)
		return nil, err
	}

	var product model.Product
	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		fmt.Println("err: =>> ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil

}
