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

// DeleteGbDevice invokes the linkvisual.DeleteGbDevice API synchronously
func (client *Client) DeleteGbDevice(request *DeleteGbDeviceRequest) (response *DeleteGbDeviceResponse, err error) {
	response = CreateDeleteGbDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGbDeviceWithChan invokes the linkvisual.DeleteGbDevice API asynchronously
func (client *Client) DeleteGbDeviceWithChan(request *DeleteGbDeviceRequest) (<-chan *DeleteGbDeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteGbDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGbDevice(request)
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

// DeleteGbDeviceWithCallback invokes the linkvisual.DeleteGbDevice API asynchronously
func (client *Client) DeleteGbDeviceWithCallback(request *DeleteGbDeviceRequest, callback func(response *DeleteGbDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGbDeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteGbDevice(request)
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

// DeleteGbDeviceRequest is the request struct for api DeleteGbDevice
type DeleteGbDeviceRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// DeleteGbDeviceResponse is the response struct for api DeleteGbDevice
type DeleteGbDeviceResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteGbDeviceRequest creates a request to invoke DeleteGbDevice API
func CreateDeleteGbDeviceRequest() (request *DeleteGbDeviceRequest) {
	request = &DeleteGbDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "DeleteGbDevice", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGbDeviceResponse creates a response to parse from DeleteGbDevice response
func CreateDeleteGbDeviceResponse() (response *DeleteGbDeviceResponse) {
	response = &DeleteGbDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
