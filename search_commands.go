package redis

import (
	"context"

	"github.com/redis/go-redis/v9/internal/proto"
)

type SearchCmdable interface {
	FT_List(ctx context.Context) *StringSliceCmd
	FTAggregate(ctx context.Context, index string, query string) *MapStringInterfaceCmd
	FTAggregateWithArgs(ctx context.Context, index string, query string, options *FTAggregateOptions) *AggregateCmd
	FTAliasAdd(ctx context.Context, index string, alias string) *StatusCmd
	FTAliasDel(ctx context.Context, alias string) *StatusCmd
	FTAliasUpdate(ctx context.Context, index string, alias string) *StatusCmd
	FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *StatusCmd
	FTConfigGet(ctx context.Context, option string) *MapMapStringInterfaceCmd
	FTConfigSet(ctx context.Context, option string, value interface{}) *StatusCmd
	FTCreate(ctx context.Context, index string, options *FTCreateOptions, schema ...*FieldSchema) *StatusCmd
	FTCursorDel(ctx context.Context, index string, cursorId int) *StatusCmd
	FTCursorRead(ctx context.Context, index string, cursorId int, count int) *MapStringInterfaceCmd
	FTDictAdd(ctx context.Context, dict string, term ...interface{}) *IntCmd
	FTDictDel(ctx context.Context, dict string, term ...interface{}) *IntCmd
	FTDictDump(ctx context.Context, dict string) *StringSliceCmd
	FTDropIndex(ctx context.Context, index string) *StatusCmd
	FTDropIndexWithArgs(ctx context.Context, index string, options *FTDropIndexOptions) *StatusCmd
	FTExplain(ctx context.Context, index string, query string) *StringCmd
	FTExplainWithArgs(ctx context.Context, index string, query string, options *FTExplainOptions) *StringCmd
	FTHybrid(ctx context.Context, index string, searchExpr string, vectorField string, vectorData Vector) *FTHybridCmd
	FTHybridWithArgs(ctx context.Context, index string, options *FTHybridOptions) *FTHybridCmd
	FTInfo(ctx context.Context, index string) *FTInfoCmd
	FTSpellCheck(ctx context.Context, index string, query string) *FTSpellCheckCmd
	FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *FTSpellCheckOptions) *FTSpellCheckCmd
	FTSearch(ctx context.Context, index string, query string) *FTSearchCmd
	FTSearchWithArgs(ctx context.Context, index string, query string, options *FTSearchOptions) *FTSearchCmd
	FTSynDump(ctx context.Context, index string) *FTSynDumpCmd
	FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *StatusCmd
	FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *FTSynUpdateOptions, terms []interface{}) *StatusCmd
	FTTagVals(ctx context.Context, index string, field string) *StringSliceCmd
}

type FTCreateOptions struct {
	OnHash          bool
	OnJSON          bool
	Prefix          []interface{}
	Filter          string
	DefaultLanguage string
	LanguageField   string
	Score           float64
	ScoreField      string
	PayloadField    string
	MaxTextFields   int
	NoOffsets       bool
	Temporary       int
	NoHL            bool
	NoFields        bool
	NoFreqs         bool
	StopWords       []interface{}
	SkipInitialScan bool
}

type FieldSchema struct {
	FieldName         string
	As                string
	FieldType         SearchFieldType
	Sortable          bool
	UNF               bool
	NoStem            bool
	NoIndex           bool
	PhoneticMatcher   string
	Weight            float64
	Separator         string
	CaseSensitive     bool
	WithSuffixtrie    bool
	VectorArgs        *FTVectorArgs
	GeoShapeFieldType string
	IndexEmpty        bool
	IndexMissing      bool
}

type FTVectorArgs struct {
	FlatOptions   *FTFlatOptions
	HNSWOptions   *FTHNSWOptions
	VamanaOptions *FTVamanaOptions
}

