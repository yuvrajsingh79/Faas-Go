#!/bin/bash 
#### port forward 

 while [ 1 ]
    do 
    export PODS=$(kubectl get pods --namespace monitoring -l "app=grafana,release=grafana" --field-selector=status.phase==Running)
    if [ ! -z "$PODS" ]; then
        break;
    fi
    done 
​
​
kubectl port-forward $POD_NAME 3000
​
​
echo "grafana url= http://localhost:3000/login"
