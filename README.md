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
kubectl get pods -l env=prod

## Agregar label a un pod
kubectl label pods podtest env=prod

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

## Generar un pod temporal y correr en su terminal
kubectl run --rm -ti podtest --image=nginx:alpine -- sh
```

### Con respecto a los namespaces
```bash
## Cambiar a otro namespace
kubectl config set-context --current --namespace=<namespace>

# EJEMPLO
federico@federico-PC:~/Documentos/Github/kubernetes-udemy/service$ kubectl run --rm -ti podtest --image=nginx:alpine -- sh
If you don't see a command prompt, try pressing enter.
/ # curl
curl: try 'curl --help' or 'curl --manual' for more information
/ # apk add -U curl
fetch https://dl-cdn.alpinelinux.org/alpine/v3.15/main/aarch64/APKINDEX.tar.gz
fetch https://dl-cdn.alpinelinux.org/alpine/v3.15/community/aarch64/APKINDEX.tar.gz
OK: 24 MiB in 42 packages
/ # apk add -U vim
(1/3) Installing xxd (8.2.4173-r0)
(2/3) Installing lua5.3-libs (5.3.6-r1)
(3/3) Installing vim (8.2.4173-r0)
Executing busybox-1.34.1-r4.trigger
OK: 52 MiB in 45 packages
/ # curl 10.43.126.207:8080
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
/ # 

## Otra forma de llamar dentro del servidor es por el nombre del servicio

/ # curl mi-servicio:8080
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

### Con respecto a los deployments
```bash
## Ver los deployments
kubectl get deployments

## Ver como fue ejecutado un deployment- si los cambios se estan realizando se muestra en tiempo real
kubectl rollout status deployment deployment-test

## Ver la historia de un deployment
kubectl rollout history deployment deployment-test

## Ver un cambio en particular
kubectl rollout history deployment deployment-test --revision=4

## Volver a una revision anterior
kubectl rollout undo deployment deployment-test --to-revision=6
```

### Con respecto a los servicios
```bash
## Ver los servicios de un label
kubectl get services -l env=prod

NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
mi-servicio   ClusterIP   10.43.126.207   <none>        8080/TCP   4m1s

# Type ClusterIP es el tipo por defecto, si no se especifica, se crea un ClusterIP. Nos da un IP interno del servidor.

## Describir un servicio
kubectl describe service mi-servicio

Name:              mi-servicio
Namespace:         default
Labels:            app=front
Annotations:       <none>
Selector:          app=front
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.43.126.207  # IP interno del servidor
IPs:               10.43.126.207
Port:              <unset>  8080/TCP  # Puerto de escucha
TargetPort:        80/TCP  # Puerto de escucha de los pods
Endpoints:         10.42.0.84:80,10.42.0.85:80,10.42.1.13:80  # IPs de los pods
Session Affinity:  None
Events:            <none>
```
- **ClusterIP** es el mas usado, es una ip virtual, permanente en el tiempo que kubernetes le asigna al servicio. Es interna, no se puede acceder desde internet.
- **NodePort** expone un puerto a nivel del nodo y permite el ingreso desde fuera del cluster.
- **LoadBalancer** el master se comporta como balanceador y manda peticiones a el servicio expuesto en alguno de los nodos.