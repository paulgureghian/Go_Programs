Let's consider a scenario, where you have a log file. A standard logfile format has basically 3 fields Timestamp and Log
Level followed by a message of some sort.

Just imagine, in a real-world application there are hundreds or thousands of logs of such kind and you are tasked to
find a way to organize this log file and fetch logs that are specifically indicating errors that occurred in the
application. So to solve this problem, we will create a Go Application which will go through all the logs in the log
file and will separate out logs with ERROR log level.
                                                                                                                        