// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const FaceFusionAPIVersion = "2018-12-01"

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewDescribeMaterialListRequest() (request *DescribeMaterialListRequest) {
	request = &DescribeMaterialListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("facefusion", FaceFusionAPIVersion, "DescribeMaterialList")

	return
}

func NewDescribeMaterialListResponse() (response *DescribeMaterialListResponse) {
	response = &DescribeMaterialListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeMaterialList
// 通常通过腾讯云人脸融合的控制台可以查看到素材相关的参数数据，可以满足使用。本接口返回活动的素材数据，包括素材状态等。用于用户通过Api查看素材相关数据，方便使用。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) DescribeMaterialList(request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
	return c.DescribeMaterialListWithContext(context.Background(), request)
}

// DescribeMaterialList
// 通常通过腾讯云人脸融合的控制台可以查看到素材相关的参数数据，可以满足使用。本接口返回活动的素材数据，包括素材状态等。用于用户通过Api查看素材相关数据，方便使用。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) DescribeMaterialListWithContext(ctx context.Context, request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
	if request == nil {
		request = NewDescribeMaterialListRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeMaterialList require credential")
	}

	request.SetContext(ctx)

	response = NewDescribeMaterialListResponse()
	err = c.Send(request, response)
	return
}

func NewFaceFusionRequest() (request *FaceFusionRequest) {
	request = &FaceFusionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("facefusion", FaceFusionAPIVersion, "FaceFusion")

	return
}

func NewFaceFusionResponse() (response *FaceFusionResponse) {
	response = &FaceFusionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FaceFusion
// 本接口用于人脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//
//	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FaceFusion(request *FaceFusionRequest) (response *FaceFusionResponse, err error) {
	return c.FaceFusionWithContext(context.Background(), request)
}

// FaceFusion
// 本接口用于人脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//
//	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FaceFusionWithContext(ctx context.Context, request *FaceFusionRequest) (response *FaceFusionResponse, err error) {
	if request == nil {
		request = NewFaceFusionRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("FaceFusion require credential")
	}

	request.SetContext(ctx)

	response = NewFaceFusionResponse()
	err = c.Send(request, response)
	return
}

func NewFaceFusionLiteRequest() (request *FaceFusionLiteRequest) {
	request = &FaceFusionLiteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("facefusion", FaceFusionAPIVersion, "FaceFusionLite")

	return
}

func NewFaceFusionLiteResponse() (response *FaceFusionLiteResponse) {
	response = &FaceFusionLiteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FaceFusionLite
// 人脸融合活动专用版，不推荐使用。人脸融合接口建议使用[人脸融合](https://cloud.tencent.com/document/product/670/31061)或[选脸融合](https://cloud.tencent.com/document/product/670/37736)接口
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ACTIVITYSTATUSINVALID = "FailedOperation.ActivityStatusInvalid"
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"
//	FAILEDOPERATION_FACEFEATUREFAILED = "FailedOperation.FaceFeatureFailed"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACEPOSEFAILED = "FailedOperation.FacePoseFailed"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESHAPEINVALID = "FailedOperation.FaceShapeInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEBACKENDSERVERFAULT = "FailedOperation.FuseBackendServerFault"
//	FAILEDOPERATION_FUSEDETECTNOFACE = "FailedOperation.FuseDetectNoFace"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEEXCEEDFIVEHUNDREDKB = "FailedOperation.ImageSizeExceedFiveHundredKB"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"
//	FAILEDOPERATION_MATERIALVALUEEXCEED = "FailedOperation.MaterialValueExceed"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	INVALIDPARAMETERVALUE_ENGINEVALUEERROR = "InvalidParameterValue.EngineValueError"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
func (c *Client) FaceFusionLite(request *FaceFusionLiteRequest) (response *FaceFusionLiteResponse, err error) {
	return c.FaceFusionLiteWithContext(context.Background(), request)
}

// FaceFusionLite
// 人脸融合活动专用版，不推荐使用。人脸融合接口建议使用[人脸融合](https://cloud.tencent.com/document/product/670/31061)或[选脸融合](https://cloud.tencent.com/document/product/670/37736)接口
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ACTIVITYSTATUSINVALID = "FailedOperation.ActivityStatusInvalid"
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"
//	FAILEDOPERATION_FACEFEATUREFAILED = "FailedOperation.FaceFeatureFailed"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACEPOSEFAILED = "FailedOperation.FacePoseFailed"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESHAPEINVALID = "FailedOperation.FaceShapeInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEBACKENDSERVERFAULT = "FailedOperation.FuseBackendServerFault"
//	FAILEDOPERATION_FUSEDETECTNOFACE = "FailedOperation.FuseDetectNoFace"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEEXCEEDFIVEHUNDREDKB = "FailedOperation.ImageSizeExceedFiveHundredKB"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"
//	FAILEDOPERATION_MATERIALVALUEEXCEED = "FailedOperation.MaterialValueExceed"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	INVALIDPARAMETERVALUE_ENGINEVALUEERROR = "InvalidParameterValue.EngineValueError"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
func (c *Client) FaceFusionLiteWithContext(ctx context.Context, request *FaceFusionLiteRequest) (response *FaceFusionLiteResponse, err error) {
	if request == nil {
		request = NewFaceFusionLiteRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("FaceFusionLite require credential")
	}

	request.SetContext(ctx)

	response = NewFaceFusionLiteResponse()
	err = c.Send(request, response)
	return
}

func NewFuseFaceRequest() (request *FuseFaceRequest) {
	request = &FuseFaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("facefusion", FaceFusionAPIVersion, "FuseFace")

	return
}

func NewFuseFaceResponse() (response *FuseFaceResponse) {
	response = &FuseFaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FuseFace
// 本接口用于单脸、多脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">选脸融合接入指引</a>。
//
// 未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FuseFace(request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
	return c.FuseFaceWithContext(context.Background(), request)
}

// FuseFace
// 本接口用于单脸、多脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">选脸融合接入指引</a>。
//
// 未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//	FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//	FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//	FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//	FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//	FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//	FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//	FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//	FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//	FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//	INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//	INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FuseFaceWithContext(ctx context.Context, request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
	if request == nil {
		request = NewFuseFaceRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("FuseFace require credential")
	}

	request.SetContext(ctx)

	response = NewFuseFaceResponse()
	err = c.Send(request, response)
	return
}