type FTFlatOptions struct {
	Type            string
	Dim             int
	DistanceMetric  string
	InitialCapacity int
	BlockSize       int
}

type FTHNSWOptions struct {
	Type                   string
	Dim                    int
	DistanceMetric         string
	InitialCapacity        int
	MaxEdgesPerNode        int
	MaxAllowedEdgesPerNode int
	EFRunTime              int
	Epsilon                float64
}

type FTVamanaOptions struct {
	Type                   string
	Dim                    int
	DistanceMetric         string
	Compression            string
	ConstructionWindowSize int
	GraphMaxDegree         int
	SearchWindowSize       int
	Epsilon                float64
	TrainingThreshold      int
	ReduceDim              int
}

type FTDropIndexOptions struct {
	DeleteDocs bool
}

type SpellCheckTerms struct {
	Include    bool
	Exclude    bool
	Dictionary string
}

type FTExplainOptions struct {
	Dialect string
}

type FTSynUpdateOptions struct {
	SkipInitialScan bool
}

type SearchAggregator int

const (
	SearchInvalid = SearchAggregator(iota)
	SearchAvg
	SearchSum
	SearchMin
	SearchMax
	SearchCount
	SearchCountDistinct
	SearchCountDistinctish
	SearchStdDev
	SearchQuantile
	SearchToList
	SearchFirstValue
	SearchRandomSample
)

func (a SearchAggregator) String() string { _ = "STUB: not implemented"; return "" }

type SearchFieldType int

const (
	SearchFieldTypeInvalid = SearchFieldType(iota)
	SearchFieldTypeNumeric
	SearchFieldTypeTag
	SearchFieldTypeText
	SearchFieldTypeGeo
	SearchFieldTypeVector
	SearchFieldTypeGeoShape
)

func (t SearchFieldType) String() string { _ = "STUB: not implemented"; return "" }

type FTAggregateReducer struct {
	Reducer SearchAggregator
	Args    []interface{}
	As      string
}

type FTAggregateGroupBy struct {
	Fields []interface{}
	Reduce []FTAggregateReducer
}

type FTAggregateSortBy struct {
	FieldName string
	Asc       bool
	Desc      bool
}

type FTAggregateApply struct {
	Field string
	As    string
}

type FTAggregateLoad struct {
	Field string
	As    string
}

type FTAggregateWithCursor struct {
	Count   int
	MaxIdle int
}

type FTAggregateSortByStep struct {
	Fields []FTAggregateSortBy
	Max    int
}

type FTAggregateStep struct {
	Load    *FTAggregateLoad
	Apply   *FTAggregateApply
	GroupBy *FTAggregateGroupBy
	SortBy  *FTAggregateSortByStep
}

type FTAggregateOptions struct {
	Verbatim bool
	LoadAll  bool
	Timeout  int

	Scorer string

	AddScores bool

	Steps []FTAggregateStep

	LimitOffset       int
	Limit             int
	Filter            string
	WithCursor        bool
	WithCursorOptions *FTAggregateWithCursor
	Params            map[string]interface{}

	DialectVersion int

	Load []FTAggregateLoad

	GroupBy []FTAggregateGroupBy

	SortBy []FTAggregateSortBy

	SortByMax int

	Apply []FTAggregateApply
}

type FTSearchFilter struct {
	FieldName interface{}
	Min       interface{}
	Max       interface{}
}

type FTSearchGeoFilter struct {
	FieldName string
	Longitude float64
	Latitude  float64
	Radius    float64
	Unit      string
}

type FTSearchReturn struct {
	FieldName string
	As        string
}

type FTSearchSortBy struct {
	FieldName string
	Asc       bool
	Desc      bool
}

