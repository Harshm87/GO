package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type data struct {
	ID        string `json:"id"`
	Subject   string `json:"sub"`
	Completed bool   `json:"completed"`
}

var datas = []data{
	{ID: "1", Subject: "Maths", Completed: false},
	{ID: "2", Subject: "English", Completed: false},
	{ID: "3", Subject: "Science", Completed: false},
}

func getdatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)
}

func adddata(c *gin.Context) {
	var newdata data

	if err := c.BindJSON(&newdata); err != nil {
		return
	}
	datas = append(datas, newdata)

	c.IndentedJSON(http.StatusCreated, newdata)
}

func getdata(c *gin.Context) {
	id := c.Param("id")
	data, err := getdatabyId(id)

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Data not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func getdatabyId(id string)(*data, error){
	for i, d := range datas {
		if d.ID == id {
			return &datas[i], nil
		}
	}
	return nil,errors.New("data not found")
}

func main() {
	router := gin.Default()
	router.GET("/datas", getdatas)
	router.GET("/datas/:id", getdatas)
	router.POST("/datas", adddata)
	router.Run("localhost:9090")
}
