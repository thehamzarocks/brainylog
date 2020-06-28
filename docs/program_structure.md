# Program Structure

A **BrainyLog** command runs a shell script that parses the command-line arguments and determines the appropriate program to run. The programs are typically written in Go and precompiled for execution.

>$ brainylog -v
>##### Displays information about the brainylog installation, including version
>##### Runs brainyLogVersion

>$ brainylog -a Some Information
>##### Adds the log to the log.bl file
>##### Runs brainyLogAdd

>$ brainylog -g Some
>##### Gets the log from the log.bl file
>##### Runs brainyLogGet

**BrainyLog** commands are quick to execute but concurrent execution would be nice to have.


