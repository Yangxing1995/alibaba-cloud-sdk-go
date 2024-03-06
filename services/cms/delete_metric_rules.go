package cms

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

// DeleteMetricRules invokes the cms.DeleteMetricRules API synchronously
func (client *Client) DeleteMetricRules(request *DeleteMetricRulesRequest) (response *DeleteMetricRulesResponse, err error) {
	response = CreateDeleteMetricRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMetricRulesWithChan invokes the cms.DeleteMetricRules API asynchronously
func (client *Client) DeleteMetricRulesWithChan(request *DeleteMetricRulesRequest) (<-chan *DeleteMetricRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteMetricRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMetricRules(request)
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

// DeleteMetricRulesWithCallback invokes the cms.DeleteMetricRules API asynchronously
func (client *Client) DeleteMetricRulesWithCallback(request *DeleteMetricRulesRequest, callback func(response *DeleteMetricRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMetricRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteMetricRules(request)
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

// DeleteMetricRulesRequest is the request struct for api DeleteMetricRules
type DeleteMetricRulesRequest struct {
	*requests.RpcRequest
	Id *[]string `position:"Query" name:"Id"  type:"Repeated"`
}

// DeleteMetricRulesResponse is the response struct for api DeleteMetricRules
type DeleteMetricRulesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteMetricRulesRequest creates a request to invoke DeleteMetricRules API
func CreateDeleteMetricRulesRequest() (request *DeleteMetricRulesRequest) {
	request = &DeleteMetricRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMetricRules", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMetricRulesResponse creates a response to parse from DeleteMetricRules response
func CreateDeleteMetricRulesResponse() (response *DeleteMetricRulesResponse) {
	response = &DeleteMetricRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
