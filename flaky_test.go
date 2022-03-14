package flakygo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type flakyTestSuite struct {
	suite.Suite
}

func Test_flakyTestSuite(t *testing.T) {
	suite.Run(t, new(flakyTestSuite))
}

func (f *flakyTestSuite) Test_flaky5() {
	f.Assert().False(flaky(5))
}

func (f *flakyTestSuite) Test_flaky10() {
	f.Assert().False(flaky(10))
}

func (f *flakyTestSuite) Test_flaky100() {
	f.Assert().False(flaky(100))
}
