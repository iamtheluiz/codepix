package repository

import (
	"fmt"

	"github.com/iamtheluiz/codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

// Register => Repository Function
func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

// Save => Repository Function
func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

// Find => Repository Function
func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &transaction, nil
}
