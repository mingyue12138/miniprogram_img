package apis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simretina/core/config"
	"strings"

	"go.pfgit.cn/letsgo/xdev"
)

var HostPort string

func Login(HostPort string, loginMes *config.LoginMes) {
	//http
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/System/Login")
	url := host.String()
	confPath := "./configs/loginMes.json"
	loginMes, err := config.ParseJsonLoginMes(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}

	// xdev.Log.Info("登录接口消息体", loginMes)
	jsonStr, _ := json.Marshal(loginMes)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	xdev.Log.Info(string(body))
	//websocket
	// fmt.Println(resp)
}

func Logout(HostPort string, deviceID *config.DeviceID) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/System/Keepalive")
	url := host.String()
	confPath := "./configs/loginMes.json"
	loginMes, err := config.ParseJsonLoginMes(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	deviceID = new(config.DeviceID)
	deviceID.DeviceID = loginMes.DeviceID
	jsonStr, _ := json.Marshal(deviceID)
	// xdev.Log.Info()
	// s := `{"deviceID": "44030100001400000001"}`
	// b := []byte(s)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	// xdev.Log.Info(b)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	xdev.Log.Info(string(body))
	// xdev.Log.Info(resp)
}
func Keepalive(HostPort string, deviceID *config.DeviceID) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/System/Keepalive")
	url := host.String()
	confPath := "./configs/loginMes.json"
	loginMes, err := config.ParseJsonLoginMes(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	deviceID = new(config.DeviceID)
	deviceID.DeviceID = loginMes.DeviceID
	jsonStr, _ := json.Marshal(deviceID)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	if err != nil {
		xdev.Log.Error(err)
	}
	// xdev.Log.Info(b)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	xdev.Log.Info(string(body))
}

