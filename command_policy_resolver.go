package redis

import (
	"context"

	"github.com/redis/go-redis/v9/internal/routing"
)

type (
	module      = string
	commandName = string
)

var defaultPolicies = map[module]map[commandName]*routing.CommandPolicy{
	"ft": {
		"create": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"search": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"aggregate": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"dictadd": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"dictdump": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"dictdel": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"suglen": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultHashSlot,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"cursor": {
			Request:  routing.ReqSpecial,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"sugadd": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultHashSlot,
		},
		"sugget": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultHashSlot,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"sugdel": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultHashSlot,
		},
		"spellcheck": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"explain": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"explaincli": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"aliasadd": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"aliasupdate": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"aliasdel": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"aliaslist": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"info": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"tagvals": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"syndump": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"synupdate": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"profile": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
			Tips: map[string]string{
				routing.ReadOnlyCMD: "",
			},
		},
		"alter": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"dropindex": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
		"drop": {
			Request:  routing.ReqDefault,
			Response: routing.RespDefaultKeyless,
		},
	},
}

func defaultPolicyKeyless(name string) bool { _ = "STUB: not implemented"; return false }

type CommandInfoResolveFunc func(ctx context.Context, cmd Cmder) *routing.CommandPolicy

type commandInfoResolver struct {
	resolveFunc      CommandInfoResolveFunc
	fallBackResolver *commandInfoResolver
}

func NewCommandInfoResolver(resolveFunc CommandInfoResolveFunc) *commandInfoResolver {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultCommandPolicyResolver() *commandInfoResolver { _ = "STUB: not implemented"; return nil }

func (r *commandInfoResolver) GetCommandPolicy(ctx context.Context, cmd Cmder) *routing.CommandPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *commandInfoResolver) SetFallbackResolver(fallbackResolver *commandInfoResolver) {
	_ = "STUB: not implemented"
	return
}
