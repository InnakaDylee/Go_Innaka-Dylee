package main

type User struct { 
	ID       int
	Username int
	Password int
}	

type UserService struct {
	Attribute []User
}

func (users UserService) GetAllUsers() []User {
	return users.Attribute
}

func (users UserService) GetUserById(UserID int) User {
	for _, user := range users.Attribute {
		if UserID == user.ID {
			return user
		}
	}
	return User{}
}
