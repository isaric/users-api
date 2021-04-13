package group

import (
	"users-api/models"
	"users-api/util"
)

func saveGroup(g models.Group) models.Group {
	util.Database.Save(&g)
	return g
}

func deleteGroup(id string) (models.Group, bool) {
	var group models.Group
	result := util.Database.Delete(&group, id).RowsAffected
	deleted := result != 0
	return group,deleted
}

func findGroup(id string) (models.Group, bool) {
	var group models.Group
	err := util.Database.First(&group, id).Error
	found := err == nil
	return group, found
}

func findAllGroups() []models.Group {
	var groups []models.Group
	util.Database.Find(&groups)
	return groups
}