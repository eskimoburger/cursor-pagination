package services

import (
	"encoding/base64"
	"encoding/json"

	"github.com/eskimoburger/cursor-pagination/entities"
	_userRepository "github.com/eskimoburger/cursor-pagination/repositories"
)

type UserService interface {
	FetchUsers(cursor entities.Cursor, limit int) ([]entities.User, entities.Cursor, error)
	EncodeCursor(cursor entities.Cursor) string
	DecodeCursor(encodedCursor string) (entities.Cursor, error)
}

type userServiceImp struct {
	userRepository _userRepository.UserRepository
}

func NewUserService(
	userRepository _userRepository.UserRepository,
) UserService {
	return &userServiceImp{
		userRepository: userRepository,
	}
}

func (s *userServiceImp) FetchUsers(cursor entities.Cursor, limit int) ([]entities.User, entities.Cursor, error) {
	users, nextCursor, err := s.userRepository.FetchUsers(cursor, limit)
	if err != nil {
		return nil, entities.Cursor{}, err
	}
	return users, nextCursor, nil
}

func (s *userServiceImp) EncodeCursor(cursor entities.Cursor) string {
	cursorBytes, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(cursorBytes)
}

func (s *userServiceImp) DecodeCursor(encodedCursor string) (entities.Cursor, error) {
	var cursor entities.Cursor
	cursorBytes, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return cursor, err
	}
	err = json.Unmarshal(cursorBytes, &cursor)
	return cursor, err
}
