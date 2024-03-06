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

// NebulaTaskDetail is a nested struct in mpaas response
type NebulaTaskDetail struct {
	Atomic           int                `json:"Atomic" xml:"Atomic"`
	TaskType         int                `json:"TaskType" xml:"TaskType"`
	Platform         string             `json:"Platform" xml:"Platform"`
	BaseInfoId       int64              `json:"BaseInfoId" xml:"BaseInfoId"`
	DownloadUrl      string             `json:"DownloadUrl" xml:"DownloadUrl"`
	TaskVersion      int64              `json:"TaskVersion" xml:"TaskVersion"`
	GreyNum          int                `json:"GreyNum" xml:"GreyNum"`
	GreyConfigInfo   string             `json:"GreyConfigInfo" xml:"GreyConfigInfo"`
	PublishPeriod    int                `json:"PublishPeriod" xml:"PublishPeriod"`
	SyncType         int                `json:"SyncType" xml:"SyncType"`
	SourceId         string             `json:"SourceId" xml:"SourceId"`
	ProductVersion   string             `json:"ProductVersion" xml:"ProductVersion"`
	WorkspaceId      string             `json:"WorkspaceId" xml:"WorkspaceId"`
	GreyEndtimeData  string             `json:"GreyEndtimeData" xml:"GreyEndtimeData"`
	PublishType      int                `json:"PublishType" xml:"PublishType"`
	TaskStatus       int                `json:"TaskStatus" xml:"TaskStatus"`
	GreyUrl          string             `json:"GreyUrl" xml:"GreyUrl"`
	SyncResult       string             `json:"SyncResult" xml:"SyncResult"`
	Status           int                `json:"Status" xml:"Status"`
	Percent          int                `json:"Percent" xml:"Percent"`
	PublishMode      int                `json:"PublishMode" xml:"PublishMode"`
	TaskName         string             `json:"TaskName" xml:"TaskName"`
	QuickRollback    int                `json:"QuickRollback" xml:"QuickRollback"`
	SourceType       string             `json:"SourceType" xml:"SourceType"`
	GreyEndtime      string             `json:"GreyEndtime" xml:"GreyEndtime"`
	IssueDesc        string             `json:"IssueDesc" xml:"IssueDesc"`
	GmtModified      string             `json:"GmtModified" xml:"GmtModified"`
	GreyEndtimeStr   string             `json:"GreyEndtimeStr" xml:"GreyEndtimeStr"`
	BizType          string             `json:"BizType" xml:"BizType"`
	Creator          string             `json:"Creator" xml:"Creator"`
	Modifier         string             `json:"Modifier" xml:"Modifier"`
	AppId            string             `json:"AppId" xml:"AppId"`
	ProductId        string             `json:"ProductId" xml:"ProductId"`
	FullRepair       int                `json:"FullRepair" xml:"FullRepair"`
	ReleaseVersion   string             `json:"ReleaseVersion" xml:"ReleaseVersion"`
	FileSize         string             `json:"FileSize" xml:"FileSize"`
	GmtModifiedStr   string             `json:"GmtModifiedStr" xml:"GmtModifiedStr"`
	UpgradeNoticeNum int64              `json:"UpgradeNoticeNum" xml:"UpgradeNoticeNum"`
	AppCode          string             `json:"AppCode" xml:"AppCode"`
	Memo             string             `json:"Memo" xml:"Memo"`
	Cronexpress      int                `json:"Cronexpress" xml:"Cronexpress"`
	SourceName       string             `json:"SourceName" xml:"SourceName"`
	WhitelistIds     string             `json:"WhitelistIds" xml:"WhitelistIds"`
	GmtCreate        string             `json:"GmtCreate" xml:"GmtCreate"`
	PackageId        int64              `json:"PackageId" xml:"PackageId"`
	OssPath          string             `json:"OssPath" xml:"OssPath"`
	UpgradeProgress  string             `json:"UpgradeProgress" xml:"UpgradeProgress"`
	Id               int64              `json:"Id" xml:"Id"`
	ExtraData        string             `json:"ExtraData" xml:"ExtraData"`
	RuleJsonList     []RuleJsonListItem `json:"RuleJsonList" xml:"RuleJsonList"`
}
