# math-print-maker

# build
Run this in the same directory of `main.go`.
```
go build -o printmaker
```

# usage
```
printmaker <option>
printmaker <problem type> <option>
```
## option
- `--help`
- `--h`

## problem type
- `decimal`
    - Generate decimal addition and subtraction problems and answers in pdf format.
- `matrix`
    - Generate matrix multiplication problems and answers in pdf format.
- `quadratic`
    - Generate quadratic equation problems and answers in pdf format.

### option
#### common
- `-c <int>`, `--column <int>`   Number of columns (default 3)
- `-s <int>`, `--size <int>`     Number of problems (default 100)
- `-h`, `--help`         help for quadratic
#### `decimal`
- `-m`, `--minus`        Generate subtraction problems
- `-p`, `--plus`         Generate addition problems

#### `matrix`
- `-d <int>`, `--dim <int>`      Dimension flag (default 3)

#### `quadratic`
- none

### example
- Generate 300 decimal addition and subtraction problems in 2 columns format.
    ```
    printmaker decimal --size 300 --column 2
    ```

- Generate 100 decimal addition problems in 3 columns format.
    ```
    printmaker decimal plus
    ```

- Generate 100 decimal subtraction problems in 3 columns format.
    ```
    printmaker decimal m
    ```

- Generate 300 2D matrix multiplication problems in 3 columns format.
    ```
    printmaker matrix --dim 2
    ```

- Generate 200 quadratic equation problems in 3 columns format.
    ```
    printmaker quadratic -s 200
    ```