package controller

import (
	"database/sql"
	"log"
	"net/http"
	"project/models"
	"github.com/gin-gonic/gin"
)


type Vegs struct {
	Vegs []models.Veg `json:"vegs"`
}


func CreateVeg(c *gin.Context, db *sql.DB) {
	var vegs Vegs

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&vegs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(vegs.Vegs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No vegs provided"})
		return
	}

	var errors []error

	for _, veg := range vegs.Vegs {
		insertQuery := "INSERT INTO vegetable (veg_nameTH, veg_nameENG, veg_price, veg_stock) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, veg.Veg_nameTH, veg.Veg_nameENG, veg.Veg_price, veg.Veg_stock)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data", "details": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vegs created successfully"})
}


func GetVegs(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM vegetable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var vegs []models.Veg
	for rows.Next() {
		var veg models.Veg
		err := rows.Scan(&veg.Veg_id, &veg.Veg_nameTH, &veg.Veg_nameENG, &veg.Veg_price, &veg.Veg_stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		vegs = append(vegs, veg)
	}

	c.JSON(http.StatusOK, vegs)
}

func GetVeg(c *gin.Context, db *sql.DB) {
	vegID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var veg models.Veg
	err := db.QueryRow("SELECT * FROM vegetable WHERE veg_id = ?", vegID).Scan(&veg.Veg_id, &veg.Veg_nameTH, &veg.Veg_nameENG, &veg.Veg_price, &veg.Veg_stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, veg)
}

func UpdateVeg(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var veg models.Veg

	if err := c.ShouldBindJSON(&veg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE vegetable SET veg_nameTH=?, veg_nameENG=?, veg_price=?, veg_stock=? WHERE veg_id=?"
	_, err := db.Exec(updateQuery, veg.Veg_nameTH, veg.Veg_nameENG, veg.Veg_price, veg.Veg_stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Veg updated successfully"})
}

func DeleteVeg(c *gin.Context, db *sql.DB) {
	vegID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM vegetable WHERE veg_id = ?"
	_, err := db.Exec(deleteQuery, vegID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	resetQuery := "ALTER TABLE vegetable AUTO_INCREMENT = 1"
	_, err = db.Exec(resetQuery)
	if err != nil {
		log.Printf("Error resetting auto-increment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error resetting auto-increment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vegetable deleted successfully"})
}