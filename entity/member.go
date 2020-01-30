package entity

//Member is a struct representing member
type Member struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//NewMember creates an instance of Member.
func NewMember() Member {
	return Member{}
}
