package user

type User interface {
	GetName() string
	GetID() uint
}

type DefaultUser struct {
	Name     string
	FullName string
	ID       uint
	Email    string
	Admin    bool
}

func (d *DefaultUser) GetName() string {
	return d.Name
}

func (d *DefaultUser) GetID() uint {
	return d.ID
}
