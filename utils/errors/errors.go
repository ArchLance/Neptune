package errors

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type RequestErr struct {
	Err error
}
type ParamErr struct {
	Err error
}
type DbErr struct {
	Err error
}
type LogicErr struct {
	Err error
}
type NotFoundErr struct {
	Err error
}
type ExistErr struct {
	Err error
}

func (e RequestErr) Error() string {
	return e.Err.Error()
}

func (e ParamErr) Error() string {
	return e.Err.Error()
}

func (e DbErr) Error() string {
	return e.Err.Error()
}

func (e LogicErr) Error() string {
	return e.Err.Error()
}
func (e NotFoundErr) Error() string {
	return e.Err.Error()
}
func (e ExistErr) Error() string {
	return e.Err.Error()
}
