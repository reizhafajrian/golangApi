package users
 type UserFormatter struct {
	 Id int `json:"id"`
	 Name string `json:"name"`
	 Occupation string `json:"occupation"`
	 Email string `json:"email"`
	 Token string `json:"token"`
 }

 func FormateUser(user User,token string)UserFormatter{
	 formatter:=UserFormatter{
		 Id: user.Id,
		 Name: user.Name,
		 Occupation: user.Occupation,
		 Email: user.Email,
		 Token: token,
	 }
	 return formatter
 }