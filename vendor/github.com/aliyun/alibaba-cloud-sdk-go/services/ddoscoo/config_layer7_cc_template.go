package ddoscoo

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

// ConfigLayer7CCTemplate invokes the ddoscoo.ConfigLayer7CCTemplate API synchronously
func (client *Client) ConfigLayer7CCTemplate(request *ConfigLayer7CCTemplateRequest) (response *ConfigLayer7CCTemplateResponse, err error) {
	response = CreateConfigLayer7CCTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigLayer7CCTemplateWithChan invokes the ddoscoo.ConfigLayer7CCTemplate API asynchronously
func (client *Client) ConfigLayer7CCTemplateWithChan(request *ConfigLayer7CCTemplateRequest) (<-chan *ConfigLayer7CCTemplateResponse, <-chan error) {
	responseChan := make(chan *ConfigLayer7CCTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigLayer7CCTemplate(request)
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

// ConfigLayer7CCTemplateWithCallback invokes the ddoscoo.ConfigLayer7CCTemplate API asynchronously
func (client *Client) ConfigLayer7CCTemplateWithCallback(request *ConfigLayer7CCTemplateRequest, callback func(response *ConfigLayer7CCTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigLayer7CCTemplateResponse
		var err error
		defer close(result)
		response, err = client.ConfigLayer7CCTemplate(request)
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

// ConfigLayer7CCTemplateRequest is the request struct for api ConfigLayer7CCTemplate
type ConfigLayer7CCTemplateRequest struct {
	*requests.RpcRequest
	Template        string `position:"Query" name:"Template"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// ConfigLayer7CCTemplateResponse is the response struct for api ConfigLayer7CCTemplate
type ConfigLayer7CCTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigLayer7CCTemplateRequest creates a request to invoke ConfigLayer7CCTemplate API
func CreateConfigLayer7CCTemplateRequest() (request *ConfigLayer7CCTemplateRequest) {
	request = &ConfigLayer7CCTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "ConfigLayer7CCTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateConfigLayer7CCTemplateResponse creates a response to parse from ConfigLayer7CCTemplate response
func CreateConfigLayer7CCTemplateResponse() (response *ConfigLayer7CCTemplateResponse) {
	response = &ConfigLayer7CCTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
