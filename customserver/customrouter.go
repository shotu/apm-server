package customserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CustomRouter struct {
}

// Custom routes registeration
func NewCustomRouter(router *mux.Router) *mux.Router {
	customRouter := CustomRouter{}
	router.HandleFunc("/", customRouter.createUserHandler).Methods("GET")
	router.HandleFunc("/jobs", customRouter.createKafkaPostHandler).Methods("POST")
	return router
}

func (ur *CustomRouter) createKafkaPostHandler(w http.ResponseWriter, r *http.Request) {

	jobsPostHandler(w, r)

}

func (ur *CustomRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// user, err := decodeUser(r)
	// if err != nil {
	// 	Error(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }

	// err = ur.userService.Create(&user)
	// if err != nil {
	// 	Error(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	details := make(map[string]string)
	details["elastic_service_status"] = "OK"
	Json(w, http.StatusOK, details)
}

// func

// func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Println(vars)
// 	username := vars["username"]

// 	user, err := ur.userService.GetByUsername(username)
// 	if err != nil {
// 		Error(w, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	Json(w, http.StatusOK, user)
// }

// func decodeUser(ur *http.Request) (root.User, error) {
// 	var u root.User
// 	if r.Body == nil {
// 		return u, errors.New("no request body")
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&u)
// 	return u, err
// }
