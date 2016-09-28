#include <sys/param.h>
#include <sys/jail.h>

#ifndef JAILWRAPPER_H_
#define JAILWRAPPER_H_
    typedef struct 
    {
        jail bsd_jail;
        char* user;
        int pid;
    } JailWrapper;

    struct JailWrapper* new_jail_wrapper(char* cmd);

    void set_jail_pid(struct JailWrapper* wrapper, int pid);

    void set_jail_user(struct JailWrapper* wrapper, char* user);

    bool destroy(struct JailWrapper* wrapper);
#endif
