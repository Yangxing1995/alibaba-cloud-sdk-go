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

// SubmitTextTo3DAvatarVideoTask invokes the avatar.SubmitTextTo3DAvatarVideoTask API synchronously
func (client *Client) SubmitTextTo3DAvatarVideoTask(request *SubmitTextTo3DAvatarVideoTaskRequest) (response *SubmitTextTo3DAvatarVideoTaskResponse, err error) {
	response = CreateSubmitTextTo3DAvatarVideoTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitTextTo3DAvatarVideoTaskWithChan invokes the avatar.SubmitTextTo3DAvatarVideoTask API asynchronously
func (client *Client) SubmitTextTo3DAvatarVideoTaskWithChan(request *SubmitTextTo3DAvatarVideoTaskRequest) (<-chan *SubmitTextTo3DAvatarVideoTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitTextTo3DAvatarVideoTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitTextTo3DAvatarVideoTask(request)
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

// SubmitTextTo3DAvatarVideoTaskWithCallback invokes the avatar.SubmitTextTo3DAvatarVideoTask API asynchronously
func (client *Client) SubmitTextTo3DAvatarVideoTaskWithCallback(request *SubmitTextTo3DAvatarVideoTaskRequest, callback func(response *SubmitTextTo3DAvatarVideoTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitTextTo3DAvatarVideoTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitTextTo3DAvatarVideoTask(request)
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

// SubmitTextTo3DAvatarVideoTaskRequest is the request struct for api SubmitTextTo3DAvatarVideoTask
type SubmitTextTo3DAvatarVideoTaskRequest struct {
	*requests.RpcRequest
	App               SubmitTextTo3DAvatarVideoTaskApp        `position:"Query" name:"App"  type:"Struct"`
	AudioInfo         SubmitTextTo3DAvatarVideoTaskAudioInfo  `position:"Query" name:"AudioInfo"  type:"Struct"`
	AvatarInfo        SubmitTextTo3DAvatarVideoTaskAvatarInfo `position:"Query" name:"AvatarInfo"  type:"Struct"`
	Title             string                                  `position:"Query" name:"Title"`
	ExtParams         string                                  `position:"Query" name:"ExtParams"`
	VideoInfo         SubmitTextTo3DAvatarVideoTaskVideoInfo  `position:"Query" name:"VideoInfo"  type:"Struct"`
	CallbackParams    string                                  `position:"Query" name:"CallbackParams"`
	TenantId          requests.Integer                        `position:"Query" name:"TenantId"`
	Callback          requests.Boolean                        `position:"Query" name:"Callback"`
	ExtParamsCLS      string                                  `position:"Query" name:"ExtParams_CLS"`
	Text              string                                  `position:"Query" name:"Text"`
	CallbackParamsCLS string                                  `position:"Query" name:"CallbackParams_CLS"`
}

// SubmitTextTo3DAvatarVideoTaskApp is a repeated param struct in SubmitTextTo3DAvatarVideoTaskRequest
type SubmitTextTo3DAvatarVideoTaskApp struct {
	AppId string `name:"AppId"`
}

// SubmitTextTo3DAvatarVideoTaskAudioInfo is a repeated param struct in SubmitTextTo3DAvatarVideoTaskRequest
type SubmitTextTo3DAvatarVideoTaskAudioInfo struct {
	Voice      string `name:"Voice"`
	Volume     string `name:"Volume"`
	SpeechRate string `name:"SpeechRate"`
	PitchRate  string `name:"PitchRate"`
	SampleRate string `name:"SampleRate"`
}

// SubmitTextTo3DAvatarVideoTaskAvatarInfo is a repeated param struct in SubmitTextTo3DAvatarVideoTaskRequest
type SubmitTextTo3DAvatarVideoTaskAvatarInfo struct {
	Code         string `name:"Code"`
	Locate       string `name:"Locate"`
	Angle        string `name:"Angle"`
	IndustryCode string `name:"IndustryCode"`
}

// SubmitTextTo3DAvatarVideoTaskVideoInfo is a repeated param struct in SubmitTextTo3DAvatarVideoTaskRequest
type SubmitTextTo3DAvatarVideoTaskVideoInfo struct {
	IsAlpha            string                                              `name:"IsAlpha"`
	BackgroundImageUrl string                                              `name:"BackgroundImageUrl"`
	IsSubtitles        string                                              `name:"IsSubtitles"`
	SubtitleEmbedded   string                                              `name:"SubtitleEmbedded"`
	SubtitleStyle      SubmitTextTo3DAvatarVideoTaskVideoInfoSubtitleStyle `name:"SubtitleStyle" type:"Struct"`
	Resolution         string                                              `name:"Resolution"`
	AlphaFormat        string                                              `name:"AlphaFormat"`
}

// SubmitTextTo3DAvatarVideoTaskVideoInfoSubtitleStyle is a repeated param struct in SubmitTextTo3DAvatarVideoTaskRequest
type SubmitTextTo3DAvatarVideoTaskVideoInfoSubtitleStyle struct {
	Color        string `name:"Color"`
	Size         string `name:"Size"`
	Name         string `name:"Name"`
	Y            string `name:"Y"`
	OutlineColor string `name:"OutlineColor"`
}

// SubmitTextTo3DAvatarVideoTaskResponse is the response struct for api SubmitTextTo3DAvatarVideoTask
type SubmitTextTo3DAvatarVideoTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitTextTo3DAvatarVideoTaskRequest creates a request to invoke SubmitTextTo3DAvatarVideoTask API
func CreateSubmitTextTo3DAvatarVideoTaskRequest() (request *SubmitTextTo3DAvatarVideoTaskRequest) {
	request = &SubmitTextTo3DAvatarVideoTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "SubmitTextTo3DAvatarVideoTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitTextTo3DAvatarVideoTaskResponse creates a response to parse from SubmitTextTo3DAvatarVideoTask response
func CreateSubmitTextTo3DAvatarVideoTaskResponse() (response *SubmitTextTo3DAvatarVideoTaskResponse) {
	response = &SubmitTextTo3DAvatarVideoTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
