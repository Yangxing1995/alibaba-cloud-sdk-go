package avatar

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

// SubmitTextTo2DAvatarVideoTask invokes the avatar.SubmitTextTo2DAvatarVideoTask API synchronously
func (client *Client) SubmitTextTo2DAvatarVideoTask(request *SubmitTextTo2DAvatarVideoTaskRequest) (response *SubmitTextTo2DAvatarVideoTaskResponse, err error) {
	response = CreateSubmitTextTo2DAvatarVideoTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitTextTo2DAvatarVideoTaskWithChan invokes the avatar.SubmitTextTo2DAvatarVideoTask API asynchronously
func (client *Client) SubmitTextTo2DAvatarVideoTaskWithChan(request *SubmitTextTo2DAvatarVideoTaskRequest) (<-chan *SubmitTextTo2DAvatarVideoTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitTextTo2DAvatarVideoTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitTextTo2DAvatarVideoTask(request)
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

// SubmitTextTo2DAvatarVideoTaskWithCallback invokes the avatar.SubmitTextTo2DAvatarVideoTask API asynchronously
func (client *Client) SubmitTextTo2DAvatarVideoTaskWithCallback(request *SubmitTextTo2DAvatarVideoTaskRequest, callback func(response *SubmitTextTo2DAvatarVideoTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitTextTo2DAvatarVideoTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitTextTo2DAvatarVideoTask(request)
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

// SubmitTextTo2DAvatarVideoTaskRequest is the request struct for api SubmitTextTo2DAvatarVideoTask
type SubmitTextTo2DAvatarVideoTaskRequest struct {
	*requests.RpcRequest
	App        SubmitTextTo2DAvatarVideoTaskApp        `position:"Query" name:"App"  type:"Struct"`
	VideoInfo  SubmitTextTo2DAvatarVideoTaskVideoInfo  `position:"Query" name:"VideoInfo"  type:"Struct"`
	AudioInfo  SubmitTextTo2DAvatarVideoTaskAudioInfo  `position:"Query" name:"AudioInfo"  type:"Struct"`
	AvatarInfo SubmitTextTo2DAvatarVideoTaskAvatarInfo `position:"Query" name:"AvatarInfo"  type:"Struct"`
	TenantId   requests.Integer                        `position:"Query" name:"TenantId"`
	Text       string                                  `position:"Query" name:"Text"`
	Title      string                                  `position:"Query" name:"Title"`
}

// SubmitTextTo2DAvatarVideoTaskApp is a repeated param struct in SubmitTextTo2DAvatarVideoTaskRequest
type SubmitTextTo2DAvatarVideoTaskApp struct {
	AppId string `name:"AppId"`
}

// SubmitTextTo2DAvatarVideoTaskVideoInfo is a repeated param struct in SubmitTextTo2DAvatarVideoTaskRequest
type SubmitTextTo2DAvatarVideoTaskVideoInfo struct {
	IsAlpha            string `name:"IsAlpha"`
	BackgroundImageUrl string `name:"BackgroundImageUrl"`
	IsSubtitles        string `name:"IsSubtitles"`
	Resolution         string `name:"Resolution"`
	AlphaFormat        string `name:"AlphaFormat"`
}

// SubmitTextTo2DAvatarVideoTaskAudioInfo is a repeated param struct in SubmitTextTo2DAvatarVideoTaskRequest
type SubmitTextTo2DAvatarVideoTaskAudioInfo struct {
	Voice      string `name:"Voice"`
	Volume     string `name:"Volume"`
	SpeechRate string `name:"SpeechRate"`
	PitchRate  string `name:"PitchRate"`
}

// SubmitTextTo2DAvatarVideoTaskAvatarInfo is a repeated param struct in SubmitTextTo2DAvatarVideoTaskRequest
type SubmitTextTo2DAvatarVideoTaskAvatarInfo struct {
	Code string `name:"Code"`
}

// SubmitTextTo2DAvatarVideoTaskResponse is the response struct for api SubmitTextTo2DAvatarVideoTask
type SubmitTextTo2DAvatarVideoTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitTextTo2DAvatarVideoTaskRequest creates a request to invoke SubmitTextTo2DAvatarVideoTask API
func CreateSubmitTextTo2DAvatarVideoTaskRequest() (request *SubmitTextTo2DAvatarVideoTaskRequest) {
	request = &SubmitTextTo2DAvatarVideoTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "SubmitTextTo2DAvatarVideoTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitTextTo2DAvatarVideoTaskResponse creates a response to parse from SubmitTextTo2DAvatarVideoTask response
func CreateSubmitTextTo2DAvatarVideoTaskResponse() (response *SubmitTextTo2DAvatarVideoTaskResponse) {
	response = &SubmitTextTo2DAvatarVideoTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
