This is the implementation of an Online Wallet using Threshold Secret Sharing.

./cmd: 
    - Here is the Api of the wallet implemented. This is not too interesting.
    - A dealer and a shareholder can compiled with the following command:
	```bash
	go build -o dealer cmd/dealer/main.go && go build -o shareholder cmd/shareholder/main.go
	```

./pkg: 
    - In this directory the algorithms relevant for Secret sharing are implemented.

./pkg/benchmarks: 
    - In this directory, the benchmarks are implemented. They can be executed with the command 
	```bash
	go test -bench . > benchmark.txt
	```
	
./pkg/cheaterDetection:
    - implements Herzberg et al. Cheater detection algorithm. This algorithm
       can not be used with redistribution

./pkg/feldman: 
    - implements Feldman's VSS
	
./pkg/redistribute: 
    - implements the redistribution algorithm
	
./pkg/secretSharing:
    - here the code is defined that is used by multiple other packages (for
      example the central datastructures)

./pkg/shamir: 
    - implements shamir's secret sharing algorithm

./pkg/vsr: 
    - implements the vsr-algorithm
