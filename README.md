# ABAG - GO task
This repo contains the implementation for the alternating trees problem specified in the task.

## Setting Up
**Note**: Uses [Go 1.17+](https://golang.org/dl/)
Ensure that you have Golang installed on your system

## Testing

```
go test
```

~ Test cases are specified in the "abag_test.go" file. 

**Note**: Ensure that test coverage is always at 100%
## Examples

```
gardenRow := []int{1,2,3,4,5}
count := abag.Solution(gardenRow)
fmt.Println(count)
```

The "solution" method returns the possible number of trees to cut down for the garden to be aesthetically pleasing.

## Go DOC
```
package abag // import "personal/abag"

Package abag provides the solution function for the alternating trees
problem.

FUNCTIONS

func Solution(a []int) int
    Solution returns the number of ways of cutting out one tree if given an
    array a consisting of n integers, where A[K] denotes the height of the K-th
    tree, so that the remaining trees are aesthetically pleasing. If it is not
    possible to achieve the desired result, it returns −1. If
    the trees are already aesthetically pleasing without any removal then it returns 0.
```
