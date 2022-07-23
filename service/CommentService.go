package service

import (
	orm "simpleRestAPI/dao"
	"simpleRestAPI/model"

	"github.com/google/uuid"
)

func Insert(comment model.Comment) (addComment *model.Comment, err error) {
	comment.Uuid = uuid.New().String()
	result := orm.Db.Create(&comment)
	if result.Error != nil {
		err = result.Error
		return nil, err
	}
	return &comment, nil
}

func FindById(id string) (res *model.Comment, err error) {
	err = orm.Db.Where("uuid=?", id).First(&res).Error
	if err != nil {
		return
	}
	return
}
func UpdateById(id string, comment model.Comment) (res *model.Comment, err error) {
	data, err := FindById(id)
	if data == nil || err != nil {
		return
	}
	err = orm.Db.Where("uuid=?", data.Uuid).Updates(&comment).Error
	if err != nil {
		return
	}
	return
}

func DeleteById(id string) (res *model.Comment, err error) {
	data, err := FindById(id)
	if data == nil || err != nil {
		return
	}
	if err = orm.Db.Where("uuid=?", data.Uuid).Delete(&res).Error; err != nil {
		return
	}
	return
}
