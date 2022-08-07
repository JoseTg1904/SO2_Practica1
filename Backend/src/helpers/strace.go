package helpers

import (
	"api/src/globals"
	"fmt"
	"log"
	"syscall"

	seccomp "github.com/seccomp/libseccomp-golang"
)

type syscallCounter []int

const maxSyscalls = 303

func (s syscallCounter) init() syscallCounter {
	s = make(syscallCounter, maxSyscalls)
	return s
}

func (s syscallCounter) getName(syscallID uint64) string {
	name, _ := seccomp.ScmpSyscall(syscallID).GetName()
	return name
}





func concatString(name string, id uint64) {
	content := fmt.Sprintf("%s() = %d", name, id)
	WriteFileSyscalls(content)
}

/*
	cleanVariableGlobals = limpia las variables
*/
func cleanVariableGlobals() {
	globals.Stringstrace = ""	
	globals.MapStrace = make(map[string]int)
	globals.ArrayStrace2 = ""
	globals.FlagDieStrace = false
	globals.ChangePID = true
}

/*
	Attach = se encarga de vinculares a un proceso y leer los syscalls
*/

func Attach(pid int) {
	log.Println("---- Empieza Attachar ----")

	cleanVariableGlobals()

	var err error
	var wstat syscall.WaitStatus
	var regs syscall.PtraceRegs
	var ss syscallCounter
	ss = ss.init()
	exit := true



	err = syscall.PtraceAttach(pid)
	if err != nil {
		log.Printf("Error Attach = %s\n", err.Error())
		globals.ChangePID = false
		goto fatal
	}

	_, err = syscall.Wait4(pid, &wstat, 0, nil)
	if err != nil {
		log.Printf("Wait4 %d err %s\n", pid, err)
		globals.ChangePID = false
		goto fatal
	}

	err = syscall.PtraceSetOptions(pid, syscall.PTRACE_O_TRACESYSGOOD)
	if err != nil {
		fmt.Printf("PtraceSetOptions Error = %s\n", err.Error())
		globals.ChangePID = false
		goto fatal
	}


	for !globals.FlagDieStrace {

		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
            	log.Printf("Regs Error = %s\n", err.Error())
				globals.FlagDieStrace = true
            	break
			}
			name := ss.getName(regs.Orig_rax)

			// fmt.Printf("name: %s, id: %d \n", name, regs.Orig_rax)
			concatString(name, regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
            log.Printf("Syscall 2 Error = %s\n", err.Error())
				globals.FlagDieStrace = true
			break
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			log.Printf("Wait4 = %d Error = %s\n", pid, err)
			globals.FlagDieStrace = true
			break
		}

		exit = !exit
	}

	fatal:
		syscall.Kill(pid, 18)
		err = syscall.PtraceDetach(pid)
		if err != nil {
			log.Println("TERMINANDO PROCESO, DETACH = ", err)
		}
}