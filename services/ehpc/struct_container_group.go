package ehpc

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

// ContainerGroup is a nested struct in ehpc response
type ContainerGroup struct {
	ContainerGroupId      string             `json:"ContainerGroupId" xml:"ContainerGroupId"`
	ContainerGroupName    string             `json:"ContainerGroupName" xml:"ContainerGroupName"`
	Status                string             `json:"Status" xml:"Status"`
	InstanceType          string             `json:"InstanceType" xml:"InstanceType"`
	SpotStrategy          string             `json:"SpotStrategy" xml:"SpotStrategy"`
	SpotPriceLimit        float64            `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	Cpu                   float64            `json:"Cpu" xml:"Cpu"`
	Memory                float64            `json:"Memory" xml:"Memory"`
	CreationTime          string             `json:"CreationTime" xml:"CreationTime"`
	SucceededTime         string             `json:"SucceededTime" xml:"SucceededTime"`
	ExpiredTime           string             `json:"ExpiredTime" xml:"ExpiredTime"`
	FailedTime            string             `json:"FailedTime" xml:"FailedTime"`
	Discount              int64              `json:"Discount" xml:"Discount"`
	EniInstanceId         string             `json:"EniInstanceId" xml:"EniInstanceId"`
	EphemeralStorage      int64              `json:"EphemeralStorage" xml:"EphemeralStorage"`
	InternetIp            string             `json:"InternetIp" xml:"InternetIp"`
	IntranetIp            string             `json:"IntranetIp" xml:"IntranetIp"`
	Ipv6Address           string             `json:"Ipv6Address" xml:"Ipv6Address"`
	RamRoleName           string             `json:"RamRoleName" xml:"RamRoleName"`
	RegionId              string             `json:"RegionId" xml:"RegionId"`
	ResourceGroupId       string             `json:"ResourceGroupId" xml:"ResourceGroupId"`
	RestartPolicy         string             `json:"RestartPolicy" xml:"RestartPolicy"`
	SecurityGroupId       string             `json:"SecurityGroupId" xml:"SecurityGroupId"`
	TenantEniInstanceId   string             `json:"TenantEniInstanceId" xml:"TenantEniInstanceId"`
	TenantEniIp           string             `json:"TenantEniIp" xml:"TenantEniIp"`
	TenantSecurityGroupId string             `json:"TenantSecurityGroupId" xml:"TenantSecurityGroupId"`
	TenantVSwitchId       string             `json:"TenantVSwitchId" xml:"TenantVSwitchId"`
	VSwitchId             string             `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                 string             `json:"VpcId" xml:"VpcId"`
	ZoneId                string             `json:"ZoneId" xml:"ZoneId"`
	DnsConfig             DnsConfig          `json:"DnsConfig" xml:"DnsConfig"`
	EciSecurityContext    EciSecurityContext `json:"EciSecurityContext" xml:"EciSecurityContext"`
	Containers            []Container        `json:"Containers" xml:"Containers"`
	Volumes               []Volume           `json:"Volumes" xml:"Volumes"`
	Events                []Event            `json:"Events" xml:"Events"`
	HostAliases           []HostAlias        `json:"HostAliases" xml:"HostAliases"`
	InitContainers        []InitContainer    `json:"InitContainers" xml:"InitContainers"`
	Tags                  []Tag              `json:"Tags" xml:"Tags"`
}
