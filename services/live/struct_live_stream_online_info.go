package live

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

// LiveStreamOnlineInfo is a nested struct in live response
type LiveStreamOnlineInfo struct {
	AudioDataRate int    `json:"AudioDataRate" xml:"AudioDataRate"`
	FrameRate     int    `json:"FrameRate" xml:"FrameRate"`
	PublishUrl    string `json:"PublishUrl" xml:"PublishUrl"`
	StreamName    string `json:"StreamName" xml:"StreamName"`
	AudioCodecId  int    `json:"AudioCodecId" xml:"AudioCodecId"`
	Height        int    `json:"Height" xml:"Height"`
	VideoDataRate int    `json:"VideoDataRate" xml:"VideoDataRate"`
	DomainName    string `json:"DomainName" xml:"DomainName"`
	TranscodeId   string `json:"TranscodeId" xml:"TranscodeId"`
	TranscodeDrm  string `json:"TranscodeDrm" xml:"TranscodeDrm"`
	PublishDomain string `json:"PublishDomain" xml:"PublishDomain"`
	PublishTime   string `json:"PublishTime" xml:"PublishTime"`
	AppName       string `json:"AppName" xml:"AppName"`
	PublishType   string `json:"PublishType" xml:"PublishType"`
	VideoCodecId  int    `json:"VideoCodecId" xml:"VideoCodecId"`
	Transcoded    string `json:"Transcoded" xml:"Transcoded"`
	Width         int    `json:"Width" xml:"Width"`
	ClientIp      string `json:"ClientIp" xml:"ClientIp"`
	ServerIp      string `json:"ServerIp" xml:"ServerIp"`
	StreamUrlArgs string `json:"StreamUrlArgs" xml:"StreamUrlArgs"`
}
