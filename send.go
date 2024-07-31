package test

type Send interface {
	Execute(data string) (int, error)
}

type SendFunc func(data string) (int, error)

func (f SendFunc) Execute(data string) (int, error) {
	return f(data)
}
