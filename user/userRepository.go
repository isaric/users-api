package user

import (
	"users-api/models"
	"users-api/util"
)

func saveUser(u models.User) models.User {
	util.Database.Save(&u)
	return u
}

func deleteUser(id string) models.User {
	var user models.User
	util.Database.Delete(&user, id)
	return user
}

func findUser(id string) models.User {
	var user models.User
	util.Database.First(&user, id)
	return user
}

func findAllUsers() []models.User {
	var users []models.User
	util.Database.Find(&users)
	return users
}