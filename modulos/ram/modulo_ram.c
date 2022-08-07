#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>
#include <linux/mmzone.h>
#include <linux/vmstat.h>

static int write_ram(struct seq_file *archivo, void *v) {
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
}

static int my_proc_open(struct inode *inode, struct file *file) {
	return single_open(file, write_ram, NULL);	
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

static int __init ram_mod_init(void) {
	proc_create("modulo_ram", 0, NULL, &my_fops);
	printk(KERN_INFO "@modulo_ram iniciado");
    return 0;
}

static void __exit ram_mod_exit(void) {
	remove_proc_entry("modulo_ram", NULL);
	printk(KERN_INFO "@modulo_ram finalizado");
}

module_init(ram_mod_init);
module_exit(ram_mod_exit);

MODULE_LICENSE("Practica 1 - Sopes 2");
MODULE_AUTHOR("Jose Alonzo");
MODULE_DESCRIPTION("Modulo que muestra la ram total y en uso en megas");
MODULE_VERSION("v1");