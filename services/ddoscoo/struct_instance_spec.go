package ddoscoo

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

// InstanceSpec is a nested struct in ddoscoo response
type InstanceSpec struct {
	BaseBandwidth    int    `json:"BaseBandwidth" xml:"BaseBandwidth"`
	QpsLimit         int    `json:"QpsLimit" xml:"QpsLimit"`
	BandwidthMbps    int    `json:"BandwidthMbps" xml:"BandwidthMbps"`
	ElasticBw        int    `json:"ElasticBw" xml:"ElasticBw"`
	DefenseCount     int    `json:"DefenseCount" xml:"DefenseCount"`
	SiteLimit        int    `json:"SiteLimit" xml:"SiteLimit"`
	PortLimit        int    `json:"PortLimit" xml:"PortLimit"`
	ElasticBandwidth int    `json:"ElasticBandwidth" xml:"ElasticBandwidth"`
	FunctionVersion  string `json:"FunctionVersion" xml:"FunctionVersion"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	DomainLimit      int    `json:"DomainLimit" xml:"DomainLimit"`
	ElasticBwModel   string `json:"ElasticBwModel" xml:"ElasticBwModel"`
	CpsLimit         int64  `json:"CpsLimit" xml:"CpsLimit"`
	ConnLimit        int64  `json:"ConnLimit" xml:"ConnLimit"`
	RealLimitBw      int64  `json:"RealLimitBw" xml:"RealLimitBw"`
	ElasticQpsMode   string `json:"ElasticQpsMode" xml:"ElasticQpsMode"`
	ElasticQps       int64  `json:"ElasticQps" xml:"ElasticQps"`
}
