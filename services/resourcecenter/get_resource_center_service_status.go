package resourcecenter

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

// GetResourceCenterServiceStatus invokes the resourcecenter.GetResourceCenterServiceStatus API synchronously
func (client *Client) GetResourceCenterServiceStatus(request *GetResourceCenterServiceStatusRequest) (response *GetResourceCenterServiceStatusResponse, err error) {
	response = CreateGetResourceCenterServiceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceCenterServiceStatusWithChan invokes the resourcecenter.GetResourceCenterServiceStatus API asynchronously
func (client *Client) GetResourceCenterServiceStatusWithChan(request *GetResourceCenterServiceStatusRequest) (<-chan *GetResourceCenterServiceStatusResponse, <-chan error) {
	responseChan := make(chan *GetResourceCenterServiceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceCenterServiceStatus(request)
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

// GetResourceCenterServiceStatusWithCallback invokes the resourcecenter.GetResourceCenterServiceStatus API asynchronously
func (client *Client) GetResourceCenterServiceStatusWithCallback(request *GetResourceCenterServiceStatusRequest, callback func(response *GetResourceCenterServiceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceCenterServiceStatusResponse
		var err error
		defer close(result)
		response, err = client.GetResourceCenterServiceStatus(request)
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

// GetResourceCenterServiceStatusRequest is the request struct for api GetResourceCenterServiceStatus
type GetResourceCenterServiceStatusRequest struct {
	*requests.RpcRequest
}

// GetResourceCenterServiceStatusResponse is the response struct for api GetResourceCenterServiceStatus
type GetResourceCenterServiceStatusResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ServiceStatus string `json:"ServiceStatus" xml:"ServiceStatus"`
	InitialStatus string `json:"InitialStatus" xml:"InitialStatus"`
}

// CreateGetResourceCenterServiceStatusRequest creates a request to invoke GetResourceCenterServiceStatus API
func CreateGetResourceCenterServiceStatusRequest() (request *GetResourceCenterServiceStatusRequest) {
	request = &GetResourceCenterServiceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "GetResourceCenterServiceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetResourceCenterServiceStatusResponse creates a response to parse from GetResourceCenterServiceStatus response
func CreateGetResourceCenterServiceStatusResponse() (response *GetResourceCenterServiceStatusResponse) {
	response = &GetResourceCenterServiceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
