package cdn

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

// DescribeCdnWafDomain invokes the cdn.DescribeCdnWafDomain API synchronously
func (client *Client) DescribeCdnWafDomain(request *DescribeCdnWafDomainRequest) (response *DescribeCdnWafDomainResponse, err error) {
	response = CreateDescribeCdnWafDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnWafDomainWithChan invokes the cdn.DescribeCdnWafDomain API asynchronously
func (client *Client) DescribeCdnWafDomainWithChan(request *DescribeCdnWafDomainRequest) (<-chan *DescribeCdnWafDomainResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnWafDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnWafDomain(request)
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

// DescribeCdnWafDomainWithCallback invokes the cdn.DescribeCdnWafDomain API asynchronously
func (client *Client) DescribeCdnWafDomainWithCallback(request *DescribeCdnWafDomainRequest, callback func(response *DescribeCdnWafDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnWafDomainResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnWafDomain(request)
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

// DescribeCdnWafDomainRequest is the request struct for api DescribeCdnWafDomain
type DescribeCdnWafDomainRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	DomainName      string `position:"Query" name:"DomainName"`
}

// DescribeCdnWafDomainResponse is the response struct for api DescribeCdnWafDomain
type DescribeCdnWafDomainResponse struct {
	*responses.BaseResponse
	TotalCount    int            `json:"TotalCount" xml:"TotalCount"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	OutPutDomains []OutPutDomain `json:"OutPutDomains" xml:"OutPutDomains"`
}

// CreateDescribeCdnWafDomainRequest creates a request to invoke DescribeCdnWafDomain API
func CreateDescribeCdnWafDomainRequest() (request *DescribeCdnWafDomainRequest) {
	request = &DescribeCdnWafDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnWafDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnWafDomainResponse creates a response to parse from DescribeCdnWafDomain response
func CreateDescribeCdnWafDomainResponse() (response *DescribeCdnWafDomainResponse) {
	response = &DescribeCdnWafDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
