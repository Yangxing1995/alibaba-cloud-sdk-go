package linkvisual

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

// DeleteLocalFileUploadJob invokes the linkvisual.DeleteLocalFileUploadJob API synchronously
func (client *Client) DeleteLocalFileUploadJob(request *DeleteLocalFileUploadJobRequest) (response *DeleteLocalFileUploadJobResponse, err error) {
	response = CreateDeleteLocalFileUploadJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLocalFileUploadJobWithChan invokes the linkvisual.DeleteLocalFileUploadJob API asynchronously
func (client *Client) DeleteLocalFileUploadJobWithChan(request *DeleteLocalFileUploadJobRequest) (<-chan *DeleteLocalFileUploadJobResponse, <-chan error) {
	responseChan := make(chan *DeleteLocalFileUploadJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLocalFileUploadJob(request)
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

// DeleteLocalFileUploadJobWithCallback invokes the linkvisual.DeleteLocalFileUploadJob API asynchronously
func (client *Client) DeleteLocalFileUploadJobWithCallback(request *DeleteLocalFileUploadJobRequest, callback func(response *DeleteLocalFileUploadJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLocalFileUploadJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteLocalFileUploadJob(request)
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

// DeleteLocalFileUploadJobRequest is the request struct for api DeleteLocalFileUploadJob
type DeleteLocalFileUploadJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteLocalFileUploadJobResponse is the response struct for api DeleteLocalFileUploadJob
type DeleteLocalFileUploadJobResponse struct {
	*responses.BaseResponse
	Code         string                 `json:"Code" xml:"Code"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
}

// CreateDeleteLocalFileUploadJobRequest creates a request to invoke DeleteLocalFileUploadJob API
func CreateDeleteLocalFileUploadJobRequest() (request *DeleteLocalFileUploadJobRequest) {
	request = &DeleteLocalFileUploadJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "DeleteLocalFileUploadJob", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLocalFileUploadJobResponse creates a response to parse from DeleteLocalFileUploadJob response
func CreateDeleteLocalFileUploadJobResponse() (response *DeleteLocalFileUploadJobResponse) {
	response = &DeleteLocalFileUploadJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
