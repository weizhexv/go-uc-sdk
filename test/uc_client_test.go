package test

import (
	dgctx "dghire.com/libs/go-common/context"
	userclient "dghire.com/libs/go-uc-sdk"
	"fmt"
	"net/http"
	"testing"
)

var setup = initUcClient()

func initUcClient() bool {
	userclient.InitUcClient(http.DefaultClient, "localhost", 9999)
	return true
}

func newDgCtx() *dgctx.DgContext {
	return &dgctx.DgContext{
		TraceId: "safjkslafss",
	}
}

func TestGetUserInfo(t *testing.T) {
	info, err := userclient.GetUserInfo(newDgCtx(), 38)
	if err != nil {
		t.Errorf("catch err: %v", err)
	}
	fmt.Printf("ret: %v", info)
}

func TestGetUserInfos(t *testing.T) {
	infos, err := userclient.GetUserInfos(newDgCtx(), []int64{1, 2, 24, 35, 36, 37})
	if err != nil {
		t.Errorf("catch err: %v", err)
	}
	fmt.Printf("ret: %v", infos)
}

func TestGetIdNameMap(t *testing.T) {
	mp, err := userclient.GetUserIdNameMap(newDgCtx(), []int64{1, 2, 24, 35, 36, 37})
	if err != nil {
		t.Errorf("catch err: %v", err)
	}
	fmt.Printf("ret: %v", mp)
}

func TestGetIdNamePairMap(t *testing.T) {
	mp, err := userclient.GetUserIdNamePairMap(newDgCtx(), []int64{1, 2, 24, 35, 36, 37})
	if err != nil {
		t.Errorf("catch err: %v", err)
	}
	fmt.Printf("ret: %v", mp)
}
