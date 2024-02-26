#include <stdlib.h>
#include <stdio.h>

void bubbleSort(int arr[], size_t len) {
    len = len / sizeof(arr[0]);
    for (size_t i = 0; i < len; i++) {
        for (size_t j = 0; j < len - 1 - i; j++){
            if (arr[j] > arr[j+1]) {
                int tmp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = tmp;
            }
        }
    }
}
