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

// ContactTemplate is a nested struct in domain response
type ContactTemplate struct {
	CCountry                string `json:"CCountry" xml:"CCountry"`
	TelExt                  string `json:"TelExt" xml:"TelExt"`
	UpdateTime              string `json:"UpdateTime" xml:"UpdateTime"`
	RegType                 string `json:"RegType" xml:"RegType"`
	CreateTime              string `json:"CreateTime" xml:"CreateTime"`
	EProvince               string `json:"EProvince" xml:"EProvince"`
	UserId                  string `json:"UserId" xml:"UserId"`
	CVenu                   string `json:"CVenu" xml:"CVenu"`
	TelArea                 string `json:"TelArea" xml:"TelArea"`
	ContactTemplateId       int64  `json:"ContactTemplateId" xml:"ContactTemplateId"`
	AuditStatus             string `json:"AuditStatus" xml:"AuditStatus"`
	CProvince               string `json:"CProvince" xml:"CProvince"`
	PostalCode              string `json:"PostalCode" xml:"PostalCode"`
	CCity                   string `json:"CCity" xml:"CCity"`
	EVenu                   string `json:"EVenu" xml:"EVenu"`
	ECompany                string `json:"ECompany" xml:"ECompany"`
	CCompany                string `json:"CCompany" xml:"CCompany"`
	EName                   string `json:"EName" xml:"EName"`
	TelMain                 string `json:"TelMain" xml:"TelMain"`
	Email                   string `json:"Email" xml:"Email"`
	DefaultTemplate         bool   `json:"DefaultTemplate" xml:"DefaultTemplate"`
	EmailVerificationStatus int    `json:"EmailVerificationStatus" xml:"EmailVerificationStatus"`
	ECity                   string `json:"ECity" xml:"ECity"`
	CName                   string `json:"CName" xml:"CName"`
}
