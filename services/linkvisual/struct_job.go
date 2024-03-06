package linkvisual

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

// Job is a nested struct in linkvisual response
type Job struct {
	Status       int    `json:"Status" xml:"Status"`
	Type         int    `json:"Type" xml:"Type"`
	Progress     int    `json:"Progress" xml:"Progress"`
	RecordType   int    `json:"RecordType" xml:"RecordType"`
	BeginTime    int    `json:"BeginTime" xml:"BeginTime"`
	Url          string `json:"Url" xml:"Url"`
	FileName     string `json:"FileName" xml:"FileName"`
	EndTime      int    `json:"EndTime" xml:"EndTime"`
	StreamType   int    `json:"StreamType" xml:"StreamType"`
	JobId        string `json:"JobId" xml:"JobId"`
	JobErrorCode int    `json:"JobErrorCode" xml:"JobErrorCode"`
	IotId        string `json:"IotId" xml:"IotId"`
}
