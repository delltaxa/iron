package main

var v_client string
var stop bool = false


var help_text string = ``+Fore["BLUE"]+`[show]   `+Fore["RED"]+`         `+Fore["GREEN"]+` --> `+Fore["RESET"]+`display all clients
`+Fore["BLUE"]+`[exploit]`+Fore["RED"]+` <remote>`+Fore["GREEN"]+` --> `+Fore["RESET"]+`Wait for remote to connect (`+Fore["BLUE"]+`nA=any`+Fore["RESET"]+`)
`+Fore["BLUE"]+`[mkexploit]`+Fore["RED"]+`       `+Fore["GREEN"]+` --> `+Fore["RESET"]+`Generate exploit for this listener`+"\n\n"