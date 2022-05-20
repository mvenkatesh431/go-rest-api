## Simple GO Rest API:

This project contains a simple GOLang rest API service. This service is useful to test the Kubernetes and Docker.

We also packaged it as a Docker Image. We used mutli stage build to create the small docker image. and **The size of the image is very small. It's just 12.4 MB.**
You can pull it from docker hub using below command.
```
docker pull mvenkatesh431/k8s-test-container
```
## API Details:
We are using the Gorilla Mux and Following routes are implemented

- `/info` - Returns the info about the request.
- `/health-check` - Returns JSON response as `True` if the service is running fine. We can also add custom delay by setting the `HEALTHDELAY` Environment variable.
- `/env` - Returns present container environment variables in JSON format.
- `/quote` - Returns a Random Quote from the `quotes.json` file.


## How to Run:

Clone the repo and run the `main.go` using the `go run`

```
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 17:35:37 Simple API Server(1.0.0) running on 0.0.0.0:10000
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

## Running using the Docker:


### `/info` Endpoint output:
```
{"Endpoint":"/info","Host":"127.0.0.1:10000","Method":"GET","RemoteIP":"127.0.0.1:54326","Version":"1.0.0"}
```

### Curl Example:
```
PS C:\Users\mvenk> curl 127.0.0.1:10000/info

