package controller

import (
	"database/sql"
	"log"
	"net/http"
	"project/models"
	"github.com/gin-gonic/gin"
)

// var sizes []models.Size

func CreateBread(c *gin.Context, db *sql.DB) {
	var bread models.Bread

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&bread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	insertQuery := "INSERT INTO bread (bread_nameTH, bread_nameENG, bread_price, bread_stock) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, bread.Bread_nameTH, bread.Bread_nameENG, bread.Bread_price, bread.Bread_stock)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bread created successfully"})
}

func GetBreads(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM bread")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var breads []models.Bread
	for rows.Next() {
		var bread models.Bread
		err := rows.Scan(&bread.Bread_id, &bread.Bread_nameTH, &bread.Bread_nameENG, &bread.Bread_price, &bread.Bread_stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		breads = append(breads, bread)
	}

	c.JSON(http.StatusOK, breads)
}

func GetBread(c *gin.Context, db *sql.DB) {
	breadID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var bread models.Bread
	err := db.QueryRow("SELECT * FROM bread WHERE bread_id = ?", breadID).Scan(&bread.Bread_id, &bread.Bread_nameTH, &bread.Bread_nameENG, &bread.Bread_price, &bread.Bread_stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, bread)
}

func UpdateBread(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var bread models.Bread

	if err := c.ShouldBindJSON(&bread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE bread SET bread_nameTH=?, bread_nameENG=?, bread_price=?, bread_stock=? WHERE bread_id=?"
	_, err := db.Exec(updateQuery, bread.Bread_nameTH, bread.Bread_nameENG, bread.Bread_price, bread.Bread_stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bread updated successfully"})
}

func DeleteBread(c *gin.Context, db *sql.DB) {
	breadID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM bread WHERE bread_id= ?"
	_, err := db.Exec(deleteQuery, breadID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bread deleted successfully"})
}
