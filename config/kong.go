package config

import (
	"micro-gin/utils"
)

type Kong struct {
	Protocol      string `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	Host          string `mapstructure:"host" json:"host" yaml:"host"`
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	UpstreamsPath string `mapstructure:"upstreams_path" json:"upstreams_path" yaml:"upstreams_path"`
	targetsPath   string `mapstructure:"targets_path" json:"targets_path" yaml:"targets_path"`
	servicesPath  string `mapstructure:"services_path" json:"services_path" yaml:"services_path"`
	routesPath    string `mapstructure:"routes_path" json:"routes_path" yaml:"routes_path"`
}

//TODO 完善下面方法

func (k *Kong) url() string {
	return k.Protocol + "://" + k.Host + ":" + k.Port
}
func (k *Kong) AddUpstreams(upstreamName string) {
	//url := k.url() + k.UpstreamsPath
	//params := map[string]string{
	//	"name": upstreamName,
	//}
	//result,err := utils.HttpDo(url, "GET", params, nil, nil)
	//if err != nil{
	//	global.App.Log.Error(err.Error())
	//	return err
	//}
}

func (k *Kong) AddTarget(upstreamName, target string, weigth int) {

}

func (k *Kong) AddServs() {

}

func (k *Kong) GetServs(upstreamName string) (interface{}, error) {
	url := k.url() + k.servicesPath
	params := map[string]string{
		"name": upstreamName,
	}
	result, err := utils.HttpDo(url, "GET", params, nil, nil)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (k *Kong) AddRouters() {

}
