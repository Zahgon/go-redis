package redis

import "context"

type ACLCmdable interface {
	ACLDryRun(ctx context.Context, username string, command ...interface{}) *StringCmd

	ACLLog(ctx context.Context, count int64) *ACLLogCmd
	ACLLogReset(ctx context.Context) *StatusCmd

	ACLGenPass(ctx context.Context, bit int) *StringCmd

	ACLSetUser(ctx context.Context, username string, rules ...string) *StatusCmd
	ACLDelUser(ctx context.Context, username string) *IntCmd
	ACLUsers(ctx context.Context) *StringSliceCmd
	ACLWhoAmI(ctx context.Context) *StringCmd
	ACLList(ctx context.Context) *StringSliceCmd

	ACLCat(ctx context.Context) *StringSliceCmd
	ACLCatArgs(ctx context.Context, options *ACLCatArgs) *StringSliceCmd
}

type ACLCatArgs struct {
	Category string
}

func (c cmdable) ACLDryRun(ctx context.Context, username string, command ...interface{}) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLLog(ctx context.Context, count int64) *ACLLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLLogReset(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ACLDelUser(ctx context.Context, username string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLSetUser(ctx context.Context, username string, rules ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLGenPass(ctx context.Context, bit int) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLUsers(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLWhoAmI(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ACLList(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) ACLCat(ctx context.Context) *StringSliceCmd { _ = "STUB: not implemented"; return nil }

func (c cmdable) ACLCatArgs(ctx context.Context, options *ACLCatArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}
