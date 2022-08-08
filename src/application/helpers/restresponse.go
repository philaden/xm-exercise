package helpers

type (
	RestResponse struct {
		IsSuccessStatusCode bool
		StatusCode          int
		Result              string
		ResultByteAray      []byte
		Error               *error
	}
)

func BaseRestResponse(isSuccessStatusCode bool, statusCode int, result string, err error) RestResponse {
	return RestResponse{
		IsSuccessStatusCode: isSuccessStatusCode,
		StatusCode:          statusCode,
		Result:              result,
		Error:               &err,
	}
}

func (response RestResponse) BaseRestResponseByte(isSuccessStatusCode bool, statusCode int, result []byte, err error) RestResponse {
	return RestResponse{
		IsSuccessStatusCode: isSuccessStatusCode,
		StatusCode:          statusCode,
		ResultByteAray:      result,
		Error:               &err,
	}
}
