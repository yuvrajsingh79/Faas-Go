Run the below steps to setup a multi node kubernetes cluster:

- git clone https://github.com/yuvrajsingh79/Faas-Go.git
- cd k8s-setup
- vagrant up

After all the nodes are up and running(check using vagrant status) run
`vagrant ssh k8s-master`

Do a git clone of this repo inside the master node and perform the below steps:

- cd Faas-Go
- chmod +x helm-istio-grafana.sh
- ./helm-istio-grafana.sh

To install serverless framework(openwhisk) and create env go and function and trigger
I am using serverless framework with openwhisk provider, also tested in IBM Blumix

- chmod +x installServerless.sh
- ./installServerless.sh

To port forward grafana and print url

- chmod +x grafanaPortForward.sh
- ./grafanaPortForward.sh

DO kubectl get svc -o wide | grep ingress

- take the port
- Do a POST request to ip:port(ingress)/v1
