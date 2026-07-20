package push

type NotificationHandlerContext struct {
	Client interface{}

	ConnPool interface{}

	PubSub interface{}

	Conn interface{}

	IsBlocking bool
}
