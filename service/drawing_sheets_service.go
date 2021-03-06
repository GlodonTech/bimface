package service

import (
	"bimface/bean"
	"bimface/config"
	"bimface/http"
	"bimface/utils"
	"fmt"
)

const getDrawingSheetsURI string = "data/v2/files/%d/drawingsheets"

//DrawingSheetsService ***
type DrawingSheetsService struct {
	AbstractService    //base class
	AccessTokenService *AccessTokenService
}

//NewDrawingSheetsService ***
func NewDrawingSheetsService(serviceClient *http.ServiceClient, endpoint *config.Endpoint,
	credential *config.Credential, accessTokenService *AccessTokenService) *DrawingSheetsService {
	o := &DrawingSheetsService{
		AbstractService: AbstractService{
			Endpoint:      endpoint,
			ServiceClient: serviceClient, //http.NewServiceClient(),
		},
		AccessTokenService: accessTokenService,
	}

	return o
}

//---------------------------------------------------------------------

func (o *DrawingSheetsService) getDrawingSheetsURI(fileID int64) string {
	return fmt.Sprintf(o.Endpoint.APIHost+getDrawingSheetsURI, fileID)
}

//GetDrawingSheets 文件转换相关: 获取文件转换的二维图纸信息
//http://static.bimface.com/book/restful/articles/api/translate/get-drawingsheets.html
//http API return json data such as:
/*
[
	{
		"portsAndViews": null,
		"viewInfo": {
			"cropBox": [
				-30480,
				-30480,
				-304800,
				30480,
				30480,
				-30.480000000000004
			],
			"elevation": 0,
			"id": "382617",
			"levelId": "",
			"name": "A102 - Plans",
			"outline": [
				2.4999999998688107,
				-1.0000000000133762,
				842.4999999998686,
				592.9999999999862
			],
			"preview": {
				"height": 724,
				"path": "m.bimface.com/e766f8bb69915a822b009e0fee6f2280/resource/thumbnails/382617/382617.png",
				"width": 1024
			},
			"thumbnails": [
				"m.bimface.com/e766f8bb69915a822b009e0fee6f2280/resource/thumbnails/382617/382617.96x96.png"
			],
			"viewPoint": {
				"origin": [
					0,
					0,
					0
				],
				"rightDirection": [
					1,
					0,
					0
				],
				"scale": 1,
				"upDirection": [
					0,
					1,
					0
				],
				"viewDirection": [
					0,
					0,
					1
				]
			},
			"viewType": "DrawingSheet"
		}
	},

	{...same as above ...},
	{...same as above ...}
]
*/
func (o *DrawingSheetsService) GetDrawingSheets(fileID int64) (*map[string]interface{}, *utils.Error) {
	accessToken, err := o.AccessTokenService.Get()
	if err != nil {
		return nil, err
	}

	headers := http.NewHeaders()
	headers.AddOAuth2Header(accessToken.Token)

	resp := o.ServiceClient.Get(o.getDrawingSheetsURI(fileID), headers.Header)

	var result *bean.GeneralResponse
	result, err = http.RespToGeneralResponse(resp)

	if err == nil {
		return result.Data.(*map[string]interface{}), nil
	}

	return nil, err
}
