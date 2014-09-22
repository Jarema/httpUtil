/*
This package is an utility for HTTP API REST JSON servers.
It handles errors, logs data and provides automatic json encoding
It's fully compatible with http/net package, so it can also be used with popular gorilla/mux library

To use it with mux:

	r := mux.NewRouter()
	r.Handle("/test", httpUtil.Handler(testHandle))

the testHandle should look like this:

	func test(w http.ResponseWriter, r *http.Request) (interface{}, *httpUtil.HandlerError) {
		res := User{1, "John", "Snow"}
		return res, nil
	}

where the User is for example:

	type User struct {
		Firstname string
		Lastname string
	}

You should return struct/map as frist parameter, and error as second.

*/
package httpUtil
