package users

import "time"

type User struct{
	Id int
	Name string
	Occupation string
	Email string
	Password string
	Avatar_file string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time 
}