package oms

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

// GetReadyFlagByProxy invokes the oms.GetReadyFlagByProxy API synchronously
func (client *Client) GetReadyFlagByProxy(request *GetReadyFlagByProxyRequest) (response *GetReadyFlagByProxyResponse, err error) {
	response = CreateGetReadyFlagByProxyResponse()
	err = client.DoAction(request, response)
	return
}

// GetReadyFlagByProxyWithChan invokes the oms.GetReadyFlagByProxy API asynchronously
func (client *Client) GetReadyFlagByProxyWithChan(request *GetReadyFlagByProxyRequest) (<-chan *GetReadyFlagByProxyResponse, <-chan error) {
	responseChan := make(chan *GetReadyFlagByProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetReadyFlagByProxy(request)
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

// GetReadyFlagByProxyWithCallback invokes the oms.GetReadyFlagByProxy API asynchronously
func (client *Client) GetReadyFlagByProxyWithCallback(request *GetReadyFlagByProxyRequest, callback func(response *GetReadyFlagByProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetReadyFlagByProxyResponse
		var err error
		defer close(result)
		response, err = client.GetReadyFlagByProxy(request)
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

// GetReadyFlagByProxyRequest is the request struct for api GetReadyFlagByProxy
type GetReadyFlagByProxyRequest struct {
	*requests.RpcRequest
	Filter         string           `position:"Query" name:"Filter"`
	DomainCode     string           `position:"Query" name:"DomainCode"`
	DataType       string           `position:"Query" name:"DataType"`
	NextToken      string           `position:"Query" name:"NextToken"`
	CompressEnable requests.Boolean `position:"Query" name:"CompressEnable"`
	ApiType        string           `position:"Query" name:"ApiType"`
	MaxResult      requests.Integer `position:"Query" name:"MaxResult"`
}

// GetReadyFlagByProxyResponse is the response struct for api GetReadyFlagByProxy
type GetReadyFlagByProxyResponse struct {
	*responses.BaseResponse
	DataType   string `json:"DataType" xml:"DataType"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	Data       string `json:"Data" xml:"Data"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
	ApiType    string `json:"ApiType" xml:"ApiType"`
	Compressed bool   `json:"Compressed" xml:"Compressed"`
}

// CreateGetReadyFlagByProxyRequest creates a request to invoke GetReadyFlagByProxy API
func CreateGetReadyFlagByProxyRequest() (request *GetReadyFlagByProxyRequest) {
	request = &GetReadyFlagByProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "GetReadyFlagByProxy", "", "")
	request.Method = requests.POST
	return
}

// CreateGetReadyFlagByProxyResponse creates a response to parse from GetReadyFlagByProxy response
func CreateGetReadyFlagByProxyResponse() (response *GetReadyFlagByProxyResponse) {
	response = &GetReadyFlagByProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
