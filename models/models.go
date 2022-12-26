package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100); not null"`
	Observations *string `gorm:"type:varchar(100)"`
	Price        uint    `gorm:"not null"`
	Invoiceitems []Invoiceitem
}

type Invoiceheader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100); not null"`
	Invoiceitems []Invoiceitem
}

type Invoiceitem struct {
	gorm.Model
	InvoiceheaderID uint
	ProductID       uint
}
