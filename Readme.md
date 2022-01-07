Steps to run and test the solution :-

1. Install Golang on your machine. (recommended 1.16)
2. Open terminal and go to project directory snakesAndLadders where go.mod file is present.
3. Run "go mod tidy" and then "go mod vendor" to install all dependencies.
  (I have provided my vendor directory as well but suggest you to perform this step anyway).
4. Change value in dev.config.yml for whatever values you want to run simulation for.
5. Run "go run main.go" from terminal in snakesAndLadders directory to get output on terminal.
6. Run "go test ./..." to run all the test packages added in the solution. You can change mock values and run test again.