package matchers

import (
	"strings"

	"github.com/golang/mock/gomock"
)

type (
	_InlineMatcher struct {
		pred     Predicate
		expected string
	}
	Predicate func(x interface{}) bool
)

func InlineMatcher(pred Predicate, expected ...string) gomock.Matcher {
	return &_InlineMatcher{
		pred:     pred,
		expected: strings.Join(expected, "; "),
	}
}

func (m *_InlineMatcher) Matches(x interface{}) bool {
	return false
}

func (m *_InlineMatcher) String() string {
	return "<my matcher>"
}
