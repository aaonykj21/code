package controller

import (
	"database/sql"
	"log"
	"net/http"
	"project/models"
	"github.com/gin-gonic/gin"
)


type Toppings struct {
	Toppings []models.Topping `json:"toppings"`
}


func CreateTopping(c *gin.Context, db *sql.DB) {
	var toppings Toppings

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&toppings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(toppings.Toppings) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No toppings provided"})
		return
	}

	var errors []error

	for _, topping := range toppings.Toppings {
		insertQuery := "INSERT INTO topping (Topping_nameTH, Topping_nameENG, Topping_price, Topping_stock) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, topping.Topping_nameTH, topping.Topping_nameENG, topping.Topping_price, topping.Topping_stock)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data", "details": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toppings created successfully"})
}


func GetToppings(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM topping")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var toppings []models.Topping
	for rows.Next() {
		var topping models.Topping
		err := rows.Scan(&topping.Topping_id, &topping.Topping_nameTH, &topping.Topping_nameENG, &topping.Topping_price, &topping.Topping_stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		toppings = append(toppings, topping)
	}

	c.JSON(http.StatusOK, toppings)
}

func GetTopping(c *gin.Context, db *sql.DB) {
	toppingID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var topping models.Topping
	err := db.QueryRow("SELECT * FROM topping WHERE Topping_ID = ?", toppingID).Scan(&topping.Topping_id, &topping.Topping_nameTH, &topping.Topping_nameENG, &topping.Topping_price, &topping.Topping_stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, topping)
}

func UpdateTopping(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var topping models.Topping

	if err := c.ShouldBindJSON(&topping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE topping SET Topping_name_th=?, Topping_name_en=?, Topping_price=?, Topping_Stock=? WHERE Topping_ID=?"
	_, err := db.Exec(updateQuery, topping.Topping_nameTH, topping.Topping_nameENG, topping.Topping_price, topping.Topping_stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Topping updated successfully"})
}

func DeleteTopping(c *gin.Context, db *sql.DB) {
	toppingID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM topping WHERE Topping_ID = ?"
	_, err := db.Exec(deleteQuery, toppingID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	resetQuery := "ALTER TABLE topping AUTO_INCREMENT = 1"
	_, err = db.Exec(resetQuery)
	if err != nil {
		log.Printf("Error resetting auto-increment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error resetting auto-increment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Topping deleted successfully"})
}