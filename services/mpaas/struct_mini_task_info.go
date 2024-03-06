package mpaas

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

// MiniTaskInfo is a nested struct in mpaas response
type MiniTaskInfo struct {
	Status          string `json:"Status" xml:"Status"`
	PublishMode     int64  `json:"PublishMode" xml:"PublishMode"`
	AppCode         string `json:"AppCode" xml:"AppCode"`
	Memo            string `json:"Memo" xml:"Memo"`
	GreyEndtime     string `json:"GreyEndtime" xml:"GreyEndtime"`
	GreyNum         int64  `json:"GreyNum" xml:"GreyNum"`
	GreyConfigInfo  string `json:"GreyConfigInfo" xml:"GreyConfigInfo"`
	GmtModified     string `json:"GmtModified" xml:"GmtModified"`
	ProductVersion  string `json:"ProductVersion" xml:"ProductVersion"`
	GreyEndtimeData string `json:"GreyEndtimeData" xml:"GreyEndtimeData"`
	PublishType     int64  `json:"PublishType" xml:"PublishType"`
	TaskStatus      int64  `json:"TaskStatus" xml:"TaskStatus"`
	WhitelistIds    string `json:"WhitelistIds" xml:"WhitelistIds"`
	Platform        string `json:"Platform" xml:"Platform"`
	GmtCreate       string `json:"GmtCreate" xml:"GmtCreate"`
	PackageId       int64  `json:"PackageId" xml:"PackageId"`
	Id              int64  `json:"Id" xml:"Id"`
}
