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

// DeleteClusterNodepool invokes the cs.DeleteClusterNodepool API synchronously
func (client *Client) DeleteClusterNodepool(request *DeleteClusterNodepoolRequest) (response *DeleteClusterNodepoolResponse, err error) {
	response = CreateDeleteClusterNodepoolResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteClusterNodepoolWithChan invokes the cs.DeleteClusterNodepool API asynchronously
func (client *Client) DeleteClusterNodepoolWithChan(request *DeleteClusterNodepoolRequest) (<-chan *DeleteClusterNodepoolResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterNodepoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteClusterNodepool(request)
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

// DeleteClusterNodepoolWithCallback invokes the cs.DeleteClusterNodepool API asynchronously
func (client *Client) DeleteClusterNodepoolWithCallback(request *DeleteClusterNodepoolRequest, callback func(response *DeleteClusterNodepoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterNodepoolResponse
		var err error
		defer close(result)
		response, err = client.DeleteClusterNodepool(request)
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

// DeleteClusterNodepoolRequest is the request struct for api DeleteClusterNodepool
type DeleteClusterNodepoolRequest struct {
	*requests.RoaRequest
	ClusterId  string `position:"Path" name:"ClusterId"`
	NodepoolId string `position:"Path" name:"NodepoolId"`
}

// DeleteClusterNodepoolResponse is the response struct for api DeleteClusterNodepool
type DeleteClusterNodepoolResponse struct {
	*responses.BaseResponse
}

// CreateDeleteClusterNodepoolRequest creates a request to invoke DeleteClusterNodepool API
func CreateDeleteClusterNodepoolRequest() (request *DeleteClusterNodepoolRequest) {
	request = &DeleteClusterNodepoolRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DeleteClusterNodepool", "/clusters/[ClusterId]/nodepools/[NodepoolId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteClusterNodepoolResponse creates a response to parse from DeleteClusterNodepool response
func CreateDeleteClusterNodepoolResponse() (response *DeleteClusterNodepoolResponse) {
	response = &DeleteClusterNodepoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
