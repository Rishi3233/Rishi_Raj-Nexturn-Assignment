package handlers

import (
	"database/sql"

	"log"
	"net/http"
	"ecommerce_inventory_service/database"
	"ecommerce_inventory_service/models"
	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind the JSON request to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert product into the database
	stmt, err := database.DB.Prepare(`
		INSERT INTO products (name, description, price, stock, category_id)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductByID handles retrieving a product by its ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Query the product by ID
	row := database.DB.QueryRow("SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?", id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch product"})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles updating an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Bind the JSON request to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare update statement
	stmt, err := database.DB.Prepare(`
		UPDATE products
		SET name = ?, description = ?, price = ?, stock = ?, category_id = ?
		WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct handles deleting a product by ID
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// Prepare delete statement
	stmt, err := database.DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
