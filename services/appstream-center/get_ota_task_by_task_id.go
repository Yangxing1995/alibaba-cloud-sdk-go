package appstream_center

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

// GetOtaTaskByTaskId invokes the appstream_center.GetOtaTaskByTaskId API synchronously
func (client *Client) GetOtaTaskByTaskId(request *GetOtaTaskByTaskIdRequest) (response *GetOtaTaskByTaskIdResponse, err error) {
	response = CreateGetOtaTaskByTaskIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetOtaTaskByTaskIdWithChan invokes the appstream_center.GetOtaTaskByTaskId API asynchronously
func (client *Client) GetOtaTaskByTaskIdWithChan(request *GetOtaTaskByTaskIdRequest) (<-chan *GetOtaTaskByTaskIdResponse, <-chan error) {
	responseChan := make(chan *GetOtaTaskByTaskIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOtaTaskByTaskId(request)
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

// GetOtaTaskByTaskIdWithCallback invokes the appstream_center.GetOtaTaskByTaskId API asynchronously
func (client *Client) GetOtaTaskByTaskIdWithCallback(request *GetOtaTaskByTaskIdRequest, callback func(response *GetOtaTaskByTaskIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOtaTaskByTaskIdResponse
		var err error
		defer close(result)
		response, err = client.GetOtaTaskByTaskId(request)
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

// GetOtaTaskByTaskIdRequest is the request struct for api GetOtaTaskByTaskId
type GetOtaTaskByTaskIdRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Body" name:"TaskId"`
}

// GetOtaTaskByTaskIdResponse is the response struct for api GetOtaTaskByTaskId
type GetOtaTaskByTaskIdResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Code          string `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ReleaseNote   string `json:"ReleaseNote" xml:"ReleaseNote"`
	OtaVersion    string `json:"OtaVersion" xml:"OtaVersion"`
	TaskStartTime string `json:"TaskStartTime" xml:"TaskStartTime"`
}

// CreateGetOtaTaskByTaskIdRequest creates a request to invoke GetOtaTaskByTaskId API
func CreateGetOtaTaskByTaskIdRequest() (request *GetOtaTaskByTaskIdRequest) {
	request = &GetOtaTaskByTaskIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "GetOtaTaskByTaskId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOtaTaskByTaskIdResponse creates a response to parse from GetOtaTaskByTaskId response
func CreateGetOtaTaskByTaskIdResponse() (response *GetOtaTaskByTaskIdResponse) {
	response = &GetOtaTaskByTaskIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
