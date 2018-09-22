package main

import (
	"net/rpc"
	"log"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliParams struct {
	Room, User, Input string;
}

// var msg string

func main() {

	//connect remote RPC service.
	// Dial was used, if Http, use DialHTTP, rest the same
	rpc, err := rpc.Dial("tcp", "127.0.0.1:6002");
	if err != nil {
		log.Fatal(err);
	}
	ret := "";

	//list all chatrooms
	fmt.Println("Current chatroom list:");
	err3 := rpc.Call("Rect.Perimeter", "list", &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);

    //room number
	var room string
	fmt.Print("Enter the room you want to connet or create (4 digits such as 1001): ")
	fmt.Scanln(&room)
	fmt.Println("Room number: "+room)


	//type in user number
	var user string
	fmt.Print("Enter the user number you want to be(2 digits such as 01): ")
	fmt.Scanln(&user)
	fmt.Println("User number: "+user)

	//instruction to recieve chat history for current room.
	fmt.Println("Type anything to receive the current room chat history...")

	for
	{
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text (type exit to leave): ")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
		if strings.TrimSpace(input) == "exit" {break}
	
		//RPC, the third parameter should be a pointer
		err2 := rpc.Call("Rect.Area", cliParams{room, user, input}, &ret);

		if err2 != nil {
			log.Fatal(err2);
		}
		fmt.Println("chatlog: "+ret);

	}



fmt.Println("Bye! See you next time!");
fmt.Println("***********************");
}

//modified from: https://www.oudahe.com/p/41463/