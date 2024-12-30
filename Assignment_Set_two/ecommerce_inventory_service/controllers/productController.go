package controllers

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
	var newProduct models.Product

	// Bind the incoming JSON payload to the product struct
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Insert the product into the database
	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Stock, newProduct.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting product"})
		log.Println(err)
		return
	}

	// Return the created product
	c.JSON(http.StatusCreated, newProduct)
}

// GetProductByID retrieves a product by its ID
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving product"})
		}
		return
	}

	// Return the product found
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates a product's details
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct models.Product

	// Bind the incoming JSON payload to the product struct
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Prepare and execute the update query
	query := `UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? WHERE id = ?`
	_, err := database.DB.Exec(query, updatedProduct.Name, updatedProduct.Description, updatedProduct.Price, updatedProduct.Stock, updatedProduct.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct deletes a product from the inventory
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM products WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting product"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
