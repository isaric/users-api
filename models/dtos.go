package models

import (
	"gorm.io/gorm"
)

type UserDTO struct {
    ID 	     uint   `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Group GroupDTO  `json:"GroupDTO"`
}

type GroupDTO struct {
	ID uint `json:"ID"`
	Name string `json:"Name"`
}

func PopulateDTOToUser(dto UserDTO) User {
	return User {
		Model: gorm.Model{ID: dto.ID},
		Name: dto.Name,
		Email: dto.Email,
		Password: dto.Password,
		Group: PopulateDTOToGroup(dto.Group),
	}
}

func PopulateGroupDTOList(groups []Group) []GroupDTO {
	var result []GroupDTO = make([]GroupDTO, 0)
	for i := range groups {
		var current = PopulateGroupToDTO(groups[i])
		result = append(result, current)
	}
	return result
}

func PopulateUserDTOList(users []User) []UserDTO {
	var result []UserDTO = make([]UserDTO, 0)
	for i := range users {
		var current = PopulateUserToDTO(users[i])
		result = append(result, current)
	}
	return result
}

func PopulateUserToDTO(u User) UserDTO {
	return UserDTO {
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Group:  PopulateGroupToDTO(u.Group),
	}
}

func PopulateDTOToGroup(dto GroupDTO) Group {
	return Group {
		Model: gorm.Model{ID: dto.ID},
		Name: dto.Name,
	}
}

func PopulateGroupToDTO(g Group) GroupDTO {
	return GroupDTO {
		ID: g.ID,
		Name: g.Name,
	}
}
