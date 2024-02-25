#include <stdio.h>
#include <stdlib.h>

// Just a test array, there are probably much better way's to set up 
// unit testing in C, but this is just for learning purposes.
int testArray[13] = {2, 5, 6, 18, 20, 21, 22, 23, 24, 25, 26, 27, 28};

int linearSearch(int needle, int haystack[], size_t size) {
    for (size_t i = 0; i < size; i++) {
        if (haystack[i] == needle) {
            return i;
        }
    }
    return -1;
}

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

int binarySearch(int needle, int haystack[], size_t size) {
    size_t start = 0;
    // I don't know if dividing by sizeof(int) is hacky or not
    // I'm sure learning a lot about how c works and how convenient
    // modern languages are.
    size_t end = size / sizeof(int);
    while (start + 1 < end) {
        size_t mid = ((end - start) / 2) + start;
        int check = haystack[mid];
        // printf("start: %lu end: %lu mid: %lu value at mid is %i\n", start, end, mid, check);
        if (check == needle) {
            return mid;
        } else if (check < needle) {
            start = mid;
        } else {
            end = mid;
        }
    }
    return -1;
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

int main(){
    testLinearSearchFound();
    testLinearSearchNotFound();
    testBinarySearchFound();
    testBinarySearchNotFound();
    return 0;
}
