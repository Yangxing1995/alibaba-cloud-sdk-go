package cloudfw

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

// DescribeControlPolicy invokes the cloudfw.DescribeControlPolicy API synchronously
func (client *Client) DescribeControlPolicy(request *DescribeControlPolicyRequest) (response *DescribeControlPolicyResponse, err error) {
	response = CreateDescribeControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeControlPolicyWithChan invokes the cloudfw.DescribeControlPolicy API asynchronously
func (client *Client) DescribeControlPolicyWithChan(request *DescribeControlPolicyRequest) (<-chan *DescribeControlPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeControlPolicy(request)
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

// DescribeControlPolicyWithCallback invokes the cloudfw.DescribeControlPolicy API asynchronously
func (client *Client) DescribeControlPolicyWithCallback(request *DescribeControlPolicyRequest, callback func(response *DescribeControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeControlPolicy(request)
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

// DescribeControlPolicyRequest is the request struct for api DescribeControlPolicy
type DescribeControlPolicyRequest struct {
	*requests.RpcRequest
	Release     string `position:"Query" name:"Release"`
	Destination string `position:"Query" name:"Destination"`
	Description string `position:"Query" name:"Description"`
	Source      string `position:"Query" name:"Source"`
	AclUuid     string `position:"Query" name:"AclUuid"`
	AclAction   string `position:"Query" name:"AclAction"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	PageSize    string `position:"Query" name:"PageSize"`
	IpVersion   string `position:"Query" name:"IpVersion"`
	Lang        string `position:"Query" name:"Lang"`
	Direction   string `position:"Query" name:"Direction"`
	CurrentPage string `position:"Query" name:"CurrentPage"`
	Proto       string `position:"Query" name:"Proto"`
}

// DescribeControlPolicyResponse is the response struct for api DescribeControlPolicy
type DescribeControlPolicyResponse struct {
	*responses.BaseResponse
	PageNo     string     `json:"PageNo" xml:"PageNo"`
	PageSize   string     `json:"PageSize" xml:"PageSize"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount string     `json:"TotalCount" xml:"TotalCount"`
	Policys    []DataItem `json:"Policys" xml:"Policys"`
}

// CreateDescribeControlPolicyRequest creates a request to invoke DescribeControlPolicy API
func CreateDescribeControlPolicyRequest() (request *DescribeControlPolicyRequest) {
	request = &DescribeControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeControlPolicyResponse creates a response to parse from DescribeControlPolicy response
func CreateDescribeControlPolicyResponse() (response *DescribeControlPolicyResponse) {
	response = &DescribeControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
