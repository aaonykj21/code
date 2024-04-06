package controller

import (
	"database/sql"
	"log"
	"net/http"
	"project/models"
	"github.com/gin-gonic/gin"
)


type Meats struct {
	Meats []models.Meat `json:"meats"`
}


func CreateMeat(c *gin.Context, db *sql.DB) {
	var meats Meats

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&meats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(meats.Meats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No meats provided"})
		return
	}

	var errors []error

	for _, meat := range meats.Meats {
		insertQuery := "INSERT INTO meat (meat_nameTH, meat_nameENG, meat_price, meat_stock) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, meat.Meat_nameTH, meat.Meat_nameENG, meat.Meat_price, meat.Meat_stock)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data", "details": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meats created successfully"})
}


func GetMeats(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM meat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var meats []models.Meat
	for rows.Next() {
		var meat models.Meat
		err := rows.Scan(&meat.Meat_id, &meat.Meat_nameTH, &meat.Meat_nameENG, &meat.Meat_price, &meat.Meat_stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		meats = append(meats, meat)
	}

	c.JSON(http.StatusOK, meats)
}

func GetMeat(c *gin.Context, db *sql.DB) {
	meatID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var meat models.Meat
	err := db.QueryRow("SELECT * FROM meat WHERE meat_id = ?", meatID).Scan(&meat.Meat_id, &meat.Meat_nameTH, &meat.Meat_nameENG, &meat.Meat_price, &meat.Meat_stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, meat)
}

func UpdateMeat(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var meat models.Meat

	if err := c.ShouldBindJSON(&meat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE meat SET meat_nameTH=?, meat_nameENG=?, meat_price=?, meat_stock=? WHERE meat_id=?"
	_, err := db.Exec(updateQuery, meat.Meat_nameTH, meat.Meat_nameENG, meat.Meat_price, meat.Meat_stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meat updated successfully"})
}

func DeleteMeat(c *gin.Context, db *sql.DB) {
	meatID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM topping WHERE Topping_ID = ?"
	_, err := db.Exec(deleteQuery, meatID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	resetQuery := "ALTER TABLE meat AUTO_INCREMENT = 1"
	_, err = db.Exec(resetQuery)
	if err != nil {
		log.Printf("Error resetting auto-increment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error resetting auto-increment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meat deleted successfully"})
}