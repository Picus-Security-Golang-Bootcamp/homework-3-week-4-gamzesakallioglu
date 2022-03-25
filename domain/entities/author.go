package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Id   int    `gorm:"primaryKey,autoIncrement"`
	Name string `gorm:"varchar(150)"`
}

type Authors []Author

// Adding Hooks
func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %v\nNAME: %s\n", a.Id, a.Name)
}

func (a *Author) BeforeDelete(db *gorm.DB) (err error) {
	fmt.Printf("%s is being deleted", a.Name)
	return nil
}
