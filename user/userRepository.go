package user

import (
	"users-api/models"
	"users-api/util"
)

func saveUser(u models.User) models.User {
	util.Database.Save(&u)
	return u
}

func deleteUser(id string) (models.User, bool) {
	var user models.User
	err := util.Database.Delete(&user, id).RowsAffected
	deleted := err != 0
	return user, deleted
}

func findUser(id string) (models.User, bool) {
	var user models.User
	err := util.Database.First(&user, id).Error
	result := err == nil
	return user, result
}

func findAllUsers() []models.User {
	var users []models.User
	util.Database.Find(&users)
	return users
}