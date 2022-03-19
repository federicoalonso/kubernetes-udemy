# kubernetes-udemy

Repositorio con la información de los cursos de Kubernetes hechos en udemy.

## Comandos

### Recursos de Kubctl
```bash
## Ver los recursos de kubctl, todo lo que podemos preguntar
kubectl api-resources

## Ver versiones disponibles para la api
kubectl api-versions

## Ver los nodos del cluster
kubectl get nodes
```

### Con respecto a los pods
```bash
## Crear un pod, no se utiliza, se crean deployments
kubectl run podtest --image=nginx:alpine

## Crearlo desde un archivo
kubectl apply -f pod.yaml 

## Eliminarlo desde el archivo
kubectl delete -f pod.yaml

## Se ven los pods
kubectl get pods
kubectl get pods --all-namespaces
kubectl get pods -n default # del namespace default

## Ver con detalles
kubectl get pods -o wide

## Ver manifiesto de la creacion de un pod
kubectl get pod podtest -o yaml

## Ver todos los detalles de un pod
kubectl describe pod podtest

## Ver los logs de un pod
kubectl logs podtest

## Ver los logs de un contenedor de un pod
kubectl logs podtest2 -c web2

## Elimina los pods, uno o varios
kubectl delete pod podtest1 podtest2

## Ingresar a la terminal de un pod
kubectl exec -it podtest -- /bin/bash

## Instalar en el contenedor sh
apk add -U curl
```

### Con respecto a los namespaces
```bash
## Cambiar a otro namespace
kubectl config set-context --current --namespace=<namespace>
```