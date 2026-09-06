package push

import (
	"errors"
)

const (
	ReasonHandlerNil       = "handler cannot be nil"
	ReasonHandlerExists    = "cannot overwrite existing handler"
	ReasonHandlerProtected = "handler is protected"

	ReasonPushNotificationsDisabled = "push notifications are disabled"
)

type ProcessorType string

const (
	ProcessorTypeProcessor     = ProcessorType("processor")
	ProcessorTypeVoidProcessor = ProcessorType("void_processor")
	ProcessorTypeCustom        = ProcessorType("custom")
)

type ProcessorOperation string

const (
	ProcessorOperationProcess    = ProcessorOperation("process")
	ProcessorOperationRegister   = ProcessorOperation("register")
	ProcessorOperationUnregister = ProcessorOperation("unregister")
	ProcessorOperationUnknown    = ProcessorOperation("unknown")
)

var (
	ErrHandlerNil = errors.New(ReasonHandlerNil)
)

func ErrHandlerExists(pushNotificationName string) error { _ = "STUB: not implemented"; return nil }

func ErrProtectedHandler(pushNotificationName string) error { _ = "STUB: not implemented"; return nil }

func ErrVoidProcessorRegister(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}

func ErrVoidProcessorUnregister(pushNotificationName string) error {
	_ = "STUB: not implemented"
	return nil
}

type HandlerError struct {
	Operation            ProcessorOperation
	PushNotificationName string
	Reason               string
	Err                  error
}

func (e *HandlerError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *HandlerError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func NewHandlerError(operation ProcessorOperation, pushNotificationName, reason string, err error) *HandlerError {
	_ = "STUB: not implemented"
	return nil
}

type ProcessorError struct {
	ProcessorType        ProcessorType
	Operation            ProcessorOperation
	PushNotificationName string
	Reason               string
	Err                  error
}

func (e *ProcessorError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ProcessorError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func NewProcessorError(processorType ProcessorType, operation ProcessorOperation, pushNotificationName, reason string, err error) *ProcessorError {
	_ = "STUB: not implemented"
	return nil
}

func IsHandlerNilError(err error) bool { _ = "STUB: not implemented"; return false }

func IsHandlerExistsError(err error) bool { _ = "STUB: not implemented"; return false }

func IsProtectedHandlerError(err error) bool { _ = "STUB: not implemented"; return false }

func IsVoidProcessorError(err error) bool { _ = "STUB: not implemented"; return false }
