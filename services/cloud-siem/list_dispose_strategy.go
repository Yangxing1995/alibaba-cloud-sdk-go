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

// ListDisposeStrategy invokes the cloud_siem.ListDisposeStrategy API synchronously
func (client *Client) ListDisposeStrategy(request *ListDisposeStrategyRequest) (response *ListDisposeStrategyResponse, err error) {
	response = CreateListDisposeStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// ListDisposeStrategyWithChan invokes the cloud_siem.ListDisposeStrategy API asynchronously
func (client *Client) ListDisposeStrategyWithChan(request *ListDisposeStrategyRequest) (<-chan *ListDisposeStrategyResponse, <-chan error) {
	responseChan := make(chan *ListDisposeStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDisposeStrategy(request)
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

// ListDisposeStrategyWithCallback invokes the cloud_siem.ListDisposeStrategy API asynchronously
func (client *Client) ListDisposeStrategyWithCallback(request *ListDisposeStrategyRequest, callback func(response *ListDisposeStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDisposeStrategyResponse
		var err error
		defer close(result)
		response, err = client.ListDisposeStrategy(request)
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

// ListDisposeStrategyRequest is the request struct for api ListDisposeStrategy
type ListDisposeStrategyRequest struct {
	*requests.RpcRequest
	EntityIdentity  string           `position:"Body" name:"EntityIdentity"`
	PlaybookName    string           `position:"Body" name:"PlaybookName"`
	PlaybookTypes   string           `position:"Body" name:"PlaybookTypes"`
	StartTime       requests.Integer `position:"Body" name:"StartTime"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
	OrderField      string           `position:"Body" name:"OrderField"`
	Order           string           `position:"Body" name:"Order"`
	SophonTaskId    string           `position:"Body" name:"SophonTaskId"`
	EffectiveStatus requests.Integer `position:"Body" name:"EffectiveStatus"`
	EndTime         requests.Integer `position:"Body" name:"EndTime"`
	CurrentPage     requests.Integer `position:"Body" name:"CurrentPage"`
	PlaybookUuid    string           `position:"Body" name:"PlaybookUuid"`
	EntityType      string           `position:"Body" name:"EntityType"`
}

// ListDisposeStrategyResponse is the response struct for api ListDisposeStrategy
type ListDisposeStrategyResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListDisposeStrategyRequest creates a request to invoke ListDisposeStrategy API
func CreateListDisposeStrategyRequest() (request *ListDisposeStrategyRequest) {
	request = &ListDisposeStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListDisposeStrategy", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDisposeStrategyResponse creates a response to parse from ListDisposeStrategy response
func CreateListDisposeStrategyResponse() (response *ListDisposeStrategyResponse) {
	response = &ListDisposeStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
