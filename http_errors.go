package apiutils

type ErrBadRequest struct{ Msg string }

func (e *ErrBadRequest) Error() string { return e.Msg }

func NewErrBadRequest(msg string) *ErrBadRequest {
	return &ErrBadRequest{Msg: msg}
}

type ErrNotFound struct{ Msg string }

func (e *ErrNotFound) Error() string { return e.Msg }

func NewErrNotFound(msg string) *ErrNotFound {
	return &ErrNotFound{Msg: msg}
}

type ErrConflict struct{ Msg string }

func (e *ErrConflict) Error() string { return e.Msg }

func NewErrConflict(msg string) *ErrConflict {
	return &ErrConflict{Msg: msg}
}

type ErrForbidden struct{ Msg string }

func (e *ErrForbidden) Error() string { return e.Msg }

func NewErrForbidden(msg string) *ErrForbidden {
	return &ErrForbidden{Msg: msg}
}

type ErrUnauthorized struct{ Msg string }

func (e *ErrUnauthorized) Error() string { return e.Msg }

func NewErrUnauthorized(msg string) *ErrUnauthorized {
	return &ErrUnauthorized{Msg: msg}
}
