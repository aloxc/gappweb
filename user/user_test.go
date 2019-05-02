package user

import "testing"

func TestGetUser(t *testing.T) {
	var id int32 = 1
	user := &User{
		Id: id,
	}
	GetUser(user)
}
func TestGetUserByIdAndPassowrd(t *testing.T) {
	var id int32 = 1
	user := &User{
		Id:       id,
		Password: "111111",
	}
	GetUser(user)
}
func TestRegister(t *testing.T) {
	var userName = "aloxc"
	var password = "111111"
	var level = User_Level_OK
	Register(userName, password, level)
}
func BenchmarkGetUserById(b *testing.B) {
}
