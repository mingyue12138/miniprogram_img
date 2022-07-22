package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LoginMes struct {
	DeviceID        string `json:"deviceID"`
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	Vendor          string `json:"vendor"`
	DeviceSN        string `json:"deviceSn"`
	DeviceType      string `json:"deviceType"`
	DeviceName      string `json:"deviceName"`
	OSName          string `json:"osName"`
	OSVersion       string `json:"osVersion"`
	SoftwareVersion string `json:"softwareVersion"`
	DeviceCPU       string `json:"deviceCpu"`
	DeviceGPU       string `json:"deviceGpu"`
	DeviceRAM       string `json:"deviceRam"`
	DeviceDisk      string `json:"deviceDisk"`
	IPAddr          string `json:"ipAddr"`
	Port            string `json:"port"`
	Longitude       string `json:"longitude"`
	Latitude        string `json:"latitude"`
	Location        string `json:"location"`
}

func ParseJsonLoginMes(f *string) (*LoginMes, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var loginMes LoginMes
	err = json.Unmarshal(b, &loginMes)
	return &loginMes, err
}

type DeviceID struct {
	DeviceID string `json:"deviceID"`
}
type DeviceStatus struct {
	CurrentTime  string `json:"currentTime"`
	DeviceID     string `json:"deviceID"`
	DeviceStatus string `json:"deviceStatus"`
	Forward      string `json:"forward"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

func ParseJsonDeviceStatus(f *string) (*DeviceStatus, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var deviceStatus DeviceStatus
	err = json.Unmarshal(b, &deviceStatus)
	return &deviceStatus, err
}

type AlgModels struct {
	AddTime       string `json:"addTime"`
	CompressType  string `json:"compressType"`
	ModelDesc     string `json:"modelDesc"`
	ModelFileName string `json:"modelFileName"`
	ModelId       int32  `json:"modelId"`
	ModelURL      string `json:"modelURL"`
	ModelVersion  string `json:"modelVersion"`
	UpdateType    string `json:"updateType"`
}
type Alg struct {
	AlgDesc     string      `json:"algDesc"`
	AlgExecFile string      `json:"algExecFile"`
	AlgFileName string      `json:"algFileName"`
	AlgFileURL  string      `json:"algFileURL"`
	AlgId       int32       `json:"algId"`
	AlgModels   []AlgModels `json:"algModels"`
	AlgRate     int32       `json:"algRate"`
	UpdateTime  string      `json:"updateTime"`
}
type AlgModelUpdate struct {
	Alg           []Alg  `json:"alg"`
	ConfigId      int32  `json:"configId"`
	DeviceId      string `json:"deviceId"`
	OutputFeature string `json:"outputFeature"`
	OutputResult  string `json:"outputResult"`
}

func ParseJsonAlgModelUpdate(f *string) (*AlgModelUpdate, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var algModelUpdate AlgModelUpdate
	err = json.Unmarshal(b, &algModelUpdate)
	return &algModelUpdate, err
}

type AlgResultQuery struct {
	AlgId        int32  `json:"algId"`
	AlgVersion   string `json:"algVersion"`
	DeviceId     string `json:"deviceId"`
	EndTime      string `json:"endTime"`
	LatitudeMax  string `json:"latitudeMax"`
	LatitudeMin  string `json:"latitudeMin"`
	LongitudeMax string `json:"longitudeMax"`
	LongitudeMin string `json:"longitudeMin"`
	StartTime    string `json:"startTime"`
	Type         int32  `json:"type"`
}

func ParseJsonAlgResultQuery(f *string) (*AlgResultQuery, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var algResultQuery AlgResultQuery
	err = json.Unmarshal(b, &algResultQuery)
	return &algResultQuery, err
}

type BaseConfig struct {
	DeviceId            string `json:"deviceId"`
	FeatureStreamAddr   string `json:"featureStreamAddr"`
	HeartBeatInterval   int32  `json:"heartBeatInterval"`
	Height              int32  `json:"height"`
	IpAddr              string `json:"ipAddr"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	VideoEncodingFormat string `json:"videoEncodingFormat"`
	VideoStreamProtocol string `json:"videoStreamProtocol"`
	VideoStream_Addr    string `json:"videoStream_Addr"`
	Width               int32  `json:"width"`
}

func ParseJsonBaseConfig(f *string) (*BaseConfig, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var baseConfig BaseConfig
	err = json.Unmarshal(b, &baseConfig)
	return &baseConfig, err
}

//TODO:
type DeviceRestart struct {
	Delay    string `json:"delay"`
	DeviceId string `json:"deviceId"`
}

func ParseJsonDeviceRestart(f *string) (*DeviceRestart, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var deviceRestart DeviceRestart
	err = json.Unmarshal(b, &deviceRestart)
	return &deviceRestart, err
}

