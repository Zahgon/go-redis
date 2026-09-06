package auth

type ReAuthCredentialsListener struct {
	reAuth func(credentials Credentials) error
	onErr  func(err error)
}

func (c *ReAuthCredentialsListener) OnNext(credentials Credentials) {
	_ = "STUB: not implemented"
	return
}

func (c *ReAuthCredentialsListener) OnError(err error) { _ = "STUB: not implemented"; return }

func NewReAuthCredentialsListener(reAuth func(credentials Credentials) error, onErr func(err error)) *ReAuthCredentialsListener {
	_ = "STUB: not implemented"
	return nil
}

var _ CredentialsListener = (*ReAuthCredentialsListener)(nil)
