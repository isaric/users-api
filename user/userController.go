package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"users-api/models"
	"users-api/util"
)

func SaveSingleUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user models.UserDTO
	err := json.Unmarshal(reqBody, &user)
	util.HandleJsonError(w, err)
	util.Encode(w, models.PopulateUserToDTO(saveUser(models.PopulateDTOToUser(user))))
}
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateUserToDTO(deleteUser(util.GetId(r))))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateUserDTOList(findAllUsers()))
}

func FindSingleUser(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateUserToDTO(findUser(util.GetId(r))))
}
