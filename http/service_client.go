package http

import (
	"github.com/imroc/req"
)

//NewServiceClient ***
func NewServiceClient() *ServiceClient {
	o := &ServiceClient{
		request: req.New(),
	}
	//o.request.Header.Set("Content-Type", "application/json")

	return o
}

//----------------------------------------------------------------------

//ServiceClient ***
type ServiceClient struct {
	request *req.Req
}

// the v ...interface{} params can be
// Headers : req.Header, http.Header
// Params  : req.Param, req.QueryParam, url.Values
// Body    : *req.bodyJson, *req.bodyXml, string, []byte, bytes.Buffer
// Files   : req.FileUpload, []FileUpload
// Cookie  : http.Cookie
// Callback: req.UploadProgress, req.DownloadProgress

//Get ***
func (o *ServiceClient) Get(url string, v ...interface{}) *req.Resp {
	resp, err := o.request.Get(url, v...)
	if err == nil {
		return resp
	}
	return nil
}

//Post ***
func (o *ServiceClient) Post(url string, v ...interface{}) *req.Resp {
	resp, err := o.request.Post(url, v...)
	if err == nil {
		return resp
	}
	return nil
}

//Put ***
func (o *ServiceClient) Put(url string, v ...interface{}) *req.Resp {
	resp, err := o.request.Put(url, v...)
	if err == nil {
		return resp
	}
	return nil
}

//Delete ***
func (o *ServiceClient) Delete(url string, v ...interface{}) *req.Resp {
	resp, err := o.request.Delete(url, v...)
	if err == nil {
		return resp
	}
	return nil
}

//Head ***
func (o *ServiceClient) Head(url string, v ...interface{}) *req.Resp {
	resp, err := o.request.Head(url, v...)
	if err == nil {
		return resp
	}
	return nil
}