type DeviceStatusQuery struct {
	DeviceId     string `json:"deviceId"`
	DeviceStatus string `json:"deviceStatus"`
	LatitudeMax  string `json:"latitudeMax"`
	LatitudeMin  string `json:"latitudeMin"`
	LongitudeMax string `json:"longitudeMax"`
	LongitudeMin string `json:"longitudeMin"`
}

func ParseJsonDeviceStatusQuery(f *string) (*DeviceStatusQuery, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var deviceStatusQuery DeviceStatusQuery
	err = json.Unmarshal(b, &deviceStatusQuery)
	return &deviceStatusQuery, err
}

type Extend struct {
	CommandArg1 string `json:"commandArg1"`
	DeviceId    string `json:"deviceId"`
}

func ParseJsonExtend(f *string) (*Extend, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var extend Extend
	err = json.Unmarshal(b, &extend)
	return &extend, err
}

type AlgModel struct {
	AddTime       string `json:"addTime"`
	CompressType  string `json:"compressType"`
	ModelDesc     string `json:"modelDesc"`
	ModelFileName string `json:"modelFileName"`
	ModelId       int32  `json:"modelId"`
	ModelURL      string `json:"modelURL"`
	ModelVersion  string `json:"modelVersion"`
	UpdateType    string `json:"updateType"`
}
type FunctionDefineAlg struct {
	AlgDesc          string     `json:"algDesc"`
	AlgExecFile      string     `json:"algExecFile"`
	AlgExpireTime    string     `json:"algExpireTime"`
	AlgFileName      string     `json:"algFileName"`
	AlgFileURL       string     `json:"algFileURL"`
	AlgId            int32      `json:"algId"`
	AlgOutputFeature string     `json:"algOutputFeature"`
	AlgOutputResult  string     `json:"algOutputResult"`
	AlgRate          int32      `json:"algRate"`
	AlgState         string     `json:"algState"`
	AlgVersion       string     `json:"algVersion"`
	AlgModels        []AlgModel `json:"algModels"`
}
type FunctionDefine struct {
	Alg           []FunctionDefineAlg `json:"alg"`
	ConfigId      int32               `json:"configId"`
	DeviceId      string              `json:"deviceId"`
	OutputFeature string              `json:"outputFeature"`
	OutputResult  string              `json:"outputResult"`
	OutputVideo   string              `json:"outputVideo"`
	VideoType     string              `json:"videoType"`
	VideoURL      string              `json:"videoURL"`
}

func ParseJsonFunctionDefine(f *string) (*FunctionDefine, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var functionDefine FunctionDefine
	err = json.Unmarshal(b, &functionDefine)
	return &functionDefine, err
}

type ImageQuery struct {
	DeviceId     string `json:"deviceId"`
	EndTime      string `json:"endTime"`
	LatitudeMax  string `json:"latitudeMax"`
	LatitudeMin  string `json:"latitudeMin"`
	LongitudeMax string `json:"longitudeMax"`
	LongitudeMin string `json:"longitudeMin"`
	StartTime    string `json:"startTime"`
}

func ParseJsonImageQuery(f *string) (*ImageQuery, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var imageQuery ImageQuery
	err = json.Unmarshal(b, &imageQuery)
	return &imageQuery, err
}

type Transcode struct {
	ConfigId      int32  `json:"configId"`
	DeviceId      string `json:"deviceId"`
	OprType       string `json:"oprType"`
	OutputBitRate int32  `json:"outputBitRate"`
	OutputCodec   string `json:"outputCodec"`
	OutputFormat  string `json:"outputFormat"`
	OutputFps     int32  `json:"outputFps"`
	OutputHeight  int32  `json:"outputHeight"`
	OutputPath    string `json:"outputPath"`
	OutputWidth   int32  `json:"outputWidth"`
	TaskUUID      string `json:"taskUUID"`
	TimeInterval  int32  `json:"timeInterval"`
	VideoType     string `json:"videoType"`
	VideoURL      string `json:"videoURL"`
}

func ParseJsonTranscode(f *string) (*Transcode, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var transcode Transcode
	err = json.Unmarshal(b, &transcode)
	return &transcode, err
}

type VideoQuery struct {
	DeviceId     string `json:"deviceId"`
	EndTime      string `json:"endTime"`
	LatitudeMax  string `json:"latitudeMax"`
	LatitudeMin  string `json:"latitudeMin"`
	LongitudeMax string `json:"longitudeMax"`
	LongitudeMin string `json:"longitudeMin"`
	StartTime    string `json:"startTime"`
}

func ParseJsonVideoQuery(f *string) (*VideoQuery, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var videoQuery VideoQuery
	err = json.Unmarshal(b, &videoQuery)
	return &videoQuery, err
}

