#include <stdlib.h>
#include <stdio.h>
#include <string.h>

void *pointerToArray(int argc, char *argv)
{
    char **ret = malloc(argc * sizeof(char *));
    int i;
    char *t = argv;
    for (i = 0; i < argc; i++)
    {
        ret[i] = t;
        t = strchr(t, '\n');
        *t = '\0';
        t = t + 1;
    }
    return (void *)ret;
}
