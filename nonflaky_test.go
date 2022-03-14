package flakygo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type nonflakyTestSuite struct {
	suite.Suite
}

func Test_nonflakyTestSuite(t *testing.T) {
	suite.Run(t, new(nonflakyTestSuite))
}

func (f *nonflakyTestSuite) Test_nonflaky5() {
	f.Assert().Equal(25, nonflaky(5))
}

func (f *nonflakyTestSuite) Test_nonflaky10() {
	f.Assert().Equal(100, nonflaky(10))
}

func (f *nonflakyTestSuite) Test_nonflaky100() {
	f.Assert().Equal(10000, nonflaky(100))
}
