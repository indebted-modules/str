package str_test

import (
	"testing"

	"github.com/indebted-modules/str"
	"github.com/stretchr/testify/suite"
)

type StringsSuite struct {
	suite.Suite
}

func TestStringsSuite(t *testing.T) {
	suite.Run(t, new(StringsSuite))
}

func (s *StringsSuite) TestContains() {
	collection := []string{"A", "B"}

	s.True(str.Contains(collection, "A"))
	s.True(str.Contains(collection, "B"))

	s.False(str.Contains(collection, "C"))
	s.False(str.Contains(collection, ""))
	s.False(str.Contains([]string{}, "A"))
	s.False(str.Contains([]string{}, ""))
}

func (s *StringsSuite) TestJoinIgnoreEmpty() {
	s.Equal("", str.JoinIgnoreEmpty(" ", "", "", ""))
	s.Equal("Jane", str.JoinIgnoreEmpty(" ", "Jane", "", ""))
	s.Equal("Jane Jackson", str.JoinIgnoreEmpty(" ", "Jane", "Jackson", ""))
	s.Equal("Jane Jackson Doe", str.JoinIgnoreEmpty(" ", "Jane", "Jackson", "Doe"))
	s.Equal("Jane Doe", str.JoinIgnoreEmpty(" ", "Jane", "", "Doe"))
	s.Equal("Doe", str.JoinIgnoreEmpty(" ", "", "", "Doe"))
	s.Equal("Jackson Doe", str.JoinIgnoreEmpty(" ", "", "Jackson", "Doe"))
	s.Equal("Jackson", str.JoinIgnoreEmpty(" ", "", "Jackson", ""))
	s.Equal("Jane-Jackson-Doe", str.JoinIgnoreEmpty("-", "Jane", "Jackson", "Doe"))
}
