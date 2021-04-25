package model

type (
	Event interface {
		value() string
	}

	register struct{}
	delete   struct{}
	entry    struct{}
	exit     struct{}
)

func Register() Event {
	return register{}
}

func Delete() Event {
	return delete{}
}

func Entry() Event {
	return entry{}
}

func Exit() Event {
	return exit{}
}

func (r register) value() string {
	return "register"
}

func (d delete) value() string {
	return "delete"
}

func (e entry) value() string {
	return "entry"
}

func (e exit) value() string {
	return "exit"
}
