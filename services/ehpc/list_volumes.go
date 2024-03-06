package ehpc

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

// ListVolumes invokes the ehpc.ListVolumes API synchronously
func (client *Client) ListVolumes(request *ListVolumesRequest) (response *ListVolumesResponse, err error) {
	response = CreateListVolumesResponse()
	err = client.DoAction(request, response)
	return
}

// ListVolumesWithChan invokes the ehpc.ListVolumes API asynchronously
func (client *Client) ListVolumesWithChan(request *ListVolumesRequest) (<-chan *ListVolumesResponse, <-chan error) {
	responseChan := make(chan *ListVolumesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVolumes(request)
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

// ListVolumesWithCallback invokes the ehpc.ListVolumes API asynchronously
func (client *Client) ListVolumesWithCallback(request *ListVolumesRequest, callback func(response *ListVolumesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVolumesResponse
		var err error
		defer close(result)
		response, err = client.ListVolumes(request)
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

// ListVolumesRequest is the request struct for api ListVolumes
type ListVolumesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListVolumesResponse is the response struct for api ListVolumes
type ListVolumesResponse struct {
	*responses.BaseResponse
	PageSize   int                  `json:"PageSize" xml:"PageSize"`
	RequestId  string               `json:"RequestId" xml:"RequestId"`
	PageNumber int                  `json:"PageNumber" xml:"PageNumber"`
	TotalCount int                  `json:"TotalCount" xml:"TotalCount"`
	Volumes    VolumesInListVolumes `json:"Volumes" xml:"Volumes"`
}

// CreateListVolumesRequest creates a request to invoke ListVolumes API
func CreateListVolumesRequest() (request *ListVolumesRequest) {
	request = &ListVolumesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListVolumes", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListVolumesResponse creates a response to parse from ListVolumes response
func CreateListVolumesResponse() (response *ListVolumesResponse) {
	response = &ListVolumesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
