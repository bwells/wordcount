#include <stdio.h>
#include "uthash.h"

typedef struct {
    char objectId[11];
    int count;

    UT_hash_handle hh;
} user;

int main(int argc, char *argv[]) {
    char *line = NULL;
    FILE *f_in;
    size_t linecap = 0;
    ssize_t linelen;
    ssize_t compare_len;
    const char match_line[] = "            \"objectId\": ";

    char objectId[11];

    user *users = NULL;
    user *s, *s_pool;
    unsigned int loc;

    if (argc == 2) {
        f_in = fopen(argv[1], "r");
        if (f_in == NULL) {
            fprintf(stderr, "File open failure\n");
            return 1;
        }
    }
    else {
        fprintf(stderr, "Supply an input file, dummy\n");
        return 1;
    }

    loc = 0;
    objectId[10] = '\0';
    compare_len = strlen(match_line);
    s_pool = (user*) malloc(1000000 * sizeof(user));
    while ((linelen = getline(&line, &linecap, f_in)) > 0) {
        if (!strncmp(match_line, line, compare_len)) {
            strncpy(objectId, line + compare_len+1, 10);

            HASH_FIND_STR(users, objectId, s);
            if (s)
                s->count += 1;
            else {
                s = s_pool + loc;
                strcpy(s->objectId, objectId);
                s->count = 1;
                HASH_ADD_STR(users, objectId, s);
                loc++;
            }
        }
    }

    printf("There are %u unique users\n", HASH_COUNT(users));

    return 0;
}