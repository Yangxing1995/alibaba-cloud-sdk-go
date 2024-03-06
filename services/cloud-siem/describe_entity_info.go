package cloud_siem

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

// DescribeEntityInfo invokes the cloud_siem.DescribeEntityInfo API synchronously
func (client *Client) DescribeEntityInfo(request *DescribeEntityInfoRequest) (response *DescribeEntityInfoResponse, err error) {
	response = CreateDescribeEntityInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEntityInfoWithChan invokes the cloud_siem.DescribeEntityInfo API asynchronously
func (client *Client) DescribeEntityInfoWithChan(request *DescribeEntityInfoRequest) (<-chan *DescribeEntityInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeEntityInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEntityInfo(request)
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

// DescribeEntityInfoWithCallback invokes the cloud_siem.DescribeEntityInfo API asynchronously
func (client *Client) DescribeEntityInfoWithCallback(request *DescribeEntityInfoRequest, callback func(response *DescribeEntityInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEntityInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeEntityInfo(request)
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

// DescribeEntityInfoRequest is the request struct for api DescribeEntityInfo
type DescribeEntityInfoRequest struct {
	*requests.RpcRequest
	EntityIdentity string           `position:"Body" name:"EntityIdentity"`
	EntityId       requests.Integer `position:"Body" name:"EntityId"`
	SophonTaskId   string           `position:"Body" name:"SophonTaskId"`
	IncidentUuid   string           `position:"Body" name:"IncidentUuid"`
}

// DescribeEntityInfoResponse is the response struct for api DescribeEntityInfo
type DescribeEntityInfoResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeEntityInfoRequest creates a request to invoke DescribeEntityInfo API
func CreateDescribeEntityInfoRequest() (request *DescribeEntityInfoRequest) {
	request = &DescribeEntityInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeEntityInfo", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEntityInfoResponse creates a response to parse from DescribeEntityInfo response
func CreateDescribeEntityInfoResponse() (response *DescribeEntityInfoResponse) {
	response = &DescribeEntityInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
