package eventbridge

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

// UpdateEventStreaming invokes the eventbridge.UpdateEventStreaming API synchronously
func (client *Client) UpdateEventStreaming(request *UpdateEventStreamingRequest) (response *UpdateEventStreamingResponse, err error) {
	response = CreateUpdateEventStreamingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEventStreamingWithChan invokes the eventbridge.UpdateEventStreaming API asynchronously
func (client *Client) UpdateEventStreamingWithChan(request *UpdateEventStreamingRequest) (<-chan *UpdateEventStreamingResponse, <-chan error) {
	responseChan := make(chan *UpdateEventStreamingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEventStreaming(request)
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

// UpdateEventStreamingWithCallback invokes the eventbridge.UpdateEventStreaming API asynchronously
func (client *Client) UpdateEventStreamingWithCallback(request *UpdateEventStreamingRequest, callback func(response *UpdateEventStreamingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEventStreamingResponse
		var err error
		defer close(result)
		response, err = client.UpdateEventStreaming(request)
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

// UpdateEventStreamingRequest is the request struct for api UpdateEventStreaming
type UpdateEventStreamingRequest struct {
	*requests.RpcRequest
	Sink               UpdateEventStreamingSink          `position:"Body" name:"Sink"  type:"Struct"`
	Transforms         *[]UpdateEventStreamingTransforms `position:"Body" name:"Transforms"  type:"Json"`
	Description        string                            `position:"Body" name:"Description"`
	FilterPattern      string                            `position:"Body" name:"FilterPattern"`
	Source             UpdateEventStreamingSource        `position:"Body" name:"Source"  type:"Struct"`
	RunOptions         UpdateEventStreamingRunOptions    `position:"Body" name:"RunOptions"  type:"Struct"`
	EventStreamingName string                            `position:"Body" name:"EventStreamingName"`
	Tag                string                            `position:"Body" name:"Tag"`
}

// UpdateEventStreamingTransforms is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingTransforms struct {
	Arn string `name:"Arn"`
}

// UpdateEventStreamingSink is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSink struct {
	SinkRabbitMQParameters   UpdateEventStreamingSinkSinkRabbitMQParameters   `name:"SinkRabbitMQParameters" type:"Struct"`
	SinkMNSParameters        UpdateEventStreamingSinkSinkMNSParameters        `name:"SinkMNSParameters" type:"Struct"`
	SinkKafkaParameters      UpdateEventStreamingSinkSinkKafkaParameters      `name:"SinkKafkaParameters" type:"Struct"`
	SinkFnfParameters        UpdateEventStreamingSinkSinkFnfParameters        `name:"SinkFnfParameters" type:"Struct"`
	SinkFcParameters         UpdateEventStreamingSinkSinkFcParameters         `name:"SinkFcParameters" type:"Struct"`
	SinkPrometheusParameters UpdateEventStreamingSinkSinkPrometheusParameters `name:"SinkPrometheusParameters" type:"Struct"`
	SinkSLSParameters        UpdateEventStreamingSinkSinkSLSParameters        `name:"SinkSLSParameters" type:"Struct"`
	SinkRocketMQParameters   UpdateEventStreamingSinkSinkRocketMQParameters   `name:"SinkRocketMQParameters" type:"Struct"`
}

// UpdateEventStreamingSource is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSource struct {
	SourceMQTTParameters       UpdateEventStreamingSourceSourceMQTTParameters       `name:"SourceMQTTParameters" type:"Struct"`
	SourceRocketMQParameters   UpdateEventStreamingSourceSourceRocketMQParameters   `name:"SourceRocketMQParameters" type:"Struct"`
	SourceSLSParameters        UpdateEventStreamingSourceSourceSLSParameters        `name:"SourceSLSParameters" type:"Struct"`
	SourcePrometheusParameters UpdateEventStreamingSourceSourcePrometheusParameters `name:"SourcePrometheusParameters" type:"Struct"`
	SourceDTSParameters        UpdateEventStreamingSourceSourceDTSParameters        `name:"SourceDTSParameters" type:"Struct"`
	SourceKafkaParameters      UpdateEventStreamingSourceSourceKafkaParameters      `name:"SourceKafkaParameters" type:"Struct"`
	SourceMNSParameters        UpdateEventStreamingSourceSourceMNSParameters        `name:"SourceMNSParameters" type:"Struct"`
	SourceRabbitMQParameters   UpdateEventStreamingSourceSourceRabbitMQParameters   `name:"SourceRabbitMQParameters" type:"Struct"`
}

// UpdateEventStreamingRunOptions is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingRunOptions struct {
	BatchWindow     UpdateEventStreamingRunOptionsBatchWindow     `name:"BatchWindow" type:"Struct"`
	RetryStrategy   UpdateEventStreamingRunOptionsRetryStrategy   `name:"RetryStrategy" type:"Struct"`
	DeadLetterQueue UpdateEventStreamingRunOptionsDeadLetterQueue `name:"DeadLetterQueue" type:"Struct"`
	MaximumTasks    string                                        `name:"MaximumTasks"`
	ErrorsTolerance string                                        `name:"ErrorsTolerance"`
}

// UpdateEventStreamingSinkSinkRabbitMQParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParameters struct {
	RoutingKey      UpdateEventStreamingSinkSinkRabbitMQParametersRoutingKey      `name:"RoutingKey" type:"Struct"`
	QueueName       UpdateEventStreamingSinkSinkRabbitMQParametersQueueName       `name:"QueueName" type:"Struct"`
	VirtualHostName UpdateEventStreamingSinkSinkRabbitMQParametersVirtualHostName `name:"VirtualHostName" type:"Struct"`
	InstanceId      UpdateEventStreamingSinkSinkRabbitMQParametersInstanceId      `name:"InstanceId" type:"Struct"`
	TargetType      UpdateEventStreamingSinkSinkRabbitMQParametersTargetType      `name:"TargetType" type:"Struct"`
	MessageId       UpdateEventStreamingSinkSinkRabbitMQParametersMessageId       `name:"MessageId" type:"Struct"`
	Exchange        UpdateEventStreamingSinkSinkRabbitMQParametersExchange        `name:"Exchange" type:"Struct"`
	Body            UpdateEventStreamingSinkSinkRabbitMQParametersBody            `name:"Body" type:"Struct"`
	Properties      UpdateEventStreamingSinkSinkRabbitMQParametersProperties      `name:"Properties" type:"Struct"`
}

// UpdateEventStreamingSinkSinkMNSParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkMNSParameters struct {
	QueueName      UpdateEventStreamingSinkSinkMNSParametersQueueName      `name:"QueueName" type:"Struct"`
	IsBase64Encode UpdateEventStreamingSinkSinkMNSParametersIsBase64Encode `name:"IsBase64Encode" type:"Struct"`
	Body           UpdateEventStreamingSinkSinkMNSParametersBody           `name:"Body" type:"Struct"`
}

// UpdateEventStreamingSinkSinkKafkaParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParameters struct {
	InstanceId UpdateEventStreamingSinkSinkKafkaParametersInstanceId `name:"InstanceId" type:"Struct"`
	Acks       UpdateEventStreamingSinkSinkKafkaParametersAcks       `name:"Acks" type:"Struct"`
	Topic      UpdateEventStreamingSinkSinkKafkaParametersTopic      `name:"Topic" type:"Struct"`
	SaslUser   UpdateEventStreamingSinkSinkKafkaParametersSaslUser   `name:"SaslUser" type:"Struct"`
	Value      UpdateEventStreamingSinkSinkKafkaParametersValue      `name:"Value" type:"Struct"`
	Key        UpdateEventStreamingSinkSinkKafkaParametersKey        `name:"Key" type:"Struct"`
}

// UpdateEventStreamingSinkSinkFnfParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFnfParameters struct {
	Input         UpdateEventStreamingSinkSinkFnfParametersInput         `name:"Input" type:"Struct"`
	ExecutionName UpdateEventStreamingSinkSinkFnfParametersExecutionName `name:"ExecutionName" type:"Struct"`
	RoleName      UpdateEventStreamingSinkSinkFnfParametersRoleName      `name:"RoleName" type:"Struct"`
	FlowName      UpdateEventStreamingSinkSinkFnfParametersFlowName      `name:"FlowName" type:"Struct"`
}

// UpdateEventStreamingSinkSinkFcParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParameters struct {
	InvocationType UpdateEventStreamingSinkSinkFcParametersInvocationType `name:"InvocationType" type:"Struct"`
	FunctionName   UpdateEventStreamingSinkSinkFcParametersFunctionName   `name:"FunctionName" type:"Struct"`
	Qualifier      UpdateEventStreamingSinkSinkFcParametersQualifier      `name:"Qualifier" type:"Struct"`
	ServiceName    UpdateEventStreamingSinkSinkFcParametersServiceName    `name:"ServiceName" type:"Struct"`
	Body           UpdateEventStreamingSinkSinkFcParametersBody           `name:"Body" type:"Struct"`
	Concurrency    UpdateEventStreamingSinkSinkFcParametersConcurrency    `name:"Concurrency" type:"Struct"`
}

