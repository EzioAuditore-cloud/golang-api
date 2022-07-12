package models

import (
	"fmt"
	orm "project/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) FindUser(username string) (u User, err error) {
	if err = orm.Db.Where("username=?", username).First(&u).Error; err != nil {
		return
	}
	return
}

func GetUsers(username string) (users []User, err error) {
	if err = orm.Db.Where("username=?", username).First(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) Users() (users []User, err error) {
	if err = orm.Db.Find(&users).Error; err != nil {
		fmt.Println("find error ")
		return
	}
	return
}

func (user User) Insert() (id int64, err error) {
	result := orm.Db.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Destroy(id int64) (Result User, err error) {
	if err = orm.Db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Db.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

func (user *User) Update(id int64) (updateUser User, err error) {

	orm.Db.Where("id = ?", id).First(&updateUser)
	fmt.Println("update User", updateUser.ID, updateUser.Username)
	fmt.Println("new data", user.Username, user.Password, user.ID)

	if err = orm.Db.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}
