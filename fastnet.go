package fastnet

import "github.com/valyala/fasthttp"

func Get(URI string) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)
	return
}

func Post(URI string, body []byte, contentType string) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType(contentType) // e.g. "application/json"
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	return
}

func GetByClient(URI string, client *fasthttp.Client) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = client.Do(req, resp)
	return
}

func PostByClient(URI string, body []byte, contentType string, client *fasthttp.Client) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType(contentType) // e.g. "application/json"
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	return
}

func PostJSON(URI string, json []byte, postFunc func(string, []byte, string) (resp *fasthttp.Response, err error)) (resp *fasthttp.Response, err error) {
	return postFunc(URI, json, "application/json")
}
