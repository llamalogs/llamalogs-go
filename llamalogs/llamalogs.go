package llamalogs

import "fmt"

type LogArgs struct {
	Sender   string
	Receiver string
	Message  string
	IsError  bool
}

func Init(accountKey string, graphName string) {

}

func Log(args LogArgs) {
	fmt.Println("from the log llama!")
	p := fmt.Sprintf("log sender %s", args.Sender)
	fmt.Println(p)
}

func Hello() {
	fmt.Println("from the llama!")
}

func Bye() {
	fmt.Println("bye from the llama!")
}
