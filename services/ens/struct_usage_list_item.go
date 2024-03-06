package ens

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

// UsageListItem is a nested struct in ens response
type UsageListItem struct {
	LanRxBw          int64  `json:"LanRxBw" xml:"LanRxBw"`
	LanTxBw          int64  `json:"LanTxBw" xml:"LanTxBw"`
	Point            int64  `json:"Point" xml:"Point"`
	PointTs          string `json:"PointTs" xml:"PointTs"`
	StorageUsageByte int64  `json:"StorageUsageByte" xml:"StorageUsageByte"`
	WanRxBw          int64  `json:"WanRxBw" xml:"WanRxBw"`
	WanTxBw          int64  `json:"WanTxBw" xml:"WanTxBw"`
}
