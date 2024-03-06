package dataworks_public

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

// Api is a nested struct in dataworks_public response
type Api struct {
	FolderId            int64               `json:"FolderId" xml:"FolderId"`
	TenantId            int64               `json:"TenantId" xml:"TenantId"`
	ModifiedTime        string              `json:"ModifiedTime" xml:"ModifiedTime"`
	CreatedTime         string              `json:"CreatedTime" xml:"CreatedTime"`
	CreatorId           string              `json:"CreatorId" xml:"CreatorId"`
	ResponseContentType int                 `json:"ResponseContentType" xml:"ResponseContentType"`
	OperatorId          string              `json:"OperatorId" xml:"OperatorId"`
	GroupId             string              `json:"GroupId" xml:"GroupId"`
	Description         string              `json:"Description" xml:"Description"`
	Timeout             int                 `json:"Timeout" xml:"Timeout"`
	ApiPath             string              `json:"ApiPath" xml:"ApiPath"`
	Status              int                 `json:"Status" xml:"Status"`
	ApiMode             int                 `json:"ApiMode" xml:"ApiMode"`
	ApiId               int64               `json:"ApiId" xml:"ApiId"`
	ProjectId           int64               `json:"ProjectId" xml:"ProjectId"`
	RequestMethod       int                 `json:"RequestMethod" xml:"RequestMethod"`
	ApiName             string              `json:"ApiName" xml:"ApiName"`
	VisibleRange        int                 `json:"VisibleRange" xml:"VisibleRange"`
	Protocols           []int               `json:"Protocols" xml:"Protocols"`
	WizardDetails       WizardDetails       `json:"WizardDetails" xml:"WizardDetails"`
	RegistrationDetails RegistrationDetails `json:"RegistrationDetails" xml:"RegistrationDetails"`
	ScriptDetails       ScriptDetails       `json:"ScriptDetails" xml:"ScriptDetails"`
}
