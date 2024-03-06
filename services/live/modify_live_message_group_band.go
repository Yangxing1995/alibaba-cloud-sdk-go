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

// ModifyLiveMessageGroupBand invokes the live.ModifyLiveMessageGroupBand API synchronously
func (client *Client) ModifyLiveMessageGroupBand(request *ModifyLiveMessageGroupBandRequest) (response *ModifyLiveMessageGroupBandResponse, err error) {
	response = CreateModifyLiveMessageGroupBandResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLiveMessageGroupBandWithChan invokes the live.ModifyLiveMessageGroupBand API asynchronously
func (client *Client) ModifyLiveMessageGroupBandWithChan(request *ModifyLiveMessageGroupBandRequest) (<-chan *ModifyLiveMessageGroupBandResponse, <-chan error) {
	responseChan := make(chan *ModifyLiveMessageGroupBandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLiveMessageGroupBand(request)
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

// ModifyLiveMessageGroupBandWithCallback invokes the live.ModifyLiveMessageGroupBand API asynchronously
func (client *Client) ModifyLiveMessageGroupBandWithCallback(request *ModifyLiveMessageGroupBandRequest, callback func(response *ModifyLiveMessageGroupBandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLiveMessageGroupBandResponse
		var err error
		defer close(result)
		response, err = client.ModifyLiveMessageGroupBand(request)
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

// ModifyLiveMessageGroupBandRequest is the request struct for api ModifyLiveMessageGroupBand
type ModifyLiveMessageGroupBandRequest struct {
	*requests.RpcRequest
	BannedAll    requests.Boolean `position:"Query" name:"BannedAll"`
	GroupId      string           `position:"Query" name:"GroupId"`
	ExceptUsers  *[]string        `position:"Query" name:"ExceptUsers"  type:"Repeated"`
	DataCenter   string           `position:"Query" name:"DataCenter"`
	AppId        string           `position:"Query" name:"AppId"`
	BannnedUsers *[]string        `position:"Query" name:"BannnedUsers"  type:"Repeated"`
}

// ModifyLiveMessageGroupBandResponse is the response struct for api ModifyLiveMessageGroupBand
type ModifyLiveMessageGroupBandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLiveMessageGroupBandRequest creates a request to invoke ModifyLiveMessageGroupBand API
func CreateModifyLiveMessageGroupBandRequest() (request *ModifyLiveMessageGroupBandRequest) {
	request = &ModifyLiveMessageGroupBandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyLiveMessageGroupBand", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLiveMessageGroupBandResponse creates a response to parse from ModifyLiveMessageGroupBand response
func CreateModifyLiveMessageGroupBandResponse() (response *ModifyLiveMessageGroupBandResponse) {
	response = &ModifyLiveMessageGroupBandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
