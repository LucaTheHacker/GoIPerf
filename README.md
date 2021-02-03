# GoIPerf
An easier way to run an IPerf3 test from your Go application.

## Why GoIPerf executes commands on CLI?
Since I wanted to use the official binary to run the tests, I've preferred to use this solution.

## Requirements
In order to use GoIPerf you'll need to have [IPerf3](https://iperf.fr/iperf-download.php) installed on your system available under the command "iperf3".  
If you want/need to use a custom binary is it possible to change the location by setting the variable Location to the bin path like in the example below.

    GoIPerf.Location = "/usr/bin/iperf3"

## How to install?
    
    go get github.com/LucaTheHacker/GoIPerf
