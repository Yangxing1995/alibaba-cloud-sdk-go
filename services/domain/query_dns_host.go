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

// QueryDnsHost invokes the domain.QueryDnsHost API synchronously
func (client *Client) QueryDnsHost(request *QueryDnsHostRequest) (response *QueryDnsHostResponse, err error) {
	response = CreateQueryDnsHostResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDnsHostWithChan invokes the domain.QueryDnsHost API asynchronously
func (client *Client) QueryDnsHostWithChan(request *QueryDnsHostRequest) (<-chan *QueryDnsHostResponse, <-chan error) {
	responseChan := make(chan *QueryDnsHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDnsHost(request)
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

// QueryDnsHostWithCallback invokes the domain.QueryDnsHost API asynchronously
func (client *Client) QueryDnsHostWithCallback(request *QueryDnsHostRequest, callback func(response *QueryDnsHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDnsHostResponse
		var err error
		defer close(result)
		response, err = client.QueryDnsHost(request)
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

// QueryDnsHostRequest is the request struct for api QueryDnsHost
type QueryDnsHostRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDnsHostResponse is the response struct for api QueryDnsHost
type QueryDnsHostResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	DnsHostList []DnsHost `json:"DnsHostList" xml:"DnsHostList"`
}

// CreateQueryDnsHostRequest creates a request to invoke QueryDnsHost API
func CreateQueryDnsHostRequest() (request *QueryDnsHostRequest) {
	request = &QueryDnsHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDnsHost", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDnsHostResponse creates a response to parse from QueryDnsHost response
func CreateQueryDnsHostResponse() (response *QueryDnsHostResponse) {
	response = &QueryDnsHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
