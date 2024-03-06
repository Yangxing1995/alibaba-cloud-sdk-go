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

// LiveStreamMonitorInfo is a nested struct in live response
type LiveStreamMonitorInfo struct {
	MonitorConfig      string        `json:"MonitorConfig" xml:"MonitorConfig"`
	Status             int           `json:"Status" xml:"Status"`
	MonitorId          string        `json:"MonitorId" xml:"MonitorId"`
	Domain             string        `json:"Domain" xml:"Domain"`
	CallbackUrl        string        `json:"CallbackUrl" xml:"CallbackUrl"`
	AudioFrom          int           `json:"AudioFrom" xml:"AudioFrom"`
	DingTalkWebHookUrl string        `json:"DingTalkWebHookUrl" xml:"DingTalkWebHookUrl"`
	MonitorName        string        `json:"MonitorName" xml:"MonitorName"`
	StopTime           string        `json:"StopTime" xml:"StopTime"`
	StartTime          string        `json:"StartTime" xml:"StartTime"`
	OutputTemplate     string        `json:"OutputTemplate" xml:"OutputTemplate"`
	Region             string        `json:"Region" xml:"Region"`
	OutputUrls         OutputUrls    `json:"OutputUrls" xml:"OutputUrls"`
	InputList          []InputConfig `json:"InputList" xml:"InputList"`
}
