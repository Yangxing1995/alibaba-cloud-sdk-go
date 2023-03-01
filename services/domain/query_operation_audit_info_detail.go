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

// QueryOperationAuditInfoDetail invokes the domain.QueryOperationAuditInfoDetail API synchronously
func (client *Client) QueryOperationAuditInfoDetail(request *QueryOperationAuditInfoDetailRequest) (response *QueryOperationAuditInfoDetailResponse, err error) {
	response = CreateQueryOperationAuditInfoDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryOperationAuditInfoDetailWithChan invokes the domain.QueryOperationAuditInfoDetail API asynchronously
func (client *Client) QueryOperationAuditInfoDetailWithChan(request *QueryOperationAuditInfoDetailRequest) (<-chan *QueryOperationAuditInfoDetailResponse, <-chan error) {
	responseChan := make(chan *QueryOperationAuditInfoDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryOperationAuditInfoDetail(request)
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

// QueryOperationAuditInfoDetailWithCallback invokes the domain.QueryOperationAuditInfoDetail API asynchronously
func (client *Client) QueryOperationAuditInfoDetailWithCallback(request *QueryOperationAuditInfoDetailRequest, callback func(response *QueryOperationAuditInfoDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryOperationAuditInfoDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryOperationAuditInfoDetail(request)
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

// QueryOperationAuditInfoDetailRequest is the request struct for api QueryOperationAuditInfoDetail
type QueryOperationAuditInfoDetailRequest struct {
	*requests.RpcRequest
	AuditRecordId requests.Integer `position:"Query" name:"AuditRecordId"`
	Lang          string           `position:"Query" name:"Lang"`
}

// QueryOperationAuditInfoDetailResponse is the response struct for api QueryOperationAuditInfoDetail
type QueryOperationAuditInfoDetailResponse struct {
	*responses.BaseResponse
	AuditInfo    string `json:"AuditInfo" xml:"AuditInfo"`
	AuditStatus  int    `json:"AuditStatus" xml:"AuditStatus"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	BusinessName string `json:"BusinessName" xml:"BusinessName"`
	AuditType    int    `json:"AuditType" xml:"AuditType"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`
	UpdateTime   int64  `json:"UpdateTime" xml:"UpdateTime"`
	Id           string `json:"Id" xml:"Id"`
	Remark       string `json:"Remark" xml:"Remark"`
}

// CreateQueryOperationAuditInfoDetailRequest creates a request to invoke QueryOperationAuditInfoDetail API
func CreateQueryOperationAuditInfoDetailRequest() (request *QueryOperationAuditInfoDetailRequest) {
	request = &QueryOperationAuditInfoDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryOperationAuditInfoDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryOperationAuditInfoDetailResponse creates a response to parse from QueryOperationAuditInfoDetail response
func CreateQueryOperationAuditInfoDetailResponse() (response *QueryOperationAuditInfoDetailResponse) {
	response = &QueryOperationAuditInfoDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
