#include <net/inet.h>

#include "jailwrapper.h"

char** split_cmd(char* cmd, char delim)
{
    char* result[sizeof(cmd)];
    int index = 0;
    for(int i = 0; i < sizeof(cmd); i++)
    {
        if(cmd[i] == delim)
        {
            result[index] = 
        }
    }
}

pid_t jexec(char* cmd)
{
    pid_t pid = fork();
    if(pid == -1)
        return pid;
    else if(pid == 0){
        execve()
    }
}

struct JailWrapper* new_jail_wrapper(char* cmd)
{
    struct jail *_jail = (struct jail*) calloc(6, sizeof(jail));
    struct in_addr *i_addr = (struct in_addr*) calloc(1, sizeof(in_addr));
    inet_aton("0.0.0.0", i_addr);
    _jail->version = "10.2";
    _jail->path = "/tmp/";
    _jail->hostname = "spawnd";
    _jail->jailname = "spawnd";
    _jail->ip4s = 1;
    _jail->ip4 = i_addr;
    int jid = jail(_jail);
    struct JailWrapper *jail_wrapper = (struct JailWrapper*) calloc(3, JailWrapper);
    jail_wrapper->bsd_jail = _jail;
    jail_wrapper->user = getuid();

}
