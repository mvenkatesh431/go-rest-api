## Simple GO Rest API:

This project contains a simple GOLang rest API service. This service is useful to test the Kubernetes and Docker.

We are using the Gorilla Mux and Following routes are implemented

- `/info` - Returns the info about the request.
- `/health-check` - Returns JSON response as `True` if the service is running fine.
- `/env` - Returns present container environment variables in JSON format.
- `/quote` - Returns a Random Quote from the `quotes.json` file.


## How to Run:

Clone the repo and run the `main.go` using the `go run`

```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 17:35:37 Simple API Server(1.0) running on 127.0.0.1:10000
2022/05/07 17:35:52 Entering /info endpoint, returning info
2022/05/07 17:43:10 Entering /info endpoint, returning info
2022/05/07 17:47:42 Entering /health-check endpoint
2022/05/07 17:49:04 Entering /env endpoint, Returning complete Env variables
2022/05/07 17:49:27 Entering /env endpoint, Returning complete Env variables
2022/05/07 17:49:56 Entering /health-check endpoint
2022/05/07 17:54:52 Entering /quote endpoint, Returning a Random Quote
2022/05/07 17:54:52 Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John Lennon
....
```

By default server listens on `10000` Port on `localhost`. Goto the `127.0.0.1:10000/info` from your browser or by using the `curl`.

#### `/info` Endpoint output:
```
{"Endpoint":"/info","Host":"127.0.0.1:10000","Method":"GET","RemoteIP":"127.0.0.1:54326","Version":"1.0"}
```

### Curl Example:
```
PS C:\Users\mvenk> curl 127.0.0.1:10000/info

StatusCode        : 200
StatusDescription : OK
Content           : {"Endpoint":"/info","Host":"127.0.0.1:10000","Method":"GET","RemoteIP":"127.0.0.1:54417","Version":"1.0"}
```

### `/health-check` endpoint:
This `/health-check` route, Returns `healthy:true` if server is running fine. So we can see the status of the server by monitoring this route. This can be used with the Kubernetes Readiness and Liveness probes

```
PS C:\Users\mvenk> curl 127.0.0.1:10000/health-check
StatusCode        : 200
StatusDescription : OK
Content           : {"healthy":true}
...
```

### `/env` endpoint:

This `/env` route dumps the all environment variables

```
PS C:\Users\mvenk> curl 127.0.0.1:10000/env
StatusCode        : 200
StatusDescription : OK
Content           : {"Version":"1.0","env":"ALLUSERSPROFILE=C:\\ProgramData,
                    APPDATA=C:\\Users\\mvenk\\AppData\\Roaming,
                    CommonProgramFiles=C:\\Progr...
```

### `/quote` endpoint:

This `/quote` route will generate random quote. We use the `quotes.json` file, which contains the list of quotes with author names. You can pass your own `json` quotes file to override the quotes.

```
PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0"}


PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0"}  

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"If my mind can conceive it, if my heart can believe it, then I can achieve it. -
                    Muhammad Ali","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The best and most beautiful things in the world cannot be seen or even touched - they
                    must be felt with the heart. - Helen Keller","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The only way to learn a new programming language is by writing programs in it. - Dennis
                    Ritchie","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert
                    Einstein","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert
                    Einstein","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Any fool can write code that a computer can understand. Good programmers write code that
                    humans can understand. - Martin Fowler","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The only true wisdom is in knowing you know nothing. - Socrates","Version":"1.0"}

PS C:\Users\mvenk> 

```


