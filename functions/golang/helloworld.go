package main

func Main(rq *Request) (interface{}, *Response) {
	return map[string]string{"message": "Hello, " + rq.Args["name"]}, nil
}
