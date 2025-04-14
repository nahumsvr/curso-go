package services

import "fmt"

type User struct {
	ID     int    `json:"id"`
	Nombre string `json:"name"`
	Email  string `json:"email"`
}

type UserService struct {
	users []User
}

func NewUserService() *UserService {
	return &UserService{
		users: []User{},
	}
}

func (s *UserService) GetAll() []User {
	return s.users
}

func (s *UserService) Create(user User) User {
	user.ID = len(s.users) + 1 // Assign a new ID
	s.users = append(s.users, user)
	return user
}

func (s *UserService) Update(id int, user User) (User, error) {
	for i, u := range s.users {
		if u.ID == id {
			s.users[i].Nombre = user.Nombre
			s.users[i].Email = user.Email
			return s.users[i], nil
		}
	}
	return User{}, fmt.Errorf("user with ID %d not found", id)
}

func (s *UserService) Delete(id int) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}
