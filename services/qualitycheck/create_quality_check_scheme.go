package qualitycheck

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

// CreateQualityCheckScheme invokes the qualitycheck.CreateQualityCheckScheme API synchronously
func (client *Client) CreateQualityCheckScheme(request *CreateQualityCheckSchemeRequest) (response *CreateQualityCheckSchemeResponse, err error) {
	response = CreateCreateQualityCheckSchemeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQualityCheckSchemeWithChan invokes the qualitycheck.CreateQualityCheckScheme API asynchronously
func (client *Client) CreateQualityCheckSchemeWithChan(request *CreateQualityCheckSchemeRequest) (<-chan *CreateQualityCheckSchemeResponse, <-chan error) {
	responseChan := make(chan *CreateQualityCheckSchemeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQualityCheckScheme(request)
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

// CreateQualityCheckSchemeWithCallback invokes the qualitycheck.CreateQualityCheckScheme API asynchronously
func (client *Client) CreateQualityCheckSchemeWithCallback(request *CreateQualityCheckSchemeRequest, callback func(response *CreateQualityCheckSchemeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQualityCheckSchemeResponse
		var err error
		defer close(result)
		response, err = client.CreateQualityCheckScheme(request)
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

// CreateQualityCheckSchemeRequest is the request struct for api CreateQualityCheckScheme
type CreateQualityCheckSchemeRequest struct {
	*requests.RpcRequest
	JsonStr       string           `position:"Query" name:"jsonStr"`
	BaseMeAgentId requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// CreateQualityCheckSchemeResponse is the response struct for api CreateQualityCheckScheme
type CreateQualityCheckSchemeResponse struct {
	*responses.BaseResponse
	Data           int64                              `json:"Data" xml:"Data"`
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Success        bool                               `json:"Success" xml:"Success"`
	Code           string                             `json:"Code" xml:"Code"`
	Message        string                             `json:"Message" xml:"Message"`
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Messages       MessagesInCreateQualityCheckScheme `json:"Messages" xml:"Messages"`
}

// CreateCreateQualityCheckSchemeRequest creates a request to invoke CreateQualityCheckScheme API
func CreateCreateQualityCheckSchemeRequest() (request *CreateQualityCheckSchemeRequest) {
	request = &CreateQualityCheckSchemeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "CreateQualityCheckScheme", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateQualityCheckSchemeResponse creates a response to parse from CreateQualityCheckScheme response
func CreateCreateQualityCheckSchemeResponse() (response *CreateQualityCheckSchemeResponse) {
	response = &CreateQualityCheckSchemeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
