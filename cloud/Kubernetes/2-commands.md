### Commands  
---
##### Version
```
kubectl version 
    --> check version
```

##### CRUD
```
kubectl run test --image=nginx 
    --> start an nginx pod named test

kubectl create deployment test --image=nginx
    --> creates an nginx deployment named test 

kubectl get <resource>
    --> retrieves specified Kubernetes resources

kubectl get <resource> -o yaml
    --> retrieves the YAML format for the Kubernetes resource

kubectl describe <resource> <name>
    --> describe a Kubernetes resource in more detail
    
kubectl rollout status <deployment-name> 
    --> waits until the deployment is ready 

kubectl delete <resource> <name>
    --> deletes specified Kubernetes resource 
```

##### Scale
```
kubectl scale deployment test --replicas 3
    --> create 3 replica Pods in the test deployment 
```

##### Expose
```
kubectl expose deployment <deployment name> --port=8888
    --> start ClusterIP
    --> Deployment listens on port 8888

kubectl expose deployment <deployment name> --port=8888 --name=test --type=NodePort 
    --> start NodePort named test 
    --> Deployment listens on port 8888, Node listens to external sources on port 30000 - 32767
    --> you can curl into the high ports (30000 - 32767) to access the K8s cluster 
```

<br />

### YAML configurations
---
##### Apply
```
kubectl apply -f filename.yml
    --> applies changes based on YAML file configs 

kubectl diff -f filename.yml
    --> finds diff of YAML file and previously applied configuration
```

##### Reference
```
kubectl api-resources
    --> quick lookup for names, versions, and kind for YAML file 

kubectl explain <resource> --recursive
    --> quick way to remind yourself how to build YAML files for a specific kind 

kubectl explain services.spec.<subfield1>.<subfield2>
    --> add and remove subfields with dot notations to modify what you want to see 
```

##### Helm Charts
```
Helm Charts are a clean way to create, manage, and apply multiple YAML files. 

Chart.yaml
   - contains the dependencies for the Chart 

values.yaml
   - defines values to write into templates 

templates
   - series of YAML files that will be kubectl applied by Helm 
```
