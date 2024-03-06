package live

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

// CreateLiveMessageGroup invokes the live.CreateLiveMessageGroup API synchronously
func (client *Client) CreateLiveMessageGroup(request *CreateLiveMessageGroupRequest) (response *CreateLiveMessageGroupResponse, err error) {
	response = CreateCreateLiveMessageGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLiveMessageGroupWithChan invokes the live.CreateLiveMessageGroup API asynchronously
func (client *Client) CreateLiveMessageGroupWithChan(request *CreateLiveMessageGroupRequest) (<-chan *CreateLiveMessageGroupResponse, <-chan error) {
	responseChan := make(chan *CreateLiveMessageGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLiveMessageGroup(request)
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

// CreateLiveMessageGroupWithCallback invokes the live.CreateLiveMessageGroup API asynchronously
func (client *Client) CreateLiveMessageGroupWithCallback(request *CreateLiveMessageGroupRequest, callback func(response *CreateLiveMessageGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveMessageGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateLiveMessageGroup(request)
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

// CreateLiveMessageGroupRequest is the request struct for api CreateLiveMessageGroup
type CreateLiveMessageGroupRequest struct {
	*requests.RpcRequest
	GroupId        string    `position:"Query" name:"GroupId"`
	GroupInfo      string    `position:"Query" name:"GroupInfo"`
	CreatorId      string    `position:"Query" name:"CreatorId"`
	DataCenter     string    `position:"Query" name:"DataCenter"`
	GroupName      string    `position:"Query" name:"GroupName"`
	AppId          string    `position:"Query" name:"AppId"`
	Administrators *[]string `position:"Query" name:"Administrators"  type:"Repeated"`
}

// CreateLiveMessageGroupResponse is the response struct for api CreateLiveMessageGroup
type CreateLiveMessageGroupResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	GroupId       string `json:"GroupId" xml:"GroupId"`
	AlreadyExists bool   `json:"AlreadyExists" xml:"AlreadyExists"`
	AlreadyDelete bool   `json:"AlreadyDelete" xml:"AlreadyDelete"`
}

// CreateCreateLiveMessageGroupRequest creates a request to invoke CreateLiveMessageGroup API
func CreateCreateLiveMessageGroupRequest() (request *CreateLiveMessageGroupRequest) {
	request = &CreateLiveMessageGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateLiveMessageGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLiveMessageGroupResponse creates a response to parse from CreateLiveMessageGroup response
func CreateCreateLiveMessageGroupResponse() (response *CreateLiveMessageGroupResponse) {
	response = &CreateLiveMessageGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
