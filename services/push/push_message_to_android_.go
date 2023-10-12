package push

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

// PushMessageToAndroid invokes the push.PushMessageToAndroid API synchronously
func (client *Client) PushMessageToAndroid(request *PushMessageToAndroidRequest) (response *PushMessageToAndroidResponse, err error) {
	response = CreatePushMessageToAndroidResponse()
	err = client.DoAction(request, response)
	return
}

// PushMessageToAndroidWithChan invokes the push.PushMessageToAndroid API asynchronously
func (client *Client) PushMessageToAndroidWithChan(request *PushMessageToAndroidRequest) (<-chan *PushMessageToAndroidResponse, <-chan error) {
	responseChan := make(chan *PushMessageToAndroidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushMessageToAndroid(request)
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

// PushMessageToAndroidWithCallback invokes the push.PushMessageToAndroid API asynchronously
func (client *Client) PushMessageToAndroidWithCallback(request *PushMessageToAndroidRequest, callback func(response *PushMessageToAndroidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushMessageToAndroidResponse
		var err error
		defer close(result)
		response, err = client.PushMessageToAndroid(request)
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

// PushMessageToAndroidRequest is the request struct for api PushMessageToAndroid
type PushMessageToAndroidRequest struct {
	*requests.RpcRequest
	StoreOffline requests.Boolean `position:"Query" name:"StoreOffline"`
	Title        string           `position:"Query" name:"Title"`
	Body         string           `position:"Query" name:"Body"`
	JobKey       string           `position:"Query" name:"JobKey"`
	Target       string           `position:"Query" name:"Target"`
	AppKey       requests.Integer `position:"Query" name:"AppKey"`
	TargetValue  string           `position:"Query" name:"TargetValue"`
}

// PushMessageToAndroidResponse is the response struct for api PushMessageToAndroid
type PushMessageToAndroidResponse struct {
	*responses.BaseResponse
	MessageId string `json:"MessageId" xml:"MessageId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePushMessageToAndroidRequest creates a request to invoke PushMessageToAndroid API
func CreatePushMessageToAndroidRequest() (request *PushMessageToAndroidRequest) {
	request = &PushMessageToAndroidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "PushMessageToAndroid", "", "")
	request.Method = requests.POST
	return
}

// CreatePushMessageToAndroidResponse creates a response to parse from PushMessageToAndroid response
func CreatePushMessageToAndroidResponse() (response *PushMessageToAndroidResponse) {
	response = &PushMessageToAndroidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
