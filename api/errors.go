package api

import "fmt"

var (

)

type httpError struct {
	Err error
	URL string
}

type RequestCreationError struct {
	httpError
}

type RequestExecutionError struct {
	httpError
}

type ResponseDecodingError struct {
	httpError
}

func (r RequestCreationError) Error() string {
	return fmt.Sprintf("an error occurred while creating the request for url %s: %s", r.URL, r.Err.Error())
}

func (r RequestExecutionError) Error() string {
	return fmt.Sprintf("an error occurred while executing the request for url %s: %s", r.URL, r.Err.Error())
}

func (r ResponseDecodingError) Error() string {
	return fmt.Sprintf("an error occurred while decoding the response for url %s: %s", r.URL, r.Err.Error())
}

func (r RequestCreationError) Is(tgt error) bool {
	_, ok := tgt.(RequestCreationError)
	return ok
}

func (r RequestExecutionError) Is(tgt error) bool {
	_, ok := tgt.(RequestExecutionError)
	return ok
}

func (r ResponseDecodingError) Is(tgt error) bool {
	_, ok := tgt.(ResponseDecodingError)
	return ok
}

func NewRequestCreationError(url string, err error) RequestCreationError {
	return RequestCreationError{
		httpError{
			Err: err,
			URL: url,
		},
	}
}

func NewRequestExecutionError(url string, err error) RequestExecutionError {
	return RequestExecutionError{
		httpError{
			Err: err,
			URL: url,
		},
	}
}

func NewResponseDecodingError(url string, err error) ResponseDecodingError {
	return ResponseDecodingError{
		httpError{
			Err: err,
			URL: url,
		},
	}
}
