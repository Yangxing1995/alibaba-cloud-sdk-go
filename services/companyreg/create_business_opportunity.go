package companyreg

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

// CreateBusinessOpportunity invokes the companyreg.CreateBusinessOpportunity API synchronously
func (client *Client) CreateBusinessOpportunity(request *CreateBusinessOpportunityRequest) (response *CreateBusinessOpportunityResponse, err error) {
	response = CreateCreateBusinessOpportunityResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBusinessOpportunityWithChan invokes the companyreg.CreateBusinessOpportunity API asynchronously
func (client *Client) CreateBusinessOpportunityWithChan(request *CreateBusinessOpportunityRequest) (<-chan *CreateBusinessOpportunityResponse, <-chan error) {
	responseChan := make(chan *CreateBusinessOpportunityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBusinessOpportunity(request)
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

// CreateBusinessOpportunityWithCallback invokes the companyreg.CreateBusinessOpportunity API asynchronously
func (client *Client) CreateBusinessOpportunityWithCallback(request *CreateBusinessOpportunityRequest, callback func(response *CreateBusinessOpportunityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBusinessOpportunityResponse
		var err error
		defer close(result)
		response, err = client.CreateBusinessOpportunity(request)
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

// CreateBusinessOpportunityRequest is the request struct for api CreateBusinessOpportunity
type CreateBusinessOpportunityRequest struct {
	*requests.RpcRequest
	Mobile      string           `position:"Query" name:"Mobile"`
	Source      requests.Integer `position:"Query" name:"Source"`
	VCode       string           `position:"Query" name:"VCode"`
	ContactName string           `position:"Query" name:"ContactName"`
	BizType     string           `position:"Query" name:"BizType"`
}

// CreateBusinessOpportunityResponse is the response struct for api CreateBusinessOpportunity
type CreateBusinessOpportunityResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateCreateBusinessOpportunityRequest creates a request to invoke CreateBusinessOpportunity API
func CreateCreateBusinessOpportunityRequest() (request *CreateBusinessOpportunityRequest) {
	request = &CreateBusinessOpportunityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "CreateBusinessOpportunity", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBusinessOpportunityResponse creates a response to parse from CreateBusinessOpportunity response
func CreateCreateBusinessOpportunityResponse() (response *CreateBusinessOpportunityResponse) {
	response = &CreateBusinessOpportunityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
