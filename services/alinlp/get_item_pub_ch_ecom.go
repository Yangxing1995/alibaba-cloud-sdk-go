package alinlp

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

// GetItemPubChEcom invokes the alinlp.GetItemPubChEcom API synchronously
func (client *Client) GetItemPubChEcom(request *GetItemPubChEcomRequest) (response *GetItemPubChEcomResponse, err error) {
	response = CreateGetItemPubChEcomResponse()
	err = client.DoAction(request, response)
	return
}

// GetItemPubChEcomWithChan invokes the alinlp.GetItemPubChEcom API asynchronously
func (client *Client) GetItemPubChEcomWithChan(request *GetItemPubChEcomRequest) (<-chan *GetItemPubChEcomResponse, <-chan error) {
	responseChan := make(chan *GetItemPubChEcomResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetItemPubChEcom(request)
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

// GetItemPubChEcomWithCallback invokes the alinlp.GetItemPubChEcom API asynchronously
func (client *Client) GetItemPubChEcomWithCallback(request *GetItemPubChEcomRequest, callback func(response *GetItemPubChEcomResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetItemPubChEcomResponse
		var err error
		defer close(result)
		response, err = client.GetItemPubChEcom(request)
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

// GetItemPubChEcomRequest is the request struct for api GetItemPubChEcom
type GetItemPubChEcomRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	ImageUrl    string `position:"Body" name:"ImageUrl"`
	Text        string `position:"Body" name:"Text"`
}

// GetItemPubChEcomResponse is the response struct for api GetItemPubChEcom
type GetItemPubChEcomResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetItemPubChEcomRequest creates a request to invoke GetItemPubChEcom API
func CreateGetItemPubChEcomRequest() (request *GetItemPubChEcomRequest) {
	request = &GetItemPubChEcomRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetItemPubChEcom", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetItemPubChEcomResponse creates a response to parse from GetItemPubChEcom response
func CreateGetItemPubChEcomResponse() (response *GetItemPubChEcomResponse) {
	response = &GetItemPubChEcomResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
