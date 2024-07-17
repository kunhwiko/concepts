### Route53 Basics
---
##### Route53
```
Route53 is AWS's highly available and scalable authoritative DNS service.
```

##### Route53 Properties & APIs
```
All Route53 properties are listed here: 
  * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_Route53.html
All Route53 APIs are listed here: 
  * https://docs.aws.amazon.com/Route53/latest/APIReference/Welcome.html
```

##### Hosted Zones
```
Hosted zones are analogous to a traditional DNS zone file, representing a collection of records that can be managed 
together and belonging to a single parent domain name. All record sets within a hosted zone must have the hosted zone's 
domain name. As an example, the amazon.com hosted zone may contain records named www.amazon.com, and www.aws.amazon.com.
```

### Route53 Configurations
---
##### TTL
```
TTL refers to how long a DNS record should be cached by the local DNS resolver. Higher TTL means less traffic to Route53,
but possibly outdated records. 
```

##### Accessibility
```
Hosted zones can be public for routing internet traffic or private for routing traffic within a VPC.
```

### Record Types & Route Policies
---
##### Alias Records
```
Alias records look similar to CNAME records, but carry the following differences:
  * Alias records are A or AAAA records that point to other AWS resources.
  * Unline CNAME records, alias records can map root domains to a different resource.
  * Alias records support health checks.
```

##### Simple Route Policy
```
Simple records are records with standard key to value mappings. Records can specify multiple values for a given key, and 
when multiple values are returned, a random value is selected by the client. Simple record sets do not support health checks.  
```

##### Weighted Route Policy
```
Weighted records allow traffic to be distributed based on a weight value. Health checks are supported for weighted 
records.
```