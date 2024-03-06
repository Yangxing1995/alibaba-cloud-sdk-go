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

// GetLeniPrivateIpAddress invokes the eflo.GetLeniPrivateIpAddress API synchronously
func (client *Client) GetLeniPrivateIpAddress(request *GetLeniPrivateIpAddressRequest) (response *GetLeniPrivateIpAddressResponse, err error) {
	response = CreateGetLeniPrivateIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// GetLeniPrivateIpAddressWithChan invokes the eflo.GetLeniPrivateIpAddress API asynchronously
func (client *Client) GetLeniPrivateIpAddressWithChan(request *GetLeniPrivateIpAddressRequest) (<-chan *GetLeniPrivateIpAddressResponse, <-chan error) {
	responseChan := make(chan *GetLeniPrivateIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLeniPrivateIpAddress(request)
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

// GetLeniPrivateIpAddressWithCallback invokes the eflo.GetLeniPrivateIpAddress API asynchronously
func (client *Client) GetLeniPrivateIpAddressWithCallback(request *GetLeniPrivateIpAddressRequest, callback func(response *GetLeniPrivateIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLeniPrivateIpAddressResponse
		var err error
		defer close(result)
		response, err = client.GetLeniPrivateIpAddress(request)
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

// GetLeniPrivateIpAddressRequest is the request struct for api GetLeniPrivateIpAddress
type GetLeniPrivateIpAddressRequest struct {
	*requests.RpcRequest
	IpName                    string `position:"Body" name:"IpName"`
	ElasticNetworkInterfaceId string `position:"Body" name:"ElasticNetworkInterfaceId"`
}

// GetLeniPrivateIpAddressResponse is the response struct for api GetLeniPrivateIpAddress
type GetLeniPrivateIpAddressResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetLeniPrivateIpAddressRequest creates a request to invoke GetLeniPrivateIpAddress API
func CreateGetLeniPrivateIpAddressRequest() (request *GetLeniPrivateIpAddressRequest) {
	request = &GetLeniPrivateIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetLeniPrivateIpAddress", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLeniPrivateIpAddressResponse creates a response to parse from GetLeniPrivateIpAddress response
func CreateGetLeniPrivateIpAddressResponse() (response *GetLeniPrivateIpAddressResponse) {
	response = &GetLeniPrivateIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
