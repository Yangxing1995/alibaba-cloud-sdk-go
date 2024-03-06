package ddosbgp

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

// DescribeOnDemandInstanceStatus invokes the ddosbgp.DescribeOnDemandInstanceStatus API synchronously
func (client *Client) DescribeOnDemandInstanceStatus(request *DescribeOnDemandInstanceStatusRequest) (response *DescribeOnDemandInstanceStatusResponse, err error) {
	response = CreateDescribeOnDemandInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOnDemandInstanceStatusWithChan invokes the ddosbgp.DescribeOnDemandInstanceStatus API asynchronously
func (client *Client) DescribeOnDemandInstanceStatusWithChan(request *DescribeOnDemandInstanceStatusRequest) (<-chan *DescribeOnDemandInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeOnDemandInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOnDemandInstanceStatus(request)
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

// DescribeOnDemandInstanceStatusWithCallback invokes the ddosbgp.DescribeOnDemandInstanceStatus API asynchronously
func (client *Client) DescribeOnDemandInstanceStatusWithCallback(request *DescribeOnDemandInstanceStatusRequest, callback func(response *DescribeOnDemandInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOnDemandInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeOnDemandInstanceStatus(request)
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

// DescribeOnDemandInstanceStatusRequest is the request struct for api DescribeOnDemandInstanceStatus
type DescribeOnDemandInstanceStatusRequest struct {
	*requests.RpcRequest
	SourceIp       string    `position:"Query" name:"SourceIp"`
	InstanceIdList *[]string `position:"Query" name:"InstanceIdList"  type:"Repeated"`
}

// DescribeOnDemandInstanceStatusResponse is the response struct for api DescribeOnDemandInstanceStatus
type DescribeOnDemandInstanceStatusResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Instances []Instance `json:"Instances" xml:"Instances"`
}

// CreateDescribeOnDemandInstanceStatusRequest creates a request to invoke DescribeOnDemandInstanceStatus API
func CreateDescribeOnDemandInstanceStatusRequest() (request *DescribeOnDemandInstanceStatusRequest) {
	request = &DescribeOnDemandInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeOnDemandInstanceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeOnDemandInstanceStatusResponse creates a response to parse from DescribeOnDemandInstanceStatus response
func CreateDescribeOnDemandInstanceStatusResponse() (response *DescribeOnDemandInstanceStatusResponse) {
	response = &DescribeOnDemandInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
