### Simple GO Rest API

This project contains a simple GOLang rest API service. This service is useful to test the Kubernetes and Docker.

We are using the Gorilla Mux and Following routes are implemented

- `/health-check` - Returns JSON response as `True` if the service is running fine.
- `/env` - Returns present container environment variables in JSON format.
- `/info` - Returns the info about the request.