StatusCode        : 200
StatusDescription : OK
Content           : {"Endpoint":"/info","Host":"127.0.0.1:10000","Method":"GET","RemoteIP":"127.0.0.1:54417","Version":"1.0.0"}
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
Content           : {"Version":"1.0.0","env":"ALLUSERSPROFILE=C:\\ProgramData,
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
                    Lennon","Version":"1.0.0"}


PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0.0"}  

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"If my mind can conceive it, if my heart can believe it, then I can achieve it. -
                    Muhammad Ali","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The best and most beautiful things in the world cannot be seen or even touched - they
                    must be felt with the heart. - Helen Keller","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The only way to learn a new programming language is by writing programs in it. - Dennis
                    Ritchie","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert
                    Einstein","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert
                    Einstein","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Any fool can write code that a computer can understand. Good programmers write code that
                    humans can understand. - Martin Fowler","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"Yeah We alll shine on, Like the Moon, and the Stars, And the SUN. - John
                    Lennon","Version":"1.0.0"}

PS C:\Users\mvenk> curl 127.0.0.1:10000/quote
StatusCode        : 200
StatusDescription : OK
Content           : {"Quote":"The only true wisdom is in knowing you know nothing. - Socrates","Version":"1.0.0"}

PS C:\Users\mvenk> 

```

## Changing the Port Number:

We can change the Port number where API server is running by setting the `PORT` environment variable.

By default server runs on `10000` Port.

```
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 20:02:39 Simple API Server(1.0.0) running on 0.0.0.0:10000
2022/05/07 20:02:44 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:02:44 Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert Einstein
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```

Set the `PORT` Environment variable and run the `go run` 

```
PS C:\Users\mvenk\go\src\go-rest-api> $env:PORT="32000"

PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go 
2022/05/07 20:03:33 Simple API Server(1.0.0) running on 0.0.0.0:32000
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
2022/05/11 00:41:26 Simple API Server(1.0.0) running on 0.0.0.0:10000
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
2022/05/11 00:43:43 Simple API Server(1.0.0) running on 0.0.0.0:10000
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
2022/05/07 20:20:42 Simple API Server(1.0.0) running on 0.0.0.0:10000
2022/05/07 20:20:52 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:20:52 Life is like riding a bicycle. To keep your balance, you must keep moving. - Albert Einstein
2022/05/07 20:20:58 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:20:58 Any fool can write code that a computer can understand. Good programmers write code that humans can understand. - Martin Fowler
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> $env:QUOTESFILE="./test.json"  
PS C:\Users\mvenk\go\src\go-rest-api> 
PS C:\Users\mvenk\go\src\go-rest-api> go run .\main.go
2022/05/07 20:21:15 Simple API Server(1.0.0) running on 0.0.0.0:10000
2022/05/07 20:21:17 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:21:17 Believe you can and you're halfway there. - Theodore Roosevelt
2022/05/07 20:21:24 Entering /quote endpoint, Invoked from 127.0.0.1:56541 , using ./quotes.json as QuotesSource
2022/05/07 20:21:24 I hated every minute of training, but I said, 'Don't quit. Suffer now and live the rest of your life as a champion. - Muhammad Ali
exit status 0xc000013a
PS C:\Users\mvenk\go\src\go-rest-api> 
```
> Here I have added a `test.json` file on the same folder, So I have used relative Path. Please make sure to specify the correct path for file (or use the Absolute path)


# Running on Docker:

You can create the docker container using the `docker run` like below.

```
PS C:\Users\mvenk> docker run -d -p 10000:10000 mvenkatesh431/k8s-test-container
fb32856d9d4fa2b9d779ecd802f48da12ce8a5a4fcf2afac86585bee4f0d57bc
PS C:\Users\mvenk>

PS C:\Users\mvenk\go\src\k8s-test-container> docker ps
CONTAINER ID   IMAGE                              COMMAND    CREATED          STATUS          PORTS                      NAMES
bb11cb7a12fb   mvenkatesh431/k8s-test-container   "./main"   22 seconds ago   Up 22 seconds   0.0.0.0:10000->10000/tcp   festive_edison
PS C:\Users\mvenk\go\src\k8s-test-container> 
```
By default server listens on `10000` Port on `all Interfaces (0.0.0.0)` of node. We have exposed the port `10000` to the host system. So that I can test with `curl` or any other app.

## Changing the runtime behavior on Docker:

### Changing the Port Number:

We can change the Port number where API server is running by setting the `PORT` environment variable. By default server runs on `10000` Port.

Here is an example to change the default port number using the docker environment variable. Use the `-e` option and pass the desired port number to `PORT` environment variable.

```
PS C:\Users\mvenk\go\src\k8s-test-container> docker run -d -e PORT='32000' -p 32000:32000 mvenkatesh431/k8s-test-container
b9dbfdda77dfd365ae6ff71a8117c8d29669e81620af281efdcec227976f142a
PS C:\Users\mvenk\go\src\k8s-test-container> 
```

You can check the logs to see where the server is listening.
```
PS C:\Users\mvenk\go\src\k8s-test-container> docker logs b9db
2022/05/16 18:30:35 Simple API Server(1.0.0) running on 0.0.0.0:32000
PS C:\Users\mvenk\go\src\k8s-test-container> 
```
Here it is listening on `0.0.0.0:32000` 

Make a call to any of the API endpoint or route.
```
PS C:\Users\mvenk\go\src\k8s-test-container> curl 127.0.0.1:32000/env 
StatusCode        : 200
StatusDescription : OK
Content           : {"Version":"1.0.0","env":"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin, HOSTNAME=b9dbfdda77df, PORT=32000, HOME=/root, 
                    VERSION=1.0.0"}
```

### Health-check Delayed response:

We can delay the `/health-check` endpoint response by setting the  **`HEALTHDELAY`** Environment variable. The configurable health-check API is useful to test the `Readiness and Liveness Probes of K8s`

Here I am setting the `5 Seconds` delay by passing the `5` to `HEALTHDELAY` env.

```
PS C:\Users\mvenk\go\src\k8s-test-container> docker run -d -e HEALTHDELAY='5' -p 10000:10000 mvenkatesh431/k8s-test-container
fb40641992d51228d30db145cb1bed593d48cd676407c38998c63ac2d0024fd1

PS C:\Users\mvenk\go\src\k8s-test-container> docker ps
CONTAINER ID   IMAGE                             COMMAND    CREATED         STATUS         PORTS                      NAMES
fb40641992d5   mvenkatesh431/k8s-test-container   "./main"   3 seconds ago   Up 2 seconds   0.0.0.0:10000->10000/tcp   friendly_jepsen
PS C:\Users\mvenk\go\src\k8s-test-container> 
```

Now try to access the `/health-check` API. and you will get the delayed response ( You are response delayed by 5 seconds).

```
PS C:\Users\mvenk\go\src\k8s-test-container> curl 127.0.0.1:10000/health-check
StatusCode        : 200
StatusDescription : OK
Content           : {"healthy":true}
```

We can also verify the delay by looking into the docker container logs. 

```
PS C:\Users\mvenk\go\src\k8s-test-container> docker logs fb40
2022/05/16 18:39:50 Simple API Server(1.0.0) running on 0.0.0.0:10000
2022/05/16 18:40:08 Entering /health-check endpoint, Invoked from 172.17.0.1:45816
2022/05/16 18:40:13 Added 5 seconds delay, Serving Requst Now
PS C:\Users\mvenk\go\src\k8s-test-container> 
```
We can see the request was received at ` 18:40:08` and we delayed it by 5 seconds and served the request at ` 18:40:13`.

# Running on Kubernetes:

Create a pod definition file and specify the container image as `mvenkatesh431/k8s-test-container`. Here is an example.

*simple-pod.yaml*
```
apiVersion: v1
kind: Pod
metadata: 
  name: myapi
  labels:
    type: myapi
spec:
  containers:
    - name: test-api-server
      image: mvenkatesh431/k8s-test-container
```

Let's create the pod using the `kubectl create` command, You can also use the `kubectl apply` command,

```
PS D:\Venkey\K8s> kubectl get pods
No resources found in default namespace.
PS D:\Venkey\K8s>
PS D:\Venkey\K8s> kubectl create -f .\simple-pod.yaml
pod/myapp created
PS D:\Venkey\K8s>
```

let's check if the pod is created using the kubectl get pods
```
PS D:\Venkey\K8s> kubectl create -f .\simple-pod.yaml
pod/myapi created
PS D:\Venkey\K8s>
PS D:\Venkey\K8s> kubectl get pods
NAME    READY   STATUS    RESTARTS   AGE
myapi   1/1     Running   0          37s
PS D:\Venkey\K8s>
```

Now the pod is created. You can also check the pod logs using the `kubectl log` command like below.
```
PS D:\Venkey\K8s> kubectl logs myapi
2022/05/14 10:46:38 Simple API Server(1.0.0) running on 0.0.0.0:10000
PS D:\Venkey\K8s>
```

Get more details about the pod using the `kubectl describe pod` command.
```
PS D:\Venkey\K8s> kubectl describe pod myapi
Name:         myapi
Namespace:    default
Priority:     0
Node:         minikube/192.168.49.2
Start Time:   Sat, 14 May 2022 16:16:24 +0530
Labels:       type=myapi
Annotations:  <none>
Status:       Running
IP:           172.17.0.3
IPs:
  IP:  172.17.0.3
....
....
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  42s   default-scheduler  Successfully assigned default/myapi to minikube
  Normal  Pulling    41s   kubelet            Pulling image "mvenkatesh431/k8s-test-container"
  Normal  Pulled     28s   kubelet            Successfully pulled image "mvenkatesh431/k8s-test-container" in 12.9612576s
  Normal  Created    28s   kubelet            Created container test-api-server
  Normal  Started    28s   kubelet            Started container test-api-server
PS D:\Venkey\K8s>
```

I am using `minikube` based setup, so I am login into `minikube node` to test the API.
```
PS D:\Venkey\K8s> minikube ssh
                         _             _
            _         _ ( )           ( )
  ___ ___  (_)  ___  (_)| |/')  _   _ | |_      __
/' _ ` _ `\| |/' _ `\| || , <  ( ) ( )| '_`\  /'__`\
| ( ) ( ) || || ( ) || || |\`\ | (_) || |_) )(  ___/
(_) (_) (_)(_)(_) (_)(_)(_) (_)`\___/'(_,__/'`\____)

$
```

Check all API endpoints.

```
docker@minikube:~$ curl 172.17.0.3:10000/info
{"Endpoint":"/info","Host":"172.17.0.3:10000","Method":"GET","RemoteIP":"172.17.0.1:42070","Version":"1.0.0"}
docker@minikube:~$
docker@minikube:~$ curl 172.17.0.3:10000/env
{"Version":"1.0.0","env":"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin, HOSTNAME=myapi, KUBERNETES_SERVICE_PORT=443, KUBERNETES_SERVICE_PORT_HTTPS=443, KUBERNETES_PORT=tcp://10.96.0.1:443, KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443, KUBERNETES_PORT_443_TCP_PROTO=tcp, KUBERNETES_PORT_443_TCP_PORT=443, KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1, KUBERNETES_SERVICE_HOST=10.96.0.1, HOME=/root, VERSION=1.0.0"}
docker@minikube:~$
docker@minikube:~$ curl 172.17.0.3:10000/health-check
{"healthy":true}
docker@minikube:~$
docker@minikube:~$ curl 172.17.0.3:10000/quote
{"Quote":"The best and most beautiful things in the world cannot be seen or even touched - they must be felt with the heart. - Helen Keller","Version":"1.0.0"}
docker@minikube:~$
docker@minikube:~$ curl 172.17.0.3:10000/quote
{"Quote":"Programming isn't about what you know; it's about what you can figure out. - Chris Pine","Version":"1.0.0"}
docker@minikube:~$
docker@minikube:~$ curl 172.17.0.3:10000/quote
{"Quote":"The only way to learn a new programming language is by writing programs in it. - Dennis Ritchie","Version":"1.0.0"}
docker@minikube:~$
docker@minikube:~$
```
