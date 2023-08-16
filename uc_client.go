package userclient

import (
	"dghire.com/libs/go-common/cmodel"
	"dghire.com/libs/go-common/constants"
	dgctx "dghire.com/libs/go-common/context"
	"dghire.com/libs/go-common/result"
	dglogger "dghire.com/libs/go-logger"
	"dghire.com/libs/go-uc-sdk/utils"
	"errors"
	"net/http"
	"strconv"
)

var httpClient *http.Client
var baseUrl string

func getUserInfoUrl(uid int64) string {
	return baseUrl + "/uc/internal/user/info" + "?" + "uid=" + strconv.FormatInt(uid, 10)
}

func getUserInfosUrl(uids []int64) string {
	return baseUrl + "/uc/internal/user/infos" + "?" + "uids=" + utils.JoinInt64Arr(uids)
}

func InitUcClient(hc *http.Client, ucHost string, ucPort int) {
	if hc == nil {
		panic("init uc client blank httpClient")
	}
	if len(ucHost) == 0 {
		panic("init uc client blank ucHost")
	}
	if ucPort <= 0 {
		ucPort = 8080
	}
	httpClient = hc
	baseUrl = "http://" + ucHost + ":" + strconv.Itoa(ucPort)
}

func GetUserInfo(dc *dgctx.DgContext, uid int64) (*UserInfo, error) {
	dglogger.Debugf(dc, "get user info by uid: %d", uid)

	req, err := http.NewRequest("GET", getUserInfoUrl(uid), nil)
	if err != nil {
		dglogger.Errorf(dc, "new getUserInfo request err: %v", err)
		return nil, err
	}

	req.Header.Set(constants.TraceId, dc.TraceId)

	res, err := httpClient.Do(req)
	if err != nil {
		dglogger.Errorf(dc, "http client request getUserInfo err: %v", err)
		return nil, err
	}

	ret := new(result.Result[*UserInfo])
	err = utils.ResToObj(res, ret)
	if err != nil {
		dglogger.Errorf(dc, "res to obj err: %v", err)
		return nil, err
	}
	if ret.Success {
		return ret.Data, nil
	}
	return nil, errors.New(ret.Message)
}

func GetUserInfos(dc *dgctx.DgContext, uids []int64) ([]*UserInfo, error) {
	dglogger.Debugf(dc, "get user infos by: %v", uids)
	if len(uids) == 0 {
		return []*UserInfo{}, nil
	}

	req, err := http.NewRequest("GET", getUserInfosUrl(uids), nil)
	if err != nil {
		dglogger.Errorf(dc, "new getUserInfos request err: %v", err)
		return nil, err
	}

	req.Header.Set(constants.TraceId, dc.TraceId)

	res, err := httpClient.Do(req)
	if err != nil {
		dglogger.Errorf(dc, "http client request getUserInfos err: %v", err)
		return nil, err
	}

	ret := new(result.Result[[]*UserInfo])
	err = utils.ResToObj(res, ret)
	if err != nil {
		dglogger.Errorf(dc, "res to obj err: %v", err)
		return nil, err
	}
	if ret.Success {
		return ret.Data, nil
	}
	return nil, errors.New(ret.Message)
}

func GetUserIdNameMap(dc *dgctx.DgContext, uids []int64) (map[int64]string, error) {
	dglogger.Debugf(dc, "get user id name map by: %v", uids)

	if len(uids) == 0 {
		return map[int64]string{}, nil
	}

	infos, err := GetUserInfos(dc, uids)
	if err != nil {
		dglogger.Errorf(dc, "get user infos err: %v", err)
		return nil, err
	}

	if len(infos) == 0 {
		return map[int64]string{}, nil
	}

	mp := make(map[int64]string)
	for _, info := range infos {
		mp[info.Uid] = info.Name
	}
	return mp, nil
}

func GetUserIdNamePairMap(dc *dgctx.DgContext, uids []int64) (map[int64]*cmodel.IdNamePair, error) {
	dglogger.Debugf(dc, "get user idNamePair map by: %v", uids)

	if len(uids) == 0 {
		return map[int64]*cmodel.IdNamePair{}, nil
	}

	infos, err := GetUserInfos(dc, uids)
	if err != nil {
		dglogger.Errorf(dc, "get user infos err: %v", err)
		return nil, err
	}

	if len(infos) == 0 {
		return map[int64]*cmodel.IdNamePair{}, nil
	}

	mp := make(map[int64]*cmodel.IdNamePair)
	for _, info := range infos {
		mp[info.Uid] = &cmodel.IdNamePair{
			Id:   info.Uid,
			Name: info.Name,
		}
	}
	return mp, nil
}
