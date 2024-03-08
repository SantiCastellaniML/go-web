package internal

type Owner struct {
	ID       int
	Name     string
	LastName string
	Tasks    []Task
}

type OwnerRepository interface {
}
