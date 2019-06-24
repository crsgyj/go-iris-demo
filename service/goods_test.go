package service

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func TestDateScore(t *testing.T) {
	score := dateScore(time.Now())

	result := assert.So(score, should.BeGreaterThan, 20190623 /*yesterday*/)

	if result.Error() != nil {
		t.Fail()
	}
}
