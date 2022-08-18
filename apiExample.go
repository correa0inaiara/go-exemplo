package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // _ serve para fazer o go ignorar o import
	"github.com/labstack/echo/v4"
)

type Car struct {
	Name string `json: "name"`
	Price float64 `json: "price"`
}

var cars []Car // [] slice é como se fosse um array, que vai crescendo automaticamente

func loadCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Mercedes", Price: 200})
	cars = append(cars, Car{Name: "Porsche", Price: 200})
}

func main() {
	loadCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}