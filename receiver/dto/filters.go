package dto

type Filters struct {
	Page                int64
	PageSize            int64
	Match               string
	DateGeneratedBefore int64
	DateGeneratedAfter  int64
	LengthLess          int64
	LengthGreater       int64
}
