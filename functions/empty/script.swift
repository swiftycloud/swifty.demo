struct Message: Encodable {
	var msg: String
}

func Main(rq: Request) -> (Encodable, Response?) {
	let result = Message(msg: "Hello, world")
	return ( result, nil )
}
