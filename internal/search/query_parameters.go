package search

import (
	"strconv"
)

type QueryParameters interface {
	GetLimit() uint64
	GetOffset() uint64
	GetParameters() map[string]interface{}
}

func NewQueryParameters(values map[string][]string) QueryParameters {
	qp := queryParameters{
		limit:  10,
		offset: 0,
		values: values,
	}

	if val := values["limit"]; len(val) != 0 {
		if limit, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			qp.limit = limit
		}
	}

	if val := values["offset"]; len(val) != 0 {
		if offset, err := strconv.ParseUint(val[0], 10, 64); err == nil {
			qp.offset = offset
		}
	}

	return &qp
}

type queryParameters struct {
	limit  uint64
	offset uint64
	values map[string][]string
}

func (rp *queryParameters) GetParameters() map[string]interface{} {
	params := make(map[string]interface{})
	for key, val := range rp.values {

		if len(val) > 1 {
			params[key] = val
		}

		if len(val) == 1 {
			params[key] = val[0]
		}

	}
	return params
}

func (rp *queryParameters) GetLimit() uint64 {
	return rp.limit
}

func (rp *queryParameters) GetOffset() uint64 {
	return rp.offset
}
