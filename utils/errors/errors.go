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
type LogicError struct {
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

func (e LogicError) Error() string {
	return e.Err.Error()
}
