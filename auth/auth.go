package auth

type StreamingCredentialsProvider interface {
	Subscribe(listener CredentialsListener) (Credentials, UnsubscribeFunc, error)
}

type UnsubscribeFunc func() error

type CredentialsListener interface {
	OnNext(credentials Credentials)
	OnError(err error)
}

type Credentials interface {
	BasicAuth() (username string, password string)

	RawCredentials() string
}

type basicAuth struct {
	username string
	password string
}

func (b *basicAuth) RawCredentials() string { _ = "STUB: not implemented"; return "" }

func (b *basicAuth) BasicAuth() (username string, password string) {
	_ = "STUB: not implemented"
	return "", ""
}

func NewBasicCredentials(username, password string) Credentials {
	_ = "STUB: not implemented"
	return *new(Credentials)
}
