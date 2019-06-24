package userutils

import (
	"errors"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	utils := &Utils{ctx: nil}
	limit := 30
	for i := 0; i < 1000; i++ {
		r := utils.RandInt(limit)
		if r > limit || r < 0 {
			t.Error(errors.New("r > limit || r < 0"))
			break
		}
	}
}

func TestInt64ToString(_ *testing.T) {
	t := time.Now().Unix() * 1000
	s := strconv.Itoa(int(t))
	log.Println(t)
	log.Println(s)
	// 1561376296333
	// 1561376453890194300
}

func TestRandString(t *testing.T) {
	utils := &Utils{ctx: nil}
	for i := 5; i < 30; i++ {
		randStr := utils.RandString(i)
		if len(randStr) != i {
			t.Error(errors.New("len(randStr) != i， len(randStr)：" + strconv.Itoa(len(randStr))))
			break
		}
	}

}
