package components

type Component interface {
	Component() (string, error)
}
