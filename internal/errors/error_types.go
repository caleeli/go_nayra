package errors

import "github.com/pkg/errors"

type RuntimeError struct {
	error
}

type StorageError struct {
	error
}

type ServiceError struct {
	error
}

func WrapEngineError(err error, format string, args ...interface{}) error {
	return &RuntimeError{errors.Wrapf(err, format, args...)}
}

func WrapStorageError(err error, format string, args ...interface{}) error {
	return &StorageError{errors.Wrapf(err, format, args...)}
}

func WrapServiceError(err error, format string, args ...interface{}) error {
	return &ServiceError{errors.Wrapf(err, format, args...)}
}
