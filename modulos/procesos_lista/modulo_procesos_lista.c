#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>
#include <linux/fs.h>
#include <linux/list.h>
#include <linux/slab.h>
#include <linux/string.h>
#include <linux/types.h>
#include <linux/mm.h>

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

static int my_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, write_process_list, NULL);
}

static ssize_t my_proc_write(struct file *file, const char __user *buffer, size_t count, loff_t *f_pos) {
    return 0;
}

static struct file_operations my_fops = {
    .owner = THIS_MODULE,
    .open = my_proc_open,
    .release = single_release,
    .read = seq_read,
    .llseek = seq_lseek,
    .write = my_proc_write};

static int __init process_list_mod_init(void) {
    proc_create("modulo_procesos_lista", 0, NULL, &my_fops);
    printk(KERN_INFO "@modulo_procesos_lista iniciado");
    return 0;
}

static void __exit process_list_mod_exit(void) {
    remove_proc_entry("modulo_procesos_lista", NULL);
    printk(KERN_INFO "@modulo_procesos_lista finalizado");
}

module_init(process_list_mod_init);
module_exit(process_list_mod_exit);

MODULE_LICENSE("Practica 1 - Sopes 2");
MODULE_AUTHOR("Jose Alonzo");
MODULE_DESCRIPTION("Modulo que lista los procesos con sus hijos");
MODULE_VERSION("v1");