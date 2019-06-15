# FuzzyWuzzyAPI
Fuzzing APIs with Go

## No longer using Golang for this project.  making the code public so I can talk about expierence I have with Golang. 

## Golang will not allow for illegal characters to be sent over HTTP requests.  Which is the entire reason for this project. I could build a raw socket function that could send illegal characters, but it wouldn't be supported on Windows natively without including adittional drivers.  Therefore, it defeats the purpose of the project. 


This project is meant to provide a interactive console that will allow for the automation of API testing.  
Mainly because I am lazy and want to automate my testing. 


### Why Golang?
Golang is that new hotness and is a great tool to cross-compile for multiple platforms.  Plus I wanted to learn Golang.  

### Why Did You Build It Like This?
idk.  Because i'm a n00b.  If you can do better, feel free to clone and PR.  Or just go use another tool.  

### Building
Building from Go.  Follow the insturctions online to setup Golang in your environment
then `go get github.com/ryanvillarreal/FuzzyWuzzyAPI`.  Once the package has been downloaded
you should be able to `cd ~/go/github/ryanvillarreal/FuzzyWuzzy && go build FuzzyWuzzyAPI`

## Phase 1.0
### Current To-Do List:
* ~~cli~~
* ~~Basic Request~~
* ~~Proxy Support~~
* ~~Loading Burp Request~~
* ~~Editing current requests~~
* ~~Viewing current requests~~
* ~~Payload loading~~
* Payload adding/saving
* ~~Attacking~~
* ~~Multi-Threading~~ ...Sorta
* Further HTTP/HTTP2 Configuration
* Logging
* Verbose Error Messages
* Specific Parameter Testing
* Swagger Parsing/Setup
* Progress Saving
* Statisitics
* ~~Add Banner... because why not?~~

## Phase 2.0
Refactor code to be more modular/easy to debug
