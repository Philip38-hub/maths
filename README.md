## Math
### scope
* Math is a program that reads a dataset from a .txt file on the terminal passed as an argument, then prints out the average, median, standard deviation and variance of the data set.
* Math only takes values of length of 10 digits each. Any value of greater number of digits will result to measure of central tendancies which is out of range(negative variance)


*Example of a data set* 


    189


    113


    121


    114


    145


    110


    ...

### Packages
`strconv` for string manipulation


`os` to read from the terminal and to open the file

`bufio` to read the file line by line

`fmt` to print to the terminal

`math` to allow usage of functions such as `math.Sqrt`

### Functions
`func Median(x []float64) float64` determines the median of the data set


`func Average(x []float64) float64` determines the average/ mean of the data set


`func StandardDev(x []float64) float64` determines the standard deviation od the data set

`func Variance(x []float64) float64` determines the variance of the data set


`func Sort(x []float64) []float64` sorts the values of the data set in ascending order. This helps in determining median


### Test case I
#### Data set
    6
    7
    3
    8
    10
    2  

#### Output  
    Average: 6
    Median: 7
    Standard Deviation: 3
    Variance: 8

### runnig the program
* To run the program, open a terminal. navigate to ../mainF
* Run `go run main.go {filename}.txt`