type FTSearchOptions struct {
	NoContent    bool
	Verbatim     bool
	NoStopWords  bool
	WithScores   bool
	WithPayloads bool
	WithSortKeys bool
	Filters      []FTSearchFilter
	GeoFilter    []FTSearchGeoFilter
	InKeys       []interface{}
	InFields     []interface{}
	Return       []FTSearchReturn
	Slop         int
	Timeout      int
	InOrder      bool
	Language     string
	Expander     string

	Scorer          string
	ExplainScore    bool
	Payload         string
	SortBy          []FTSearchSortBy
	SortByWithCount bool
	LimitOffset     int
	Limit           int

	CountOnly bool
	Params    map[string]interface{}

	DialectVersion int
}

type FTHybridCombineMethod string

const (
	FTHybridCombineRRF      FTHybridCombineMethod = "RRF"
	FTHybridCombineLinear   FTHybridCombineMethod = "LINEAR"
	FTHybridCombineFunction FTHybridCombineMethod = "FUNCTION"
)

type FTHybridSearchExpression struct {
	Query        string
	Scorer       string
	ScorerParams []interface{}
	YieldScoreAs string
}

type FTHybridVectorMethod = string

const (
	KNN   FTHybridCombineMethod = "KNN"
	RANGE FTHybridCombineMethod = "RANGE"
)

type FTHybridVectorExpression struct {
	VectorField string
	VectorData  Vector

	VectorParamName string
	Method          FTHybridVectorMethod
	MethodParams    []interface{}

	ShardKRatio  float64
	Filter       string
	YieldScoreAs string
}

type FTHybridCombineOptions struct {
	Method       FTHybridCombineMethod
	Count        int
	Window       int
	Constant     float64
	Alpha        float64
	Beta         float64
	YieldScoreAs string
}

type FTHybridGroupBy struct {
	Count        int
	Fields       []string
	ReduceFunc   string
	ReduceCount  int
	ReduceParams []interface{}
}

type FTHybridApply struct {
	Expression string
	AsField    string
}

type FTHybridWithCursor struct {
	Count   int
	MaxIdle int
}

type FTHybridOptions struct {
	CountExpressions  int
	SearchExpressions []FTHybridSearchExpression
	VectorExpressions []FTHybridVectorExpression
	Combine           *FTHybridCombineOptions
	Load              []string
	GroupBy           *FTHybridGroupBy
	Apply             []FTHybridApply
	SortBy            []FTSearchSortBy
	Filter            string
	LimitOffset       int
	Limit             int
	Params            map[string]interface{}
	ExplainScore      bool
	Timeout           int
	WithCursor        bool
	WithCursorOptions *FTHybridWithCursor
}

type FTSynDumpResult struct {
	Term     string
	Synonyms []string
}

type FTSynDumpCmd struct {
	baseCmd
	val []FTSynDumpResult
}

type FTAggregateResult struct {
	Total    int
	Rows     []AggregateRow
	Warnings []string
}

type AggregateRow struct {
	Fields map[string]interface{}
}

type AggregateCmd struct {
	baseCmd
	val *FTAggregateResult
}

type FTInfoResult struct {
	IndexErrors              IndexErrors
	Attributes               []FTAttribute
	BytesPerRecordAvg        string
	Cleaning                 int
	CursorStats              CursorStats
	DialectStats             map[string]int
	DocTableSizeMB           float64
	FieldStatistics          []FieldStatistic
	GCStats                  GCStats
	GeoshapesSzMB            float64
	HashIndexingFailures     int
	IndexDefinition          IndexDefinition
	IndexName                string
	IndexOptions             []string
	Indexing                 int
	InvertedSzMB             float64
	KeyTableSizeMB           float64
	MaxDocID                 int
	NumDocs                  int
	NumRecords               int
	NumTerms                 int
	NumberOfUses             int
	OffsetBitsPerRecordAvg   string
	OffsetVectorsSzMB        float64
	OffsetsPerTermAvg        string
	PercentIndexed           float64
	RecordsPerDocAvg         string
	SortableValuesSizeMB     float64
	TagOverheadSzMB          float64
	TextOverheadSzMB         float64
	TotalIndexMemorySzMB     float64
	TotalIndexingTime        int
	TotalInvertedIndexBlocks int
	VectorIndexSzMB          float64
}

