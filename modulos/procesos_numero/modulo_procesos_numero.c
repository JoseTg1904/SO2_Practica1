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

static int write_process_number(struct seq_file *archivo, void *v) {
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
}

static int my_proc_open(struct inode *inode, struct file *file) {
    return single_open(file, write_process_number, NULL);
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
    .write = my_proc_write
};

static int __init process_number_mod_init(void) {
    proc_create("modulo_procesos_numero", 0, NULL, &my_fops);
    printk(KERN_INFO "@modulo_procesos_numero iniciado");
    return 0;
}

static void __exit process_number_mod_exit(void) {
    remove_proc_entry("modulo_procesos_numero", NULL);
    printk(KERN_INFO "@modulo_procesos_numero finalizado");
}

module_init(process_number_mod_init);
module_exit(process_number_mod_exit);

MODULE_LICENSE("Practica 1 - Sopes 2");
MODULE_AUTHOR("Jose Alonzo");
MODULE_DESCRIPTION("Modulo que lista la cantidad de procesos en ejecucion, suspendidos, detenidos y zombies");
MODULE_VERSION("v1");