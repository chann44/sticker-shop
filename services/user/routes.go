package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chann44/go-shop/services/auth"
	"github.com/chann44/go-shop/types"
	"github.com/chann44/go-shop/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

	router.HandleFunc("/user/{user_id}", h.handleGetUser).Methods("GET")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.store.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, user)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// parse and validate payload
	var paylaod types.UserRegisterPayload

	if err := utils.ParseJons(r, &paylaod); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if user exist in the databse
	_, err := h.store.GetUserByEmail(paylaod.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this email %s already exist", paylaod.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(paylaod.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(types.User{
		FirstName: paylaod.FirstName,
		LastName:  paylaod.LastName,
		Email:     paylaod.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	// create a new user in db
	utils.WriteJson(w, http.StatusCreated, nil)

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
