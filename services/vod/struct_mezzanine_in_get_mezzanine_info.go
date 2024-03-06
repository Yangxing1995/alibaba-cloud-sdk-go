package vod

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// MezzanineInGetMezzanineInfo is a nested struct in vod response
type MezzanineInGetMezzanineInfo struct {
	CreationTime      string        `json:"CreationTime" xml:"CreationTime"`
	Status            string        `json:"Status" xml:"Status"`
	FileURL           string        `json:"FileURL" xml:"FileURL"`
	VideoId           string        `json:"VideoId" xml:"VideoId"`
	Height            int64         `json:"Height" xml:"Height"`
	Bitrate           string        `json:"Bitrate" xml:"Bitrate"`
	FileName          string        `json:"FileName" xml:"FileName"`
	OutputType        string        `json:"OutputType" xml:"OutputType"`
	PreprocessStatus  string        `json:"PreprocessStatus" xml:"PreprocessStatus"`
	Width             int64         `json:"Width" xml:"Width"`
	Size              int64         `json:"Size" xml:"Size"`
	CRC64             string        `json:"CRC64" xml:"CRC64"`
	Duration          string        `json:"Duration" xml:"Duration"`
	Fps               string        `json:"Fps" xml:"Fps"`
	StorageClass      string        `json:"StorageClass" xml:"StorageClass"`
	RestoreStatus     string        `json:"RestoreStatus" xml:"RestoreStatus"`
	RestoreExpiration string        `json:"RestoreExpiration" xml:"RestoreExpiration"`
	AudioStreamList   []AudioStream `json:"AudioStreamList" xml:"AudioStreamList"`
	VideoStreamList   []VideoStream `json:"VideoStreamList" xml:"VideoStreamList"`
}
