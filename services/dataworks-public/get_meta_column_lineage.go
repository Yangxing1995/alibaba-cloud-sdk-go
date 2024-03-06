package dataworks_public

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

// GetMetaColumnLineage invokes the dataworks_public.GetMetaColumnLineage API synchronously
func (client *Client) GetMetaColumnLineage(request *GetMetaColumnLineageRequest) (response *GetMetaColumnLineageResponse, err error) {
	response = CreateGetMetaColumnLineageResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaColumnLineageWithChan invokes the dataworks_public.GetMetaColumnLineage API asynchronously
func (client *Client) GetMetaColumnLineageWithChan(request *GetMetaColumnLineageRequest) (<-chan *GetMetaColumnLineageResponse, <-chan error) {
	responseChan := make(chan *GetMetaColumnLineageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaColumnLineage(request)
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

// GetMetaColumnLineageWithCallback invokes the dataworks_public.GetMetaColumnLineage API asynchronously
func (client *Client) GetMetaColumnLineageWithCallback(request *GetMetaColumnLineageRequest, callback func(response *GetMetaColumnLineageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaColumnLineageResponse
		var err error
		defer close(result)
		response, err = client.GetMetaColumnLineage(request)
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

// GetMetaColumnLineageRequest is the request struct for api GetMetaColumnLineage
type GetMetaColumnLineageRequest struct {
	*requests.RpcRequest
	DataSourceType string           `position:"Query" name:"DataSourceType"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	ColumnName     string           `position:"Query" name:"ColumnName"`
	ColumnGuid     string           `position:"Query" name:"ColumnGuid"`
	DatabaseName   string           `position:"Query" name:"DatabaseName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	TableName      string           `position:"Query" name:"TableName"`
	Direction      string           `position:"Query" name:"Direction"`
}

// GetMetaColumnLineageResponse is the response struct for api GetMetaColumnLineage
type GetMetaColumnLineageResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string                     `json:"RequestId" xml:"RequestId"`
	Success        bool                       `json:"Success" xml:"Success"`
	ErrorCode      string                     `json:"ErrorCode" xml:"ErrorCode"`
	Data           DataInGetMetaColumnLineage `json:"Data" xml:"Data"`
}

// CreateGetMetaColumnLineageRequest creates a request to invoke GetMetaColumnLineage API
func CreateGetMetaColumnLineageRequest() (request *GetMetaColumnLineageRequest) {
	request = &GetMetaColumnLineageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMetaColumnLineage", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMetaColumnLineageResponse creates a response to parse from GetMetaColumnLineage response
func CreateGetMetaColumnLineageResponse() (response *GetMetaColumnLineageResponse) {
	response = &GetMetaColumnLineageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
