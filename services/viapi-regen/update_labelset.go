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

// UpdateLabelset invokes the viapi_regen.UpdateLabelset API synchronously
func (client *Client) UpdateLabelset(request *UpdateLabelsetRequest) (response *UpdateLabelsetResponse, err error) {
	response = CreateUpdateLabelsetResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLabelsetWithChan invokes the viapi_regen.UpdateLabelset API asynchronously
func (client *Client) UpdateLabelsetWithChan(request *UpdateLabelsetRequest) (<-chan *UpdateLabelsetResponse, <-chan error) {
	responseChan := make(chan *UpdateLabelsetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLabelset(request)
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

// UpdateLabelsetWithCallback invokes the viapi_regen.UpdateLabelset API asynchronously
func (client *Client) UpdateLabelsetWithCallback(request *UpdateLabelsetRequest, callback func(response *UpdateLabelsetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLabelsetResponse
		var err error
		defer close(result)
		response, err = client.UpdateLabelset(request)
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

// UpdateLabelsetRequest is the request struct for api UpdateLabelset
type UpdateLabelsetRequest struct {
	*requests.RpcRequest
	Description string           `position:"Body" name:"Description"`
	Id          requests.Integer `position:"Body" name:"Id"`
	TagUserList string           `position:"Body" name:"TagUserList"`
	UserOssUrl  string           `position:"Body" name:"UserOssUrl"`
	ObjectKey   string           `position:"Body" name:"ObjectKey"`
	Name        string           `position:"Body" name:"Name"`
}

// UpdateLabelsetResponse is the response struct for api UpdateLabelset
type UpdateLabelsetResponse struct {
	*responses.BaseResponse
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Code      string               `json:"Code" xml:"Code"`
	Data      DataInUpdateLabelset `json:"Data" xml:"Data"`
}

// CreateUpdateLabelsetRequest creates a request to invoke UpdateLabelset API
func CreateUpdateLabelsetRequest() (request *UpdateLabelsetRequest) {
	request = &UpdateLabelsetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "UpdateLabelset", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLabelsetResponse creates a response to parse from UpdateLabelset response
func CreateUpdateLabelsetResponse() (response *UpdateLabelsetResponse) {
	response = &UpdateLabelsetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
