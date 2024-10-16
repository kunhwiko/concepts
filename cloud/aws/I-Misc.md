### CloudFront
---
##### CloudFront
```
CloudFront is a global AWS CDN service that caches content at edge locations. As the service is global, CloudFront can
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