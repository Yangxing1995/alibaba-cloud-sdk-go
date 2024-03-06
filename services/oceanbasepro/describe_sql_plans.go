package oceanbasepro

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

// DescribeSQLPlans invokes the oceanbasepro.DescribeSQLPlans API synchronously
func (client *Client) DescribeSQLPlans(request *DescribeSQLPlansRequest) (response *DescribeSQLPlansResponse, err error) {
	response = CreateDescribeSQLPlansResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLPlansWithChan invokes the oceanbasepro.DescribeSQLPlans API asynchronously
func (client *Client) DescribeSQLPlansWithChan(request *DescribeSQLPlansRequest) (<-chan *DescribeSQLPlansResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLPlans(request)
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

// DescribeSQLPlansWithCallback invokes the oceanbasepro.DescribeSQLPlans API asynchronously
func (client *Client) DescribeSQLPlansWithCallback(request *DescribeSQLPlansRequest, callback func(response *DescribeSQLPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLPlansResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLPlans(request)
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

// DescribeSQLPlansRequest is the request struct for api DescribeSQLPlans
type DescribeSQLPlansRequest struct {
	*requests.RpcRequest
	SQLId    string `position:"Body" name:"SQLId"`
	TenantId string `position:"Body" name:"TenantId"`
}

// DescribeSQLPlansResponse is the response struct for api DescribeSQLPlans
type DescribeSQLPlansResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SQLPlans  []Data `json:"SQLPlans" xml:"SQLPlans"`
}

// CreateDescribeSQLPlansRequest creates a request to invoke DescribeSQLPlans API
func CreateDescribeSQLPlansRequest() (request *DescribeSQLPlansRequest) {
	request = &DescribeSQLPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeSQLPlans", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLPlansResponse creates a response to parse from DescribeSQLPlans response
func CreateDescribeSQLPlansResponse() (response *DescribeSQLPlansResponse) {
	response = &DescribeSQLPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
