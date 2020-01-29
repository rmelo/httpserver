package entity

//Member is a struct representing member
type Member struct {
	Name  string
	Email string
}

//NewMember creates an instance of Member.
func NewMember() interface{} {
	return &Member{}
}
