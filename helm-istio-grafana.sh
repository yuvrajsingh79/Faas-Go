#!/bin/bash  
cluster_name=$(kubectl config current-context)
echo "-------checking helm is install or not---------"
hash helm  &> /dev/null
if [ $? -eq 1 ]; then
  
  echo >&2 "installing helm in ${cluster_name} "
  
  wget https://get.helm.sh/helm-v2.14.2-linux-amd64.tar.gz #download your desired version
  
  tar -zxvf helm-v2.14.2-linux-amd64.tar.gz ##unpack it
  mkdir ~/bin &&  mv linux-amd64/helm ~/bin/helm ##Find the helm binary in the unpacked directory, and move it to its desired destination
 echo "------helm version installed------"
  helm -v
  
else
    echo "helm is installed"
    helm -v
fi
echo "------initialzing helm in cluster-----" 
    helm init ####initializing#####
    echo "check tiller pod "
    export POD_NAME=$(kubectl get pods --namespace kube-system -l "name=tiller" -o jsonpath="{.items[0].metadata.name}")
   
    while [ 1 ]
    do 
    export POD_STATUS=$(kubectl get pods -n kube-system $POD_NAME -o jsonpath="Name: {.metadata.name} Status: {.status.phase}" | awk  '{print $4}')
    if [ $POD_STATUS == 'Running' ]; then
        break;
    fi
    done
​
echo "########installing prometheus and grafana#######"
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
​
echo "####install prometheus#####"
helm install --name prometheus stable/prometheus --set server.persistentVolume.enable=false
echo "check prometheus pod "

​
while [ 1 ]
    do 
    PODS=$(kubectl get pods -l app=prometheus --field-selector=status.phase==Running) 
    if [ ! -z "$PODS" ]; then
        break;
    fi
    done 
​
​
echo "####install grafana#####"
helm install --name grafana stable/grafana
​
echo "###check grafana pod###"

    while [ 1 ]
    do 
    export PODS=$(kubectl get pods -l app=grafana --field-selector=status.phase==Running)
      if [ ! -z "$PODS" ]; then
        break;
    fi
    done 
echo "########installing nginx#########"
helm install --name nginx-ingress ​stable/nginx-ingress --set controller.service.type=NodePort
​
​
echo "###check ingress pod###"

   
    while [ 1 ]
    do 
    export PODS=$(kubectl get pods -l app=nginx-ingress --field-selector=status.phase==Running)
    if [ ! -z "$PODS" ]; then
        break;
    fi
    done 

echo "####installing istio####"
helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.1.7/charts/
echo "#### istio added to helm repo###"
helm repo list | grep istio.io

echo "#### checking istio added or not ####"
while [ 1 ]
    do 
    helm repo list | grep istio.io
      if [ $? -eq 0 ]; then
        break;
    fi
    done 

echo "####istio init ####"
helm install --name istio-init --namespace istio-system istio.io/istio-init

while [ 1 ]
echo "###checking for all 53 crds of istio to get created###"
    do 
    export CRD=$(kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l)
      if [ $CRD == 53 ]; then
        break;
    fi
    done
echo "###istio install####"
helm install --name istio --namespace istio-system --set grafana.enabled=true istio.io/istio
echo "####check istio pods are running#####"
while [ 1 ]
    do 
    export PODS=$(kubectl get pods -n istio-system --field-selector=status.phase==Running)
    if [ ! -z "$PODS" ]; then
        break;
    fi
    done 
echo "###done###"