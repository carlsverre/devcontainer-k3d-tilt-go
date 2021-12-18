# devcontainer-k3d-tilt-go
Example of developing go microservices locally using tilt inside a devcontainer and deployed into Kubernetes (k3d)

## tutorial

1. Open this repo in locally in a vscode devcontainer or on Github Codespaces (not tested)
2. Once the devcontainer is setup without errors, run `tilt up` in the terminal
3. In another terminal run `curl localhost:8080` and see the response "Hello World!"
4. Modify `main.go` to return a different message
5. Within a second or so run `curl localhost:8080` to see your updated message

## debugging

* If you have issues with the k3d cluster, run `k3d cluster delete` followed by `k3d cluster create --registry-create k3d-registry:127.0.0.1:5432`
* If you have issues with tilt not reloading the code type `s` in the tilt console to see the log or open the tilt UI
