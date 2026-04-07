package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (sayHello *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}