// UpdateEventStreamingSinkSinkPrometheusParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParameters struct {
	VSwitchId         UpdateEventStreamingSinkSinkPrometheusParametersVSwitchId         `name:"VSwitchId" type:"Struct"`
	Password          UpdateEventStreamingSinkSinkPrometheusParametersPassword          `name:"Password" type:"Struct"`
	Data              UpdateEventStreamingSinkSinkPrometheusParametersData              `name:"Data" type:"Struct"`
	VpcId             UpdateEventStreamingSinkSinkPrometheusParametersVpcId             `name:"VpcId" type:"Struct"`
	SecurityGroupId   UpdateEventStreamingSinkSinkPrometheusParametersSecurityGroupId   `name:"SecurityGroupId" type:"Struct"`
	AuthorizationType UpdateEventStreamingSinkSinkPrometheusParametersAuthorizationType `name:"AuthorizationType" type:"Struct"`
	NetworkType       UpdateEventStreamingSinkSinkPrometheusParametersNetworkType       `name:"NetworkType" type:"Struct"`
	URL               UpdateEventStreamingSinkSinkPrometheusParametersURL               `name:"URL" type:"Struct"`
	Username          UpdateEventStreamingSinkSinkPrometheusParametersUsername          `name:"Username" type:"Struct"`
}

// UpdateEventStreamingSinkSinkSLSParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParameters struct {
	RoleName UpdateEventStreamingSinkSinkSLSParametersRoleName `name:"RoleName" type:"Struct"`
	Project  UpdateEventStreamingSinkSinkSLSParametersProject  `name:"Project" type:"Struct"`
	Topic    UpdateEventStreamingSinkSinkSLSParametersTopic    `name:"Topic" type:"Struct"`
	Body     UpdateEventStreamingSinkSinkSLSParametersBody     `name:"Body" type:"Struct"`
	LogStore UpdateEventStreamingSinkSinkSLSParametersLogStore `name:"LogStore" type:"Struct"`
}

// UpdateEventStreamingSinkSinkRocketMQParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParameters struct {
	InstanceId UpdateEventStreamingSinkSinkRocketMQParametersInstanceId `name:"InstanceId" type:"Struct"`
	Keys       UpdateEventStreamingSinkSinkRocketMQParametersKeys       `name:"Keys" type:"Struct"`
	Topic      UpdateEventStreamingSinkSinkRocketMQParametersTopic      `name:"Topic" type:"Struct"`
	Body       UpdateEventStreamingSinkSinkRocketMQParametersBody       `name:"Body" type:"Struct"`
	Properties UpdateEventStreamingSinkSinkRocketMQParametersProperties `name:"Properties" type:"Struct"`
	Tags       UpdateEventStreamingSinkSinkRocketMQParametersTags       `name:"Tags" type:"Struct"`
}

// UpdateEventStreamingSourceSourceMQTTParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceMQTTParameters struct {
	InstanceId string `name:"InstanceId"`
	RegionId   string `name:"RegionId"`
	Topic      string `name:"Topic"`
}

// UpdateEventStreamingSourceSourceRocketMQParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceRocketMQParameters struct {
	InstanceSecurityGroupId string `name:"InstanceSecurityGroupId"`
	Offset                  string `name:"Offset"`
	GroupID                 string `name:"GroupID"`
	InstanceUsername        string `name:"InstanceUsername"`
	AuthType                string `name:"AuthType"`
	InstancePassword        string `name:"InstancePassword"`
	InstanceVSwitchIds      string `name:"InstanceVSwitchIds"`
	InstanceNetwork         string `name:"InstanceNetwork"`
	InstanceId              string `name:"InstanceId"`
	InstanceEndpoint        string `name:"InstanceEndpoint"`
	InstanceVpcId           string `name:"InstanceVpcId"`
	RegionId                string `name:"RegionId"`
	Topic                   string `name:"Topic"`
	InstanceType            string `name:"InstanceType"`
	Tag                     string `name:"Tag"`
	Timestamp               string `name:"Timestamp"`
}

// UpdateEventStreamingSourceSourceSLSParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceSLSParameters struct {
	RoleName string `name:"RoleName"`
}

// UpdateEventStreamingSourceSourcePrometheusParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourcePrometheusParameters struct {
	DataType  string `name:"DataType"`
	ClusterId string `name:"ClusterId"`
	Labels    string `name:"Labels"`
}

// UpdateEventStreamingSourceSourceDTSParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceDTSParameters struct {
	BrokerUrl      string `name:"BrokerUrl"`
	Password       string `name:"Password"`
	InitCheckPoint string `name:"InitCheckPoint"`
	Topic          string `name:"Topic"`
	TaskId         string `name:"TaskId"`
	Sid            string `name:"Sid"`
	Username       string `name:"Username"`
}

// UpdateEventStreamingSourceSourceKafkaParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceKafkaParameters struct {
	InstanceId      string `name:"InstanceId"`
	ConsumerGroup   string `name:"ConsumerGroup"`
	RegionId        string `name:"RegionId"`
	VSwitchIds      string `name:"VSwitchIds"`
	VpcId           string `name:"VpcId"`
	ValueDataType   string `name:"ValueDataType"`
	SecurityGroupId string `name:"SecurityGroupId"`
	Topic           string `name:"Topic"`
	OffsetReset     string `name:"OffsetReset"`
	Network         string `name:"Network"`
}

// UpdateEventStreamingSourceSourceMNSParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceMNSParameters struct {
	QueueName      string `name:"QueueName"`
	RegionId       string `name:"RegionId"`
	IsBase64Decode string `name:"IsBase64Decode"`
}

// UpdateEventStreamingSourceSourceRabbitMQParameters is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSourceSourceRabbitMQParameters struct {
	QueueName       string `name:"QueueName"`
	VirtualHostName string `name:"VirtualHostName"`
	InstanceId      string `name:"InstanceId"`
	RegionId        string `name:"RegionId"`
}

// UpdateEventStreamingRunOptionsBatchWindow is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingRunOptionsBatchWindow struct {
	CountBasedWindow string `name:"CountBasedWindow"`
	TimeBasedWindow  string `name:"TimeBasedWindow"`
}

// UpdateEventStreamingRunOptionsRetryStrategy is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingRunOptionsRetryStrategy struct {
	PushRetryStrategy        string `name:"PushRetryStrategy"`
	MaximumRetryAttempts     string `name:"MaximumRetryAttempts"`
	MaximumEventAgeInSeconds string `name:"MaximumEventAgeInSeconds"`
}

// UpdateEventStreamingRunOptionsDeadLetterQueue is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingRunOptionsDeadLetterQueue struct {
	Arn string `name:"Arn"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersRoutingKey is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersRoutingKey struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersQueueName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersQueueName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersVirtualHostName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersVirtualHostName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersInstanceId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersInstanceId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersTargetType is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersTargetType struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersMessageId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersMessageId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersExchange is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersExchange struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersBody is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersBody struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRabbitMQParametersProperties is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRabbitMQParametersProperties struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkMNSParametersQueueName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkMNSParametersQueueName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkMNSParametersIsBase64Encode is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkMNSParametersIsBase64Encode struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkMNSParametersBody is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkMNSParametersBody struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersInstanceId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersInstanceId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersAcks is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersAcks struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersTopic is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersTopic struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersSaslUser is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersSaslUser struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersValue is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersValue struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkKafkaParametersKey is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkKafkaParametersKey struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFnfParametersInput is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFnfParametersInput struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFnfParametersExecutionName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFnfParametersExecutionName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFnfParametersRoleName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFnfParametersRoleName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFnfParametersFlowName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFnfParametersFlowName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersInvocationType is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersInvocationType struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersFunctionName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersFunctionName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersQualifier is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersQualifier struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersServiceName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersServiceName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersBody is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersBody struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkFcParametersConcurrency is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkFcParametersConcurrency struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersVSwitchId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersVSwitchId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersPassword is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersPassword struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersData is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersData struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersVpcId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersVpcId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersSecurityGroupId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersSecurityGroupId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersAuthorizationType is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersAuthorizationType struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersNetworkType is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersNetworkType struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersURL is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersURL struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkPrometheusParametersUsername is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkPrometheusParametersUsername struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkSLSParametersRoleName is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParametersRoleName struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkSLSParametersProject is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParametersProject struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkSLSParametersTopic is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParametersTopic struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkSLSParametersBody is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParametersBody struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkSLSParametersLogStore is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkSLSParametersLogStore struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersInstanceId is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersInstanceId struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersKeys is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersKeys struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersTopic is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersTopic struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersBody is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersBody struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersProperties is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersProperties struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingSinkSinkRocketMQParametersTags is a repeated param struct in UpdateEventStreamingRequest
type UpdateEventStreamingSinkSinkRocketMQParametersTags struct {
	Template string `name:"Template"`
	Form     string `name:"Form"`
	Value    string `name:"Value"`
}

// UpdateEventStreamingResponse is the response struct for api UpdateEventStreaming
type UpdateEventStreamingResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateEventStreamingRequest creates a request to invoke UpdateEventStreaming API
func CreateUpdateEventStreamingRequest() (request *UpdateEventStreamingRequest) {
	request = &UpdateEventStreamingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "UpdateEventStreaming", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateEventStreamingResponse creates a response to parse from UpdateEventStreaming response
func CreateUpdateEventStreamingResponse() (response *UpdateEventStreamingResponse) {
	response = &UpdateEventStreamingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