const AgeChangerAPIVersion = "2020-03-04"

func NewCancelFaceMorphJobRequest() (request *CancelFaceMorphJobRequest) {
	request = &CancelFaceMorphJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "CancelFaceMorphJob")

	return
}

func NewCancelFaceMorphJobResponse() (response *CancelFaceMorphJobResponse) {
	response = &CancelFaceMorphJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CancelFaceMorphJob
// 撤销人像渐变任务请求
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//	FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CancelFaceMorphJob(request *CancelFaceMorphJobRequest) (response *CancelFaceMorphJobResponse, err error) {
	return c.CancelFaceMorphJobWithContext(context.Background(), request)
}

// CancelFaceMorphJob
// 撤销人像渐变任务请求
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//	FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CancelFaceMorphJobWithContext(ctx context.Context, request *CancelFaceMorphJobRequest) (response *CancelFaceMorphJobResponse, err error) {
	if request == nil {
		request = NewCancelFaceMorphJobRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("CancelFaceMorphJob require credential")
	}

	request.SetContext(ctx)

	response = NewCancelFaceMorphJobResponse()
	err = c.Send(request, response)
	return
}

func NewChangeAgePicRequest() (request *ChangeAgePicRequest) {
	request = &ChangeAgePicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "ChangeAgePic")

	return
}

func NewChangeAgePicResponse() (response *ChangeAgePicResponse) {
	response = &ChangeAgePicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ChangeAgePic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸变老或变年轻的图片，支持实现人脸不同年龄的变化。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ChangeAgePic(request *ChangeAgePicRequest) (response *ChangeAgePicResponse, err error) {
	return c.ChangeAgePicWithContext(context.Background(), request)
}

// ChangeAgePic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸变老或变年轻的图片，支持实现人脸不同年龄的变化。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ChangeAgePicWithContext(ctx context.Context, request *ChangeAgePicRequest) (response *ChangeAgePicResponse, err error) {
	if request == nil {
		request = NewChangeAgePicRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("ChangeAgePic require credential")
	}

	request.SetContext(ctx)

	response = NewChangeAgePicResponse()
	err = c.Send(request, response)
	return
}

func NewFaceCartoonPicRequest() (request *FaceCartoonPicRequest) {
	request = &FaceCartoonPicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "FaceCartoonPic")

	return
}

func NewFaceCartoonPicResponse() (response *FaceCartoonPicResponse) {
	response = &FaceCartoonPicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FaceCartoonPic
// 输入一张人脸照片，生成个性化的二次元动漫形象，可用于打造个性头像、趣味活动、特效类应用等场景，提升社交娱乐的体验。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"
//	FAILEDOPERATION_FACESHAPEFAILED = "FailedOperation.FaceShapeFailed"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) FaceCartoonPic(request *FaceCartoonPicRequest) (response *FaceCartoonPicResponse, err error) {
	return c.FaceCartoonPicWithContext(context.Background(), request)
}

