package globals

import "api/src/structs"

/* VARIABLES GLOBALES */

var DataRAM structs.ResponseRAM
var Stringstrace string
var MapStrace = make(map[string]int)
// var ArrayStrace []structs.Syscall
var ArrayStrace2 string
var PID int = -1
var ChangePID bool = false
var FlagDieStrace bool = true
var Path string = "syscalls.txt"
var FlagActiveMap bool = false
var FlagActiveArray bool = false