func GetPushStreamAddress(HostPort string, deviceID *config.DeviceID) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/getPushStreamAddress")
	url := host.String()

	confPath := "./configs/loginMes.json"
	loginMes, err := config.ParseJsonLoginMes(&confPath)
	if err != nil {
		xdev.Log.Info(err)
	}
	deviceID = new(config.DeviceID)
	deviceID.DeviceID = loginMes.DeviceID
	jsonStr, _ := json.Marshal(deviceID)
	xdev.Log.Info(jsonStr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		xdev.Log.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func DeviceStatus(HostPort string, deviceStatus *config.DeviceStatus) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/DeviceStatus")
	url := host.String()

	confPath := "./configs/deviceStatus.json"
	deviceStatus, err := config.ParseJsonDeviceStatus(&confPath)
	//deviceStatus.CurrentTime = time.Now().Format("2006-01-02 15:04:05.99")
	if err != nil {
		xdev.Log.Error(err)
	}
	xdev.Log.Infoln(deviceStatus)
	jsonStr, _ := json.Marshal(deviceStatus)
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(jsonStr))
	if err != nil {
		xdev.Log.Error(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}

//数据上报
func Data(HostPort string) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/Feature/Data")
	url := host.String()
	confPath := "./configs/testImage.jpg"
	srcByte, err := ioutil.ReadFile(confPath)
	if err != nil {
		xdev.Log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(srcByte))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func FeatureInfo(HostPort string, featureInfo *config.FeatureInfo) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/FeatureInfo")
	url := host.String()

	confPath := "./configs/FeatureInfo.json"
	featureInfo, err := config.ParseJsonFeatureInfo(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(featureInfo)
	// xdev.Log.Info(featureInfo)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func ImageInfo(HostPort string, imageInfo *config.ImageInfo) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/ImageInfo")
	url := host.String()

	confPath := "./configs/ImageInfo.json"
	imageInfo, err := config.ParseJsonImageInfo(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(ImageInfo)
	// xdev.Log.Info(featureInfo)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func ImagesData(HostPort string) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/Images/Data")
	url := host.String()
	confPath := "./configs/testImage.jpg"
	srcByte, err := ioutil.ReadFile(confPath)
	if err != nil {
		xdev.Log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(srcByte))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func ModelQuery(HostPort string, modelQuery *config.ModelQuery) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/ModelQuery")
	url := host.String()

	confPath := "./configs/ModelQuery.json"
	modelQuery, err := config.ParseJsonModelQuery(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(modelQuery)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func AlgStatus(HostPort string, algStatus *config.AlgStatus) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/ResultInfo/AlgStatus/{ConfigID}/{AlgID}")
	url := host.String()

	confPath := "./configs/AlgStatus.json"
	algStatus, err := config.ParseJsonAlgStatus(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	// xdev.Log.Infoln(algStatus)
	jsonStr, _ := json.Marshal(algStatus)
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(jsonStr))
	if err != nil {
		xdev.Log.Error(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func TranscodeStatus(HostPort string, transcodeStatus *config.TranscodeStatus) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/ResultInfo/TranscodeStatus/{TaskID}")
	url := host.String()

	confPath := "./configs/TranscodeStatus.json"
	transcodeStatus, err := config.ParseJsonTranscodeStatus(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	// xdev.Log.Infoln(algStatus)
	jsonStr, _ := json.Marshal(transcodeStatus)
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(jsonStr))
	if err != nil {
		xdev.Log.Error(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func ExtendData(HostPort string, ExtendData *config.ExtendData) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/System/ExtendData")
	url := host.String()

	confPath := "./configs/ExtendData.json"
	extendData, err := config.ParseJsonExtendData(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(extendData)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func VideoSliceInfo(HostPort string, videoSliceInfo *config.VideoSliceInfo) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/VideoSliceInfo")
	url := host.String()
	confPath := "./configs/VideoSliceInfo.json"
	videoSliceInfo, err := config.ParseJsonVideoSliceInfo(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(videoSliceInfo)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func VideoSlicesData(HostPort string) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/szswm/VideoSliceInfo")
	url := host.String()

	confPath := "./configs/img1.mp4"
	srcByte, err := ioutil.ReadFile(confPath)
	if err != nil {
		xdev.Log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(srcByte))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}

//控制下发
func AlgModelUpdate(HostPort string, algModelUpdate *config.AlgModelUpdate) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/algModelUpdate")
	url := host.String()

	confPath := "./configs/algModelUpdate.json"
	algModelUpdate, err := config.ParseJsonAlgModelUpdate(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(algModelUpdate)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func AlgResultQuery(HostPort string, algResultQuery *config.AlgResultQuery) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/algResultQuery")
	url := host.String()

	confPath := "./configs/algResultQuery.json"
	algResultQuery, err := config.ParseJsonAlgResultQuery(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	// algResultQuery.
	// xdev.Log.Info(algResultQuery)
	jsonStr, _ := json.Marshal(algResultQuery)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func AlgModelDelete(HostPort string, algModelDelete *config.AlgModelUpdate) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/algModelDelete")
	url := host.String()

	confPath := "./configs/algModelUpdate.json"
	algModelDelete, err := config.ParseJsonAlgModelUpdate(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(algModelDelete)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func BaseConfig(HostPort string, baseConfig *config.BaseConfig) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/baseConfig")
	url := host.String()

	confPath := "./configs/baseConfig.json"
	baseConfig, err := config.ParseJsonBaseConfig(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(baseConfig)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}

//TODO:文档和接口字段有出入
func DeviceRestart(HostPort string, deviceRestart *config.DeviceRestart) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/deviceRestart")
	url := host.String()

	confPath := "./configs/deviceRestart.json"
	deviceRestart, err := config.ParseJsonDeviceRestart(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(deviceRestart)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func DeviceShutDown(HostPort string, deviceShutDown *config.DeviceRestart) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/deviceShutDown")
	url := host.String()

	confPath := "./configs/deviceRestart.json"
	deviceShutDown, err := config.ParseJsonDeviceRestart(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(deviceShutDown)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func DeviceStatusQuery(HostPort string, deviceStatusQuery *config.DeviceStatusQuery) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/deviceStatusQuery")
	url := host.String()

	confPath := "./configs/deviceStatusQuery.json"
	deviceStatusQuery, err := config.ParseJsonDeviceStatusQuery(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(deviceStatusQuery)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func Extend(HostPort string, extend *config.Extend) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/extend")
	url := host.String()

	confPath := "./configs/extend.json"
	extend, err := config.ParseJsonExtend(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(extend)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func FunctionDefine(HostPort string, functionDefine *config.FunctionDefine) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/functionDefine")
	url := host.String()

	confPath := "./configs/functionDefine.json"
	functionDefine, err := config.ParseJsonFunctionDefine(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(functionDefine)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func ImageQuery(HostPort string, imageQuery *config.ImageQuery) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/imageQuery")
	url := host.String()

	confPath := "./configs/imageQuery.json"
	imageQuery, err := config.ParseJsonImageQuery(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(imageQuery)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func Transcode(HostPort string, transcode *config.Transcode) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/transcode")
	url := host.String()

	confPath := "./configs/transcode.json"
	transcode, err := config.ParseJsonTranscode(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(transcode)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
func VideoQuery(HostPort string, videoQuery *config.VideoQuery) {
	var host strings.Builder
	host.WriteString(HostPort)
	host.WriteString("/api/v1/cmd/szswm/videoQuery")
	url := host.String()

	confPath := "./configs/videoQuery.json"
	videoQuery, err := config.ParseJsonVideoQuery(&confPath)
	if err != nil {
		xdev.Log.Error(err)
	}
	jsonStr, _ := json.Marshal(videoQuery)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add("ds",token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	xdev.Log.Info(string(body))
}
