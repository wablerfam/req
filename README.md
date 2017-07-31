# req
Simple HTTP Client on cross platform  
#making it as a hobby!!

## Download and installation
### Download
https://github.com/wablerfam/req/releases  
### installation 
No need. There is a binary that can be executed immediately in the decompressed compressed file

## Usage
### Basic
Linux and Windows  

    /go # req http://mock:1323/static/index.html
    {"code":200,"time":0.005131622,"body":"no print","output":"no output"}

## Options
### Output response body
Linux and Windows  

    /go # req -b http://mock:1323/hello
    {"code":200,"time":0.003842074,"body":"Hello World","output":"no output"}

### Download file
Linux and Windows  

    /go # req -o output.html http://mock:1323/static/index.html
    {"code":200,"time":0.006174548,"body":"no print","output":"output.html"}
    
    /go # #easy output -O
    /go # req -O http://mock:1323/static/index.html
    {"code":200,"time":0.008993875,"body":"no print","output":"index.html"}
    /go # ls index.html
    index.html