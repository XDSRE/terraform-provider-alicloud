package dds

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

// CreateRecommendationTask invokes the dds.CreateRecommendationTask API synchronously
func (client *Client) CreateRecommendationTask(request *CreateRecommendationTaskRequest) (response *CreateRecommendationTaskResponse, err error) {
	response = CreateCreateRecommendationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRecommendationTaskWithChan invokes the dds.CreateRecommendationTask API asynchronously
func (client *Client) CreateRecommendationTaskWithChan(request *CreateRecommendationTaskRequest) (<-chan *CreateRecommendationTaskResponse, <-chan error) {
	responseChan := make(chan *CreateRecommendationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRecommendationTask(request)
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

// CreateRecommendationTaskWithCallback invokes the dds.CreateRecommendationTask API asynchronously
func (client *Client) CreateRecommendationTaskWithCallback(request *CreateRecommendationTaskRequest, callback func(response *CreateRecommendationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRecommendationTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateRecommendationTask(request)
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

// CreateRecommendationTaskRequest is the request struct for api CreateRecommendationTask
type CreateRecommendationTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// CreateRecommendationTaskResponse is the response struct for api CreateRecommendationTask
type CreateRecommendationTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateRecommendationTaskRequest creates a request to invoke CreateRecommendationTask API
func CreateCreateRecommendationTaskRequest() (request *CreateRecommendationTaskRequest) {
	request = &CreateRecommendationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CreateRecommendationTask", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRecommendationTaskResponse creates a response to parse from CreateRecommendationTask response
func CreateCreateRecommendationTaskResponse() (response *CreateRecommendationTaskResponse) {
	response = &CreateRecommendationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
