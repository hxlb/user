package repository

import (
	"github.com/hxlb/pkg/db/xorm"
	"github.com/hxlb/user/model"
)

type Repository interface {
	Find(id int32) (*model.User, error)
	Create(*model.User) error
	Update(*model.User, int64) (*model.User, error)
	FindByField(string, string, string) (*model.User, error)
}

type User struct {
}

func (repo *User) Find(id uint32) (*model.User, error) {
	var user model.User
	orm := xorm.MustDB()
	_, err := orm.Id(id).Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (repo *User) Create(user *model.User) error {
	if _, err := xorm.MustDB().InsertOne(user); err != nil {
		return err
	}
	return nil
}

func (repo *User) Update(user *model.User) (*model.User, error) {
	if _, err := xorm.MustDB().Update(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *User) FindByField(key string, value string, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	user := &model.User{}
	if _, err := xorm.MustDB().Where(key+" = ?", value).Cols(fields).Get(user); err != nil {
		return nil, err
	}
	return user, nil
}
