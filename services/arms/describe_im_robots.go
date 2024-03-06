package arms

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

// DescribeIMRobots invokes the arms.DescribeIMRobots API synchronously
func (client *Client) DescribeIMRobots(request *DescribeIMRobotsRequest) (response *DescribeIMRobotsResponse, err error) {
	response = CreateDescribeIMRobotsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIMRobotsWithChan invokes the arms.DescribeIMRobots API asynchronously
func (client *Client) DescribeIMRobotsWithChan(request *DescribeIMRobotsRequest) (<-chan *DescribeIMRobotsResponse, <-chan error) {
	responseChan := make(chan *DescribeIMRobotsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIMRobots(request)
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

// DescribeIMRobotsWithCallback invokes the arms.DescribeIMRobots API asynchronously
func (client *Client) DescribeIMRobotsWithCallback(request *DescribeIMRobotsRequest, callback func(response *DescribeIMRobotsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIMRobotsResponse
		var err error
		defer close(result)
		response, err = client.DescribeIMRobots(request)
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

// DescribeIMRobotsRequest is the request struct for api DescribeIMRobots
type DescribeIMRobotsRequest struct {
	*requests.RpcRequest
	Size      requests.Integer `position:"Query" name:"Size"`
	RobotIds  string           `position:"Query" name:"RobotIds"`
	RobotName string           `position:"Query" name:"RobotName"`
	Page      requests.Integer `position:"Query" name:"Page"`
}

// DescribeIMRobotsResponse is the response struct for api DescribeIMRobots
type DescribeIMRobotsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateDescribeIMRobotsRequest creates a request to invoke DescribeIMRobots API
func CreateDescribeIMRobotsRequest() (request *DescribeIMRobotsRequest) {
	request = &DescribeIMRobotsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DescribeIMRobots", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIMRobotsResponse creates a response to parse from DescribeIMRobots response
func CreateDescribeIMRobotsResponse() (response *DescribeIMRobotsResponse) {
	response = &DescribeIMRobotsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
