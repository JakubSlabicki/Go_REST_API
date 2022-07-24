# Welcome to Go_REST_API!

Hi! In this project I will tell You how to create simple project in **Go** with **REST** service. We will perform in this example **GET** operation and as a response we will get some random numbers with calculating the standard deviation of them. At the end we will create a **Dockerfile** and prepare for building a **Docker image** and **Docker container**.


# Setup steps

 - Oracle VM VirtualBox 6.1
 - Ubuntu 22.04 LTS
 - Go 1.18
 - Docker Desktop 4.10.1
 - Visual Studio Code 1.69


## VirtualBox

The VirtualBox is used for virtualization Linux OS, downloaded from [here](https://www.virtualbox.org/wiki/Downloads). 

## Ubuntu Desktop

To run **Docker** You need to download one of the newest OS version. For **Docker Desktop 4.10.1** You need a 64-bit version of either Ubuntu Jammy Jellyfish 22.04 (LTS) or Ubuntu Impish Indri 21.10. You can get newest iso version from [here](https://ubuntu.com/download/desktop). 
Recommended:
 -  disk space above 20 GB
 -  memory at least 4GB RAM

After complete Ubuntu installation on VirtualBox make sure Your PC has enabled **the visualization**. You can check it in **Task Manager**. If necessary enable the virtualization in BIOS setup menu. It's because **Docker** needs **KVM virtualization support**.

To enable "**Enable Nested VT-x/AMD-V**" option in VirtualBox You need to run this command in VirtualBox installation folders:

    VBoxManage modifyvm <YourVirtualBoxName> --nested-hw-virt on

## Go

You can download and install Go from [here](https://go.dev/doc/install).
To verify **Go** version run:

    $ go version

## Docker Desktop

To install **Docker Desktop** firstly You need to meet [system requirements](https://docs.docker.com/desktop/install/linux-install/#system-requirements) such as **KVM** support and **QEMU** version.

Then You can easily download software from [here](https://docs.docker.com/desktop/install/ubuntu/).

## Visual Studio Code

You can download software from [here](https://code.visualstudio.com/download).
After installation and first run try download extensions:

 - Go
 - Docker

# Program

Firstly make sure everything is running we create simple "Hello World" project.
Create new project in VS Code, create new file named for example "main.go" insert "Hello World" HTTP code:


    package main

    import (
    "fmt"
    "net/http"
    )

    func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
    }

    func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
    }

In terminal run: 

    go run main.go

Run in browser "*localhost:8080*" and check if response is served. Now You can try changing *Fprintf* argument to see responses. 
**Congratulations You created first Go project**.
My Go code is [there](https://github.com/JakubSlabicki/Go_REST_API/blob/main/main.go).

## go.mod and go.sum
To create "go.mod" file run in terminal:

    go mod init <modulename>
    
To create "go.sum" file run in terminal:

    go mod tidy

## Dockerfile

I have created files needed for creating Docker image:

 - .dockerignore
 - docker-compose.debug.yml
 - docker-compose.yml
 - Dockerfile
 
To run **Docker image** run in VS Code terminal:

    docker build --pull --rm -f "Dockerfile" -t <projectname>:<version> "."
   
If Your project is named "helloworld" and You want to build version 1.0 , You should run:

    docker build --pull --rm -f "Dockerfile" -t helloworld:1.0 "."

To run **Docker container** run in VS Code terminal:

    docker run --rm -d -p <port>:<port>/tcp <projectname>:<version>

f Your project is named "helloworld", You want to build version 1.0 with ports 8081 as in example , You should run:

    docker run --rm -d -p 8081:8081/tcp helloworld:1.0
