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
	"encoding/json"
	tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeMaterialListRequestParams struct {
	// 活动Id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeMaterialListRequest struct {
	*tchttp.BaseRequest

	// 活动Id
	ActivityId *int64 `json:"ActivityId,omitempty" name:"ActivityId"`

	// 素材Id
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 每次拉取条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMaterialListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityId")
	delete(f, "MaterialId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaterialListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaterialListResponseParams struct {
	// 素材列表数据
	MaterialInfos []*PublicMaterialInfos `json:"MaterialInfos,omitempty" name:"MaterialInfos"`

	// 素材条数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaterialListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaterialListResponseParams `json:"Response"`
}

func (r *DescribeMaterialListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaterialListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionLiteRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos"`

	// 返回图像方式（url 或 base64) ，二选一。默认url, url有效期为30天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`

	// 算法引擎参数:  1）选脸版 - youturecreat; 2）优享版 - youtu1vN； 3）畅享版 - ptu； 4）随机 - ALL;  默认为活动选择的算法
	Engine *string `json:"Engine,omitempty" name:"Engine"`
}

type FaceFusionLiteRequest struct {
	*tchttp.BaseRequest

	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos"`

	// 返回图像方式（url 或 base64) ，二选一。默认url, url有效期为30天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 请注意，鉴政服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`

	// 算法引擎参数:  1）选脸版 - youturecreat; 2）优享版 - youtu1vN； 3）畅享版 - ptu； 4）随机 - ALL;  默认为活动选择的算法
	Engine *string `json:"Engine,omitempty" name:"Engine"`
}

func (r *FaceFusionLiteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionLiteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "MergeInfos")
	delete(f, "RspImgType")
	delete(f, "CelebrityIdentify")
	delete(f, "Engine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceFusionLiteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionLiteResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 鉴政结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitempty" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FaceFusionLiteResponse struct {
	*tchttp.BaseResponse
	Response *FaceFusionLiteResponseParams `json:"Response"`
}

func (r *FaceFusionLiteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionLiteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 图片 base64 数据。请确保人脸为正脸，无旋转。若某些手机拍摄后人脸被旋转，请使用图片的 EXIF 信息对图片进行旋转处理；请勿在 base64 数据中包含头部，如“data:image/jpeg;base64,”。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 历史遗留字段，无需填写。因为融合只需提取人脸特征，不需要鉴黄。
	PornDetect *int64 `json:"PornDetect,omitempty" name:"PornDetect"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`

	// 图片Url地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type FaceFusionRequest struct {
	*tchttp.BaseRequest

	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 图片 base64 数据。请确保人脸为正脸，无旋转。若某些手机拍摄后人脸被旋转，请使用图片的 EXIF 信息对图片进行旋转处理；请勿在 base64 数据中包含头部，如“data:image/jpeg;base64,”。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 历史遗留字段，无需填写。因为融合只需提取人脸特征，不需要鉴黄。
	PornDetect *int64 `json:"PornDetect,omitempty" name:"PornDetect"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`

	// 图片Url地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *FaceFusionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "Image")
	delete(f, "PornDetect")
	delete(f, "CelebrityIdentify")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceFusionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceFusionResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 不适宜内容识别结果
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitempty" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FaceFusionResponse struct {
	*tchttp.BaseResponse
	Response *FaceFusionResponseParams `json:"Response"`
}

func (r *FaceFusionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceFusionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceInfo struct {
	// 人脸框的横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框的纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框的宽度
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框的高度
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type FaceRect struct {
	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

// Predefined struct for user
type FuseFaceRequestParams struct {
	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100]
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100]
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitempty" name:"FuseFaceDegree"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`
}

type FuseFaceRequest struct {
	*tchttp.BaseRequest

	// 活动 ID，请在人脸融合控制台查看。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 素材 ID，请在人脸融合控制台查看。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回图像方式（url 或 base64) ，二选一。url有效期为7天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 用户人脸图片、素材模板图的人脸位置信息。
	MergeInfos []*MergeInfo `json:"MergeInfos,omitempty" name:"MergeInfos"`

	// 脸型融合比例，数值越高，融合后的脸型越像素材人物。取值范围[0,100]
	// 若此参数不填写，则使用人脸融合控制台中脸型参数数值。（换脸版算法暂不支持此参数调整）
	FuseProfileDegree *int64 `json:"FuseProfileDegree,omitempty" name:"FuseProfileDegree"`

	// 五官融合比例，数值越高，融合后的五官越像素材人物。取值范围[0,100]
	// 若此参数不填写，则使用人脸融合控制台中五官参数数值。（换脸版算法暂不支持此参数调整）
	FuseFaceDegree *int64 `json:"FuseFaceDegree,omitempty" name:"FuseFaceDegree"`

	// 0表示不需要不适宜内容识别，1表示需要不适宜内容识别。默认值为0。
	// 请注意，不适宜内容识别服务开启后，您需要根据返回结果自行判断是否调整您的业务逻辑。例如提示您的用户图片非法，请更换图片。
	CelebrityIdentify *int64 `json:"CelebrityIdentify,omitempty" name:"CelebrityIdentify"`
}

func (r *FuseFaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ModelId")
	delete(f, "RspImgType")
	delete(f, "MergeInfos")
	delete(f, "FuseProfileDegree")
	delete(f, "FuseFaceDegree")
	delete(f, "CelebrityIdentify")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FuseFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FuseFaceResponseParams struct {
	// RspImgType 为 url 时，返回结果的 url， RspImgType 为 base64 时返回 base64 数据。
	FusedImage *string `json:"FusedImage,omitempty" name:"FusedImage"`

	// 不适宜内容识别结果。该数组的顺序和请求中mergeinfo的顺序一致，一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewResultSet []*FuseFaceReviewResult `json:"ReviewResultSet,omitempty" name:"ReviewResultSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FuseFaceResponse struct {
	*tchttp.BaseResponse
	Response *FuseFaceResponseParams `json:"Response"`
}

func (r *FuseFaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FuseFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FuseFaceReviewDetail struct {
	// 保留字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 人员名称
	Label *string `json:"Label,omitempty" name:"Label"`

	// 对应识别label的置信度，分数越高意味违法违规可能性越大。
	// 0到70，Suggestion建议为PASS；
	// 70到80，Suggestion建议为REVIEW；
	// 80到100，Suggestion建议为BLOCK。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
}

type FuseFaceReviewResult struct {
	// 保留字段
	Category *string `json:"Category,omitempty" name:"Category"`

	// 状态码， 0为处理成功，其他值为处理失败
	Code *string `json:"Code,omitempty" name:"Code"`

	// 对应状态码信息描述
	CodeDescription *string `json:"CodeDescription,omitempty" name:"CodeDescription"`

	// 保留字段
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 保留字段
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 审核详细内容
	DetailSet []*FuseFaceReviewDetail `json:"DetailSet,omitempty" name:"DetailSet"`
}

type MaterialFaceList struct {
	// 人脸序号
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 人脸框信息
	FaceInfo *FaceInfo `json:"FaceInfo,omitempty" name:"FaceInfo"`
}

type MergeInfo struct {
	// 输入图片base64
	Image *string `json:"Image,omitempty" name:"Image"`

	// 输入图片url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 上传的图片人脸位置信息（人脸框）
	InputImageFaceRect *FaceRect `json:"InputImageFaceRect,omitempty" name:"InputImageFaceRect"`

	// 控制台上传的素材人脸ID，不填默认取最大人脸
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" name:"TemplateFaceID"`
}

type PublicMaterialInfos struct {
	// 素材Id
	MaterialId *string `json:"MaterialId,omitempty" name:"MaterialId"`

	// 素材状态
	MaterialStatus *int64 `json:"MaterialStatus,omitempty" name:"MaterialStatus"`

	// 脸型参数P图
	BlendParamPtu *int64 `json:"BlendParamPtu,omitempty" name:"BlendParamPtu"`

	// 五官参数P图
	PositionParamPtu *int64 `json:"PositionParamPtu,omitempty" name:"PositionParamPtu"`

	// 脸型参数优图
	BlendParamYoutu *int64 `json:"BlendParamYoutu,omitempty" name:"BlendParamYoutu"`

	// 五官参数优图
	PositionParamYoutu *int64 `json:"PositionParamYoutu,omitempty" name:"PositionParamYoutu"`

	// 素材COS地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 人脸信息
	MaterialFaceList []*MaterialFaceList `json:"MaterialFaceList,omitempty" name:"MaterialFaceList"`
}

type AgeInfo struct {
	// 变化到的人脸年龄 [10,80]。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

// Predefined struct for user
type CancelFaceMorphJobRequestParams struct {
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CancelFaceMorphJobRequest struct {
	*tchttp.BaseRequest

	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *CancelFaceMorphJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFaceMorphJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFaceMorphJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelFaceMorphJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelFaceMorphJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelFaceMorphJobResponseParams `json:"Response"`
}

func (r *CancelFaceMorphJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFaceMorphJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAgePicRequestParams struct {
	// 人脸变老变年轻信息。
	// 您可以输入最多3个 AgeInfo 来实现给一张图中的最多3张人脸变老变年轻。
	AgeInfos []*AgeInfo `json:"AgeInfos,omitempty" name:"AgeInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type ChangeAgePicRequest struct {
	*tchttp.BaseRequest

	// 人脸变老变年轻信息。
	// 您可以输入最多3个 AgeInfo 来实现给一张图中的最多3张人脸变老变年轻。
	AgeInfos []*AgeInfo `json:"AgeInfos,omitempty" name:"AgeInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *ChangeAgePicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAgePicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgeInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeAgePicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeAgePicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeAgePicResponse struct {
	*tchttp.BaseResponse
	Response *ChangeAgePicResponseParams `json:"Response"`
}

func (r *ChangeAgePicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeAgePicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceCartoonPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 关闭全图动漫化，传入true（不分大小写）即关闭全图动漫化。
	DisableGlobalEffect *string `json:"DisableGlobalEffect,omitempty" name:"DisableGlobalEffect"`
}

type FaceCartoonPicRequest struct {
	*tchttp.BaseRequest

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`

	// 关闭全图动漫化，传入true（不分大小写）即关闭全图动漫化。
	DisableGlobalEffect *string `json:"DisableGlobalEffect,omitempty" name:"DisableGlobalEffect"`
}

func (r *FaceCartoonPicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceCartoonPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	delete(f, "DisableGlobalEffect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FaceCartoonPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FaceCartoonPicResponseParams struct {
	// 结果图片Base64信息。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。(默认为base64)
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FaceCartoonPicResponse struct {
	*tchttp.BaseResponse
	Response *FaceCartoonPicResponseParams `json:"Response"`
}

func (r *FaceCartoonPicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FaceCartoonPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceMorphOutput struct {
	// 人像渐变输出的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	MorphUrl *string `json:"MorphUrl,omitempty" name:"MorphUrl"`

	// 人像渐变输出的结果MD5，用于校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	MorphMd5 *string `json:"MorphMd5,omitempty" name:"MorphMd5"`

	// 人像渐变输出的结果封面图base64字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverImage *string `json:"CoverImage,omitempty" name:"CoverImage"`
}

type GenderInfo struct {
	// 选择转换方向，0：男变女，1：女变男。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

type GradientInfo struct {
	// 图片的展示时长，即单张图片静止不变的时间。GIF默认每张图片0.7s，视频默认每张图片0.5s。最大取值1s。
	Tempo *float64 `json:"Tempo,omitempty" name:"Tempo"`

	// 人像渐变的最长时间，即单张图片使用渐变特效的时间。 GIF默认值为0.5s，视频默值认为1s。最大取值1s。
	MorphTime *float64 `json:"MorphTime,omitempty" name:"MorphTime"`
}

// Predefined struct for user
type MorphFaceRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 人员人脸总数量至少2张，不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 人脸渐变参数。可调整每张图片的展示时长、人像渐变的最长时间
	GradientInfos []*GradientInfo `json:"GradientInfos,omitempty" name:"GradientInfos"`

	// 视频帧率，取值[1,25]。默认10
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 视频类型，取值0。目前仅支持MP4格式，默认为MP4格式
	OutputType *int64 `json:"OutputType,omitempty" name:"OutputType"`

	// 视频宽度，取值[128,1280]。默认值720
	OutputWidth *int64 `json:"OutputWidth,omitempty" name:"OutputWidth"`

	// 视频高度，取值[128,1280]。默认值1280
	OutputHeight *int64 `json:"OutputHeight,omitempty" name:"OutputHeight"`
}

type MorphFaceRequest struct {
	*tchttp.BaseRequest

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// 人员人脸总数量至少2张，不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitempty" name:"Images"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。jpg格式长边像素不可超过4000，其他格式图片长边像素不可超2000。
	// Url、Image必须提供一个，如果都提供，只使用 Url。图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	// 人员人脸总数量不可超过5张。
	// 若图片中包含多张人脸，只选取其中人脸面积最大的人脸。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 人脸渐变参数。可调整每张图片的展示时长、人像渐变的最长时间
	GradientInfos []*GradientInfo `json:"GradientInfos,omitempty" name:"GradientInfos"`

	// 视频帧率，取值[1,25]。默认10
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 视频类型，取值0。目前仅支持MP4格式，默认为MP4格式
	OutputType *int64 `json:"OutputType,omitempty" name:"OutputType"`

	// 视频宽度，取值[128,1280]。默认值720
	OutputWidth *int64 `json:"OutputWidth,omitempty" name:"OutputWidth"`

	// 视频高度，取值[128,1280]。默认值1280
	OutputHeight *int64 `json:"OutputHeight,omitempty" name:"OutputHeight"`
}

func (r *MorphFaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MorphFaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	delete(f, "Urls")
	delete(f, "GradientInfos")
	delete(f, "Fps")
	delete(f, "OutputType")
	delete(f, "OutputWidth")
	delete(f, "OutputHeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MorphFaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MorphFaceResponseParams struct {
	// 人像渐变任务的Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 预估处理时间，粒度为秒
	EstimatedProcessTime *int64 `json:"EstimatedProcessTime,omitempty" name:"EstimatedProcessTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MorphFaceResponse struct {
	*tchttp.BaseResponse
	Response *MorphFaceResponseParams `json:"Response"`
}

func (r *MorphFaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MorphFaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFaceMorphJobRequestParams struct {
	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type QueryFaceMorphJobRequest struct {
	*tchttp.BaseRequest

	// 人像渐变任务Job id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

func (r *QueryFaceMorphJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFaceMorphJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFaceMorphJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFaceMorphJobResponseParams struct {
	// 当前任务状态：排队中、处理中、处理失败或者处理完成
	JobStatus *string `json:"JobStatus,omitempty" name:"JobStatus"`

	// 人像渐变输出的结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceMorphOutput *FaceMorphOutput `json:"FaceMorphOutput,omitempty" name:"FaceMorphOutput"`

	// 当前任务状态码：1：排队中、3: 处理中、5: 处理失败、7:处理完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobStatusCode *int64 `json:"JobStatusCode,omitempty" name:"JobStatusCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFaceMorphJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryFaceMorphJobResponseParams `json:"Response"`
}

func (r *QueryFaceMorphJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFaceMorphJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwapGenderPicRequestParams struct {
	// 人脸转化性别信息。
	// 您可以输入最多3个 GenderInfo 来实现给一张图中的最多3张人脸转换性别。
	GenderInfos []*GenderInfo `json:"GenderInfos,omitempty" name:"GenderInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

type SwapGenderPicRequest struct {
	*tchttp.BaseRequest

	// 人脸转化性别信息。
	// 您可以输入最多3个 GenderInfo 来实现给一张图中的最多3张人脸转换性别。
	GenderInfos []*GenderInfo `json:"GenderInfos,omitempty" name:"GenderInfos"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 返回图像方式（base64 或 url ) ，二选一。url有效期为1天。
	RspImgType *string `json:"RspImgType,omitempty" name:"RspImgType"`
}

func (r *SwapGenderPicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwapGenderPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GenderInfos")
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwapGenderPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwapGenderPicResponseParams struct {
	// RspImgType 为 base64 时，返回处理后的图片 base64 数据。默认返回base64
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// RspImgType 为 url 时，返回处理后的图片 url 数据。
	ResultUrl *string `json:"ResultUrl,omitempty" name:"ResultUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SwapGenderPicResponse struct {
	*tchttp.BaseResponse
	Response *SwapGenderPicResponseParams `json:"Response"`
}

func (r *SwapGenderPicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwapGenderPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
