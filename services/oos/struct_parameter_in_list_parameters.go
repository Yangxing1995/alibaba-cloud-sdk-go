package oos

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

// ParameterInListParameters is a nested struct in oos response
type ParameterInListParameters struct {
	Type             string                 `json:"Type" xml:"Type"`
	UpdatedDate      string                 `json:"UpdatedDate" xml:"UpdatedDate"`
	UpdatedBy        string                 `json:"UpdatedBy" xml:"UpdatedBy"`
	Tags             map[string]interface{} `json:"Tags" xml:"Tags"`
	Description      string                 `json:"Description" xml:"Description"`
	CreatedBy        string                 `json:"CreatedBy" xml:"CreatedBy"`
	ResourceGroupId  string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	CreatedDate      string                 `json:"CreatedDate" xml:"CreatedDate"`
	ParameterVersion string                 `json:"ParameterVersion" xml:"ParameterVersion"`
	Name             string                 `json:"Name" xml:"Name"`
	Id               string                 `json:"Id" xml:"Id"`
	ShareType        string                 `json:"ShareType" xml:"ShareType"`
}
