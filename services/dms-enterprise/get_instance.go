package dms_enterprise

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

// GetInstance invokes the dms_enterprise.GetInstance API synchronously
func (client *Client) GetInstance(request *GetInstanceRequest) (response *GetInstanceResponse, err error) {
	response = CreateGetInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceWithChan invokes the dms_enterprise.GetInstance API asynchronously
func (client *Client) GetInstanceWithChan(request *GetInstanceRequest) (<-chan *GetInstanceResponse, <-chan error) {
	responseChan := make(chan *GetInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstance(request)
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

// GetInstanceWithCallback invokes the dms_enterprise.GetInstance API asynchronously
func (client *Client) GetInstanceWithCallback(request *GetInstanceRequest, callback func(response *GetInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetInstance(request)
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

// GetInstanceRequest is the request struct for api GetInstance
type GetInstanceRequest struct {
	*requests.RpcRequest
	Tid  requests.Integer `position:"Query" name:"Tid"`
	Sid  string           `position:"Query" name:"Sid"`
	Host string           `position:"Query" name:"Host"`
	Port requests.Integer `position:"Query" name:"Port"`
}

// GetInstanceResponse is the response struct for api GetInstance
type GetInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool     `json:"Success" xml:"Success"`
	Instance     Instance `json:"Instance" xml:"Instance"`
}

// CreateGetInstanceRequest creates a request to invoke GetInstance API
func CreateGetInstanceRequest() (request *GetInstanceRequest) {
	request = &GetInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetInstance", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetInstanceResponse creates a response to parse from GetInstance response
func CreateGetInstanceResponse() (response *GetInstanceResponse) {
	response = &GetInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
