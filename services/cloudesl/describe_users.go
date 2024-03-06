package cloudesl

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

// DescribeUsers invokes the cloudesl.DescribeUsers API synchronously
func (client *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
	response = CreateDescribeUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUsersWithChan invokes the cloudesl.DescribeUsers API asynchronously
func (client *Client) DescribeUsersWithChan(request *DescribeUsersRequest) (<-chan *DescribeUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUsers(request)
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

// DescribeUsersWithCallback invokes the cloudesl.DescribeUsers API asynchronously
func (client *Client) DescribeUsersWithCallback(request *DescribeUsersRequest, callback func(response *DescribeUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeUsers(request)
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

// DescribeUsersRequest is the request struct for api DescribeUsers
type DescribeUsersRequest struct {
	*requests.RpcRequest
	ExtraParams string           `position:"Body" name:"ExtraParams"`
	UserId      string           `position:"Body" name:"UserId"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	UserType    string           `position:"Body" name:"UserType"`
	UserName    string           `position:"Body" name:"UserName"`
}

// DescribeUsersResponse is the response struct for api DescribeUsers
type DescribeUsersResponse struct {
	*responses.BaseResponse
	ErrorMessage   string     `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	Message        string     `json:"Message" xml:"Message"`
	PageNumber     int        `json:"PageNumber" xml:"PageNumber"`
	DynamicCode    string     `json:"DynamicCode" xml:"DynamicCode"`
	Code           string     `json:"Code" xml:"Code"`
	TotalCount     int        `json:"TotalCount" xml:"TotalCount"`
	DynamicMessage string     `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	PageSize       int        `json:"PageSize" xml:"PageSize"`
	Users          []UserInfo `json:"Users" xml:"Users"`
}

// CreateDescribeUsersRequest creates a request to invoke DescribeUsers API
func CreateDescribeUsersRequest() (request *DescribeUsersRequest) {
	request = &DescribeUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeUsers", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUsersResponse creates a response to parse from DescribeUsers response
func CreateDescribeUsersResponse() (response *DescribeUsersResponse) {
	response = &DescribeUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
