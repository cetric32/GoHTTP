package GoHTTP

import (
	"io"
	"net/http"
	"time"
)

// GoHTTP struct to hold the http client and request
type GoHTTP struct {
	client  *http.Client
	request *http.Request
	headers map[string]string
}

// NewGoHTTP creates a new GoHTTP instance
func NewGoHTTP() *GoHTTP {
	return &GoHTTP{
		client:  &http.Client{},
		request: &http.Request{},
	}
}

func (g *GoHTTP) setNewRequest(method string, url string, body io.Reader) error {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return err
	}

	g.request = req
	g.addHeaders()

	return nil
}

// Get makes a GET request
func (g *GoHTTP) Get(url string) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("GET", url, nil)

	if err != nil {
		return
	}

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// Post makes a POST request
func (g *GoHTTP) Post(url string, body io.Reader) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("POST", url, body)

	if err != nil {
		return
	}

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// PostForm makes a POST request with form data
func (g *GoHTTP) PostForm(url string, data map[string][]string) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("POST", url, nil)

	if err != nil {
		return
	}

	g.request.PostForm = data

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// Put makes a PUT request
func (g *GoHTTP) Put(url string, body io.Reader) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("PUT", url, body)

	if err != nil {
		return
	}

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// Patch makes a PATCH request
func (g *GoHTTP) Patch(url string, body io.Reader) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("PATCH", url, body)

	if err != nil {
		return
	}

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// Delete makes a DELETE request
func (g *GoHTTP) Delete(url string) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	err = g.setNewRequest("DELETE", url, nil)

	if err != nil {
		return
	}

	responseData, statusCode, err = g.Do(g.request)

	if err != nil {
		return
	}

	return
}

// Do makes a request
func (g *GoHTTP) Do(req *http.Request) (responseData []byte, statusCode int, err error) {
	responseData = []byte("")

	resp, err := g.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode
	responseData, err = io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return
}

// SetTimeout sets the timeout for the request
func (g *GoHTTP) SetTimeout(timeout int) {
	g.client.Timeout = time.Duration(timeout) * time.Second
}

func (g *GoHTTP) addHeaders() {
	if g.headers == nil {
		return
	}

	if g.request == nil {
		return
	}

	for key, value := range g.headers {
		g.request.Header.Set(key, value)
	}
}

// AddHeaders adds headers to the request
func (g *GoHTTP) AddHeaders(headers map[string]string) {
	g.headers = headers
}
