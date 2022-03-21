package util

import (
	"github.com/NineSwards/go-job/util"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	home := util.GetEnv("HOME")
	t.Log(home)
	home, ok := os.LookupEnv("HOME")
	t.Log(home, ok)
}
