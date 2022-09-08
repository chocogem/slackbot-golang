package domain

type User struct {
	UserId       string    
	UserName     string   
	FirstName    string    
	LastName     string   
	ExpireDate   string    
}
type BigQueryConfig struct {
	ProjectId   string
	DatasetName string
}

func NewUser(userId string, userName string, firstName string, lastName string,expireDate string) User {
	return User{
		UserId:     userId,
		UserName:   userName,
		FirstName:  firstName,
		LastName:   lastName,
		ExpireDate: expireDate,
	}
}
