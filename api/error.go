package api

type APIError struct {
	Message string
	Err     error
}

func (e *APIError) Error() string {
	if e.Message == "" && e.Err == nil {
		return "Unknown Aperture API error."
	}
	if e.Message == "" {
		return e.Err.Error()
	}
	if e.Err == nil {
		return e.Message
	}
	return e.Message + ": " + e.Err.Error()
}

func newAPIError(resp response, err error) *APIError {
	result := &APIError{
		Err: err,
	}
	if resp.Debug != "" {
		result.Message = resp.Debug
		return result
	}
	if resp.Message != "" {
		result.Message = resp.Message
		return result
	}
	result.Message = resp.Code
	return result
}
