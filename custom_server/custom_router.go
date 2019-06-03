package custom_server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CustomRouter struct {
}

func NewUserRouter(router *mux.Router) *mux.Router {
	customRouter := CustomRouter{}

	router.HandleFunc("/", customRouter.createUserHandler).Methods("GET")
	// router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")
	return router
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
	Json(w, http.StatusOK, "")
}

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
