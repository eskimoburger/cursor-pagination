package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/eskimoburger/cursor-pagination/entities"
	_userService "github.com/eskimoburger/cursor-pagination/services"
)

type UserController interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	userService _userService.UserService
}

func NewUserController(
	userService _userService.UserService,
) UserController {
	return &userControllerImpl{
		userService,
	}
}

func (c *userControllerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	limit := 10
	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		limit, _ = strconv.Atoi(limitParam)
	}

	encodedCursor := r.URL.Query().Get("cursor")
	var cursor entities.Cursor
	var err error
	if encodedCursor != "" {
		cursor, err = c.userService.DecodeCursor(encodedCursor)
		if err != nil {
			http.Error(w, "Invalid cursor", http.StatusBadRequest)
			return
		}
	} else {
		cursor = entities.Cursor{CreatedAt: time.Time{}, ID: 0}
	}

	paginatedUsers, nextCursor, err := c.userService.FetchUsers(cursor, limit)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	response := struct {
		Users      []entities.User `json:"users"`
		NextCursor string          `json:"next_cursor"`
	}{
		Users:      paginatedUsers,
		NextCursor: c.userService.EncodeCursor(nextCursor),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
