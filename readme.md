## Simple GO Rest API:

This project contains a simple GOLang rest API service. This service is useful to test the Kubernetes and Docker.

We are using the Gorilla Mux and Following routes are implemented

- `/info` - Returns the info about the request.
- `/health-check` - Returns JSON response as `True` if the service is running fine. We can also add custom delay by setting the `HEALTHDELAY` Environment variable.
- `/env` - Returns present container environment variables in JSON format.
- `/quote` - Returns a Random Quote from the `quotes.json` file.


## How to Run:

Clone the repo and run the `main.go` using the `go run`

```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 17:35:37 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/07 17:35:52 Entering /info endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:43:10 Entering /info endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:47:42 Entering /health-check endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:49:04 Entering /env endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:49:27 Entering /env endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:49:56 Entering /health-check endpoint, Invoked from 127.0.0.1:56541
2022/05/07 17:54:52 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 17:54:52 Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John Lennon
PS C:\Users\mvenk\go\src\go-rest-api> 
....
```

By default server listens on `10000` Port on `all Interfaces (0.0.0.0)` of node. Goto the `127.0.0.1:10000/info` from your browser or by using the `curl`.

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
This `/health-check` route, Returns `healthy:true` if server is running fine. So we can see the status of the server by monitoring this route. This can be used with the Kubernetes Readiness and Liveness probes.
We can also add custom delay by setting the `HEALTHDELAY` environment variable.

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

## Changing the Port Number:

We can change the Port number where API server is running by setting the `PORT` environment variable.

By default server runs on `10000` Port.

```
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 20:02:39 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/07 20:02:44 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:02:44 Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert Einstein
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

Set the `PORT` Environment variable and run the `go run` 

```
PS C:\Users\mvenk\go\src\go-rest-api> $env:PORT="32000"

PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go 
2022/05/07 20:03:33 Simple API Server(1.0) running on 0.0.0.0:32000
2022/05/07 20:03:44 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:03:44 Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert Einstein
2022/05/07 20:03:52 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:03:52 If my mind can conceive it, if my heart can believe it, then I can achieve it. - Muhammad Ali
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

We can see the `PORT` is changed.

> I am using the Windows powershell, So I have used `$env:PORT=32000` to set the environment variable. If you are in Linux set using the `export` command.

## Health-check Delayed response:

We can delay the `/health-check` endpoint response by setting the  **`HEALTHDELAY`** Environment variable.

```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go    
2022/05/11 00:41:26 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/11 00:43:29 Entering /health-check endpoint, Invoked from 127.0.0.1:54043
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

Set the `HEALTHDELAY` env
```
PS C:\Users\mvenk\go\src\go-rest-api> $env:HEALTHDELAY="5"
```

Re-run the app
```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go    
2022/05/11 00:43:43 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/11 00:43:48 Entering /health-check endpoint, Invoked from 127.0.0.1:54049
2022/05/11 00:43:53 Added 5 seconds delay, Serving Requst Now
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

> Note that the response is served after 5 seconds at `00:43:53`

## Passing your own Quotes(`quotes.json`) file:

You can change the source Quotes fine by setting the `QUOTESFILE` environment variable. Please look at the following example.

```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 20:20:42 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/07 20:20:52 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:20:52 Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert Einstein
2022/05/07 20:20:58 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:20:58 Any fool can write code that a computer can understand. Good programmers write code that humans can understand. - Martin Fowler
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> $env:QUOTESFILE="./test.json"  
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 20:21:15 Simple API Server(1.0) running on 0.0.0.0:10000
2022/05/07 20:21:17 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:21:17 Believe you can and you're halfway there. - Theodore Roosevelt
2022/05/07 20:21:24 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:21:24 I hated every minute of training, but I said, 'Don't quit. Suffer now and live the rest of your life as a champion. - Muhammad Ali
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

> Here I have added a `test.json` file on the same folder, So I have used relative Path. Please make sure to specify the correct path for file (or use the Absolute path)
