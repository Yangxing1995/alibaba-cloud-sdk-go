package arms

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

// ListDashboardsByName invokes the arms.ListDashboardsByName API synchronously
func (client *Client) ListDashboardsByName(request *ListDashboardsByNameRequest) (response *ListDashboardsByNameResponse, err error) {
	response = CreateListDashboardsByNameResponse()
	err = client.DoAction(request, response)
	return
}

// ListDashboardsByNameWithChan invokes the arms.ListDashboardsByName API asynchronously
func (client *Client) ListDashboardsByNameWithChan(request *ListDashboardsByNameRequest) (<-chan *ListDashboardsByNameResponse, <-chan error) {
	responseChan := make(chan *ListDashboardsByNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDashboardsByName(request)
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

// ListDashboardsByNameWithCallback invokes the arms.ListDashboardsByName API asynchronously
func (client *Client) ListDashboardsByNameWithCallback(request *ListDashboardsByNameRequest, callback func(response *ListDashboardsByNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDashboardsByNameResponse
		var err error
		defer close(result)
		response, err = client.ListDashboardsByName(request)
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

// ListDashboardsByNameRequest is the request struct for api ListDashboardsByName
type ListDashboardsByNameRequest struct {
	*requests.RpcRequest
	DataSourceType   string           `position:"Query" name:"DataSourceType"`
	DashBoardName    string           `position:"Query" name:"DashBoardName"`
	ProductCode      string           `position:"Query" name:"ProductCode"`
	ClusterId        string           `position:"Query" name:"ClusterId"`
	OnlyQuery        requests.Boolean `position:"Query" name:"OnlyQuery"`
	GroupName        string           `position:"Query" name:"GroupName"`
	ClusterType      string           `position:"Query" name:"ClusterType"`
	DashBoardVersion string           `position:"Query" name:"DashBoardVersion"`
}

// ListDashboardsByNameResponse is the response struct for api ListDashboardsByName
type ListDashboardsByNameResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateListDashboardsByNameRequest creates a request to invoke ListDashboardsByName API
func CreateListDashboardsByNameRequest() (request *ListDashboardsByNameRequest) {
	request = &ListDashboardsByNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListDashboardsByName", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDashboardsByNameResponse creates a response to parse from ListDashboardsByName response
func CreateListDashboardsByNameResponse() (response *ListDashboardsByNameResponse) {
	response = &ListDashboardsByNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
