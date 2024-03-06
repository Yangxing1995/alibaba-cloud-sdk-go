package iot

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

// QueryConsumerGroupByGroupId invokes the iot.QueryConsumerGroupByGroupId API synchronously
func (client *Client) QueryConsumerGroupByGroupId(request *QueryConsumerGroupByGroupIdRequest) (response *QueryConsumerGroupByGroupIdResponse, err error) {
	response = CreateQueryConsumerGroupByGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryConsumerGroupByGroupIdWithChan invokes the iot.QueryConsumerGroupByGroupId API asynchronously
func (client *Client) QueryConsumerGroupByGroupIdWithChan(request *QueryConsumerGroupByGroupIdRequest) (<-chan *QueryConsumerGroupByGroupIdResponse, <-chan error) {
	responseChan := make(chan *QueryConsumerGroupByGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryConsumerGroupByGroupId(request)
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

// QueryConsumerGroupByGroupIdWithCallback invokes the iot.QueryConsumerGroupByGroupId API asynchronously
func (client *Client) QueryConsumerGroupByGroupIdWithCallback(request *QueryConsumerGroupByGroupIdRequest, callback func(response *QueryConsumerGroupByGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryConsumerGroupByGroupIdResponse
		var err error
		defer close(result)
		response, err = client.QueryConsumerGroupByGroupId(request)
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

// QueryConsumerGroupByGroupIdRequest is the request struct for api QueryConsumerGroupByGroupId
type QueryConsumerGroupByGroupIdRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	GroupId           string `position:"Query" name:"GroupId"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// QueryConsumerGroupByGroupIdResponse is the response struct for api QueryConsumerGroupByGroupId
type QueryConsumerGroupByGroupIdResponse struct {
	*responses.BaseResponse
	RequestId    string                            `json:"RequestId" xml:"RequestId"`
	Success      bool                              `json:"Success" xml:"Success"`
	ErrorMessage string                            `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                            `json:"Code" xml:"Code"`
	Data         DataInQueryConsumerGroupByGroupId `json:"Data" xml:"Data"`
}

// CreateQueryConsumerGroupByGroupIdRequest creates a request to invoke QueryConsumerGroupByGroupId API
func CreateQueryConsumerGroupByGroupIdRequest() (request *QueryConsumerGroupByGroupIdRequest) {
	request = &QueryConsumerGroupByGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryConsumerGroupByGroupId", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryConsumerGroupByGroupIdResponse creates a response to parse from QueryConsumerGroupByGroupId response
func CreateQueryConsumerGroupByGroupIdResponse() (response *QueryConsumerGroupByGroupIdResponse) {
	response = &QueryConsumerGroupByGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
