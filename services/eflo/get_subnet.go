package eflo

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

// GetSubnet invokes the eflo.GetSubnet API synchronously
func (client *Client) GetSubnet(request *GetSubnetRequest) (response *GetSubnetResponse, err error) {
	response = CreateGetSubnetResponse()
	err = client.DoAction(request, response)
	return
}

// GetSubnetWithChan invokes the eflo.GetSubnet API asynchronously
func (client *Client) GetSubnetWithChan(request *GetSubnetRequest) (<-chan *GetSubnetResponse, <-chan error) {
	responseChan := make(chan *GetSubnetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSubnet(request)
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

// GetSubnetWithCallback invokes the eflo.GetSubnet API asynchronously
func (client *Client) GetSubnetWithCallback(request *GetSubnetRequest, callback func(response *GetSubnetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSubnetResponse
		var err error
		defer close(result)
		response, err = client.GetSubnet(request)
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

// GetSubnetRequest is the request struct for api GetSubnet
type GetSubnetRequest struct {
	*requests.RpcRequest
	SubnetId string `position:"Body" name:"SubnetId"`
	VpdId    string `position:"Body" name:"VpdId"`
}

// GetSubnetResponse is the response struct for api GetSubnet
type GetSubnetResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetSubnetRequest creates a request to invoke GetSubnet API
func CreateGetSubnetRequest() (request *GetSubnetRequest) {
	request = &GetSubnetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetSubnet", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSubnetResponse creates a response to parse from GetSubnet response
func CreateGetSubnetResponse() (response *GetSubnetResponse) {
	response = &GetSubnetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
