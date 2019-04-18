# TECH(K)NOW DAY TAIPEI
Saturday, 20 April 2019 from 09:00 to 18:00 (CST)  
Workshop 03 - GoLang  
  
The repository is the example code of the Golang workshop
  
Links:  
[Event Details and Ticket Information](https://www.eventbrite.co.uk/e/techknow-day-taipei-saturday-april-20-9am-to-6pm-tickets-55343822864)  
  
## Let's Go for Web!
How to build a web service with golang step-by-step.  
## Slide Contents
1. Intrudoction  
2. Install  
3. Basic Knowledge  
4. RESTful API  
5. HTML Template  
6. Static Assets  
7. Post Form  
8. User Login  
9. Cookie  
10. Socket  

## Installation
  
To run the example, you need to install gin and melody
```sh
$ go get -u github.com/gin-gonic/gin
$ go get -u gopkg.in/olahol/melody.v1
```
  
## Quick Start
```sh
# run main.go and visit localhost:8080/ on browser
$ go run main.go
```
### Web routers

localhost:8080/  
localhost:8080/login  
localhost:8080/chat  
localhost:8080/ws

### File tree
```
.
├── README.md
├── Slides
│   ├── 01. Introduction.pptx
│   ├── 02. RESTful API Service.pptx
│   └── 03. User Login.pptx
├── assets
│   └── images
│       └── gopher.png
├── main.go
└── templates
    ├── chat.html.tmpl
    ├── header.html.tmpl
    ├── index.html.tmpl
    └── login.html.tmpl
```

