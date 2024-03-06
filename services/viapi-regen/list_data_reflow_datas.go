package viapi_regen

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

// ListDataReflowDatas invokes the viapi_regen.ListDataReflowDatas API synchronously
func (client *Client) ListDataReflowDatas(request *ListDataReflowDatasRequest) (response *ListDataReflowDatasResponse, err error) {
	response = CreateListDataReflowDatasResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataReflowDatasWithChan invokes the viapi_regen.ListDataReflowDatas API asynchronously
func (client *Client) ListDataReflowDatasWithChan(request *ListDataReflowDatasRequest) (<-chan *ListDataReflowDatasResponse, <-chan error) {
	responseChan := make(chan *ListDataReflowDatasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataReflowDatas(request)
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

// ListDataReflowDatasWithCallback invokes the viapi_regen.ListDataReflowDatas API asynchronously
func (client *Client) ListDataReflowDatasWithCallback(request *ListDataReflowDatasRequest, callback func(response *ListDataReflowDatasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataReflowDatasResponse
		var err error
		defer close(result)
		response, err = client.ListDataReflowDatas(request)
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

// ListDataReflowDatasRequest is the request struct for api ListDataReflowDatas
type ListDataReflowDatasRequest struct {
	*requests.RpcRequest
	StartTime   requests.Integer `position:"Body" name:"StartTime"`
	ImageName   string           `position:"Body" name:"ImageName"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	EndTime     requests.Integer `position:"Body" name:"EndTime"`
	CurrentPage requests.Integer `position:"Body" name:"CurrentPage"`
	ServiceId   requests.Integer `position:"Body" name:"ServiceId"`
	Category    string           `position:"Body" name:"Category"`
}

// ListDataReflowDatasResponse is the response struct for api ListDataReflowDatas
type ListDataReflowDatasResponse struct {
	*responses.BaseResponse
	Message   string                    `json:"Message" xml:"Message"`
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Code      string                    `json:"Code" xml:"Code"`
	Data      DataInListDataReflowDatas `json:"Data" xml:"Data"`
}

// CreateListDataReflowDatasRequest creates a request to invoke ListDataReflowDatas API
func CreateListDataReflowDatasRequest() (request *ListDataReflowDatasRequest) {
	request = &ListDataReflowDatasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "ListDataReflowDatas", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataReflowDatasResponse creates a response to parse from ListDataReflowDatas response
func CreateListDataReflowDatasResponse() (response *ListDataReflowDatasResponse) {
	response = &ListDataReflowDatasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
