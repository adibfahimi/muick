package types

import "fmt"

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

type HttpResponse struct {
	Headers map[string]string
	Version string
	Reason  string
	Body    string
	Status  int
}

func (res *HttpResponse) String() string {
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", res.Status, res.Reason)

	var headersStr string
	for key, value := range res.Headers {
		headersStr += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	bodyStr := res.Body

	return statusLine + headersStr + "\r\n" + bodyStr
}
