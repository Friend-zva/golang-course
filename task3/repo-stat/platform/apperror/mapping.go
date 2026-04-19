package apperror

import (
	"errors"
	"net/http"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func Pack(err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) {
		grpcCode := codes.Internal
		switch appErr.Code {
		case http.StatusNotFound:
			grpcCode = codes.NotFound
		case http.StatusInternalServerError:
			grpcCode = codes.Internal
		case http.StatusBadGateway:
			grpcCode = codes.Unavailable
		}

		var msg string
		if appErr.Err != nil {
			msg = appErr.Err.Error()
		} else {
			msg = appErr.Message
		}

		return status.Error(grpcCode, msg)
	}

	return status.Error(codes.Internal, err.Error())
}

func Unpack(err error) error {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return err
	}

	st, ok := status.FromError(err)
	if ok {
		errMes := errors.New(st.Message())
		switch st.Code() {
		case codes.NotFound:
			return ErrNotFound.Wrap(errMes)
		case codes.Internal:
			return ErrInternal.Wrap(errMes)
		case codes.Unavailable:
			return ErrExternal.Wrap(errMes)
		}
	}

	return ErrInternal.Wrap(err)
}
