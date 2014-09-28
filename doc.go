/*
This package is an utility for HTTP API REST JSON servers.
It handles errors, logs data and provides automatic json encoding
It's fully compatible with http/net package, so it can also be used with popular gorilla/mux library

### basic usage

To use it with mux:

	r := mux.NewRouter()
	r.Handle("/test", httpUtil.Handler(testHandle))

the testHandle should look like this:

	func test(w http.ResponseWriter, r *http.Request) (interface{}, *httpUtil.HandlerError) {
		res := User{"John", "Snow"}
		return res, nil
	}

where the User is for example:

	type User struct {
		Firstname string
		Lastname string
	}

You should return struct/map as first parameter, and error as second.
Remember that err is type HandlerError, so its expected to send Error, Message and Code.

For example:

	func test(w http.ResponseWriter, r *http.Request) (interface{}, *httpUtil.HandlerError) {
		data, err := returnedErr()s
		if err != {
			return nil, &httpUtil.HandlerError{err, "i am broken. please fix me", http.StatusServiceUnavailable}
		}
		return data, nil
	}

### setting up logger

If You will not call LogToFile(), httpUtil will log to stderr, just like default log package.
If You want to log to file, call:

	f, err := httpUtil.LogToFile("./logs", "http.log")
	if err != nil {
		//handle error
	}
	defer f.Close()

f.Close is needed to close file where logs are being pushed.
*/
package httpUtil
