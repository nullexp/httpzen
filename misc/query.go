package misc

import (
	"errors"
	"strings"
)

type QueryDefinition interface {
	GetName() string

	GetSupportedOperators() []QueryOperator

	GetType() DataType

	GetDescription() string
}

type queryDefinition struct {
	name string

	supportOperators []QueryOperator

	dataType DataType

	description string
}

func NewQueryDefinition(name string, sup []QueryOperator, typ DataType) queryDefinition {
	return queryDefinition{name: name, supportOperators: sup, dataType: typ}
}

func NewQueryDefinitionWithDescription(name string, sup []QueryOperator, typ DataType, desc string) queryDefinition {
	return queryDefinition{name: name, supportOperators: sup, dataType: typ, description: desc}
}

func (q queryDefinition) GetName() string {
	return q.name
}

func (q queryDefinition) GetSupportedOperators() []QueryOperator {
	return q.supportOperators
}

func (q queryDefinition) GetType() DataType {
	return q.dataType
}

func (q queryDefinition) GetDescription() string {
	return q.description
}

type Operand struct {
	Value any
}

func NewOperand(value any) *Operand { return &Operand{Value: value} }

type Query interface {
	GetName() string

	SetName(string)

	GetOperator() QueryOperator

	GetOperand() *Operand
}

type QueryInfo interface {
	GetQuery() []Query

	IsAnd() bool
}

func NewQueryInfo(queries []Query, isAnd bool) QueryInfo {
	return basicQueryInfo{queries: queries, isAnd: isAnd}
}

func NewSimpleAndQueryInfo(name string, op QueryOperator, model any, operand *Operand) QueryInfo {
	return NewQueryInfo([]Query{NewQuery(name, op, operand)}, true)
}

func NewMultipleAndQueryInfo(queries ...Query) QueryInfo {
	return NewQueryInfo(queries, true)
}

type basicQueryInfo struct {
	queries []Query

	isAnd bool
}

func (b basicQueryInfo) GetQuery() []Query {
	return b.queries
}

func (b basicQueryInfo) IsAnd() bool {
	return b.isAnd
}

type basicQuery struct {
	name string

	op QueryOperator

	operand *Operand
}

func NewQuery(name string, op QueryOperator, operand *Operand) Query {
	return &basicQuery{name: name, op: op, operand: operand}
}

func NewEmptyQuery(name string, op QueryOperator) Query {
	return &basicQuery{name: name, op: op}
}

func (b basicQuery) GetName() string {
	return b.name
}

func (b *basicQuery) SetName(name string) {
	b.name = name
}

func (b basicQuery) GetOperator() QueryOperator {
	return b.op
}

func (b basicQuery) GetOperand() *Operand {
	return b.operand
}

type ParameterExpectation int

const (
	ParameterExpectationZero ParameterExpectation = iota

	ParameterExpectationSingle

	ParameterExpectationMultiple
)

func ParseQueryOperator(op string) (QueryOperator, ParameterExpectation, error) {
	switch QueryOperator(op) {

	case QueryOperatorEqual:

		return QueryOperatorEqual, ParameterExpectationSingle, nil

	case QueryOperatorNotEqual:

		return QueryOperatorNotEqual, ParameterExpectationSingle, nil

	case QueryOperatorMoreThan:

		return QueryOperatorMoreThan, ParameterExpectationSingle, nil

	case QueryOperatorEqualOrMoreThan:

		return QueryOperatorEqualOrMoreThan, ParameterExpectationSingle, nil

	case QueryOperatorLessThan:

		return QueryOperatorLessThan, ParameterExpectationSingle, nil

	case QueryOperatorEqualOrLessThan:

		return QueryOperatorEqualOrLessThan, ParameterExpectationSingle, nil

	case QueryOperatorContain:

		return QueryOperatorContain, ParameterExpectationMultiple, nil

	case QueryOperatorNotContain:

		return QueryOperatorNotContain, ParameterExpectationMultiple, nil

	case QueryOperatorEmpty:

		return QueryOperatorEmpty, ParameterExpectationZero, nil

	case QueryOperatorNotEmpty:

		return QueryOperatorNotEmpty, ParameterExpectationZero, nil

	}

	return "", 0, errors.New("unknown operator")
}

var ReservedQueryMap = map[string]bool{
	QueryLimit: true, QuerySkip: true, QuerySort: true,

	QuerySearch: true, QueryDefault: true,

	QueryFields: true, QueryEmbed: true, QueryOr: true,
}

const (
	QueryLimit = "limit"

	QuerySkip = "skip"

	QuerySort = "sort"

	QueryOr = "or"

	QuerySearch = "search"

	QueryDefault = "q"

	QueryFields = "fields"

	QueryEmbed = "embed"
)

func GenerateDefaultOrQuery(query string, fields ...string) QueryInfo {
	queries := []Query{}

	parts := strings.Split(query, " ")

	for _, p := range parts {

		trimed := strings.Trim(p, " ")

		if trimed == "" {
			continue
		}

		for _, f := range fields {
			queries = append(queries, NewQuery(f, QueryOperatorEqual, NewOperand(trimed)))
		}

	}

	return NewQueryInfo(queries, false)
}
