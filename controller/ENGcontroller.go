package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"project/models"
	"github.com/gin-gonic/gin"
)

// CreateOrderDetail_th handles the creation of a new order detail in Thai
func CreateOrderDetail_eng(c *gin.Context, db *sql.DB) {
	var orderDetail_eng models.Order_detailENG

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&orderDetail_eng); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert array of topping to a comma-separated string
	meats := strings.Join(orderDetail_eng.Meat_nameENG, ",")
	vegs := strings.Join(orderDetail_eng.Veg_nameENG, ",")
	sauces := strings.Join(orderDetail_eng.Sauce_nameENG, ",")
	toppings := strings.Join(orderDetail_eng.Topping_nameENG, ",")

	// Insert order detail into the database
	insertQuery := "INSERT INTO order_detail (order_id, bread_nameENG, meat_nameENG, veg_nameENG, sauce_nameENG, topping_nameENG) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, orderDetail_eng.Order_id, orderDetail_eng.Bread_nameENG, meats, vegs, sauces, toppings)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	// Decrease the stock of size
	_, err = db.Exec("UPDATE bread SET bread_stock = bread_stock - 1 WHERE bread_nameENG = ?", orderDetail_eng.Bread_nameENG)
	if err != nil {
		log.Printf("Error updating bread stock: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating bread stock"})
		return
	}

	for _, m := range orderDetail_eng.Meat_nameENG {
		_, err = db.Exec("UPDATE meat SET meat_stock = meat_stock - 1 WHERE meat_nameENG = ?", m)
		if err != nil {
			log.Printf("Error updating meat stock: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating meat stock"})
			return
		}
	}

	for _, v := range orderDetail_eng.Veg_nameENG {
		_, err = db.Exec("UPDATE vegetable SET veg_stock = veg_stock - 1 WHERE veg_nameENG = ?", v)
		if err != nil {
			log.Printf("Error updating veg stock: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating veg stock"})
			return
		}
	}

	for _, s := range orderDetail_eng.Sauce_nameENG {
		_, err = db.Exec("UPDATE sauce SET sauce_stock = sauce_stock - 1 WHERE sauce_nameENG = ?", s)
		if err != nil {
			log.Printf("Error updating sauce stock: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating sauce stock"})
			return
		}
	}

	for _, t := range orderDetail_eng.Topping_nameENG {
		_, err = db.Exec("UPDATE topping SET topping_stock = topping_stock - 1 WHERE topping_nameENG = ?", t)
		if err != nil {
			log.Printf("Error updating topping stock: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating topping stock"})
			return
		}
	}


	c.JSON(http.StatusOK, gin.H{"message": "Order detail created successfully"})
}

// GetOrderDetail_th retrieves an order detail in Thai by its ID
func GetOrderDetail_eng(c *gin.Context, db *sql.DB) {
	detailID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var orderDetail models.Order_detailENG
	var meats string
	var vegs string
	var sauces string
	var toppings string
	err := db.QueryRow("SELECT bread_nameENG, meat_nameENG, veg_nameENG, sauce_nameENG, topping_nameENG FROM order_detail WHERE order_id = ?", detailID).Scan(&orderDetail.Bread_nameENG, &meats, &vegs, &sauces,&toppings)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
    meatSlice := strings.Split(meats, ",")
	vegSlice := strings.Split(vegs, ",")
	sauceSlice := strings.Split(sauces, ",")
	toppingSlice := strings.Split(toppings, ",")

	// Calculate the total price
	totalPrice, err := calculateTotalPrice_eng(db, orderDetail.Bread_nameENG, meatSlice , vegSlice, sauceSlice, toppingSlice)
	if err != nil {
		log.Printf("Error calculating total price: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calculating total price"})
		return
	}

	orderDetail.Sum_price = totalPrice

	// Update the Sum_Price in the table
	updateQuery := "UPDATE order_detail SET sum_price = ? WHERE order_id = ?"
	_, err = db.Exec(updateQuery, orderDetail.Sum_price, detailID)
	if err != nil {
		log.Printf("Error updating Sum_Price: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating sum_price"})
		return
	}

	c.JSON(http.StatusOK, orderDetail)
}

// calculateTotalPrice_th calculates the total price of an order in Thai
func calculateTotalPrice_eng(db *sql.DB, bread string, meat []string, veg []string, sauce []string,topping []string) (int, error) {
	var breadPrice int
	var meatPrice int = 0
	var vegPrice int = 0
	var saucePrice int = 0
	var toppingPrice int = 0 

	// Retrieve the price of each component
	err := db.QueryRow("SELECT bread_price FROM bread WHERE bread_nameENG = ?", bread).Scan(&breadPrice)
	if err != nil {
		return 0, err
	}

	for _, m := range meat {
		var price int
		err = db.QueryRow("SELECT meat_price FROM meat WHERE meat_nameENG = ?", m).Scan(&price)
		if err != nil {
			return 0, err
		}
		meatPrice += price
	}

	for _, v := range veg {
		var price int
		err = db.QueryRow("SELECT veg_price FROM veg WHERE veg_nameENG = ?", v).Scan(&price)
		if err != nil {
			return 0, err
		}
		vegPrice += price
	}

	for _, s := range sauce {
		var price int
		err = db.QueryRow("SELECT sauce_price FROM sacue WHERE sauce_nameENG = ?", s).Scan(&price)
		if err != nil {
			return 0, err
		}
		saucePrice += price
	}

	for _, t := range topping {
		var price int
		err = db.QueryRow("SELECT topping_price FROM topping WHERE topping_nameENG = ?", t).Scan(&price)
		if err != nil {
			return 0, err
		}
		toppingPrice += price
	}
	// Calculate the total price
	totalPrice := breadPrice + meatPrice + vegPrice + saucePrice + toppingPrice

	return totalPrice, nil
}

// GetOrderDetails_th retrieves all order details in Thai
func GetOrderDetails_eng(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	rows, err := db.Query("SELECT * FROM order_detail")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var orderDetails []models.Order_detailENG
	for rows.Next() {
		var orderDetail models.Order_detailENG
		err := rows.Scan(&orderDetail.Order_id, &orderDetail.Bread_nameENG, &orderDetail.Meat_nameENG, &orderDetail.Veg_nameENG, &orderDetail.Sauce_nameENG, &orderDetail.Topping_nameENG)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	c.JSON(http.StatusOK, orderDetails)
}

// UpdateOrderDetail_th updates an existing order detail in Thai
func UpdateOrderDetail_eng(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var orderDetail models.Order_detailENG

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the database
	updateQuery := "UPDATE order_detail SET bread_nameENG=?, meat_nameENG=?, veg_nameENG=?, sauce_nameENG=?, topping_nameENG, sum_price=? WHERE order_id=?"
	_, err := db.Exec(updateQuery, orderDetail.Bread_nameENG, orderDetail.Meat_nameENG, orderDetail.Veg_nameENG,orderDetail.Sauce_nameENG,orderDetail.Topping_nameENG,orderDetail.Sum_price, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail updated successfully"})
}

// DeleteOrderDetail_th deletes an existing order detail in Thai
func DeleteOrderDetail_eng(c *gin.Context, db *sql.DB) {
	detailID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM order_detail WHERE order_id = ?"
	_, err := db.Exec(deleteQuery, detailID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order detail deleted successfully"})
}