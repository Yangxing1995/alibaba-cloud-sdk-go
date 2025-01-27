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

// SaveQuickQuery invokes the cloud_siem.SaveQuickQuery API synchronously
func (client *Client) SaveQuickQuery(request *SaveQuickQueryRequest) (response *SaveQuickQueryResponse, err error) {
	response = CreateSaveQuickQueryResponse()
	err = client.DoAction(request, response)
	return
}

// SaveQuickQueryWithChan invokes the cloud_siem.SaveQuickQuery API asynchronously
func (client *Client) SaveQuickQueryWithChan(request *SaveQuickQueryRequest) (<-chan *SaveQuickQueryResponse, <-chan error) {
	responseChan := make(chan *SaveQuickQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveQuickQuery(request)
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

// SaveQuickQueryWithCallback invokes the cloud_siem.SaveQuickQuery API asynchronously
func (client *Client) SaveQuickQueryWithCallback(request *SaveQuickQueryRequest, callback func(response *SaveQuickQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveQuickQueryResponse
		var err error
		defer close(result)
		response, err = client.SaveQuickQuery(request)
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

// SaveQuickQueryRequest is the request struct for api SaveQuickQuery
type SaveQuickQueryRequest struct {
	*requests.RpcRequest
	Query       string `position:"Body" name:"Query"`
	DisplayName string `position:"Body" name:"DisplayName"`
}

// SaveQuickQueryResponse is the response struct for api SaveQuickQuery
type SaveQuickQueryResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	DyCode    string `json:"DyCode" xml:"DyCode"`
	DyMessage string `json:"DyMessage" xml:"DyMessage"`
}

// CreateSaveQuickQueryRequest creates a request to invoke SaveQuickQuery API
func CreateSaveQuickQueryRequest() (request *SaveQuickQueryRequest) {
	request = &SaveQuickQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "SaveQuickQuery", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveQuickQueryResponse creates a response to parse from SaveQuickQuery response
func CreateSaveQuickQueryResponse() (response *SaveQuickQueryResponse) {
	response = &SaveQuickQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
