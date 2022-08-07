obj-m += modulo_ram.o
all:
	make -C /lib/modules/$(shell uname -r)/build M=$(shell pwd)
modulesclean:
	make -C /lib/modules/$(shell uname -r)/build M=$(shell pwd) clean