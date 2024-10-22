package repositories

import "github.com/eskimoburger/cursor-pagination/entities"

type UserRepository interface {
	FetchUsers(cursor entities.Cursor, limit int) ([]entities.User, entities.Cursor, error)
}

type InMemoryUserRepository struct {
	users []entities.User
}

func NewInMemoryUserRepository(
	users []entities.User,
) UserRepository {
	return &InMemoryUserRepository{
		users,
	}

}

func (r *InMemoryUserRepository) FetchUsers(cursor entities.Cursor, limit int) ([]entities.User, entities.Cursor, error) {
	var result []entities.User
	var lastCursor entities.Cursor

	for _, user := range r.users {
		if (user.CreatedAt.After(cursor.CreatedAt)) || (user.CreatedAt.Equal(cursor.CreatedAt) && user.ID > cursor.ID) {
			result = append(result, user)
			if len(result) >= limit {
				lastCursor = entities.Cursor{CreatedAt: user.CreatedAt, ID: user.ID}
				break
			}
		}
	}

	return result, lastCursor, nil

}
