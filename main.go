package main

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type Car struct {
	CarsId       string `json:"carsid"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	LicencePlate string `json:"licenceplate"`
	OwnerName    string `json:"ownername"`
	Date         string `json:"date"`
	Status       string `json:"status"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	// register route greet
	app.GET("/cars", func(ctx *gofr.Context) (interface{}, error) {
		var cars []Car

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM cars;")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var car Car
			if err := rows.Scan(&car.CarsId, &car.Make, &car.Model, &car.LicencePlate, &car.OwnerName, &car.Date, &car.Status); err != nil {
				return nil, err
			}

			cars = append(cars, car)
		}

		// return the customer
		return cars, nil
	})

	app.GET("/car/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		var cars []Car
		rows, err := ctx.DB().QueryContext(ctx, "SELECT *  FROM cars WHERE carsid=?", id)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var car Car
			if err := rows.Scan(&car.CarsId, &car.Make, &car.Model, &car.LicencePlate, &car.OwnerName, &car.Date, &car.Status); err != nil {
				return nil, err
			}
			ctx.Logger.Info(car)
			cars = append(cars, car)
		}

		return cars, nil
	})

	app.POST("/car", func(ctx *gofr.Context) (interface{}, error) {

		new_car := make(map[string]string)
		if err := ctx.Bind(&new_car); err != nil {
			return nil, err
		}

		ctx.Logger.Info("Testing Info logger")
		ctx.Logger.Info(new_car)

		queryInsert := "INSERT INTO cars (carsid, make,  model, licenceplate, ownername, date, status) VALUES (?,?,?,?,?,?,?)"

		if len(new_car) != 7 {
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}

		result, err := ctx.DB().ExecContext(ctx, queryInsert, new_car["cars_id"], new_car["make"], new_car["model"], new_car["licence_plate"], new_car["owner_name"], new_car["date"], new_car["status"])
		if err != nil {
			return result, err
		}

		return nil, err

	})

	app.PUT("/car/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		new_car := make(map[string]string)
		if err := ctx.Bind(&new_car); err != nil {
			return nil, err
		}

		if len(new_car) != 6 {
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}

		queryInsert := "UPDATE cars SET make = ?, model = ?, licenceplate = ?, ownername = ?, date = ?, status = ? WHERE carsid = ?;"

		result, err := ctx.DB().ExecContext(ctx, queryInsert, new_car["make"], new_car["model"], new_car["licence_plate"], new_car["owner_name"], new_car["date"], new_car["status"], id)
		if err != nil {
			return result, err
		}

		return nil, err
	})

	app.DELETE("/car/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		_, err := ctx.DB().ExecContext(ctx, "DELETE FROM cars WHERE carsid = ?", id)
		if err != nil {
			return nil, err
		}

		return nil, err
	})

	app.Start()
}
