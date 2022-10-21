package adb

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

// ModifyElasticPlan invokes the adb.ModifyElasticPlan API synchronously
func (client *Client) ModifyElasticPlan(request *ModifyElasticPlanRequest) (response *ModifyElasticPlanResponse, err error) {
	response = CreateModifyElasticPlanResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyElasticPlanWithChan invokes the adb.ModifyElasticPlan API asynchronously
func (client *Client) ModifyElasticPlanWithChan(request *ModifyElasticPlanRequest) (<-chan *ModifyElasticPlanResponse, <-chan error) {
	responseChan := make(chan *ModifyElasticPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyElasticPlan(request)
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

// ModifyElasticPlanWithCallback invokes the adb.ModifyElasticPlan API asynchronously
func (client *Client) ModifyElasticPlanWithCallback(request *ModifyElasticPlanRequest, callback func(response *ModifyElasticPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyElasticPlanResponse
		var err error
		defer close(result)
		response, err = client.ModifyElasticPlan(request)
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

// ModifyElasticPlanRequest is the request struct for api ModifyElasticPlan
type ModifyElasticPlanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ElasticPlanTimeStart    string           `position:"Query" name:"ElasticPlanTimeStart"`
	ElasticPlanEndDay       string           `position:"Query" name:"ElasticPlanEndDay"`
	ElasticPlanWeeklyRepeat string           `position:"Query" name:"ElasticPlanWeeklyRepeat"`
	ElasticPlanEnable       requests.Boolean `position:"Query" name:"ElasticPlanEnable"`
	ElasticPlanTimeEnd      string           `position:"Query" name:"ElasticPlanTimeEnd"`
	ElasticPlanStartDay     string           `position:"Query" name:"ElasticPlanStartDay"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string           `position:"Query" name:"DBClusterId"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	ElasticPlanName         string           `position:"Query" name:"ElasticPlanName"`
	ResourcePoolName        string           `position:"Query" name:"ResourcePoolName"`
	ElasticPlanNodeNum      requests.Integer `position:"Query" name:"ElasticPlanNodeNum"`
}

// ModifyElasticPlanResponse is the response struct for api ModifyElasticPlan
type ModifyElasticPlanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyElasticPlanRequest creates a request to invoke ModifyElasticPlan API
func CreateModifyElasticPlanRequest() (request *ModifyElasticPlanRequest) {
	request = &ModifyElasticPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyElasticPlan", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyElasticPlanResponse creates a response to parse from ModifyElasticPlan response
func CreateModifyElasticPlanResponse() (response *ModifyElasticPlanResponse) {
	response = &ModifyElasticPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
