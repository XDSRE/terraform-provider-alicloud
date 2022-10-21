package cs

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

// ScaleClusterNodePool invokes the cs.ScaleClusterNodePool API synchronously
func (client *Client) ScaleClusterNodePool(request *ScaleClusterNodePoolRequest) (response *ScaleClusterNodePoolResponse, err error) {
	response = CreateScaleClusterNodePoolResponse()
	err = client.DoAction(request, response)
	return
}

// ScaleClusterNodePoolWithChan invokes the cs.ScaleClusterNodePool API asynchronously
func (client *Client) ScaleClusterNodePoolWithChan(request *ScaleClusterNodePoolRequest) (<-chan *ScaleClusterNodePoolResponse, <-chan error) {
	responseChan := make(chan *ScaleClusterNodePoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleClusterNodePool(request)
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

// ScaleClusterNodePoolWithCallback invokes the cs.ScaleClusterNodePool API asynchronously
func (client *Client) ScaleClusterNodePoolWithCallback(request *ScaleClusterNodePoolRequest, callback func(response *ScaleClusterNodePoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleClusterNodePoolResponse
		var err error
		defer close(result)
		response, err = client.ScaleClusterNodePool(request)
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

// ScaleClusterNodePoolRequest is the request struct for api ScaleClusterNodePool
type ScaleClusterNodePoolRequest struct {
	*requests.RoaRequest
	ClusterId  string `position:"Path" name:"ClusterId"`
	NodepoolId string `position:"Path" name:"NodepoolId"`
}

// ScaleClusterNodePoolResponse is the response struct for api ScaleClusterNodePool
type ScaleClusterNodePoolResponse struct {
	*responses.BaseResponse
	TaskId string `json:"task_id" xml:"task_id"`
}

// CreateScaleClusterNodePoolRequest creates a request to invoke ScaleClusterNodePool API
func CreateScaleClusterNodePoolRequest() (request *ScaleClusterNodePoolRequest) {
	request = &ScaleClusterNodePoolRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ScaleClusterNodePool", "/clusters/[ClusterId]/nodepools/[NodepoolId]", "", "")
	request.Method = requests.POST
	return
}

// CreateScaleClusterNodePoolResponse creates a response to parse from ScaleClusterNodePool response
func CreateScaleClusterNodePoolResponse() (response *ScaleClusterNodePoolResponse) {
	response = &ScaleClusterNodePoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
