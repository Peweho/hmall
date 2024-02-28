package test

import (
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	var a error
	threading.GoSafe(func() {
		a = status.Error(codes.Aborted, dtmcli.ResultFailure)
	})
	time.Sleep(time.Second)
	fmt.Println(a)
}
