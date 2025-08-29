package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)
func task1() {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("Establishing connection with localhost:8888")
	http.ListenAndServe(":8888", nil)
}
func echoHandler(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Error while reading Body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "Got: %s", body)
	fmt.Println(string(body))
}
////////////////////
func task2() {
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8888", nil)
	
}

type Person struct {
	Name string `json:"Name"`
}
func greet(writer http.ResponseWriter, request *http.Request) {
	var p Person
	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		log.Fatal("error while parsing JSON", err)
	}
	response := fmt.Sprintf("hello ,%s", p.Name)
	json.NewEncoder(writer).Encode(&Response{Message: response})
}
///////////////////////////
func task3() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run(":8888")
}
func helloHandler(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	response := fmt.Sprintf("hello, %s!", name)
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})

}
type Response struct {
	Message string `json:"Message"`
}

///////////////////////////
func task4() {
	r := gin.Default()
	r.POST("/divide", divideHandler)
	r.Run(":8888")

}
func divideHandler(c *gin.Context) {
	var obj division
	if err := c.BindJSON(&obj); err != nil {
		log.Println(Red, "error:", err, Reset)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		log.Println(Blue + "info:successfully parsed json" + Reset)
	}
	res, err := divide(obj.A, obj.B)
	if err != nil {
		log.Println(Red, err, Reset)
		c.JSON(400, gin.H{
			"error": "zero division",
		})
	} else {
		c.JSON(200, gin.H{
			"result": res,
		})
		log.Printf(Green+"OK:operation is complete. result:%v"+Reset, res)
	}

}
func divide(x int, y int) (float64, error) {
	result := float64(x) / float64(y)
	if y == 0 {
		return 0, errors.New("error:division by zero")
	} else {
		return result, nil
	}

}
type division struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	// task1()
	// task2()
	// task3()
	task4()
}
