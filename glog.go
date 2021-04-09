package glog

type GLog struct {
}

func NewLogger(
	sdkType SDKType,
	andLiveConfig conf.AndLive,
	tencentyunConfig conf.Tencentyun,
	resty *resty.Request,
) (live Live) {
	switch sdkType {
	case SDKTypeAndLive:
		live = and.NewLive(andLiveConfig, resty)
	case SDKTypeTencentyun:
		live = tencentyun.NewLive(tencentyunConfig)
	}

	return
}
