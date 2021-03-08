Golang app in a container with hot reloading and dynamic local package.

A good way to develop Go programs is to use a Docker container to run them. 
In addition, having hot reloading is a useful feature to save time and directly see code updates.
Finally, being able to switch easily between local packages and released packages is very important for micro-services development.

In this program, I will create a Go program that allows these three features.

Requirements:
The requirement to make it work:
Docker installed.

How to use:
To edit the program and the local package, execute:
$ make dev-local

To edit the program and use the released package, execute:
$ make dev
