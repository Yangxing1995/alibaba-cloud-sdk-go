package pts

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

// GetUserVpcSecurityGroup invokes the pts.GetUserVpcSecurityGroup API synchronously
func (client *Client) GetUserVpcSecurityGroup(request *GetUserVpcSecurityGroupRequest) (response *GetUserVpcSecurityGroupResponse, err error) {
	response = CreateGetUserVpcSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserVpcSecurityGroupWithChan invokes the pts.GetUserVpcSecurityGroup API asynchronously
func (client *Client) GetUserVpcSecurityGroupWithChan(request *GetUserVpcSecurityGroupRequest) (<-chan *GetUserVpcSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *GetUserVpcSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserVpcSecurityGroup(request)
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

// GetUserVpcSecurityGroupWithCallback invokes the pts.GetUserVpcSecurityGroup API asynchronously
func (client *Client) GetUserVpcSecurityGroupWithCallback(request *GetUserVpcSecurityGroupRequest, callback func(response *GetUserVpcSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserVpcSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.GetUserVpcSecurityGroup(request)
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

// GetUserVpcSecurityGroupRequest is the request struct for api GetUserVpcSecurityGroup
type GetUserVpcSecurityGroupRequest struct {
	*requests.RpcRequest
	VpcId      string           `position:"Query" name:"VpcId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// GetUserVpcSecurityGroupResponse is the response struct for api GetUserVpcSecurityGroup
type GetUserVpcSecurityGroupResponse struct {
	*responses.BaseResponse
	SecurityGroupCount int             `json:"SecurityGroupCount" xml:"SecurityGroupCount"`
	RequestId          string          `json:"RequestId" xml:"RequestId"`
	Message            string          `json:"Message" xml:"Message"`
	PageSize           int             `json:"PageSize" xml:"PageSize"`
	PageNumber         int             `json:"PageNumber" xml:"PageNumber"`
	HttpStatusCode     int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code               string          `json:"Code" xml:"Code"`
	Success            bool            `json:"Success" xml:"Success"`
	SecurityGroupList  []SecurityGroup `json:"SecurityGroupList" xml:"SecurityGroupList"`
}

// CreateGetUserVpcSecurityGroupRequest creates a request to invoke GetUserVpcSecurityGroup API
func CreateGetUserVpcSecurityGroupRequest() (request *GetUserVpcSecurityGroupRequest) {
	request = &GetUserVpcSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetUserVpcSecurityGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateGetUserVpcSecurityGroupResponse creates a response to parse from GetUserVpcSecurityGroup response
func CreateGetUserVpcSecurityGroupResponse() (response *GetUserVpcSecurityGroupResponse) {
	response = &GetUserVpcSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
