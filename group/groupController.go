package group

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"users-api/models"
	"users-api/util"
)

func SaveSingleGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var group models.GroupDTO
	err := json.Unmarshal(reqBody, &group)
	util.HandleJsonError(w, err)

	util.Encode(w, models.PopulateGroupToDTO(saveGroup(models.PopulateDTOToGroup(group))))
}

func RemoveGroup(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateGroupToDTO(deleteGroup(util.GetId(r))))
}

func ListGroups(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateGroupDTOList(findAllGroups()))
}

func FindSingleGroup(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateGroupToDTO(findGroup(util.GetId(r))))
}