type IndexErrors struct {
	IndexingFailures     int
	LastIndexingError    string
	LastIndexingErrorKey string
}

type FTAttribute struct {
	Identifier      string
	Attribute       string
	Type            string
	Weight          float64
	Sortable        bool
	NoStem          bool
	NoIndex         bool
	UNF             bool
	PhoneticMatcher string
	CaseSensitive   bool
	WithSuffixtrie  bool

	Algorithm      string
	DataType       string
	Dim            int
	DistanceMetric string
	M              int
	EFConstruction int
}

type CursorStats struct {
	GlobalIdle    int
	GlobalTotal   int
	IndexCapacity int
	IndexTotal    int
}

type FieldStatistic struct {
	Identifier  string
	Attribute   string
	IndexErrors IndexErrors
}

type GCStats struct {
	BytesCollected       int
	TotalMsRun           int
	TotalCycles          int
	AverageCycleTimeMs   string
	LastRunTimeMs        int
	GCNumericTreesMissed int
	GCBlocksDenied       int
}

type IndexDefinition struct {
	KeyType      string
	Prefixes     []string
	DefaultScore float64
}

type FTSpellCheckOptions struct {
	Distance int
	Terms    *FTSpellCheckTerms

	Dialect int
}

type FTSpellCheckTerms struct {
	Inclusion  string
	Dictionary string
	Terms      []interface{}
}

type SpellCheckResult struct {
	Term        string
	Suggestions []SpellCheckSuggestion
}

type SpellCheckSuggestion struct {
	Score      float64
	Suggestion string
}

type FTSearchResult struct {
	Total    int
	Docs     []Document
	Warnings []string
}

type Document struct {
	ID      string
	Score   *float64
	Payload *string
	SortKey *string
	Fields  map[string]string
	Error   error
}

type AggregateQuery []interface{}

