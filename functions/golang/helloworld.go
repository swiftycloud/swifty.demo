package main

func Main(rq *Request) (interface{}, *Responce) {
	return map[string]string{"message": "Hello, " + rq.Args["name"]}, nil
}
