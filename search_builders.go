package redis

import (
	"context"
)

type SearchBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	query   string
	options *FTSearchOptions
}

func (c *Client) NewSearchBuilder(ctx context.Context, index, query string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) WithScores() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) NoContent() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Verbatim() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) NoStopWords() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) WithPayloads() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) WithSortKeys() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Filter(field string, min, max interface{}) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) GeoFilter(field string, lon, lat, radius float64, unit string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) InKeys(keys ...interface{}) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) InFields(fields ...interface{}) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) ReturnFields(fields ...string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) ReturnAs(field, alias string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) Slop(slop int) *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Timeout(timeout int) *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) InOrder() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Language(lang string) *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Expander(expander string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) Scorer(scorer string) *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) ExplainScore() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Payload(payload string) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) SortBy(field string, asc bool) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) WithSortByCount() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Param(key string, value interface{}) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) ParamsMap(p map[string]interface{}) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) Dialect(version int) *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Limit(offset, count int) *SearchBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SearchBuilder) CountOnly() *SearchBuilder { _ = "STUB: not implemented"; return nil }

func (b *SearchBuilder) Run() (FTSearchResult, error) {
	_ = "STUB: not implemented"
	return *new(FTSearchResult), nil
}

type AggregateBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	query   string
	options *FTAggregateOptions
	err     error
}

func (c *Client) NewAggregateBuilder(ctx context.Context, index, query string) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) setErr(err error) { _ = "STUB: not implemented"; return }

func (b *AggregateBuilder) Verbatim() *AggregateBuilder { _ = "STUB: not implemented"; return nil }

func (b *AggregateBuilder) AddScores() *AggregateBuilder { _ = "STUB: not implemented"; return nil }

func (b *AggregateBuilder) Scorer(s string) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) LoadAll() *AggregateBuilder { _ = "STUB: not implemented"; return nil }

func (b *AggregateBuilder) Load(field string, alias ...string) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Timeout(ms int) *AggregateBuilder { _ = "STUB: not implemented"; return nil }

func (b *AggregateBuilder) Apply(field string, alias ...string) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) GroupBy(fields ...interface{}) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Reduce(fn SearchAggregator, args ...interface{}) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) ReduceAs(fn SearchAggregator, alias string, args ...interface{}) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) SortBy(field string, asc bool) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) SortByMax(max int) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Filter(expr string) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) WithCursor(count, maxIdle int) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Params(p map[string]interface{}) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Dialect(version int) *AggregateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AggregateBuilder) Run() (*FTAggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreateIndexBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	options *FTCreateOptions
	schema  []*FieldSchema
}

func (c *Client) NewCreateIndexBuilder(ctx context.Context, index string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) OnHash() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) OnJSON() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) Prefix(prefixes ...interface{}) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) Filter(filter string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) DefaultLanguage(lang string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) LanguageField(field string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) Score(score float64) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) ScoreField(field string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) PayloadField(field string) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) NoOffsets() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) Temporary(sec int) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) NoHL() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) NoFields() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) NoFreqs() *CreateIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *CreateIndexBuilder) StopWords(words ...interface{}) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) SkipInitialScan() *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) Schema(field *FieldSchema) *CreateIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CreateIndexBuilder) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }

type DropIndexBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	options *FTDropIndexOptions
}

func (c *Client) NewDropIndexBuilder(ctx context.Context, index string) *DropIndexBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *DropIndexBuilder) DeleteDocs() *DropIndexBuilder { _ = "STUB: not implemented"; return nil }

func (b *DropIndexBuilder) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }

type AliasBuilder struct {
	c      *Client
	ctx    context.Context
	alias  string
	index  string
	action string
}

func (c *Client) NewAliasBuilder(ctx context.Context, alias string) *AliasBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AliasBuilder) Action(action string) *AliasBuilder { _ = "STUB: not implemented"; return nil }

func (b *AliasBuilder) Add(index string) *AliasBuilder { _ = "STUB: not implemented"; return nil }

func (b *AliasBuilder) Del() *AliasBuilder { _ = "STUB: not implemented"; return nil }

func (b *AliasBuilder) Update(index string) *AliasBuilder { _ = "STUB: not implemented"; return nil }

func (b *AliasBuilder) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }

type ExplainBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	query   string
	options *FTExplainOptions
}

func (c *Client) NewExplainBuilder(ctx context.Context, index, query string) *ExplainBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ExplainBuilder) Dialect(d string) *ExplainBuilder { _ = "STUB: not implemented"; return nil }

func (b *ExplainBuilder) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }

type FTInfoBuilder struct {
	c     *Client
	ctx   context.Context
	index string
}

func (c *Client) NewSearchInfoBuilder(ctx context.Context, index string) *FTInfoBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *FTInfoBuilder) Run() (FTInfoResult, error) {
	_ = "STUB: not implemented"
	return *new(FTInfoResult), nil
}

type SpellCheckBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	query   string
	options *FTSpellCheckOptions
}

func (c *Client) NewSpellCheckBuilder(ctx context.Context, index, query string) *SpellCheckBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SpellCheckBuilder) Distance(d int) *SpellCheckBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SpellCheckBuilder) Terms(include bool, dictionary string, terms ...interface{}) *SpellCheckBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SpellCheckBuilder) Dialect(d int) *SpellCheckBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SpellCheckBuilder) Run() ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DictBuilder struct {
	c      *Client
	ctx    context.Context
	dict   string
	terms  []interface{}
	action string
}

func (c *Client) NewDictBuilder(ctx context.Context, dict string) *DictBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *DictBuilder) Action(action string) *DictBuilder { _ = "STUB: not implemented"; return nil }

func (b *DictBuilder) Add(terms ...interface{}) *DictBuilder { _ = "STUB: not implemented"; return nil }

func (b *DictBuilder) Del(terms ...interface{}) *DictBuilder { _ = "STUB: not implemented"; return nil }

func (b *DictBuilder) Dump() *DictBuilder { _ = "STUB: not implemented"; return nil }

func (b *DictBuilder) Run() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

type TagValsBuilder struct {
	c     *Client
	ctx   context.Context
	index string
	field string
}

func (c *Client) NewTagValsBuilder(ctx context.Context, index, field string) *TagValsBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *TagValsBuilder) Run() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

type CursorBuilder struct {
	c        *Client
	ctx      context.Context
	index    string
	cursorId int64
	count    int
	action   string
}

func (c *Client) NewCursorBuilder(ctx context.Context, index string, cursorId int64) *CursorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CursorBuilder) Action(action string) *CursorBuilder { _ = "STUB: not implemented"; return nil }

func (b *CursorBuilder) Read() *CursorBuilder { _ = "STUB: not implemented"; return nil }

func (b *CursorBuilder) Del() *CursorBuilder { _ = "STUB: not implemented"; return nil }

func (b *CursorBuilder) Count(count int) *CursorBuilder { _ = "STUB: not implemented"; return nil }

func (b *CursorBuilder) Run() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

type SynUpdateBuilder struct {
	c       *Client
	ctx     context.Context
	index   string
	groupId interface{}
	options *FTSynUpdateOptions
	terms   []interface{}
}

func (c *Client) NewSynUpdateBuilder(ctx context.Context, index string, groupId interface{}) *SynUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SynUpdateBuilder) SkipInitialScan() *SynUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SynUpdateBuilder) Terms(terms ...interface{}) *SynUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *SynUpdateBuilder) Run() (string, error) { _ = "STUB: not implemented"; return "", nil }
