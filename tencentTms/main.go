package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms/v20201229"
)

const SecretID = ""
const SecretKey = ""

func main() {
	credential := common.NewCredential(SecretID, SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tms.tencentcloudapi.com"
	client, _ := tms.NewClient(credential, "ap-guangzhou", cpf)
	request := tms.NewTextModerationRequest()
	//request.Content = common.StringPtr(base64.StdEncoding.EncodeToString([]byte("Fuck you $120 for sex")))
	request.Content = common.StringPtr(base64.StdEncoding.EncodeToString([]byte("$150")))
	request.BizType = common.StringPtr("test_tms_strategy")
	response, err := client.TextModeration(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
