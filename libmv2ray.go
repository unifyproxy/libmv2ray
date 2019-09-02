package libmv2ray

import (
	. "fmt"
	v2core "v2ray.com/core"
	"v2ray.com/core/common"
	"v2ray.com/core/common/platform"
	v2stats "v2ray.com/core/features/stats"
	v2ser "v2ray.com/core/infra/conf/serial"

	"strings"
	"sync"

	_ "v2ray.com/core/main/distro/all"
)

type V2Manager struct {
	sync.Mutex
	v2server v2core.Server
	v2manager v2stats.Manager

	main_chan chan struct{}

	config string
}


func (v2 *V2Manager) StartV2ray() {
	v2.main_chan = make(chan struct{})
	if v2.v2server == nil {
		config, err := v2ser.LoadJSONConfig(strings.NewReader(v2.config))
		common.Must(err)
		instance, err := v2core.New(config)
		common.Must(err)
		v2.v2manager = instance.GetFeature(v2stats.ManagerType()).(v2stats.Manager)
		v2.v2server = instance
	}
	err := v2.v2server.Start()
	if err != nil {
		println(err.Error())
	}

	go func () {
		select {
		case <-v2.main_chan:
			break
		}
	}()

	println(platform.GetConfigurationPath())
}


func (v2 *V2Manager)  GetNetSpeed(tag string) int64 {
	if v2.v2manager != nil {
		var str string
		arg := Sprintf("inbound>>>%s>>>traffic>>>%s", tag, str)
		c := v2.v2manager.GetCounter(arg)
		return c.Value()
	}
	return 0
}