package matchers

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/mniak/Alkanoid/domain"
)

type (
	_InlineMatcher struct {
		gomock.Matcher
		expressionName string
		getField       ExpressionFieldFunc
	}
	ExpressionFieldFunc func(x interface{}) interface{}
)

func ExpressionEquals(fieldName string, expected interface{}, fn ExpressionFieldFunc) gomock.Matcher {
	return &_InlineMatcher{
		Matcher:        gomock.Eq(expected),
		expressionName: fieldName,
		getField:       fn,
	}
}

func (m *_InlineMatcher) Matches(x interface{}) bool {
	fieldValue := m.getField(x)
	return m.Matcher.Matches(fieldValue)
}

func (m *_InlineMatcher) String() string {
	return fmt.Sprintf("%s %s", m.expressionName, m.Matcher.String())
}

func TransactionFieldEquals(fieldName string, expected interface{}, getField func(domain.Transaction) interface{}) gomock.Matcher {
	return ExpressionEquals(fieldName, expected, func(x interface{}) interface{} {
		return getField(x.(domain.Transaction))
	})
}
