package redis

import (
	"context"
)

type VectorSetCmdable interface {
	VAdd(ctx context.Context, key, element string, val Vector) *BoolCmd
	VAddWithArgs(ctx context.Context, key, element string, val Vector, addArgs *VAddArgs) *BoolCmd
	VCard(ctx context.Context, key string) *IntCmd
	VDim(ctx context.Context, key string) *IntCmd
	VEmb(ctx context.Context, key, element string, raw bool) *SliceCmd
	VGetAttr(ctx context.Context, key, element string) *StringCmd
	VInfo(ctx context.Context, key string) *MapStringInterfaceCmd
	VLinks(ctx context.Context, key, element string) *StringSliceSliceCmd
	VLinksWithScores(ctx context.Context, key, element string) *VectorScoreSliceSliceCmd
	VRandMember(ctx context.Context, key string) *StringCmd
	VRandMemberCount(ctx context.Context, key string, count int) *StringSliceCmd
	VRem(ctx context.Context, key, element string) *BoolCmd
	VSetAttr(ctx context.Context, key, element string, attr interface{}) *BoolCmd
	VClearAttributes(ctx context.Context, key, element string) *BoolCmd
	VSim(ctx context.Context, key string, val Vector) *StringSliceCmd
	VSimWithScores(ctx context.Context, key string, val Vector) *VectorScoreSliceCmd
	VSimWithArgs(ctx context.Context, key string, val Vector, args *VSimArgs) *StringSliceCmd
	VSimWithArgsWithScores(ctx context.Context, key string, val Vector, args *VSimArgs) *VectorScoreSliceCmd
	VSimWithArgsWithAttribs(ctx context.Context, key string, val Vector, args *VSimArgs) *VectorAttribSliceCmd
	VSimWithArgsWithScoresWithAttribs(ctx context.Context, key string, val Vector, args *VSimArgs) *VectorScoreAttribSliceCmd
	VRange(ctx context.Context, key, start, end string, count int64) *StringSliceCmd
	VIsMember(ctx context.Context, key, element string) *BoolCmd
}

type Vector interface {
	Value() []any
}

const (
	vectorFormatFP32   string = "FP32"
	vectorFormatValues string = "Values"
	vectorFormatF16    string = "FLOAT16"
	vectorFormatBF16   string = "BFLOAT16"
	vectorFormatF64    string = "FLOAT64"
	vectorFormatI8     string = "INT8"
	vectorFormatU8     string = "UINT8"
)

type VectorFP32 struct {
	Val []byte
}

func (v *VectorFP32) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorFP32)(nil)

type VectorFloat16 struct {
	Val []byte
}

func (v *VectorFloat16) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorFloat16)(nil)

type VectorBFloat16 struct {
	Val []byte
}

func (v *VectorBFloat16) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorBFloat16)(nil)

type VectorFloat64 struct {
	Val []byte
}

func (v *VectorFloat64) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorFloat64)(nil)

type VectorInt8 struct {
	Val []byte
}

func (v *VectorInt8) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorInt8)(nil)

type VectorUint8 struct {
	Val []byte
}

func (v *VectorUint8) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorUint8)(nil)

type VectorValues struct {
	Val []float64
}

func (v *VectorValues) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorValues)(nil)

type VectorRef struct {
	Name string
}

func (v *VectorRef) Value() []any { _ = "STUB: not implemented"; return nil }

var _ Vector = (*VectorRef)(nil)

type VectorScore struct {
	Name  string
	Score float64
}

type VectorAttrib struct {
	Name    string
	Attribs *string
}

type VectorScoreAttrib struct {
	Name    string
	Score   float64
	Attribs *string
}

func (c cmdable) VAdd(ctx context.Context, key, element string, val Vector) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

type VAddArgs struct {
	Reduce int64
	Cas    bool

	NoQuant bool
	Q8      bool
	Bin     bool

	EF      int64
	SetAttr string
	M       int64
}

func (v VAddArgs) reduce() int64 { _ = "STUB: not implemented"; return 0 }

func (v VAddArgs) appendArgs(args []any) []any { _ = "STUB: not implemented"; return nil }

func (c cmdable) VAddWithArgs(ctx context.Context, key, element string, val Vector, addArgs *VAddArgs) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VDim(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VEmb(ctx context.Context, key, element string, raw bool) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VGetAttr(ctx context.Context, key, element string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VInfo(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VLinks(ctx context.Context, key, element string) *StringSliceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VLinksWithScores(ctx context.Context, key, element string) *VectorScoreSliceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VRandMember(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VRandMemberCount(ctx context.Context, key string, count int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VRem(ctx context.Context, key, element string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSetAttr(ctx context.Context, key, element string, attr interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VClearAttributes(ctx context.Context, key, element string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSim(ctx context.Context, key string, val Vector) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSimWithScores(ctx context.Context, key string, val Vector) *VectorScoreSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type VSimArgs struct {
	Count    int64
	EF       int64
	Filter   string
	FilterEF int64
	Truth    bool
	NoThread bool
	Epsilon  float64
}

func (v VSimArgs) appendArgs(args []any) []any { _ = "STUB: not implemented"; return nil }

func (c cmdable) VSimWithArgs(ctx context.Context, key string, val Vector, simArgs *VSimArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSimWithArgsWithScores(ctx context.Context, key string, val Vector, simArgs *VSimArgs) *VectorScoreSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSimWithArgsWithAttribs(ctx context.Context, key string, val Vector, simArgs *VSimArgs) *VectorAttribSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VSimWithArgsWithScoresWithAttribs(ctx context.Context, key string, val Vector, simArgs *VSimArgs) *VectorScoreAttribSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VRange(ctx context.Context, key, start, end string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) VIsMember(ctx context.Context, key, element string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}
