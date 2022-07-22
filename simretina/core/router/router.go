package router

import (
	"simretina/core/apis"
	"simretina/core/config"
)

func InitRouter(hostPort string) {
	var (
		login             *config.LoginMes
		deviceID          *config.DeviceID
		deviceStatus      *config.DeviceStatus
		algModelUpdate    *config.AlgModelUpdate
		algResultQuery    *config.AlgResultQuery
		algModelDelete    *config.AlgModelUpdate
		baseConfig        *config.BaseConfig
		deviceRestart     *config.DeviceRestart
		deviceShutDown    *config.DeviceRestart
		deviceStatusQuery *config.DeviceStatusQuery
		extend            *config.Extend
		functionDefine    *config.FunctionDefine
		imageQuery        *config.ImageQuery
		transcode         *config.Transcode
		videoQuery        *config.VideoQuery
		featureInfo       *config.FeatureInfo
		imageInfo         *config.ImageInfo
		modelQuery        *config.ModelQuery
		algStatus         *config.AlgStatus
		transcodeStatus   *config.TranscodeStatus
		ExtendData        *config.ExtendData
		videoSliceInfo    *config.VideoSliceInfo
	)
	//设备接入路由
	apis.Login(hostPort, login)
	apis.Logout(hostPort, deviceID)
	apis.Keepalive(hostPort, deviceID)
	// deviceID = new(hostPort, config.DeviceID)
	// apis.GetPushStreamAddress(hostPort, deviceID)
	apis.DeviceStatus(hostPort, deviceStatus)
	//协议下发指令路由
	apis.AlgModelUpdate(hostPort, algModelUpdate)
	apis.AlgResultQuery(hostPort, algResultQuery)
	apis.AlgModelDelete(hostPort, algModelDelete)
	apis.BaseConfig(hostPort, baseConfig)
	apis.DeviceRestart(hostPort, deviceRestart)
	apis.DeviceShutDown(hostPort, deviceShutDown)
	apis.DeviceStatusQuery(hostPort, deviceStatusQuery)
	apis.Extend(hostPort, extend)
	apis.FunctionDefine(hostPort, functionDefine)
	apis.ImageQuery(hostPort, imageQuery)
	apis.Transcode(hostPort, transcode)
	apis.VideoQuery(hostPort, videoQuery)
	//数据上报路由
	apis.Data(hostPort)
	apis.FeatureInfo(hostPort, featureInfo)
	apis.ImageInfo(hostPort, imageInfo)
	apis.ImagesData(hostPort)
	apis.ModelQuery(hostPort, modelQuery)
	apis.AlgStatus(hostPort, algStatus)
	apis.TranscodeStatus(hostPort, transcodeStatus)
	apis.ExtendData(hostPort, ExtendData)
	apis.VideoSliceInfo(hostPort, videoSliceInfo)
	apis.VideoSlicesData(hostPort)
}
