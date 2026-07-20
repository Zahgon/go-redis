package maintnotifications

import (
	"errors"

	"github.com/redis/go-redis/v9/internal/maintnotifications/logs"
)

var (
	ErrInvalidRelaxedTimeout             = errors.New(logs.InvalidRelaxedTimeoutError())
	ErrInvalidHandoffTimeout             = errors.New(logs.InvalidHandoffTimeoutError())
	ErrInvalidHandoffWorkers             = errors.New(logs.InvalidHandoffWorkersError())
	ErrInvalidHandoffQueueSize           = errors.New(logs.InvalidHandoffQueueSizeError())
	ErrInvalidPostHandoffRelaxedDuration = errors.New(logs.InvalidPostHandoffRelaxedDurationError())
	ErrInvalidEndpointType               = errors.New(logs.InvalidEndpointTypeError())
	ErrInvalidMaintNotifications         = errors.New(logs.InvalidMaintNotificationsError())
	ErrMaxHandoffRetriesReached          = errors.New(logs.MaxHandoffRetriesReachedError())

	ErrInvalidHandoffRetries = errors.New(logs.InvalidHandoffRetriesError())
)

var (
	ErrInvalidClient = errors.New(logs.InvalidClientError())
)

var (
	ErrHandoffQueueFull = errors.New(logs.HandoffQueueFullError())
)

var (
	ErrInvalidNotification = errors.New(logs.InvalidNotificationError())
)

var (
	ErrConnectionMarkedForHandoff = errors.New(logs.ConnectionMarkedForHandoffErrorMessage)

	ErrConnectionMarkedForHandoffWithState = errors.New(logs.ConnectionMarkedForHandoffErrorMessage + " with state")

	ErrConnectionInvalidHandoffState = errors.New(logs.ConnectionInvalidHandoffStateErrorMessage)
)

var (
	ErrShutdown = errors.New(logs.ShutdownError())
)

var (
	ErrCircuitBreakerOpen = errors.New(logs.CircuitBreakerOpenErrorMessage)
)

var (
	ErrInvalidCircuitBreakerFailureThreshold = errors.New(logs.InvalidCircuitBreakerFailureThresholdError())

	ErrInvalidCircuitBreakerResetTimeout = errors.New(logs.InvalidCircuitBreakerResetTimeoutError())

	ErrInvalidCircuitBreakerMaxRequests = errors.New(logs.InvalidCircuitBreakerMaxRequestsError())
)
