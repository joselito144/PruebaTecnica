package services

import (
	"fmt"
	"time"

	"pruebaTecnica/src/core/models"

	"github.com/jinzhu/gorm"
)

type PersonServices struct {
	models.Person
}

func (PersonServices) TableName() string {
	return "Persons"
}

func (person *PersonServices) FindAllPersons(db *gorm.DB) (*[]PersonServices, error) {
	var err error
	persons := []PersonServices{}
	db.Debug().Model(&PersonServices{}).Find(&persons)

	if err != nil {
		return &[]PersonServices{}, err
	}
	return &persons, err
}

func (person *PersonServices) SavePerson(db *gorm.DB) (*PersonServices, error) {

	var err error
	person.CreationDate = time.Now()
	person.Active = true
	err = db.Debug().Create(&person).Error
	if err != nil {
		return &PersonServices{}, err
	}
	return person, nil
}

func (person *PersonServices) FindPersonByID(db *gorm.DB, pid uint64) (*PersonServices, error) {
	err := db.Debug().Model(&PersonServices{}).Where("id = ?", pid).Take(&person).Error
	if err != nil {
		return &PersonServices{}, err
	}
	return person, nil
}

func (person *PersonServices) UpdatePerson(db *gorm.DB) (*PersonServices, error) {
	var err error
	err = db.Debug().Create(&person).Error
	if err != nil {
		return &PersonServices{}, err
	}

	err = db.Debug().Model(&PersonServices{}).Where("id = ?", person.ID).Updates(person).Error

	if err != nil {
		return &PersonServices{}, err
	}
	return person, nil
}

func (person *PersonServices) UnactivedPerson(db *gorm.DB, idPerson uint64) (*PersonServices, error) {
	var err error
	person, err = person.FindPersonByID(db, idPerson)

	if err != nil {
		return &PersonServices{}, err
	}

	person.Active = false

	fmt.Printf("%v", person)

	err = db.Debug().Model(&PersonServices{}).Where("id = ?", idPerson).Updates(person).Error

	if err != nil {
		return &PersonServices{}, err
	}

	return person, nil
}
