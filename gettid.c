#include <sys/syscall.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

unsigned long cgettid(void)
{
    return syscall(SYS_gettid);
}
