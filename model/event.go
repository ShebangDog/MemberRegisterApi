package model

type (
	Event interface {
		Value() string
	}

	NoSuchEventError struct {
		msg string
	}

	register struct{}
	delete   struct{}
	entry    struct{}
	exit     struct{}
)

var eventList = []Event{Register(), Delete(), Entry(), Exit()}

func ParseEvent(str string) (Event, error) {
	for _, event := range eventList {
		if event.Value() == str {
			return event, nil
		}
	}

	return nil, &NoSuchEventError{msg: "no such event error"}
}

func Register() Event {
	return &register{}
}

func Delete() Event {
	return &delete{}
}

func Entry() Event {
	return &entry{}
}

func Exit() Event {
	return &exit{}
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

func (e *NoSuchEventError) Error() string {
	return e.msg
}
