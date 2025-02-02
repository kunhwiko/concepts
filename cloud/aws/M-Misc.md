### CloudFront
---
##### CloudFront
```
CloudFront is a global AWS CDN service that caches content at edge locations. Before caching, contents from the edge are 
also fetched using the AWS internal network, leading to even lower latencies. As the service is global, CloudFront can
be used as an endpoint that is safe against DDoS attacks.
```

##### CloudFront Origin
```
a) Origin is a location where content is stored and from where CloudFront retrieves content to serve users.
b) CloudFront can read or upload to S3 bucket as an origin. To ensure that the bucket can only be accessed by CloudFront, 
   origin access controls and bucket policies need to be configured accordingly.
c) CloudFront can also serve custom origins as an HTTP backend such as application load balancers, EC2 instances, and 
   static S3 websites. In case of EC2 instances, the instances must be publicly accessible and security groups must allow
   for known CloudFront IP ranges. If a public ALB is used instead, EC2 instances can remain private and the ALB's security
   group must allow for known CloudFront IP ranges.
```

##### CloudFront Configurations
```
a) It is possible to restrict access of contents in certain countries.
b) Pricing of data transfer varies across countries. Price classes can be set to reduce the number of edge locations to
   cheaper areas.
c) When origin contents are refreshed, the cache can be invalidated by invalidating all files or files at a specific path.
   This will bypass CloudFront's TTL settings.
```

### Global Accelerator
---
##### Global Accelerator
```
Global Accelerator leverages the AWS internal network to optimize the network path between users and applications.
Global Accelerator provides two global static Anycast IPs to route traffic directly to edge locations, which will then 
forward traffic to the destination (e.g. NLBs, EC2 instances, elastic IPs) from within the AWS internal network.
```

##### Health Checks
```
Global Accelerator can declare several endpoints and will route requests to the endpoint with the lowest latency. Health
checks can be made on each endpoint, and if an endpoint fails, traffic will be routed to the next closest healthy endpoint. 
```

### SageMaker
---
##### SageMaker
```
SageMaker is a service to build machine learning models. Unlike other services such as Polly or Rekognition which focus
on a specific aspect of ML, SageMaker is a full suite that helps with the labeling of data, building/training/tuning of 
models, and deploying models to endpoints. 
```

### WAF & Shield
---
##### Web Application Firewall (WAF)
```
WAF is a service that helps protect web applications from common Layer 7 web exploits, including IP address filtering,
geolocation filtering, and rate limiting. WAF can be deployed to ALBs, CloudFront, and API Gateway.
```

##### AWS Shield
```
Shield is a managed DDoS protection service that is activated by default and protects against SYN/UDP floods, reflection
attacks, and others. A paid advanced option is available for more complex attacks along with fee protection and access to
DDoS response team.
```

##### AWS Firewall Manager
```
Firewall Manager is a service that helps manage and apply rules in all accounts within an AWS organization. Common set 
of rules include WAF rules, AWS shield rules, security groups, and network firewall rules.
```