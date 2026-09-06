package redis

type FTAggregateCollect struct {
	FieldsAll bool

	Fields []string

	Distinct bool

	SortBy []FTAggregateSortBy

	Limit *FTAggregateCollectLimit

	As string
}

type FTAggregateCollectLimit struct {
	Offset int
	Count  int
}

func ensureAtPrefix(name string) string { _ = "STUB: not implemented"; return "" }

func buildCollectArgs(o FTAggregateCollect) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCollectReducer(o FTAggregateCollect) (FTAggregateReducer, error) {
	_ = "STUB: not implemented"
	return *new(FTAggregateReducer), nil
}

type CollectEntry = map[string]interface{}

type CollectColumn = []CollectEntry

func (r AggregateRow) Collect(alias string) (CollectColumn, error) {
	_ = "STUB: not implemented"
	return *new(CollectColumn), nil
}

func parseCollectValue(v interface{}) (CollectColumn, error) {
	_ = "STUB: not implemented"
	return *new(CollectColumn), nil
}

func parseCollectEntry(e interface{}) (CollectEntry, error) {
	_ = "STUB: not implemented"
	return *new(CollectEntry), nil
}
