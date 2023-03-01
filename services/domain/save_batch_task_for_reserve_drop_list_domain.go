package domain

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

// SaveBatchTaskForReserveDropListDomain invokes the domain.SaveBatchTaskForReserveDropListDomain API synchronously
func (client *Client) SaveBatchTaskForReserveDropListDomain(request *SaveBatchTaskForReserveDropListDomainRequest) (response *SaveBatchTaskForReserveDropListDomainResponse, err error) {
	response = CreateSaveBatchTaskForReserveDropListDomainResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForReserveDropListDomainWithChan invokes the domain.SaveBatchTaskForReserveDropListDomain API asynchronously
func (client *Client) SaveBatchTaskForReserveDropListDomainWithChan(request *SaveBatchTaskForReserveDropListDomainRequest) (<-chan *SaveBatchTaskForReserveDropListDomainResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForReserveDropListDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForReserveDropListDomain(request)
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

// SaveBatchTaskForReserveDropListDomainWithCallback invokes the domain.SaveBatchTaskForReserveDropListDomain API asynchronously
func (client *Client) SaveBatchTaskForReserveDropListDomainWithCallback(request *SaveBatchTaskForReserveDropListDomainRequest, callback func(response *SaveBatchTaskForReserveDropListDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForReserveDropListDomainResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForReserveDropListDomain(request)
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

// SaveBatchTaskForReserveDropListDomainRequest is the request struct for api SaveBatchTaskForReserveDropListDomain
type SaveBatchTaskForReserveDropListDomainRequest struct {
	*requests.RpcRequest
	Domains           *[]SaveBatchTaskForReserveDropListDomainDomains `position:"Query" name:"Domains"  type:"Repeated"`
	ContactTemplateId string                                          `position:"Query" name:"ContactTemplateId"`
}

// SaveBatchTaskForReserveDropListDomainDomains is a repeated param struct in SaveBatchTaskForReserveDropListDomainRequest
type SaveBatchTaskForReserveDropListDomainDomains struct {
	DomainName string `name:"DomainName"`
}

// SaveBatchTaskForReserveDropListDomainResponse is the response struct for api SaveBatchTaskForReserveDropListDomain
type SaveBatchTaskForReserveDropListDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForReserveDropListDomainRequest creates a request to invoke SaveBatchTaskForReserveDropListDomain API
func CreateSaveBatchTaskForReserveDropListDomainRequest() (request *SaveBatchTaskForReserveDropListDomainRequest) {
	request = &SaveBatchTaskForReserveDropListDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForReserveDropListDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForReserveDropListDomainResponse creates a response to parse from SaveBatchTaskForReserveDropListDomain response
func CreateSaveBatchTaskForReserveDropListDomainResponse() (response *SaveBatchTaskForReserveDropListDomainResponse) {
	response = &SaveBatchTaskForReserveDropListDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
