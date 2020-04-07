package util

type Auth interface {
	Action() bool
}

type UnAuth struct{}

func (UnAuth) Action() bool {
	return true
}
