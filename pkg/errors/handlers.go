package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcError(err error) error {
	appErr := &Error{}
	if errors.As(err, &appErr) {
		return status.Error(errGrpcCodeMap[appErr.Code()], appErr.Error())
	}
	return status.Error(codes.Internal, err.Error())
}

func HandleServiceError(err error) error {
	switch err {
	case ErrCategoryNotFound:
		return NewError(ErrCategoryNotFound.Error(), ErrCodeNotFound)
	case ErrServiceNotFound:
		return NewError(ErrServiceNotFound.Error(), ErrCodeNotFound)
	case ErrMasterNotFound:
		return NewError(ErrMasterNotFound.Error(), ErrCodeNotFound)
	case ErrProfileNotFound:
		return NewError(ErrProfileNotFound.Error(), ErrCodeNotFound)
	case ErrImagesNotFound:
		return NewError(ErrImagesNotFound.Error(), ErrCodeNotFound)
	case ErrSocialNetworkNotFound:
		return NewError(ErrSocialNetworkNotFound.Error(), ErrCodeNotFound)
	case ErrSocialNetworkTypeNotFound:
		return NewError(ErrSocialNetworkTypeNotFound.Error(), ErrCodeNotFound)
	case ErrClientNotFound:
		return NewError(ErrClientNotFound.Error(), ErrCodeNotFound)
	case ErrAppointmentNotFound:
		return NewError(ErrAppointmentNotFound.Error(), ErrCodeNotFound)
	case ErrTimeCellNotFound:
		return NewError(ErrTimeCellNotFound.Error(), ErrCodeNotFound)
	case ErrWorkTimeNotFound:
		return NewError(ErrWorkTimeNotFound.Error(), ErrCodeNotFound)
	case ErrValidationError:
		return NewError(ErrValidationError.Error(), ErrCodeInvalidArgument)
	default:
		return NewError(ErrUnknown.Error(), ErrCodeUnknown)
	}
}
