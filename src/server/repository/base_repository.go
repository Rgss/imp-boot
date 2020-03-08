package repository

import (
	"log"
	"strconv"
	"github.com/Rgss/imp-boot/src/common/app"
	"github.com/Rgss/imp-boot/src/common/db"
)

type BaseRepository struct {
	ModelEntity interface{}
}

// model db
//func (this *BaseRepository) DB() (*gorm.DB) {
//	return db.MysqlInstance().DB();
//}

// create
func (this *BaseRepository) Create(entity interface{}) (interface{}, error) {
	err := db.Connection().DB().Create(entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// Insert
func (this *BaseRepository) Insert(entity interface{}) (interface{}, error) {
	return this.Create(entity)
}

// delete
func (this *BaseRepository) Delete(id int, model interface{}) (bool, error) {
	res := db.Connection().DB().Where("id = ?", strconv.Itoa(id)).Limit(1).Delete(model)
	log.Printf("Delete.RowsAffected: %v\n", res.RowsAffected)

	if res.RowsAffected >= 1 {
		return true, nil
	}

	return false, res.Error
}

// DeleteBy
func (this *BaseRepository) DeleteBy(where interface{}, model interface{}) (bool, error) {
	res := db.Connection().DB().Where(where).Delete(model)
	log.Printf("DeleteBy.RowsAffected: %v\n", res.RowsAffected)

	if res.RowsAffected >= 1 {
		return true, nil
	}

	return false, res.Error
}

// find
func (this *BaseRepository) Find(id int, entity interface{}) (interface{}, error) {
	//var entity Entity
	res := db.Connection().DB().Where("id = ?", strconv.Itoa(id)).Limit(1).Find(entity)
	if res.Error != nil {
		app.Logger().Info("Find", res.Error)
		return entity, res.Error
	}

	return entity, nil
}

func (this *BaseRepository) Model(entity interface{}) *BaseRepository {
	this.ModelEntity = entity
	return this
}

// FindBy
func (this *BaseRepository) FindBy(where interface{}, entity interface{}) (interface{}, error) {
	res := db.Connection().DB().Where(where).Find(entity)
	if res.Error != nil {
		app.Logger().Info("FindBy error:", res.Error)
		return entity, res.Error
	}

	return entity, nil
}

// FindAll
func (this *BaseRepository) FindAll(page int, size int, where interface{}, order string, entities []interface{}) ([] interface{}, error) {
	// var entity []Entity
	if order == "" {
		order = "`id` DESC"
	}

	start := (page - 1) * size
	res := db.Connection().DB().Limit(size).Offset(start).Order(order).Where(where).Find(& entities)
	if res.Error != nil {
		app.Logger().Info("FindAll error:", res.Error)
		return entities, res.Error
	}

	return entities, nil
}


// update
func (this *BaseRepository) Update(id int, entity interface{}, model interface{}) (bool, error) {
	res := db.Connection().DB().Model(model).Where("id = ?", id).Updates(entity)
	if res.Error != nil {
		app.Logger().Info("Update error:", res.Error)
		return false, res.Error
	}

	log.Printf("Update.RowsAffected: %v\n", res.RowsAffected)
	if res.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}

// update by
func (this *BaseRepository) UpdateBy(where interface{}, entity interface{}, model interface{}) (bool, error) {
	res := db.Connection().DB().Model(model).Where(where).Updates(entity)
	if res.Error != nil {
		app.Logger().Info("UpdateBy error:", res.Error)
		return false, res.Error
	}

	log.Printf("UpdateBy.RowsAffected: %v\n", res.RowsAffected)
	if res.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}

func (this *BaseRepository) Exists(id int, entity interface{}) (bool, error) {
	_, err := this.Find(id, entity)
	if err != nil {
		return true, err
	}

	return false, nil
}

func (this *BaseRepository) ExistsBy(where interface{}, entity interface{}) (bool, error) {
	_, err := this.FindBy(where, entity)
	if err != nil {
		return true, err
	}

	return false, nil
}


func (this *BaseRepository) IncreaseBy(column string, step int, where interface{}, entity interface{}) {

}

func (this *BaseRepository) DecreaseBy(column string, step int, where interface{}, entity interface{}) {

}