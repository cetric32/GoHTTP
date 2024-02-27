package GoHTTP

import (
	"io"
	"net/http"
	"time"
)

// GoHTTP struct to hold the http client and request.
type GoHTTP struct {
	client  *http.Client
	request *http.Request
	headers map[string]string
}

// NewGoHTTP creates a new GoHTTP instance.
// Example:
// goHttp := GoHTTP.NewGoHTTP()
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

// Get makes a GET request.
// Example:
// result, statusCode, err := goHttp.Get("https://jsonplaceholder.typicode.com/posts/1")
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

// Post makes a POST request.
// Example:
// result, statusCode, err := goHttp.Post("https://jsonplaceholder.typicode.com/posts", nil)
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

// PostForm makes a POST request with form data.
// Example:
// result, statusCode, err := goHttp.PostForm("https://jsonplaceholder.typicode.com/posts", map[string][]string{"title": {"foo"}, "body": {"bar"}, "userId": {"1"}})
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

// Put makes a PUT request.
// Example:
// result, statusCode, err := goHttp.Put("https://jsonplaceholder.typicode.com/posts/1", nil)
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

// Patch makes a PATCH request.
// Example:
// result, statusCode, err := goHttp.Patch("https://jsonplaceholder.typicode.com/posts/1", nil)
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

// Delete makes a DELETE request.
// Example:
// result, statusCode, err := goHttp.Delete("https://jsonplaceholder.typicode.com/posts/1")
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

// Do makes a request.
// Example:
// import "net/http".
// req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil).
// result, statusCode, err := goHttp.Do(req)
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

// SetTimeout sets the timeout for the request.
// Example:
// goHttp.SetTimeout(10)
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

// AddHeaders adds headers to the request.
// Example:
// goHttp.AddHeaders(map[string]string{"Content-Type": "application/json", "Authorization": "Bearer <token>"})
func (g *GoHTTP) AddHeaders(headers map[string]string) {
	g.headers = headers
}
