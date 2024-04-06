package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"project/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetupDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "Natcha", "127.0.0.1", "3306", "sandwich_shop")
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database!")

	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	return Db
}

func main() {
	Db = SetupDB()

	r := gin.Default()

	r.POST("/createBread", func(c *gin.Context) { controller.CreateBread(c, Db) })
	r.GET("/getBreads", func(c *gin.Context) { controller.GetBreads(c, Db) })
	r.GET("/getBread/:id", func(c *gin.Context) { controller.GetBread(c, Db) })
	r.PUT("/updateBread/:id", func(c *gin.Context) { controller.UpdateBread(c, Db) })
	r.DELETE("/deleteBread/:id", func(c *gin.Context) { controller.DeleteBread(c, Db) })

	
	r.POST("/createMeat", func(c *gin.Context) { controller.CreateMeat(c, Db) })
	r.GET("/getMeats", func(c *gin.Context) { controller.GetMeats(c, Db) })
	r.GET("/getMeat/:id", func(c *gin.Context) { controller.GetMeats(c, Db) })
	r.PUT("/updateMeat/:id", func(c *gin.Context) { controller.UpdateMeat(c, Db) })
	r.DELETE("/deleteMeat/:id", func(c *gin.Context) { controller.DeleteMeat(c, Db) })

	
	r.POST("/createVeg", func(c *gin.Context) { controller.CreateVeg(c, Db) })
	r.GET("/getVegs", func(c *gin.Context) { controller.GetVegs(c, Db) })
	r.GET("/getVeg/:id", func(c *gin.Context) { controller.GetVeg(c, Db) })
	r.PUT("/updateVeg/:id", func(c *gin.Context) { controller.UpdateVeg(c, Db) })
	r.DELETE("/deleteVeg/:id", func(c *gin.Context) { controller.DeleteVeg(c, Db) }) 


	r.POST("/createSauce", func(c *gin.Context) { controller.CreateSauce(c, Db) })
	r.GET("/getSauces", func(c *gin.Context) { controller.GetSauces(c, Db) })
	r.GET("/getSauce/:id", func(c *gin.Context) { controller.GetSauce(c, Db) })
	r.PUT("/updateSauce/:id", func(c *gin.Context) { controller.UpdateSauce(c, Db) })
	r.DELETE("/deleteSauce/:id", func(c *gin.Context) { controller.DeleteSauce(c, Db) })

	
	r.POST("/createTopping", func(c *gin.Context) { controller.CreateTopping(c, Db) })
	r.GET("/getToppings", func(c *gin.Context) { controller.GetToppings(c, Db) })
	r.GET("/getTopping/:id", func(c *gin.Context) { controller.GetTopping(c, Db) })
	r.PUT("/updateTopping/:id", func(c *gin.Context) { controller.UpdateTopping(c, Db) })
	r.DELETE("/deleteTopping/:id", func(c *gin.Context) { controller.DeleteTopping(c, Db) })


	//orderdetail_en
	r.POST("/createOrderDetail-en", func(c *gin.Context) { controller.CreateOrderDetail_eng(c, Db) })
	r.GET("/getOrderDetails-en", func(c *gin.Context) { controller.GetOrderDetails_eng(c, Db) })
	r.GET("/getOrderDetail-en/:id", func(c *gin.Context) { controller.GetOrderDetail_eng(c, Db) })
	r.PUT("/updateOrderDetail-en/:id", func(c *gin.Context) { controller.UpdateOrderDetail_eng(c, Db) })
	r.DELETE("/deleteOrderDetail-en/:id", func(c *gin.Context) { controller.DeleteOrderDetail_eng(c, Db) })

	//orderdetail_th
	r.POST("/createOrderDetail-th", func(c *gin.Context) { controller.CreateOrderDetail_th(c, Db) })
	r.GET("/getOrderDetails-th", func(c *gin.Context) { controller.GetOrderDetails_th(c, Db) })
	r.GET("/getOrderDetail-th/:id", func(c *gin.Context) { controller.GetOrderDetail_th(c, Db) })
	r.PUT("/updateOrderDetail-th/:id", func(c *gin.Context) { controller.UpdateOrderDetail_th(c, Db) })
	r.DELETE("/deleteOrderDetail-th/:id", func(c *gin.Context) { controller.DeleteOrderDetail_th(c, Db) })
	r.Run(":8080")
}
