package imagerecog

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

// RecognizeScene invokes the imagerecog.RecognizeScene API synchronously
func (client *Client) RecognizeScene(request *RecognizeSceneRequest) (response *RecognizeSceneResponse, err error) {
	response = CreateRecognizeSceneResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeSceneWithChan invokes the imagerecog.RecognizeScene API asynchronously
func (client *Client) RecognizeSceneWithChan(request *RecognizeSceneRequest) (<-chan *RecognizeSceneResponse, <-chan error) {
	responseChan := make(chan *RecognizeSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeScene(request)
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

// RecognizeSceneWithCallback invokes the imagerecog.RecognizeScene API asynchronously
func (client *Client) RecognizeSceneWithCallback(request *RecognizeSceneRequest, callback func(response *RecognizeSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeSceneResponse
		var err error
		defer close(result)
		response, err = client.RecognizeScene(request)
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

// RecognizeSceneRequest is the request struct for api RecognizeScene
type RecognizeSceneRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	ImageType          requests.Integer `position:"Body" name:"ImageType"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	ImageURL           string           `position:"Body" name:"ImageURL"`
}

// RecognizeSceneResponse is the response struct for api RecognizeScene
type RecognizeSceneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeSceneRequest creates a request to invoke RecognizeScene API
func CreateRecognizeSceneRequest() (request *RecognizeSceneRequest) {
	request = &RecognizeSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imagerecog", "2019-09-30", "RecognizeScene", "", "")
	request.Method = requests.POST
	return
}

// CreateRecognizeSceneResponse creates a response to parse from RecognizeScene response
func CreateRecognizeSceneResponse() (response *RecognizeSceneResponse) {
	response = &RecognizeSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
