package group

import (
	"users-api/models"
	"users-api/util"
)

func saveGroup(g models.Group) models.Group {
	util.Database.Save(&g)
	return g
}

func deleteGroup(id string) models.Group {
	var group models.Group
	util.Database.Delete(&group, id)
	return group
}

func findGroup(id string) models.Group {
	var group models.Group
	util.Database.First(&group, id)
	return group
}

func findAllGroups() []models.Group {
	var groups []models.Group
	util.Database.Find(&groups)
	return groups
}