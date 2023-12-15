package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cafe")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func main() {
	router := gin.Default()
	ctx := context.Background()

	router.POST("/insert", func(c *gin.Context) {
		db := Connection()
		defer db.Close()

		var req struct {
			Nama   string `json:"nama"`
			Umur   int    `json:"umur"`
			Alamat string `json:"alamat"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertSql := "INSERT INTO `go-user` (nama, umur, alamat) VALUES (?, ?, ?)"
		_, err := db.ExecContext(ctx, insertSql, req.Nama, req.Umur, req.Alamat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data Successfully Inserted"})
	})

	router.GET("/select", func(c *gin.Context) {
		db := Connection()
		defer db.Close()

		selectSql := "SELECT * FROM `go-user`"
		rows, err := db.QueryContext(ctx, selectSql)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var data []gin.H
		for rows.Next() {
			var (
				id     int
				nama   string
				umur   int
				alamat string
			)

			err = rows.Scan(&id, &nama, &umur, &alamat)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			data = append(data, gin.H{"id": id, "nama": nama, "umur": umur, "alamat": alamat})
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	router.GET("/select/:id", func(c *gin.Context) {
		db := Connection()
		defer db.Close()

		id := c.Param("id")
		selectSql := "SELECT * FROM `go-user` WHERE id = ?"

		row := db.QueryRow(selectSql, id)
		var (
			resultID int
			nama     string
			umur     int
			alamat   string
		)

		if err := row.Scan(&resultID, &nama, &umur, &alamat); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result := gin.H{
			"id":     resultID,
			"nama":   nama,
			"umur":   umur,
			"alamat": alamat,
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.PUT("/update/:id", func(c *gin.Context) {
		db := Connection()
		defer db.Close()

		id := c.Param("id")
		var req struct {
			Umur int `json:"umur"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateSql := "UPDATE `go-user` SET umur = ? WHERE id = ?"
		_, err := db.ExecContext(ctx, updateSql, req.Umur, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data Successfully Updated"})
	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		db := Connection()
		defer db.Close()

		id := c.Param("id")

		deleteSql := "DELETE FROM `go-user` WHERE id = ?"
		_, err := db.ExecContext(ctx, deleteSql, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data Successfully Deleted"})
	})

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
