package redisotel

const (
	AttrDBSystemName         = "db.system.name"
	AttrDBNamespace          = "db.namespace"
	AttrDBOperationName      = "db.operation.name"
	AttrDBOperationBatchSize = "db.operation.batch.size"
	AttrDBResponseStatusCode = "db.response.status_code"

	AttrDBClientConnectionPoolName = "db.client.connection.pool.name"
	AttrDBClientConnectionState    = "db.client.connection.state"

	AttrServerAddress = "server.address"
	AttrServerPort    = "server.port"

	AttrNetworkPeerAddress = "network.peer.address"
	AttrNetworkPeerPort    = "network.peer.port"

	AttrErrorType = "error.type"

	AttrRedisClientLibrary                = "redis.client.library"
	AttrRedisClientConnectionPubSub       = "redis.client.connection.pubsub"
	AttrRedisClientConnectionCloseReason  = "redis.client.connection.close.reason"
	AttrRedisClientErrorsCategory         = "redis.client.errors.category"
	AttrRedisClientErrorsInternal         = "redis.client.errors.internal"
	AttrRedisClientOperationRetryAttempts = "redis.client.operation.retry_attempts"

	AttrRedisClientPubSubDirection = "redis.client.pubsub.direction"
	AttrRedisClientPubSubSharded   = "redis.client.pubsub.sharded"
	AttrRedisClientPubSubChannel   = "redis.client.pubsub.channel"

	AttrRedisClientStreamName          = "redis.client.stream.name"
	AttrRedisClientStreamConsumerGroup = "redis.client.stream.consumer_group"
	AttrRedisClientStreamConsumerName  = "redis.client.stream.consumer_name"

	AttrRedisClientConnectionNotification = "redis.client.connection.notification"
)

const (
	ConnectionStateIdle = "idle"
	ConnectionStateUsed = "used"
)

const (
	PubSubDirectionSent     = "sent"
	PubSubDirectionReceived = "received"
)

const (
	DBSystemRedis = "redis"
)
