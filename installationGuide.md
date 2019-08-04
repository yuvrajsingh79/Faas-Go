Run the below steps to setup a multi node kubernetes cluster:

- git clone https://github.com/yuvrajsingh79/Faas-Go.git
- cd k8s-setup
- vagrant up

verify vagrant status run `vagrant status`

connect to master node run
`vagrant ssh k8s-master`

clone this repository inside the master node and perform the below steps:

- cd Faas-Go
- chmod +x helm-istio-grafana.sh
- ./helm-istio-grafana.sh

perform port forwarding of grafana and print the url

- chmod +x grafanaPortForward.sh
- ./grafanaPortForward.sh

to install serverless framework(openwhisk) and create functions and trigger,

- chmod +x installServerless.sh
- ./installServerless.sh

run `kubectl get svc -o wide | grep ingress`

- take the port
- Do a GET request to ip:port(ingress)/ (it will list all the cluster resources)
-                                     /namespace (it will list all namespaces)
-                                     /deployment (it will list all deployments)
-                                     /service (it will list all services)

**Note:** I am using serverless framework with openwhisk provider, also tested in IBM Blumix,
the code for both serverless functions are present in `~/Faas-Go/serverless/`

gRPC client-server demonstration code is present inside `~/Faas-Go/grpc/`,
to perform gRPC demo run

- chmod +x grpc-setup.sh
- `./grpc-setup.sh`
