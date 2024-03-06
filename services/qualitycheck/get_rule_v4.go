package qualitycheck

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

// GetRuleV4 invokes the qualitycheck.GetRuleV4 API synchronously
func (client *Client) GetRuleV4(request *GetRuleV4Request) (response *GetRuleV4Response, err error) {
	response = CreateGetRuleV4Response()
	err = client.DoAction(request, response)
	return
}

// GetRuleV4WithChan invokes the qualitycheck.GetRuleV4 API asynchronously
func (client *Client) GetRuleV4WithChan(request *GetRuleV4Request) (<-chan *GetRuleV4Response, <-chan error) {
	responseChan := make(chan *GetRuleV4Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleV4(request)
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

// GetRuleV4WithCallback invokes the qualitycheck.GetRuleV4 API asynchronously
func (client *Client) GetRuleV4WithCallback(request *GetRuleV4Request, callback func(response *GetRuleV4Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleV4Response
		var err error
		defer close(result)
		response, err = client.GetRuleV4(request)
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

// GetRuleV4Request is the request struct for api GetRuleV4
type GetRuleV4Request struct {
	*requests.RpcRequest
	IsSchemeData requests.Integer `position:"Body" name:"IsSchemeData"`
	RuleId       requests.Integer `position:"Body" name:"RuleId"`
}

// GetRuleV4Response is the response struct for api GetRuleV4
type GetRuleV4Response struct {
	*responses.BaseResponse
}

// CreateGetRuleV4Request creates a request to invoke GetRuleV4 API
func CreateGetRuleV4Request() (request *GetRuleV4Request) {
	request = &GetRuleV4Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRuleV4", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleV4Response creates a response to parse from GetRuleV4 response
func CreateGetRuleV4Response() (response *GetRuleV4Response) {
	response = &GetRuleV4Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
