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

to install serverless framework(openwhisk) and create functions and trigger,
I am using serverless framework with openwhisk provider, also tested in IBM Blumix,
the code for both serverless functions are present in `~/Faas-Go/serverless/`

- chmod +x installServerless.sh
- ./installServerless.sh

perform port forwarding of grafana and print the url

- chmod +x grafanaPortForward.sh
- ./grafanaPortForward.sh

run `kubectl get svc -o wide | grep ingress`

- take the port
- Do a POST request to ip:port(ingress)/v1

gRPC client-server demonstration code is present inside `~/Faas-Go/grpc/`,
to perform gRPC demo run

- chmod +x grpc-setup.sh
- `./grpc-setup.sh`