type FeatureInfo struct {
	AlgId       int32  `json:"algId"`
	AlgVersion  string `json:"algVersion"`
	DeviceId    string `json:"deviceId"`
	FeatureData string `json:"featureData"`
	FeatureId   string `json:"featureId"`
	FeatureUrl  string `json:"featureUrl"`
	Latitude    string `json:"latitude"`
	LeftTopX    int32  `json:"leftTopX"`
	LeftTopY    int32  `json:"leftTopY"`
	Longitude   string `json:"longitude"`
	RightBtmX   int32  `json:"rightBtmX"`
	RightBtmY   int32  `json:"rightBtmY"`
	Ts          int64  `json:"ts"`
}

func ParseJsonFeatureInfo(f *string) (*FeatureInfo, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var featureInfo FeatureInfo
	err = json.Unmarshal(b, &featureInfo)
	return &featureInfo, err
}

type ImageInfo struct {
	CreateTime    string `json:"createTime"`
	DeviceId      string `json:"deviceId"`
	EntryTime     string `json:"entryTime"`
	FileFormat    string `json:"fileFormat"`
	FileSize      int32  `json:"fileSize"`
	Height        int32  `json:"height"`
	ImageId       string `json:"imageId"`
	ImageProcFlag string `json:"imageProcFlag"`
	ImageUrl      string `json:"imageUrl"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Ts            int64  `json:"ts"`
	Width         int32  `json:"width"`
}

func ParseJsonImageInfo(f *string) (*ImageInfo, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var imageInfo ImageInfo
	err = json.Unmarshal(b, &imageInfo)
	return &imageInfo, err
}

//TODO:endTime API文档为int32
type AlgModelFilter struct {
	AgType            string `json:"agType"`
	AlgComputingPower string `json:"algComputingPower"`
	AlgKeyword        string `json:"algKeyword"`
	AlgModelSize      string `json:"algModelSize"`
	AlgVRAM           string `json:"algVRAM"`
	EndTime           int32  `json:"endTime"`
	StartTime         string `json:"startTime"`
}
type ModelQuery struct {
	AlgModelFilter AlgModelFilter `json:"algModelFilter"`
	DeviceId       string         `json:"deviceId"`
	DeviceType     string         `json:"deviceType"`
}

func ParseJsonModelQuery(f *string) (*ModelQuery, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var modelQuery ModelQuery
	err = json.Unmarshal(b, &modelQuery)
	return &modelQuery, err
}

type AlgStatus struct {
	AlgStatus string `json:"algStatus"`
}

func ParseJsonAlgStatus(f *string) (*AlgStatus, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var algStatus AlgStatus
	err = json.Unmarshal(b, &algStatus)
	return &algStatus, err
}

type TranscodeStatus struct {
	TranscodeOutputUrl string `json:"transcodeOutputUrl"`
	TranscodeStatus    string `json:"transcodeStatus"`
}

func ParseJsonTranscodeStatus(f *string) (*TranscodeStatus, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var transcodeStatus TranscodeStatus
	err = json.Unmarshal(b, &transcodeStatus)
	return &transcodeStatus, err
}

//TODO：扩展信息，文档标注自定义
type ExtendedField struct {
}
type ExtendData struct {
	CurrentTime string        `json:"currentTime"`
	DeviceId    string        `json:"deviceId"`
	ExtendInfo  ExtendedField `json:"extendInfo"`
}

func ParseJsonExtendData(f *string) (*ExtendData, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var extendData ExtendData
	err = json.Unmarshal(b, &extendData)
	return &extendData, err
}

type VideoSliceInfo struct {
	AudioEncodingFormat string  `json:"audioEncodingFormat"`
	AudioFlag           int32   `json:"audioFlag"`
	BeginTime           string  `json:"beginTime"`
	ContentDescription  string  `json:"contentDescription"`
	DeviceId            string  `json:"deviceId"`
	EndTime             string  `json:"endTime"`
	EntryTime           string  `json:"entryTime"`
	FileFormat          string  `json:"fileFormat"`
	FileSize            int64   `json:"fileSize"`
	Height              int32   `json:"height"`
	OriginVideoId       string  `json:"originVideoId"`
	OriginVideoUrl      string  `json:"originVideoUrl"`
	ShotPlaceLongitude  string  `json:"shotPlaceLongitude"`
	ShotPlacetLatitude  string  `json:"shotPlacetLatitude"`
	ThumbnailUrl        string  `json:"thumbnailUrl"`
	TimeErr             int32   `json:"timeErr"`
	Title               string  `json:"title"`
	VideoEncodingFormat string  `json:"videoEncodingFormat"`
	VideoId             string  `json:"videoId"`
	VideoLen            float64 `json:"videoLen"`
	VideoUrl            string  `json:"videoUrl"`
	Width               int32   `json:"width"`
}

func ParseJsonVideoSliceInfo(f *string) (*VideoSliceInfo, error) {
	jsonFile, err := os.Open(*f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var videoSliceInfo VideoSliceInfo
	err = json.Unmarshal(b, &videoSliceInfo)
	return &videoSliceInfo, err
}
