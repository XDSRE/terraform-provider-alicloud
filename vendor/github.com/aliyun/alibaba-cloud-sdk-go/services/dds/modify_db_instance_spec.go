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

// ModifyDBInstanceSpec invokes the dds.ModifyDBInstanceSpec API synchronously
func (client *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
	response = CreateModifyDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceSpecWithChan invokes the dds.ModifyDBInstanceSpec API asynchronously
func (client *Client) ModifyDBInstanceSpecWithChan(request *ModifyDBInstanceSpecRequest) (<-chan *ModifyDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecWithCallback invokes the dds.ModifyDBInstanceSpec API asynchronously
func (client *Client) ModifyDBInstanceSpecWithCallback(request *ModifyDBInstanceSpecRequest, callback func(response *ModifyDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceSpec(request)
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

// ModifyDBInstanceSpecRequest is the request struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    string           `position:"Query" name:"DBInstanceStorage"`
	ReadonlyReplicas     string           `position:"Query" name:"ReadonlyReplicas"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	ReplicationFactor    string           `position:"Query" name:"ReplicationFactor"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	FromApp              string           `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// ModifyDBInstanceSpecResponse is the response struct for api ModifyDBInstanceSpec
type ModifyDBInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDBInstanceSpecRequest creates a request to invoke ModifyDBInstanceSpec API
func CreateModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
	request = &ModifyDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceSpec", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceSpecResponse creates a response to parse from ModifyDBInstanceSpec response
func CreateModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
	response = &ModifyDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}