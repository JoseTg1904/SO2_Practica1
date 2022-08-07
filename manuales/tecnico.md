<p style="font-size: 18px">
Universidad de San Carlos de Guatemala
<br>
Facultad de Ingeniería
<br>
Escuela de Ciencias y Sistemas
<br>
Sistemas operativos 2
<br>
Ing. Cesar Batz Saquimux
<br>
Aux. Pablo Barillas
<br>
Aux. Orfrant Rivera
</p>

<br><br><br><br>

<h1 align="center" style="font-size: 40px; font-weight: bold;">Practica 1</h1>

<br><br><br>

<div align="center">

|  Carnet   |             Nombre              |
| :-------: | :-----------------------------: |
| 201700319 |  Estanley Rafael Cóbar García   |
| 201700965 |  José Carlos I Alonzo Colocho   |
| 201700972 | Davis Francisco Edward Enriquez |
| 201709140 |   Oscar Armin Crisostomo Ruiz   |

</div>

<br><br>

<h4 align="center" style="font-size: 18px; font-weight: bold;">Guatemala 2022</h4>

---

<br><br><br><br>

---

<h1>Tabla de Contenido:</h1>

- [**1. Endpoints HTTP**](#1-endpoints-http)
  - [**_1.1. RAM_**](#11-ram)
  - [**_1.2. Informacion de los Procesos_**](#12-informacion-de-los-procesos)
  - [**_1.3. Informacion de Cada Proceso_**](#13-informacion-de-cada-proceso)
  - [**_1.4. Informacion de Strace_**](#14-informacion-de-strace)
  - [**_1.5. Informacion de Strace para Histograma_**](#15-informacion-de-strace-para-histograma)
  - [**_1.6. Informacion de Strace para llamadas del sistema especificas_**](#16-informacion-de-strace-para-llamadas-del-sistema-especificas)
  - [**_1.7. Matar a un Proceso_**](#17-matar-a-un-proceso)
  - [**_1.8. Respuestas a Peticiones Errones_**](#18-respuestas-a-peticiones-errones)
- [**2. Strace**](#2strace)
  - [**_2.1. Funcion Attach_**](#21-funcion-attach)
  - [**_2.2. Funcion GetName_**](#22-funcion-getname)
- [**3. Modulos kernel**](#3-modulos-kernel)
    - [**_3.1. RAM_**](#31-ram)
    - [**_3.2. Cantidad de procesos_**](#32-cantidad-de-procesos)
    - [**_3.3. Arbol de procesos_**](#33-arbol-de-procesos)

## <br><br><br>

***
# **1. Endpoints HTTP**

## **_1.1. RAM_**

Devuelve la información de los parametros de la memoria RAM.

- Endpoint = /ram
- Tipo de Petición = GET

Respuesta de petición JSON:

```json
{
  "total": 123,
  "consumida": 123,
  "porcentaje": 15,
  "grafica": [
    { "tiempo": "22:48:55", "consumo": 48 },
    { "tiempo": "22:49:00", "consumo": 60 },
    { "tiempo": "22:49:05", "consumo": 50 }
  ]
}
```

## **_1.2. Informacion de los Procesos_**

Devuelve la información general de los procesos.

- Endpoint = /procesos/numeros
- Tipo de Petición = GET

Respuesta de petición JSON:

```json
{
  "procesos": 50,
  "ejecucion": 20,
  "suspendidos": 10,
  "detenidos": 15,
  "zombies": 5
}
```

## **_1.3. Informacion de Cada Proceso_**

Devuelve la información de cada uno de los procesos.

- Endpoint = /procesos/lista
- Tipo de Petición = GET

Ejemplo de Respuesta de petición JSON:

```json
{
  "procesos": [
    {
      "pid": 1,
      "nombre": "systemd",
      "usuario": "root",
      "estado": "Interruptible",
      "hijos": [
        {
          "pid": 235,
          "nombre": "systemd-journal",
          "hijos": []
        },
        {
          "pid": 247,
          "nombre": "systemd-udevd",
          "hijos": []
        },
        {
          "pid": 323,
          "nombre": "haveged",
          "hijos": []
        },
        {
          "pid": 375,
          "nombre": "dbus-daemon",
          "hijos": []
        }
      ]
    },
    {
      "pid": 2,
      "nombre": "kthreadd",
      "usuario": "root",
      "estado": "Interruptible",
      "hijos": [
        {
          "pid": 3,
          "nombre": "rcu_gp",
          "hijos": []
        },
        {
          "pid": 4,
          "nombre": "rcu_par_gp",
          "hijos": []
        },
        {
          "pid": 6,
          "nombre": "kworker/0:0H",
          "hijos": []
        },
        {
          "pid": 8,
          "nombre": "mm_percpu_wq",
          "hijos": []
        }
      ]
    },
    {
      "pid": 3,
      "nombre": "rcu_gp",
      "usuario": "root",
      "estado": "Desconocido",
      "hijos": []
    },
    {
      "pid": 4,
      "nombre": "rcu_par_gp",
      "usuario": "root",
      "estado": "Desconocido",
      "hijos": []
    }
  ]
}
```

## **_1.4. Informacion de Strace_**

Retorna la información de todas las llamadas al sistema del proceso enviado.

- Endpoint = /strace/:pid
  - Donde pid es el identificador del proceso. Por ejemplo: /strace/1562
- Tipo de Petición = GET

Respuesta de petición JSON:

```json
{
  "data": " read() - id = 0\n read() - id = 0\n write() - id = 1\n poll() - id = 7\n rt_sigaction() - id = 13\n poll() - id = 7\n poll() - id = 7\n write() - id = 1\n write() - id = 1\n write() - id = 1\n rt_sigaction() - id = 13\n rt_sigaction() - id = 13\n poll() - id = 7\n poll() - id = 7\n write() - id = 1\n write() - id = 1\n rt_sigaction() - id = 13\n write() - id = 1\n read() - id = 0\n write() - id = 1\n poll() - id = 7\n rt_sigaction() - id = 13\n poll() - id = 7\n poll() - id = 7\n write() - id = 1\n rt_sigaction() - id = 13\n write() - id = 1\n read() - id = 0\n write() - id = 1\n poll() - id = 7\n rt_sigaction() - id = 13\n poll() - id = 7\n poll() - id = 7\n write() - id = 1\n rt_sigaction() - id = 13\n write() - id = 1\n"
}
```

## **_1.5. Informacion de Strace para Histograma_**

Retorna la información de todas las llamadas al sistema del proceso enviado agrupandolos por nombre de proceso y la cantidad de veces repetidas.

- Endpoint = /strace/histograma/:pid
  - Donde pid es el identificador del proceso. Por ejemplo: /strace/histograma/1562
- Tipo de Petición = GET

Respuesta de petición JSON:

```json
{
  "names": ["read", "write", "poll", "rt_sigaction"],
  "values": [7, 22, 20, 14]
}
```

## **_1.6. Informacion de Strace para llamadas del sistema especificas_**

Retorna la información de todas las llamadas al sistema del proceso enviado agrupandolos por nombre de proceso y la cantidad de veces repetidas.

- Endpoint = /strace/syscalls/:pid
  - Donde pid es el identificador del proceso. Por ejemplo: /strace/syscalls/1562
- Tipo de Petición = POST

Cuerpo del POST:

```json
{
  "syscalls": ["write", "read"]
}
```

Respuesta de petición JSON:

```json
{
  "data": "write() = 1\nread() = 0\nread() = 0\nwrite() = 1\nwrite() = 1\nwrite() = 1\nwrite() = 1\nwrite() = 1\nwrite() = 1\nwrite() = 1\nread() = 0\nwrite() = 1\nwrite() = 1\nread() = 0\nwrite() = 1\nwrite() = 1\nwrite() = 1\nread() = 0\nwrite() = 1\nwrite() = 1\nwrite() = 1\nread() = 0\nwrite() = 1\nwrite() = 1\nwrite() = 1\nread() = 0\nwrite() = 1\nwrite() = 1\nwrite() = 1\n"
}
```

## **_1.7. Matar a un Proceso_**

Se encarga de matar el proceso especifico.

- Endpoint = /kill/:pid
  - Donde pid es el identificador del proceso. Por ejemplo: /kill/1562
- Tipo de Petición = GET

Respuesta de petición JSON:

```json
{
  "message": "Proceso eliminado satisfactoriamente"
}
```

## **_1.8. Respuestas a Peticiones Errones_**

Cuando no se realiza una petición correcta o no se ejecute correctamente un endpoint se mandara como respuesta el siguiente mensaje de error:

```json
{
  "message": "Descripción del error"
}
```

<br><br>

***

# **2.Strace**

Strace es una utilidad de línea de comandos para comprobación de errores en el sistema operativo GNU/Linux. Permite vigilar las llamadas al sistema usadas por un determinado programa y todas las señales que éste recibe

## **_2.1. Funcion Attach_**

La función Attach se encarga de vincularse a un proceso y de ejecutar el strace como si estuviesemos en un terminal. Verifica si el PID enviado se puede obtener sus llamadas al sistema, en caso correcto se ejecutará un ciclo infinito que obtiene todas las llamadas al sistema y termina cuando se cambia de proceso o cuando se termina la ejecución del proceso. En caso que el proceso no permita poder obtener sus llamadas al sistema entonces no se ejecuta y se retorna ne la función.

```shell
strace -p PID
```

Función en GO:

```golang
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
```

## **_2.2. Funcion GetName_**

Recibe como parametro el id de la llamada del sistema de tipo entero y se encarga de obtener el nombre de la llamada al sistema.

```golang
func (s syscallCounter) getName(syscallID uint64) string {
	name, _ := seccomp.ScmpSyscall(syscallID).GetName()
	return name
}
```

<br><br>

# **3. Modulos kernel**

## **_3.1. RAM_**

A continuacion se muestra la porción de codigo que hace posible la lectura de la ram de la forma mas exacta posible.

```c
    struct sysinfo info; //struct que contiene cierta informacion del sistema
    si_meminfo(&info);

    long total_memoria = info.totalram * info.mem_unit;
    long memoria_libre = info.freeram * info.mem_unit;
    long memoria_buffer = info.bufferram * info.mem_unit;
    long memoria_cache = (global_node_page_state(NR_FILE_PAGES) + info.bufferram) * info.mem_unit;

    long total_mega = total_memoria / (1024*1024); //factor de conversion a MB
    long libre_mega = memoria_libre / (1024*1024);
    long buffer_mega = memoria_buffer / (1024*1024);
    long cache_mega = memoria_cache /  (1024*1024);
    long uso_mega = (total_mega) - (libre_mega + buffer_mega + cache_mega);
    long porcentaje = (uso_mega * 100) / total_mega;

    seq_printf(archivo, "{\n");
    seq_printf(archivo, "\t\"total\": %8lu,\n", (total_memoria / (1000*1000)));
    seq_printf(archivo, "\t\"consumida\": %8lu,\n", uso_mega);
    seq_printf(archivo, "\t\"porcentaje\": %lu\n", porcentaje);
    seq_printf(archivo, "}\n");
	return 0;
```

Para su realización fue unicamente necesario la estructura sysinfo, la cual nos revela ciertos datos como la memoria total, la memoria libre, la memoria en buffer y la unidad base de todas estas medidas, para el caso de la memoria cache es necesario obtener la cantidad de paginas que ocupan cache mas la memoria en buffer para obtener dicha memoria.

Para dar el calculo aproximado en MB de las unidades solicitadas se procedio a dividir cada una de las medidas dentro de su multiplo correspondiente y a hacer el calculo necesario para la memoria en uso que es la total menos la sumatoria de la libre, en buffer y en cache; brindadonos los datos necesarios para satisfacer las necesidades de este modulo.

<br>

## **_3.2. Cantidad de procesos_**

A continuacion se muestra la porción de codigo que hace posible la cuantificacion de los tipos de procesos existenes en el sistema.

```c
    struct task_struct *task; //estructura para tareas/procesos
    long totales = 0;
    long ejecucion = 0;
    long suspendios = 0;
    long detenidos = 0;
    long zombies = 0;

    for_each_process(task) {
        switch(task->state) {
            case TASK_RUNNING:
                ejecucion++; break;

            case TASK_STOPPED:
                detenidos++; break;

            default:
                suspendios++;
        }

        if(task->exit_state == EXIT_ZOMBIE) { zombies++; }

        totales++;
    }

    seq_printf(archivo, "{\n");
    seq_printf(archivo, "\t\"procesos\": %8lu,\n", totales);
    seq_printf(archivo, "\t\"ejecucion\": %8lu,\n", ejecucion);
    seq_printf(archivo, "\t\"suspendidos\": %lu,\n", suspendios);
    seq_printf(archivo, "\t\"detenidos\": %lu,\n", detenidos);
    seq_printf(archivo, "\t\"zombies\": %lu\n", zombies);
    seq_printf(archivo, "}\n");

    return 0;
```

Para su realizacion unicamente fue necesaria la estructura task_structure la cual nos revela el estado del proceso actualmente y su estado de salida, siendo esto lo unico necesario para poder realizar la funcionalidad de dicho modulo.

Para poder recorrer la lista de procesos se utilizo un metodo propio del kernel para interactuar con cada proceso existente en el sistema operativo.

<br>

## **3.3. Arbol de procesos**

A continuacion se muestra la porción de codigo que hace posible la creacion del arbol de procesos.

```c
void setear_tipo_proceso(struct seq_file *archivo, struct task_struct *task) {
    char estado[20];

    switch(task->state) {
        case TASK_RUNNING:
            strcpy(estado, "Running"); break;

        case TASK_STOPPED:
            strcpy(estado, "Stopped"); break;

        case TASK_INTERRUPTIBLE:
            strcpy(estado, "Interruptible"); break;

        case TASK_UNINTERRUPTIBLE:
            strcpy(estado, "Uninterruptible"); break;

        case TASK_TRACED:
            strcpy(estado, "Traced"); break;

        case TASK_PARKED:
            strcpy(estado, "Parked"); break;

        case TASK_DEAD:
            strcpy(estado, "Dead"); break;

        case TASK_WAKEKILL:
            strcpy(estado, "Wakekill"); break;

        case TASK_WAKING:
            strcpy(estado, "Waking"); break;

        case TASK_NOLOAD:
            strcpy(estado, "Noload"); break;

        case TASK_NEW:
            strcpy(estado, "New"); break;

        case TASK_STATE_MAX:
            strcpy(estado, "State max"); break;

        default:
            strcpy(estado, "noState");
    }

    if (strcmp(estado, "noState") == 0) {
        switch(task->exit_state) {
            case EXIT_DEAD:
                strcpy(estado, "Dead"); break;

            case EXIT_ZOMBIE:
                strcpy(estado, "Zombie"); break;

            case EXIT_TRACE:
                strcpy(estado, "Trace"); break;

            default:
                strcpy(estado, "Desconocido");
        }
    }

    seq_printf(archivo, "\"estado\": \"%s\",\n", estado);
}

void explorar_hijos(struct seq_file *archivo, struct task_struct *task) {
    struct task_struct *childtask; // estructura para obtener hijos de procesos padres
    struct list_head *list; // estructura necesaria para recorrer lista de tareas

    seq_printf(archivo, "\"hijos\": \n");
    seq_printf(archivo, "[\n");

    list_for_each(list, &task->children) {
        childtask = list_entry(list, struct task_struct, sibling);

        seq_printf(archivo, "{\n");
        seq_printf(archivo, "\"pid\": %d,\n", childtask->pid);
        seq_printf(archivo, "\"nombre\": \"%s\",\n", childtask->comm);

        if (list_empty(&childtask->children)) { seq_printf(archivo, "\"hijos\": [ ]\n"); }
        else { explorar_hijos(archivo, childtask); }

        seq_printf(archivo, "},\n");
    }
    seq_printf(archivo, "]\n");
}

static int write_process_list(struct seq_file *archivo, void *v) {
    struct task_struct *task;      // estructura para tareas/procesos

    seq_printf(archivo, "[\n");
    for_each_process(task) {
        seq_printf(archivo, "{\n");
        seq_printf(archivo, "\"pid\": %d,\n", task->pid);
        seq_printf(archivo, "\"nombre\": \"%s\",\n", task->comm);
        seq_printf(archivo, "\"usuario\": \"%d\",\n", __kuid_val(task->real_cred->uid));
        setear_tipo_proceso(archivo, task);

        if (list_empty(&task->children)) { seq_printf(archivo, "\"hijos\": [ ]\n"); }
        else { explorar_hijos(archivo, task); }

        seq_printf(archivo, "},\n");
    }
    seq_printf(archivo, "]\n");

    return 0;
}
```

Para la realización de este modulo fue necesario la utilización de varios structs que son: task_struct y list_head los cuales nos sirven para poder obtener información del proceso y poder dar inicio a una lista enlazada respectivamente.

el metodo principal del modulo realiza un recorrido similar al del modulo anterior, sin enmbargo aqui se obtiene la informacion solicitada del proceso, y se revisa si este proceso tiene hijos o no.

Ya en el proceso denominado _explorar_hijos_ se realiza un recorrido recursivo para obtener todos los posibles hijos que pueda llegar a tener un proceso.

Y por ultimo existe un metodo denominado _setear_tipo_proceso_ el cual unicamente toma el identificador que es propio del estado del proceso y lo convierte a un string para su mayor entendibilidad.