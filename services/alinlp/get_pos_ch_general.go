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

// GetPosChGeneral invokes the alinlp.GetPosChGeneral API synchronously
func (client *Client) GetPosChGeneral(request *GetPosChGeneralRequest) (response *GetPosChGeneralResponse, err error) {
	response = CreateGetPosChGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// GetPosChGeneralWithChan invokes the alinlp.GetPosChGeneral API asynchronously
func (client *Client) GetPosChGeneralWithChan(request *GetPosChGeneralRequest) (<-chan *GetPosChGeneralResponse, <-chan error) {
	responseChan := make(chan *GetPosChGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPosChGeneral(request)
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

// GetPosChGeneralWithCallback invokes the alinlp.GetPosChGeneral API asynchronously
func (client *Client) GetPosChGeneralWithCallback(request *GetPosChGeneralRequest, callback func(response *GetPosChGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPosChGeneralResponse
		var err error
		defer close(result)
		response, err = client.GetPosChGeneral(request)
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

// GetPosChGeneralRequest is the request struct for api GetPosChGeneral
type GetPosChGeneralRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	TokenizerId string `position:"Body" name:"TokenizerId"`
	Text        string `position:"Body" name:"Text"`
	OutType     string `position:"Body" name:"OutType"`
}

// GetPosChGeneralResponse is the response struct for api GetPosChGeneral
type GetPosChGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetPosChGeneralRequest creates a request to invoke GetPosChGeneral API
func CreateGetPosChGeneralRequest() (request *GetPosChGeneralRequest) {
	request = &GetPosChGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetPosChGeneral", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPosChGeneralResponse creates a response to parse from GetPosChGeneral response
func CreateGetPosChGeneralResponse() (response *GetPosChGeneralResponse) {
	response = &GetPosChGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
