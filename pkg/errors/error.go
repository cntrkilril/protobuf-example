package errors

import "errors"

type (
	Error struct {
		msg  string
		code ErrCode
	}
	ErrCode int64
)

const (
	_ = iota
	ErrCodeUnknown
	ErrCodeNotFound
	ErrCodeInvalidArgument
)

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

var _ error = &Error{}

func NewError(msg string, code ErrCode) *Error {
	return &Error{msg, code}
}

var (
	ErrUnknown                   = errors.New("что-то пошло не так")
	ErrValidationError           = errors.New("невалидные данные")
	ErrCategoryNotFound          = errors.New("категория не найдена")
	ErrServiceNotFound           = errors.New("сервис не найден")
	ErrMasterNotFound            = errors.New("мастер не найден")
	ErrImagesNotFound            = errors.New("изображение не найдено")
	ErrSocialNetworkNotFound     = errors.New("социальная сеть не найдена")
	ErrSocialNetworkTypeNotFound = errors.New("тип социальной сети не найден")
	ErrProfileNotFound           = errors.New("профиль не найден")
	ErrClientNotFound            = errors.New("клиент не найден")
	ErrAppointmentNotFound       = errors.New("запись не найдена")
	ErrTimeCellNotFound          = errors.New("ячейка времени не найдена")
	ErrWorkTimeNotFound          = errors.New("время работы не найдено")
)
