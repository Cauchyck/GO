package handler

const HelloServiceName = "handler/HelloService"


type HelloServer struct{

}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello, "+request
	return nil
}
