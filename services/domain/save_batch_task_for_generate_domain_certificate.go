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

// SaveBatchTaskForGenerateDomainCertificate invokes the domain.SaveBatchTaskForGenerateDomainCertificate API synchronously
func (client *Client) SaveBatchTaskForGenerateDomainCertificate(request *SaveBatchTaskForGenerateDomainCertificateRequest) (response *SaveBatchTaskForGenerateDomainCertificateResponse, err error) {
	response = CreateSaveBatchTaskForGenerateDomainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForGenerateDomainCertificateWithChan invokes the domain.SaveBatchTaskForGenerateDomainCertificate API asynchronously
func (client *Client) SaveBatchTaskForGenerateDomainCertificateWithChan(request *SaveBatchTaskForGenerateDomainCertificateRequest) (<-chan *SaveBatchTaskForGenerateDomainCertificateResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForGenerateDomainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForGenerateDomainCertificate(request)
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

// SaveBatchTaskForGenerateDomainCertificateWithCallback invokes the domain.SaveBatchTaskForGenerateDomainCertificate API asynchronously
func (client *Client) SaveBatchTaskForGenerateDomainCertificateWithCallback(request *SaveBatchTaskForGenerateDomainCertificateRequest, callback func(response *SaveBatchTaskForGenerateDomainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForGenerateDomainCertificateResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForGenerateDomainCertificate(request)
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

// SaveBatchTaskForGenerateDomainCertificateRequest is the request struct for api SaveBatchTaskForGenerateDomainCertificate
type SaveBatchTaskForGenerateDomainCertificateRequest struct {
	*requests.RpcRequest
	DomainNames  *[]string `position:"Query" name:"DomainNames"  type:"Json"`
	UserClientIp string    `position:"Query" name:"UserClientIp"`
	Lang         string    `position:"Query" name:"Lang"`
}

// SaveBatchTaskForGenerateDomainCertificateResponse is the response struct for api SaveBatchTaskForGenerateDomainCertificate
type SaveBatchTaskForGenerateDomainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForGenerateDomainCertificateRequest creates a request to invoke SaveBatchTaskForGenerateDomainCertificate API
func CreateSaveBatchTaskForGenerateDomainCertificateRequest() (request *SaveBatchTaskForGenerateDomainCertificateRequest) {
	request = &SaveBatchTaskForGenerateDomainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForGenerateDomainCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForGenerateDomainCertificateResponse creates a response to parse from SaveBatchTaskForGenerateDomainCertificate response
func CreateSaveBatchTaskForGenerateDomainCertificateResponse() (response *SaveBatchTaskForGenerateDomainCertificateResponse) {
	response = &SaveBatchTaskForGenerateDomainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
