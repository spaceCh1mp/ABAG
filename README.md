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
