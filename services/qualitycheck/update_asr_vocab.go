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

// UpdateAsrVocab invokes the qualitycheck.UpdateAsrVocab API synchronously
func (client *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (response *UpdateAsrVocabResponse, err error) {
	response = CreateUpdateAsrVocabResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAsrVocabWithChan invokes the qualitycheck.UpdateAsrVocab API asynchronously
func (client *Client) UpdateAsrVocabWithChan(request *UpdateAsrVocabRequest) (<-chan *UpdateAsrVocabResponse, <-chan error) {
	responseChan := make(chan *UpdateAsrVocabResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAsrVocab(request)
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

// UpdateAsrVocabWithCallback invokes the qualitycheck.UpdateAsrVocab API asynchronously
func (client *Client) UpdateAsrVocabWithCallback(request *UpdateAsrVocabRequest, callback func(response *UpdateAsrVocabResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAsrVocabResponse
		var err error
		defer close(result)
		response, err = client.UpdateAsrVocab(request)
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

// UpdateAsrVocabRequest is the request struct for api UpdateAsrVocab
type UpdateAsrVocabRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// UpdateAsrVocabResponse is the response struct for api UpdateAsrVocab
type UpdateAsrVocabResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateAsrVocabRequest creates a request to invoke UpdateAsrVocab API
func CreateUpdateAsrVocabRequest() (request *UpdateAsrVocabRequest) {
	request = &UpdateAsrVocabRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateAsrVocab", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAsrVocabResponse creates a response to parse from UpdateAsrVocab response
func CreateUpdateAsrVocabResponse() (response *UpdateAsrVocabResponse) {
	response = &UpdateAsrVocabResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
