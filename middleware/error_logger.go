package middleware

import (
	"context"
	"log"

	goa "goa.design/goa/v3/pkg"
)

//endpointに付ける
//cmd/calc/main.goに呼出しを書く

func ErrorLogger(l *log.Logger, prefix string) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		// A Goa endpoint is itself a function
		return goa.Endpoint(func(ctx context.Context, req interface{}) (interface{}, error) {
			//call the original endpoint function
			res, err := e(ctx, req)
			//Log any error

			if err != nil {
				l.Printf("MyError : %s : %s", prefix, err.Error())
			} else {
				l.Printf("MyError : %s error not exist ", prefix)
			}
			//return endpoint results
			return res, err
		})
	}
}
