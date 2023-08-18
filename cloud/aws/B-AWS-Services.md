### Computing and Hosting Services
---
##### Elastic Compute Cloud (EC2)
```
EC2 provides scalable computing capacity in AWS. It allows to launch necessary virtual servers, scale based on
requirements, configure security and networking, and manage storage.
```

##### Elastic Kubernetes Service (EKS)
```
EKS is a managed Kubernetes service and AWS's container-as-a-service offering. EKS leverages EC2 for hosting cluster 
nodes and integrates AWS infrastructure (e.g. VPC networking, IAM, availability).
```

### Storage Services
---
##### S3
```
S3 is a managed blob storage that can store data in the form of "buckets" and can be infinitely scaled.
```

### Networking Services
---
##### CloudFront
```
CloudFront is a CDN service that can cache contents, limit access to the origin server (e.g. S3, HTTP server), and limit
what resources can be fetched (e.g. specific files in S3).
```

##### Route53
```
Route53 is AWS's highly available and scalable DNS service.
```

### Miscellaneous
---
##### CloudFormation
```
CloudFormation is a service that helps declaratively model and set up AWS resources. A template is used to describe all
AWS resources required and CloudFormation takes care of configuring and provisioning resources. It abstracts away the 
need to imperatively create resources and having to think through what resource is dependent on what.
```