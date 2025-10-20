Usage
1. Build the Proxy
bash
go build -o main
2. Run the Proxy
bash
./main -listen :8080 -backend localhost:9001
-listen sets the address and port the proxy listens on (e.g., :8080).

-backend sets the backend server address to forward traffic to (e.g., localhost:9001).

3. Connect a TCP Client
Any TCP client connecting to localhost:8080 will be transparently proxied to localhost:9001.

Check this [project where I used it locally](https://github.com/Rajarshi-Misra/Bitcoin-AI-Agent)
