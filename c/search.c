#include <stdlib.h>

int linearSearch(int needle, int haystack[], size_t len) {
    len = len / sizeof(haystack[0]);
    for (size_t i = 0; i < len; i++) {
        if (haystack[i] == needle) {
            return i;
        }
    }
    return -1;
}

int binarySearch(int needle, int* haystack, size_t len) {
    size_t start = 0;
    len = len / sizeof(haystack[0]);
    while (start + 1 < len) {
        size_t mid = ((len - start) / 2) + start;
        int check = haystack[mid];
        // printf("start: %lu len: %lu mid: %lu value at mid is %i\n", start, len, mid, check);
        if (check == needle) {
            return mid;
        } else if (check < needle) {
            start = mid;
        } else {
            len = mid;
        }
    }
    return -1;
}

