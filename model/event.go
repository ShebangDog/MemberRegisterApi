package model

type (
	Event interface {
		Value() string
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

func (r register) Value() string {
	return "register"
}

func (d delete) Value() string {
	return "delete"
}

func (e entry) Value() string {
	return "entry"
}

func (e exit) Value() string {
	return "exit"
}
