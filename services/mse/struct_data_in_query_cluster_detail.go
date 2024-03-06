package mse

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

// DataInQueryClusterDetail is a nested struct in mse response
type DataInQueryClusterDetail struct {
	VpcId                string                 `json:"VpcId" xml:"VpcId"`
	CreateTime           string                 `json:"CreateTime" xml:"CreateTime"`
	IntranetAddress      string                 `json:"IntranetAddress" xml:"IntranetAddress"`
	DiskType             string                 `json:"DiskType" xml:"DiskType"`
	PubNetworkFlow       string                 `json:"PubNetworkFlow" xml:"PubNetworkFlow"`
	DiskCapacity         int64                  `json:"DiskCapacity" xml:"DiskCapacity"`
	MemoryCapacity       int64                  `json:"MemoryCapacity" xml:"MemoryCapacity"`
	ClusterAliasName     string                 `json:"ClusterAliasName" xml:"ClusterAliasName"`
	InstanceCount        int                    `json:"InstanceCount" xml:"InstanceCount"`
	IntranetPort         string                 `json:"IntranetPort" xml:"IntranetPort"`
	ClusterId            string                 `json:"ClusterId" xml:"ClusterId"`
	IntranetDomain       string                 `json:"IntranetDomain" xml:"IntranetDomain"`
	InternetDomain       string                 `json:"InternetDomain" xml:"InternetDomain"`
	PayInfo              string                 `json:"PayInfo" xml:"PayInfo"`
	InternetAddress      string                 `json:"InternetAddress" xml:"InternetAddress"`
	InstanceId           string                 `json:"InstanceId" xml:"InstanceId"`
	AclEntryList         string                 `json:"AclEntryList" xml:"AclEntryList"`
	HealthStatus         string                 `json:"HealthStatus" xml:"HealthStatus"`
	InitCostTime         int64                  `json:"InitCostTime" xml:"InitCostTime"`
	RegionId             string                 `json:"RegionId" xml:"RegionId"`
	AclId                string                 `json:"AclId" xml:"AclId"`
	Cpu                  int                    `json:"Cpu" xml:"Cpu"`
	ClusterType          string                 `json:"ClusterType" xml:"ClusterType"`
	ClusterName          string                 `json:"ClusterName" xml:"ClusterName"`
	InitStatus           string                 `json:"InitStatus" xml:"InitStatus"`
	InternetPort         string                 `json:"InternetPort" xml:"InternetPort"`
	AppVersion           string                 `json:"AppVersion" xml:"AppVersion"`
	NetType              string                 `json:"NetType" xml:"NetType"`
	ClusterVersion       string                 `json:"ClusterVersion" xml:"ClusterVersion"`
	ClusterSpecification string                 `json:"ClusterSpecification" xml:"ClusterSpecification"`
	VSwitchId            string                 `json:"VSwitchId" xml:"VSwitchId"`
	ConnectionType       string                 `json:"ConnectionType" xml:"ConnectionType"`
	MseVersion           string                 `json:"MseVersion" xml:"MseVersion"`
	ChargeType           string                 `json:"ChargeType" xml:"ChargeType"`
	OrderClusterVersion  string                 `json:"OrderClusterVersion" xml:"OrderClusterVersion"`
	Tags                 map[string]interface{} `json:"Tags" xml:"Tags"`
	ResourceGroupId      string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceModels       []InstanceModel        `json:"InstanceModels" xml:"InstanceModels"`
}
