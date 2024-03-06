package cdn

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

// CertInfo is a nested struct in cdn response
type CertInfo struct {
	CertName                string `json:"CertName" xml:"CertName"`
	CertOrg                 string `json:"CertOrg" xml:"CertOrg"`
	DomainList              string `json:"DomainList" xml:"DomainList"`
	DomainNames             string `json:"DomainNames" xml:"DomainNames"`
	CertExpireTime          string `json:"CertExpireTime" xml:"CertExpireTime"`
	CertStatus              string `json:"CertStatus" xml:"CertStatus"`
	DomainCnameStatus       string `json:"DomainCnameStatus" xml:"DomainCnameStatus"`
	CreateTime              string `json:"CreateTime" xml:"CreateTime"`
	ServerCertificateStatus string `json:"ServerCertificateStatus" xml:"ServerCertificateStatus"`
	CertUpdateTime          string `json:"CertUpdateTime" xml:"CertUpdateTime"`
	CertSubjectCommonName   string `json:"CertSubjectCommonName" xml:"CertSubjectCommonName"`
	CertCaIsLegacy          string `json:"CertCaIsLegacy" xml:"CertCaIsLegacy"`
	CertDomainName          string `json:"CertDomainName" xml:"CertDomainName"`
	CertStartTime           string `json:"CertStartTime" xml:"CertStartTime"`
	CertId                  string `json:"CertId" xml:"CertId"`
	Issuer                  string `json:"Issuer" xml:"Issuer"`
	HttpsCrt                string `json:"HttpsCrt" xml:"HttpsCrt"`
	CertType                string `json:"CertType" xml:"CertType"`
	DomainName              string `json:"DomainName" xml:"DomainName"`
	CertRegion              string `json:"CertRegion" xml:"CertRegion"`
	ServerCertificate       string `json:"ServerCertificate" xml:"ServerCertificate"`
	CertCommonName          string `json:"CertCommonName" xml:"CertCommonName"`
	Status                  string `json:"Status" xml:"Status"`
	CertExpired             string `json:"CertExpired" xml:"CertExpired"`
	CertLife                string `json:"CertLife" xml:"CertLife"`
}
