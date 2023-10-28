package qso

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type QsoRepository struct {
	database *gorm.DB
}

func (repository *QsoRepository) FindAll() []Qso {
	var qsos []Qso
	repository.database.Find(&qsos)
	return qsos
}

func (repository *Qsorepository) Find(id int) (Qso, error) {
	var qso Qso
	err := repository.database.Find(&qso, id).Error
	if qso.Callsign == "" {
		err = errors.New("Qso not found")
	}
	return qso, err
}

func (repository *QsoRepository) Create(qso Qso) (Qso, error) {
	err := repository.database.Create(&qso).Error
	if err != nil {
		return qso, err
	}

	return qso, nil
}
func (repository *QsoRepository) Save(user Qso) (Qso, error) {
	err := repository.database.Save(user).Error
	return user, err
}

func (repository *QsoRepository) Delete(id int) int64 {
	count := repository.database.Delete(&Qso{}, id).RowsAffected
	return count
}

func NewQsoRepository(database *gorm.DB) *QsoRepository {
	return &QsoRepository{
		database: database,
	}
}
