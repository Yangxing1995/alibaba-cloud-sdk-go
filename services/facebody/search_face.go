package facebody

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

// SearchFace invokes the facebody.SearchFace API synchronously
func (client *Client) SearchFace(request *SearchFaceRequest) (response *SearchFaceResponse, err error) {
	response = CreateSearchFaceResponse()
	err = client.DoAction(request, response)
	return
}

// SearchFaceWithChan invokes the facebody.SearchFace API asynchronously
func (client *Client) SearchFaceWithChan(request *SearchFaceRequest) (<-chan *SearchFaceResponse, <-chan error) {
	responseChan := make(chan *SearchFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchFace(request)
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

// SearchFaceWithCallback invokes the facebody.SearchFace API asynchronously
func (client *Client) SearchFaceWithCallback(request *SearchFaceRequest, callback func(response *SearchFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchFaceResponse
		var err error
		defer close(result)
		response, err = client.SearchFace(request)
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

// SearchFaceRequest is the request struct for api SearchFace
type SearchFaceRequest struct {
	*requests.RpcRequest
	MaxFaceNum            requests.Integer `position:"Body" name:"MaxFaceNum"`
	FormatResultToJson    requests.Boolean `position:"Query" name:"FormatResultToJson"`
	QualityScoreThreshold requests.Float   `position:"Body" name:"QualityScoreThreshold"`
	Limit                 requests.Integer `position:"Body" name:"Limit"`
	OssFile               string           `position:"Query" name:"OssFile"`
	RequestProxyBy        string           `position:"Query" name:"RequestProxyBy"`
	DbNames               string           `position:"Body" name:"DbNames"`
	DbName                string           `position:"Body" name:"DbName"`
	ImageUrl              string           `position:"Body" name:"ImageUrl"`
}

// SearchFaceResponse is the response struct for api SearchFace
type SearchFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSearchFaceRequest creates a request to invoke SearchFace API
func CreateSearchFaceRequest() (request *SearchFaceRequest) {
	request = &SearchFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "SearchFace", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchFaceResponse creates a response to parse from SearchFace response
func CreateSearchFaceResponse() (response *SearchFaceResponse) {
	response = &SearchFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