func (c cmdable) FT_List(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTAggregate(ctx context.Context, index string, query string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func validateFTAggregateOptions(options *FTAggregateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func appendFTAggregateStep(args []interface{}, step FTAggregateStep) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FTAggregateQuery(query string, options *FTAggregateOptions) (AggregateQuery, error) {
	_ = "STUB: not implemented"
	return *new(AggregateQuery), nil
}

func ProcessAggregateResult(data []interface{}) (*FTAggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAggregateCmd(ctx context.Context, args ...interface{}) *AggregateCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *AggregateCmd) SetVal(val *FTAggregateResult) { _ = "STUB: not implemented"; return }

func (cmd *AggregateCmd) Val() *FTAggregateResult { _ = "STUB: not implemented"; return nil }

func (cmd *AggregateCmd) Result() (*FTAggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *AggregateCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *AggregateCmd) RawResult() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *AggregateCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *AggregateCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFTAggregateMapRESP3(data map[interface{}]interface{}) (*FTAggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFTAggregateResultsMapRESP3(resultsData []interface{}) ([]AggregateRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFTAggregateRowMapRESP3(itemMap map[interface{}]interface{}) (AggregateRow, error) {
	_ = "STUB: not implemented"
	return *new(AggregateRow), nil
}

func (cmd *AggregateCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) FTAggregateWithArgs(ctx context.Context, index string, query string, options *FTAggregateOptions) *AggregateCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTAliasAdd(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTAliasDel(ctx context.Context, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTAliasUpdate(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTConfigGet(ctx context.Context, option string) *MapMapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTConfigSet(ctx context.Context, option string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTCreate(ctx context.Context, index string, options *FTCreateOptions, schema ...*FieldSchema) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTCursorDel(ctx context.Context, index string, cursorId int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTDictAdd(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTDictDel(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTDictDump(ctx context.Context, dict string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTDropIndex(ctx context.Context, index string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTDropIndexWithArgs(ctx context.Context, index string, options *FTDropIndexOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTExplain(ctx context.Context, index string, query string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTExplainWithArgs(ctx context.Context, index string, query string, options *FTExplainOptions) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTExplainCli(ctx context.Context, key, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseFTAttributeFromMap(attrMap map[interface{}]interface{}) FTAttribute {
	_ = "STUB: not implemented"
	return *new(FTAttribute)
}

func getMapStringKey(m map[interface{}]interface{}, key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func parseIndexErrorsRESP3(m map[interface{}]interface{}) IndexErrors {
	_ = "STUB: not implemented"
	return *new(IndexErrors)
}

func parseCursorStatsRESP3(m map[interface{}]interface{}) CursorStats {
	_ = "STUB: not implemented"
	return *new(CursorStats)
}

func parseGCStatsRESP3(m map[interface{}]interface{}) GCStats {
	_ = "STUB: not implemented"
	return *new(GCStats)
}

func parseIndexDefinitionRESP3(m map[interface{}]interface{}) IndexDefinition {
	_ = "STUB: not implemented"
	return *new(IndexDefinition)
}

func parseDialectStatsRESP3(m map[interface{}]interface{}) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func ftInfoNumString(val interface{}) string { _ = "STUB: not implemented"; return "" }

func ftInfoNumInt(val interface{}) int { _ = "STUB: not implemented"; return 0 }

func parseFTInfo(data map[string]interface{}) (FTInfoResult, error) {
	_ = "STUB: not implemented"
	return *new(FTInfoResult), nil
}

type FTInfoCmd struct {
	baseCmd
	val FTInfoResult
}

func newFTInfoCmd(ctx context.Context, args ...interface{}) *FTInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTInfoCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FTInfoCmd) SetVal(val FTInfoResult) { _ = "STUB: not implemented"; return }

func (cmd *FTInfoCmd) Result() (FTInfoResult, error) {
	_ = "STUB: not implemented"
	return *new(FTInfoResult), nil
}

func (cmd *FTInfoCmd) Val() FTInfoResult { _ = "STUB: not implemented"; return *new(FTInfoResult) }

func (cmd *FTInfoCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *FTInfoCmd) RawResult() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *FTInfoCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTInfoCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) FTInfo(ctx context.Context, index string) *FTInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTSpellCheck(ctx context.Context, index string, query string) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *FTSpellCheckOptions) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

type FTSpellCheckCmd struct {
	baseCmd
	val []SpellCheckResult
}

func newFTSpellCheckCmd(ctx context.Context, args ...interface{}) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTSpellCheckCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FTSpellCheckCmd) SetVal(val []SpellCheckResult) { _ = "STUB: not implemented"; return }

func (cmd *FTSpellCheckCmd) Result() ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSpellCheckCmd) Val() []SpellCheckResult { _ = "STUB: not implemented"; return nil }

func (cmd *FTSpellCheckCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *FTSpellCheckCmd) RawResult() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSpellCheckCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFTSpellCheckRESP3(data map[interface{}]interface{}) ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFTSpellCheck(data []interface{}) ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSpellCheckCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func parseFTSearch(data []interface{}, noContent, withScores, withPayloads, withSortKeys bool) (FTSearchResult, error) {
	_ = "STUB: not implemented"
	return *new(FTSearchResult), nil
}

type FTSearchCmd struct {
	baseCmd
	val     FTSearchResult
	options *FTSearchOptions
}

func newFTSearchCmd(ctx context.Context, options *FTSearchOptions, args ...interface{}) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTSearchCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FTSearchCmd) SetVal(val FTSearchResult) { _ = "STUB: not implemented"; return }

func (cmd *FTSearchCmd) Result() (FTSearchResult, error) {
	_ = "STUB: not implemented"
	return *new(FTSearchResult), nil
}

func (cmd *FTSearchCmd) Val() FTSearchResult {
	_ = "STUB: not implemented"
	return *new(FTSearchResult)
}

func (cmd *FTSearchCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *FTSearchCmd) RawResult() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSearchCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFTSearchMapRESP3(data map[interface{}]interface{}) (FTSearchResult, error) {
	_ = "STUB: not implemented"
	return *new(FTSearchResult), nil
}

func parseFTSearchResultsMapRESP3(resultsData []interface{}) ([]Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFTSearchDocumentMapRESP3(itemMap map[interface{}]interface{}) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

func (cmd *FTSearchCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

type FTHybridResult struct {
	TotalResults  int
	Results       []map[string]interface{}
	Warnings      []string
	ExecutionTime float64
}

type FTHybridCursorResult struct {
	SearchCursorID int
	VsimCursorID   int
}

type FTHybridCmd struct {
	baseCmd
	val        FTHybridResult
	cursorVal  *FTHybridCursorResult
	options    *FTHybridOptions
	withCursor bool
}

func newFTHybridCmd(ctx context.Context, options *FTHybridOptions, args ...interface{}) *FTHybridCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTHybridCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FTHybridCmd) SetVal(val FTHybridResult) { _ = "STUB: not implemented"; return }

func (cmd *FTHybridCmd) Result() (FTHybridResult, error) {
	_ = "STUB: not implemented"
	return *new(FTHybridResult), nil
}

func (cmd *FTHybridCmd) CursorResult() (*FTHybridCursorResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTHybridCmd) Val() FTHybridResult {
	_ = "STUB: not implemented"
	return *new(FTHybridResult)
}

func (cmd *FTHybridCmd) CursorVal() *FTHybridCursorResult { _ = "STUB: not implemented"; return nil }

func (cmd *FTHybridCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *FTHybridCmd) RawResult() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFTHybrid(data []interface{}, withCursor bool) (FTHybridResult, *FTHybridCursorResult, error) {
	_ = "STUB: not implemented"
	return *new(FTHybridResult), nil, nil
}

func (cmd *FTHybridCmd) readReply(rd *proto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTHybridCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) FTSearch(ctx context.Context, index string, query string) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

type SearchQuery []interface{}

func FTSearchQuery(query string, options *FTSearchOptions) (SearchQuery, error) {
	_ = "STUB: not implemented"
	return *new(SearchQuery), nil
}

func (c cmdable) FTSearchWithArgs(ctx context.Context, index string, query string, options *FTSearchOptions) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

func NewFTSynDumpCmd(ctx context.Context, args ...interface{}) *FTSynDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FTSynDumpCmd) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *FTSynDumpCmd) SetVal(val []FTSynDumpResult) { _ = "STUB: not implemented"; return }

func (cmd *FTSynDumpCmd) Val() []FTSynDumpResult { _ = "STUB: not implemented"; return nil }

func (cmd *FTSynDumpCmd) Result() ([]FTSynDumpResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSynDumpCmd) RawVal() interface{} { _ = "STUB: not implemented"; return nil }

func (cmd *FTSynDumpCmd) RawResult() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSynDumpCmd) readReply(rd *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func parseFTSynDumpRESP3(data map[interface{}]interface{}) ([]FTSynDumpResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSynDumpCmd) Clone() Cmder { _ = "STUB: not implemented"; return *new(Cmder) }

func (c cmdable) FTSynDump(ctx context.Context, index string) *FTSynDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *FTSynUpdateOptions, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTTagVals(ctx context.Context, index string, field string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c cmdable) FTHybrid(ctx context.Context, index string, searchExpr string, vectorField string, vectorData Vector) *FTHybridCmd {
	_ = "STUB: not implemented"
	return nil
}

func hybridVectorBlob(v Vector) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func hybridVectorBytes(blob []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func generateVectorParamName(params map[string]interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (c cmdable) FTHybridWithArgs(ctx context.Context, index string, options *FTHybridOptions) *FTHybridCmd {
	_ = "STUB: not implemented"
	return nil
}
