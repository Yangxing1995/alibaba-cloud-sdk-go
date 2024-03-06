package sas

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribePropertyScaDetail invokes the sas.DescribePropertyScaDetail API synchronously
func (client *Client) DescribePropertyScaDetail(request *DescribePropertyScaDetailRequest) (response *DescribePropertyScaDetailResponse, err error) {
	response = CreateDescribePropertyScaDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertyScaDetailWithChan invokes the sas.DescribePropertyScaDetail API asynchronously
func (client *Client) DescribePropertyScaDetailWithChan(request *DescribePropertyScaDetailRequest) (<-chan *DescribePropertyScaDetailResponse, <-chan error) {
	responseChan := make(chan *DescribePropertyScaDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertyScaDetail(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribePropertyScaDetailWithCallback invokes the sas.DescribePropertyScaDetail API asynchronously
func (client *Client) DescribePropertyScaDetailWithCallback(request *DescribePropertyScaDetailRequest, callback func(response *DescribePropertyScaDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertyScaDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertyScaDetail(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribePropertyScaDetailRequest is the request struct for api DescribePropertyScaDetail
type DescribePropertyScaDetailRequest struct {
	*requests.RpcRequest
	SearchItemSub       string           `position:"Query" name:"SearchItemSub"`
	Remark              string           `position:"Query" name:"Remark"`
	Pid                 string           `position:"Query" name:"Pid"`
	SearchItem          string           `position:"Query" name:"SearchItem"`
	Uuid                string           `position:"Query" name:"Uuid"`
	Biz                 string           `position:"Query" name:"Biz"`
	SourceIp            string           `position:"Query" name:"SourceIp"`
	PageSize            requests.Integer `position:"Query" name:"PageSize"`
	ProcessStartedStart requests.Integer `position:"Query" name:"ProcessStartedStart"`
	ProcessStartedEnd   requests.Integer `position:"Query" name:"ProcessStartedEnd"`
	Lang                string           `position:"Query" name:"Lang"`
	ScaVersion          string           `position:"Query" name:"ScaVersion"`
	SearchInfoSub       string           `position:"Query" name:"SearchInfoSub"`
	SearchInfo          string           `position:"Query" name:"SearchInfo"`
	CurrentPage         requests.Integer `position:"Query" name:"CurrentPage"`
	BizType             string           `position:"Query" name:"BizType"`
	Port                string           `position:"Query" name:"Port"`
	Name                requests.Integer `position:"Query" name:"Name"`
	ScaName             string           `position:"Query" name:"ScaName"`
	ScaNamePattern      string           `position:"Query" name:"ScaNamePattern"`
	User                string           `position:"Query" name:"User"`
}

// DescribePropertyScaDetailResponse is the response struct for api DescribePropertyScaDetail
type DescribePropertyScaDetailResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo      `json:"PageInfo" xml:"PageInfo"`
	Propertys []PropertySca `json:"Propertys" xml:"Propertys"`
}

// CreateDescribePropertyScaDetailRequest creates a request to invoke DescribePropertyScaDetail API
func CreateDescribePropertyScaDetailRequest() (request *DescribePropertyScaDetailRequest) {
	request = &DescribePropertyScaDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertyScaDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePropertyScaDetailResponse creates a response to parse from DescribePropertyScaDetail response
func CreateDescribePropertyScaDetailResponse() (response *DescribePropertyScaDetailResponse) {
	response = &DescribePropertyScaDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
