package adb

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

// StageInfosItem is a nested struct in adb response
type StageInfosItem struct {
	InputRows      int64  `json:"InputRows" xml:"InputRows"`
	InputDataSize  int64  `json:"InputDataSize" xml:"InputDataSize"`
	OutputRows     int64  `json:"OutputRows" xml:"OutputRows"`
	OutputDataSize int64  `json:"OutputDataSize" xml:"OutputDataSize"`
	PeakMemory     int64  `json:"PeakMemory" xml:"PeakMemory"`
	OperatorCost   int64  `json:"OperatorCost" xml:"OperatorCost"`
	StageId        string `json:"StageId" xml:"StageId"`
	State          string `json:"State" xml:"State"`
	Progress       string `json:"Progress" xml:"Progress"`
}
