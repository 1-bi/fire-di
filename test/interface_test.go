package test

import (
	"github.com/1-bi/fire-di/test/fixture"
	"github.com/smartystreets/gunit"
	"testing"
)

// TestInterface_case1 test case message
func TestInterface_case1(t *testing.T) {

	gunit.Run(new(fixture.InterfaceFixTure), t)
}
