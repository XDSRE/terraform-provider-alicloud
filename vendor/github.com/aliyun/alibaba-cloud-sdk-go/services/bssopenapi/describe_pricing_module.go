package bssopenapi

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

// DescribePricingModule invokes the bssopenapi.DescribePricingModule API synchronously
func (client *Client) DescribePricingModule(request *DescribePricingModuleRequest) (response *DescribePricingModuleResponse, err error) {
	response = CreateDescribePricingModuleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePricingModuleWithChan invokes the bssopenapi.DescribePricingModule API asynchronously
func (client *Client) DescribePricingModuleWithChan(request *DescribePricingModuleRequest) (<-chan *DescribePricingModuleResponse, <-chan error) {
	responseChan := make(chan *DescribePricingModuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePricingModule(request)
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

// DescribePricingModuleWithCallback invokes the bssopenapi.DescribePricingModule API asynchronously
func (client *Client) DescribePricingModuleWithCallback(request *DescribePricingModuleRequest, callback func(response *DescribePricingModuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePricingModuleResponse
		var err error
		defer close(result)
		response, err = client.DescribePricingModule(request)
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

// DescribePricingModuleRequest is the request struct for api DescribePricingModule
type DescribePricingModuleRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	ProductType      string           `position:"Query" name:"ProductType"`
}

// DescribePricingModuleResponse is the response struct for api DescribePricingModule
type DescribePricingModuleResponse struct {
	*responses.BaseResponse
	Code      string                      `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   bool                        `json:"Success" xml:"Success"`
	Data      DataInDescribePricingModule `json:"Data" xml:"Data"`
}

// CreateDescribePricingModuleRequest creates a request to invoke DescribePricingModule API
func CreateDescribePricingModuleRequest() (request *DescribePricingModuleRequest) {
	request = &DescribePricingModuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "DescribePricingModule", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePricingModuleResponse creates a response to parse from DescribePricingModule response
func CreateDescribePricingModuleResponse() (response *DescribePricingModuleResponse) {
	response = &DescribePricingModuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}