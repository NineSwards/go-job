package util

import (
	"os"
	"testing"
)
import "github.com/NineSwards/go-job/core/util"

func TestGetEnv(t *testing.T) {
	home := util.GetEnv("HOME")
	t.Log(home)
	home, ok := os.LookupEnv("HOME")
	t.Log(home, ok)
}
