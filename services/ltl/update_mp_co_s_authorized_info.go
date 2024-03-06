package ltl

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

// UpdateMPCoSAuthorizedInfo invokes the ltl.UpdateMPCoSAuthorizedInfo API synchronously
func (client *Client) UpdateMPCoSAuthorizedInfo(request *UpdateMPCoSAuthorizedInfoRequest) (response *UpdateMPCoSAuthorizedInfoResponse, err error) {
	response = CreateUpdateMPCoSAuthorizedInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMPCoSAuthorizedInfoWithChan invokes the ltl.UpdateMPCoSAuthorizedInfo API asynchronously
func (client *Client) UpdateMPCoSAuthorizedInfoWithChan(request *UpdateMPCoSAuthorizedInfoRequest) (<-chan *UpdateMPCoSAuthorizedInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateMPCoSAuthorizedInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMPCoSAuthorizedInfo(request)
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

// UpdateMPCoSAuthorizedInfoWithCallback invokes the ltl.UpdateMPCoSAuthorizedInfo API asynchronously
func (client *Client) UpdateMPCoSAuthorizedInfoWithCallback(request *UpdateMPCoSAuthorizedInfoRequest, callback func(response *UpdateMPCoSAuthorizedInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMPCoSAuthorizedInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateMPCoSAuthorizedInfo(request)
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

// UpdateMPCoSAuthorizedInfoRequest is the request struct for api UpdateMPCoSAuthorizedInfo
type UpdateMPCoSAuthorizedInfoRequest struct {
	*requests.RpcRequest
	PhaseGroupId        string                 `position:"Query" name:"PhaseGroupId"`
	ApiVersion          string                 `position:"Query" name:"ApiVersion"`
	AuthorizedPhaseList map[string]interface{} `position:"Query" name:"AuthorizedPhaseList"`
	BizChainId          string                 `position:"Query" name:"BizChainId"`
	MemberId            string                 `position:"Query" name:"MemberId"`
}

// UpdateMPCoSAuthorizedInfoResponse is the response struct for api UpdateMPCoSAuthorizedInfo
type UpdateMPCoSAuthorizedInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateMPCoSAuthorizedInfoRequest creates a request to invoke UpdateMPCoSAuthorizedInfo API
func CreateUpdateMPCoSAuthorizedInfoRequest() (request *UpdateMPCoSAuthorizedInfoRequest) {
	request = &UpdateMPCoSAuthorizedInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "UpdateMPCoSAuthorizedInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateMPCoSAuthorizedInfoResponse creates a response to parse from UpdateMPCoSAuthorizedInfo response
func CreateUpdateMPCoSAuthorizedInfoResponse() (response *UpdateMPCoSAuthorizedInfoResponse) {
	response = &UpdateMPCoSAuthorizedInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
