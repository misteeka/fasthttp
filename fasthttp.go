package fasthttp

import "github.com/valyala/fasthttp"

func Get(URI string) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	return
}

func Post(URI string, body []byte) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
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
	fasthttp.ReleaseRequest(req)
	return
}

func PostByClient(URI string, body []byte, client *fasthttp.Client) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetBody(body)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(URI)
	resp = fasthttp.AcquireResponse()
	err = client.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	return
}

func Put(URI string, body []byte, contentType string) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	req.Header.SetContentType("application/json")
	req.Header.SetMethodBytes([]byte("PUT"))
	resp = fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	return
}

func Delete(URI string, body []byte, contentType string) (resp *fasthttp.Response, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(URI)
	req.Header.SetContentType("application/json")
	req.Header.SetMethodBytes([]byte("DELETE"))
	resp = fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)

	fasthttp.ReleaseRequest(req)
	return
}

func ReleaseResponse(response *fasthttp.Response) {
	fasthttp.ReleaseResponse(response)
}
