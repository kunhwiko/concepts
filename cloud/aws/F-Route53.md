### Route53 Basics
---
##### Route53 Properties & APIs
```
All Route53 properties are listed here: 
  * https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_Route53.html
All Route53 APIs are listed here: 
  * https://docs.aws.amazon.com/Route53/latest/APIReference/Welcome.html
```

### Hosted Zones
---
##### Hosted Zones
```
Hosted zones are analogous to a traditional DNS zone file, representing a collection of records that can be managed 
together and belonging to a single parent domain name. All record sets within a hosted zone must have the hosted zone's 
domain name. As an example, the amazon.com hosted zone may contain records named www.amazon.com, and www.aws.amazon.com.
```

#### Accessibility
```
Hosted zones can be public for routing internet traffic or private for routing traffic within a VPC.
```