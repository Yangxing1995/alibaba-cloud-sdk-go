package dms_enterprise

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

// EditLogicDatabase invokes the dms_enterprise.EditLogicDatabase API synchronously
func (client *Client) EditLogicDatabase(request *EditLogicDatabaseRequest) (response *EditLogicDatabaseResponse, err error) {
	response = CreateEditLogicDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// EditLogicDatabaseWithChan invokes the dms_enterprise.EditLogicDatabase API asynchronously
func (client *Client) EditLogicDatabaseWithChan(request *EditLogicDatabaseRequest) (<-chan *EditLogicDatabaseResponse, <-chan error) {
	responseChan := make(chan *EditLogicDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditLogicDatabase(request)
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

// EditLogicDatabaseWithCallback invokes the dms_enterprise.EditLogicDatabase API asynchronously
func (client *Client) EditLogicDatabaseWithCallback(request *EditLogicDatabaseRequest, callback func(response *EditLogicDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditLogicDatabaseResponse
		var err error
		defer close(result)
		response, err = client.EditLogicDatabase(request)
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

// EditLogicDatabaseRequest is the request struct for api EditLogicDatabase
type EditLogicDatabaseRequest struct {
	*requests.RpcRequest
	LogicDbId   requests.Integer `position:"Query" name:"LogicDbId"`
	Tid         requests.Integer `position:"Query" name:"Tid"`
	Alias       string           `position:"Query" name:"Alias"`
	DatabaseIds *[]string        `position:"Query" name:"DatabaseIds"  type:"Json"`
}

// EditLogicDatabaseResponse is the response struct for api EditLogicDatabase
type EditLogicDatabaseResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateEditLogicDatabaseRequest creates a request to invoke EditLogicDatabase API
func CreateEditLogicDatabaseRequest() (request *EditLogicDatabaseRequest) {
	request = &EditLogicDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "EditLogicDatabase", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEditLogicDatabaseResponse creates a response to parse from EditLogicDatabase response
func CreateEditLogicDatabaseResponse() (response *EditLogicDatabaseResponse) {
	response = &EditLogicDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
