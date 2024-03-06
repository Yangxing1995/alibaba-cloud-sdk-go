package domain

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

// QueryTransferInResponse is a nested struct in domain response
type QueryTransferInResponse struct {
	ServerLockStatus string `json:"ServerLockStatus" xml:"ServerLockStatus"`
	LockInstanceId   string `json:"LockInstanceId" xml:"LockInstanceId"`
	UserId           string `json:"UserId" xml:"UserId"`
	GmtCreate        string `json:"GmtCreate" xml:"GmtCreate"`
	ExpireDate       string `json:"ExpireDate" xml:"ExpireDate"`
	StartDate        string `json:"StartDate" xml:"StartDate"`
	LockProductId    string `json:"LockProductId" xml:"LockProductId"`
	DomainInstanceId string `json:"DomainInstanceId" xml:"DomainInstanceId"`
	GmtModified      string `json:"GmtModified" xml:"GmtModified"`
	DomainName       string `json:"DomainName" xml:"DomainName"`
}
