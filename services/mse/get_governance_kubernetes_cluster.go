package mse

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

// GetGovernanceKubernetesCluster invokes the mse.GetGovernanceKubernetesCluster API synchronously
func (client *Client) GetGovernanceKubernetesCluster(request *GetGovernanceKubernetesClusterRequest) (response *GetGovernanceKubernetesClusterResponse, err error) {
	response = CreateGetGovernanceKubernetesClusterResponse()
	err = client.DoAction(request, response)
	return
}

// GetGovernanceKubernetesClusterWithChan invokes the mse.GetGovernanceKubernetesCluster API asynchronously
func (client *Client) GetGovernanceKubernetesClusterWithChan(request *GetGovernanceKubernetesClusterRequest) (<-chan *GetGovernanceKubernetesClusterResponse, <-chan error) {
	responseChan := make(chan *GetGovernanceKubernetesClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGovernanceKubernetesCluster(request)
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

// GetGovernanceKubernetesClusterWithCallback invokes the mse.GetGovernanceKubernetesCluster API asynchronously
func (client *Client) GetGovernanceKubernetesClusterWithCallback(request *GetGovernanceKubernetesClusterRequest, callback func(response *GetGovernanceKubernetesClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGovernanceKubernetesClusterResponse
		var err error
		defer close(result)
		response, err = client.GetGovernanceKubernetesCluster(request)
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

// GetGovernanceKubernetesClusterRequest is the request struct for api GetGovernanceKubernetesCluster
type GetGovernanceKubernetesClusterRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// GetGovernanceKubernetesClusterResponse is the response struct for api GetGovernanceKubernetesCluster
type GetGovernanceKubernetesClusterResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                               `json:"Message" xml:"Message"`
	RequestId      string                               `json:"RequestId" xml:"RequestId"`
	Code           int                                  `json:"Code" xml:"Code"`
	Success        bool                                 `json:"Success" xml:"Success"`
	Data           DataInGetGovernanceKubernetesCluster `json:"Data" xml:"Data"`
}

// CreateGetGovernanceKubernetesClusterRequest creates a request to invoke GetGovernanceKubernetesCluster API
func CreateGetGovernanceKubernetesClusterRequest() (request *GetGovernanceKubernetesClusterRequest) {
	request = &GetGovernanceKubernetesClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "GetGovernanceKubernetesCluster", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetGovernanceKubernetesClusterResponse creates a response to parse from GetGovernanceKubernetesCluster response
func CreateGetGovernanceKubernetesClusterResponse() (response *GetGovernanceKubernetesClusterResponse) {
	response = &GetGovernanceKubernetesClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
