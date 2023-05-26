package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
  )

  type Employee struct {
	ID				uint
	Name         	string
	Address      	string
	Age          	uint8
	Birthdate     	string
	Level			string
	Id_department	uint
  }

  type Department struct {
	ID                	uint
	Name				string
  }

func main() {
	dsn := "host=localhost user=postgres password=orcl dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var departements = []Department{
		{Name: "Departement 1"},
		{Name: "Departement 2"},
		{Name: "Departement 3"},
		{Name: "Departement 4"},
		{Name: "Departement 5"},
	}
	db.Create(&departements)

	var employees = []Employee {
		{
            Name:             "John",
            Address:          "Jl. Tamini",
            Age:               26,
            Birthdate:        "1995-01-01",
            Level:            "Staff",
            Id_department:    1,
        },
        {
            Name:             "Jane",
            Address:          "Jl. Tamini",
            Age:               25,
            Birthdate:        "1996-01-01",
            Level:            "Staff",
            Id_department:    2,
        },
		{
            Name:             "Jake",
            Address:          "Jl. Tamini",
            Age:               24,
            Birthdate:        "1997-01-01",
            Level:            "Staff",
            Id_department:    3,
        },
        {
            Name:             "Jess",
            Address:          "Jl. Tamini",
            Age:               23,
            Birthdate:         "1998-01-01",
            Level:            "Staff",
            Id_department:    4,
        },
		{
            Name:             "Jett",
            Address:          "Jl. Tamini",
            Age:               22,
            Birthdate:        "1999-01-01",
            Level:            "Staff",
            Id_department:    5,
        },
	}

	db.Create(&employees)
	fmt.Println("Insertion Success")
}