// FaceCartoonPic
// 输入一张人脸照片，生成个性化的二次元动漫形象，可用于打造个性头像、趣味活动、特效类应用等场景，提升社交娱乐的体验。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACEDETECTFAILED = "FailedOperation.FaceDetectFailed"
//	FAILEDOPERATION_FACESHAPEFAILED = "FailedOperation.FaceShapeFailed"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) FaceCartoonPicWithContext(ctx context.Context, request *FaceCartoonPicRequest) (response *FaceCartoonPicResponse, err error) {
	if request == nil {
		request = NewFaceCartoonPicRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("FaceCartoonPic require credential")
	}

	request.SetContext(ctx)

	response = NewFaceCartoonPicResponse()
	err = c.Send(request, response)
	return
}

func NewMorphFaceRequest() (request *MorphFaceRequest) {
	request = &MorphFaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "MorphFace")

	return
}

func NewMorphFaceResponse() (response *MorphFaceResponse) {
	response = &MorphFaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// MorphFace
// 输入2-5张人脸照片，生成一段以人脸为焦点的渐变视频或GIF图，支持自定义图片播放速度、视频每秒传输帧数，可用于短视频、表情包、创意H5等应用场景，丰富静态图片的玩法。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) MorphFace(request *MorphFaceRequest) (response *MorphFaceResponse, err error) {
	return c.MorphFaceWithContext(context.Background(), request)
}

// MorphFace
// 输入2-5张人脸照片，生成一段以人脸为焦点的渐变视频或GIF图，支持自定义图片播放速度、视频每秒传输帧数，可用于短视频、表情包、创意H5等应用场景，丰富静态图片的玩法。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) MorphFaceWithContext(ctx context.Context, request *MorphFaceRequest) (response *MorphFaceResponse, err error) {
	if request == nil {
		request = NewMorphFaceRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("MorphFace require credential")
	}

	request.SetContext(ctx)

	response = NewMorphFaceResponse()
	err = c.Send(request, response)
	return
}

func NewQueryFaceMorphJobRequest() (request *QueryFaceMorphJobRequest) {
	request = &QueryFaceMorphJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "QueryFaceMorphJob")

	return
}

func NewQueryFaceMorphJobResponse() (response *QueryFaceMorphJobResponse) {
	response = &QueryFaceMorphJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// QueryFaceMorphJob
// 查询人像渐变处理进度
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) QueryFaceMorphJob(request *QueryFaceMorphJobRequest) (response *QueryFaceMorphJobResponse, err error) {
	return c.QueryFaceMorphJobWithContext(context.Background(), request)
}

// QueryFaceMorphJob
// 查询人像渐变处理进度
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETERVALUE = "InvalidParameterValue"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) QueryFaceMorphJobWithContext(ctx context.Context, request *QueryFaceMorphJobRequest) (response *QueryFaceMorphJobResponse, err error) {
	if request == nil {
		request = NewQueryFaceMorphJobRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("QueryFaceMorphJob require credential")
	}

	request.SetContext(ctx)

	response = NewQueryFaceMorphJobResponse()
	err = c.Send(request, response)
	return
}

func NewSwapGenderPicRequest() (request *SwapGenderPicRequest) {
	request = &SwapGenderPicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("ft", AgeChangerAPIVersion, "SwapGenderPic")

	return
}

func NewSwapGenderPicResponse() (response *SwapGenderPicResponse) {
	response = &SwapGenderPicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SwapGenderPic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸性别转换的图片。男变女可实现美颜、淡妆、加刘海和长发的效果；女变男可实现加胡须、变短发的效果。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SwapGenderPic(request *SwapGenderPicRequest) (response *SwapGenderPicResponse, err error) {
	return c.SwapGenderPicWithContext(context.Background(), request)
}

// SwapGenderPic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸性别转换的图片。男变女可实现美颜、淡妆、加刘海和长发的效果；女变男可实现加胡须、变短发的效果。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//	FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//	FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//	FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//	FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//	FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//	FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//	FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//	FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//	FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//	INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//	INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//	INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//	INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//	RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//	RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//	RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//	RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//	RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//	RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//	RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//	RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//	RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//	RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//	RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SwapGenderPicWithContext(ctx context.Context, request *SwapGenderPicRequest) (response *SwapGenderPicResponse, err error) {
	if request == nil {
		request = NewSwapGenderPicRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("SwapGenderPic require credential")
	}

	request.SetContext(ctx)

	response = NewSwapGenderPicResponse()
	err = c.Send(request, response)
	return
}
