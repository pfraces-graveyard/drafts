#define MAX_LINE_LENGTH 80

#include <regex.h>
#include <stdio.h>
#include <string.h>
 
int main(void)
{
    regex_t preg;
    char line[MAX_LINE_LENGTH];
    char match[MAX_LINE_LENGTH];
    char* pattern = "{[^{}]*}";
 
    if (regcomp(&preg, pattern, 0) != 0) {
        return 1;
    }
    
    while ( fgets(line, sizeof line, stdin) != NULL) {
        regexmatch(&preg, line, match);
        printf("%s\n", match);
    }

    regfree(&preg);
    return 0;
}

regexmatch (regex_t* preg, char* line, char* match) {
    size_t nmatch = 1;
    regmatch_t pmatch[1];

    if (regexec(preg, line, nmatch, pmatch, 0) == 0) {
        int len = pmatch[0].rm_eo - pmatch[0].rm_so;
        strncpy(match, line + pmatch[0].rm_so, len);
        match[len] = '\0';
    }
}

