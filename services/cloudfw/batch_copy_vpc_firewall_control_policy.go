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

// BatchCopyVpcFirewallControlPolicy invokes the cloudfw.BatchCopyVpcFirewallControlPolicy API synchronously
func (client *Client) BatchCopyVpcFirewallControlPolicy(request *BatchCopyVpcFirewallControlPolicyRequest) (response *BatchCopyVpcFirewallControlPolicyResponse, err error) {
	response = CreateBatchCopyVpcFirewallControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// BatchCopyVpcFirewallControlPolicyWithChan invokes the cloudfw.BatchCopyVpcFirewallControlPolicy API asynchronously
func (client *Client) BatchCopyVpcFirewallControlPolicyWithChan(request *BatchCopyVpcFirewallControlPolicyRequest) (<-chan *BatchCopyVpcFirewallControlPolicyResponse, <-chan error) {
	responseChan := make(chan *BatchCopyVpcFirewallControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchCopyVpcFirewallControlPolicy(request)
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

// BatchCopyVpcFirewallControlPolicyWithCallback invokes the cloudfw.BatchCopyVpcFirewallControlPolicy API asynchronously
func (client *Client) BatchCopyVpcFirewallControlPolicyWithCallback(request *BatchCopyVpcFirewallControlPolicyRequest, callback func(response *BatchCopyVpcFirewallControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchCopyVpcFirewallControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.BatchCopyVpcFirewallControlPolicy(request)
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

// BatchCopyVpcFirewallControlPolicyRequest is the request struct for api BatchCopyVpcFirewallControlPolicy
type BatchCopyVpcFirewallControlPolicyRequest struct {
	*requests.RpcRequest
	SourceIp            string `position:"Query" name:"SourceIp"`
	Lang                string `position:"Query" name:"Lang"`
	SourceVpcFirewallId string `position:"Query" name:"SourceVpcFirewallId"`
	TargetVpcFirewallId string `position:"Query" name:"TargetVpcFirewallId"`
}

// BatchCopyVpcFirewallControlPolicyResponse is the response struct for api BatchCopyVpcFirewallControlPolicy
type BatchCopyVpcFirewallControlPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchCopyVpcFirewallControlPolicyRequest creates a request to invoke BatchCopyVpcFirewallControlPolicy API
func CreateBatchCopyVpcFirewallControlPolicyRequest() (request *BatchCopyVpcFirewallControlPolicyRequest) {
	request = &BatchCopyVpcFirewallControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "BatchCopyVpcFirewallControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchCopyVpcFirewallControlPolicyResponse creates a response to parse from BatchCopyVpcFirewallControlPolicy response
func CreateBatchCopyVpcFirewallControlPolicyResponse() (response *BatchCopyVpcFirewallControlPolicyResponse) {
	response = &BatchCopyVpcFirewallControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
