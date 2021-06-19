package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "goods/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Id string `json:"id"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.GoodsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		rs, err := s.Get(ctx, req.Id)
		return GetResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, id string) (rs string, err error) {
	request := GetRequest{Id: id}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).Rs, response.(GetResponse).Err
}
