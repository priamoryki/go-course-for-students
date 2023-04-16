package ads

type User struct {
	RepoEntity
	Nickname string
	Email    string
}

func (user *User) HasName(name string) bool {
	return user.Nickname == name
}
