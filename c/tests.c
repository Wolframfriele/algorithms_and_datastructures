#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "search.c"
#include "sort.c"

// Test Search

int testArray[13] = {2, 5, 6, 18, 20, 21, 22, 23, 24, 25, 26, 27, 28};

void testLinearSearchFound() {
    int result = linearSearch(6, testArray, sizeof(testArray));
    if (result == 2) {
        printf("testLinearSearchFound: Ok\n");
    } else {
        printf("testLinearSearchFound: Failed\n");
    }
}

void testLinearSearchNotFound() {
    int result = linearSearch(3, testArray, sizeof(testArray));
    if (result == -1) {
        printf("testLinearSearchNotFound: Ok\n");
    } else {
        printf("testLinearSearchNotFound: Failed\n");
    }
}

void testBinarySearchFound() {
    int result = binarySearch(6, testArray, sizeof(testArray));
    if (result == 2) {
        printf("testBinarySearchFound: Ok\n");
    } else {
        printf("testBinarySearchFound: Failed\n");
    }
}

void testBinarySearchNotFound() {
    int result = binarySearch(3, testArray, sizeof(testArray));
    if (result == -1) {
        printf("testBinarySearchNotFound: Ok\n");
    } else {
        printf("testBinarySearchNotFound: Failed\n");
    }
}

// Test Sorting

void testBubbleSort() {
    int result[5] = {5, 3, 2, 6, 4};
    bubbleSort(result, sizeof(result));
    int want[5] = {2, 3, 4, 5, 6};
    size_t maxi = sizeof(result) / sizeof(result[0]);
    bool succes = true;
    for (size_t i = 0;i < maxi; i++){
        if (result[i] != want[i]) {
            succes = false;
            printf("got %i, want %i\n", result[i], want[i]);
        }
    }
    if (succes) {
        printf("testBubbleSort: Ok\n");
    } else {
        printf("testBubbleSort: Failed\n");
    }
}

int main() {
    testLinearSearchFound();
    testLinearSearchNotFound();
    testBinarySearchFound();
    testBinarySearchNotFound();
    testBubbleSort();
    return 0;
}

