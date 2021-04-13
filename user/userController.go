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

// swagger:route GET /profile
//
// Gets profile of user
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss
//
//     Security:
//       api_key:
//       oauth: read, write
//
//     Responses:
//       default: genericError
//       200: someResponse
//       422: validationError
func ListUsers(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateUserDTOList(findAllUsers()))
}

// FindSingleUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.UserDTO
// @Router /user/{id} [get]
func FindSingleUser(w http.ResponseWriter, r *http.Request) {
	util.Encode(w, models.PopulateUserToDTO(findUser(util.GetId(r))))
}
