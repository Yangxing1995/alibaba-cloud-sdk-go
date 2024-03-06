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

// CreateProduceForPartner invokes the companyreg.CreateProduceForPartner API synchronously
func (client *Client) CreateProduceForPartner(request *CreateProduceForPartnerRequest) (response *CreateProduceForPartnerResponse, err error) {
	response = CreateCreateProduceForPartnerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProduceForPartnerWithChan invokes the companyreg.CreateProduceForPartner API asynchronously
func (client *Client) CreateProduceForPartnerWithChan(request *CreateProduceForPartnerRequest) (<-chan *CreateProduceForPartnerResponse, <-chan error) {
	responseChan := make(chan *CreateProduceForPartnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProduceForPartner(request)
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

// CreateProduceForPartnerWithCallback invokes the companyreg.CreateProduceForPartner API asynchronously
func (client *Client) CreateProduceForPartnerWithCallback(request *CreateProduceForPartnerRequest, callback func(response *CreateProduceForPartnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProduceForPartnerResponse
		var err error
		defer close(result)
		response, err = client.CreateProduceForPartner(request)
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

// CreateProduceForPartnerRequest is the request struct for api CreateProduceForPartner
type CreateProduceForPartnerRequest struct {
	*requests.RpcRequest
	BizType string `position:"Query" name:"BizType"`
	ExtInfo string `position:"Query" name:"ExtInfo"`
	BizId   string `position:"Query" name:"BizId"`
}

// CreateProduceForPartnerResponse is the response struct for api CreateProduceForPartner
type CreateProduceForPartnerResponse struct {
	*responses.BaseResponse
	BizId     string `json:"BizId" xml:"BizId"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateProduceForPartnerRequest creates a request to invoke CreateProduceForPartner API
func CreateCreateProduceForPartnerRequest() (request *CreateProduceForPartnerRequest) {
	request = &CreateProduceForPartnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "CreateProduceForPartner", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProduceForPartnerResponse creates a response to parse from CreateProduceForPartner response
func CreateCreateProduceForPartnerResponse() (response *CreateProduceForPartnerResponse) {
	response = &CreateProduceForPartnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
