package main

import (
	"github.com/LeMonsalve/go-gorm/models"
	"github.com/LeMonsalve/go-gorm/storage"
)

func main() {
	storage.New()

	invoice := models.Invoiceheader{
		Client: "Juan Jose Monsalve",
		Invoiceitems: []models.Invoiceitem{
			{ProductID: 4},
			{ProductID: 5},
		},
	}

	storage.DB().Create(&invoice)
}
