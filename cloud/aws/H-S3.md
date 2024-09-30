### S3 Basics
---
##### S3
```
a) S3 is an infinitely scaling storage that store objects (files) into buckets. 
b) S3 buckets are created at the region level but must have a globally unique name across all regions in all accounts.
c) S3 objects can be versioned, encrypted, and tagged with metadata. These objects can be accessed through a URL that 
   comprises of s3://<bucket-name>/<key-name>. The key name is the full path to the object from within a bucket.
```

##### S3 Versioning & Replication
```
a) S3 allows for versioning at the bucket level, making it possible to revert changes or restore deleted objects.
b) S3 buckets support cross-region and same-region replication if versioning is on and proper IAM permissions are set.
   When replication is enabled, new objects are replicated automatically. Replications cannot be chained for multiple 
   buckets (e.g. A replicates to B, then B replicates to C). For existing objects, a separate batch operation job must be 
   executed to replicate objects.
```

##### S3 Storage Class
```
a) S3 has various storage class tiers. S3 standard general purpose has the highest availability at a higher average cost.
   Data that needs to be accessed less frequently can be stored in lower tier storage classes (e.g. Glacier) at a lower
   cost, but with higher latency and costs per object retrieval.
b) It is possible to transition objects between storage classes. Lifecycle rules can be used to automate the process
   of moving objects between storage classes or delete objects/versions based on certain rules (e.g. age).
```

##### S3 Performance Optimization
```
a) S3 supports multi-part upload, which parallelizes uploads for large objects.
b) S3 supports byte range fetches to download a specific portion of an object. This can be used for concurrent fetches
   for different byte ranges of the same object and improve retry times.
c) S3 supports transferring files through AWS edge locations to increase transfer speed.
```

##### S3 Additional Features
```
a) S3 events can be created to deliver notifications to Eventbridge, SQS, Lambda when an object is created, removed, 
   restored, replicated etc. 
b) By default, the bucket owner pays for all S3 storage and data transfer costs. It is possible to set requestors to
   pay the cost of data transfer if they are authenticated in AWS.
c) S3 Storage Lens provides metrics for storage usage, cost efficiency (e.g. cheaper storage class), data protection
   (e.g. versioning, encryption), ownership, event notifications, performance, and activity trends. Metrics can be
   exported for analytical purposes to help identify inefficiencies and anomalies.
d) S3 can be used to host static websites.
```

### S3 Security
---
##### S3 Security Basics
```
a) Security can be set at the user level (e.g. IAM policies) or at the resource level (e.g. bucket policies, object ACL).
b) A user can access an S3 object if the user has sufficient IAM permissions or there is a resource-level policy that 
   allows for access. If there is an explicit deny rule, the user's access will always be denied.
c) Supports encryption in transit through HTTPS endpoints to protect in-flight unencrypted data. Bucket policies can be 
   used to always enforce users to use secure transport.
```

##### S3 Encryption
```
a) SSE-S3: Server-side encryption with Amazon S3 managed keys, this is the default encryption method.
b) SSE-KMS: Server-side encryption with KMS keys. This allows for user control over keys and auditing of key usage on
   CloudTrail. Users need access to both the object and KMS keys to upload/download the object. Note that any KMS 
   quotas such as requests per second will become applicable. 
c) SSE-C: Server-side encryption with customer provided keys. Keys are provided by the user and are not stored by AWS.
   Users must provide the encryption key in the header of the request when uploading/downloading objects. Using HTTPS is
   mandatory for SSE-C.
d) Client-side encryption: Encrypt and decrypts data from the client-side.
```

##### Bucket Policy
```
Bucket policies are used to specify what API calls are allowed or denied on a bucket/object for specified principals.
These policies can be used to allow for cross account access or to allow public access. They can also be used to force
users to always encrypt data with a specific mechanism (e.g. SSE-C) as policies are evaluated before encryption.
```

##### S3 CORS
```
If a client makes a cross origin request on an S3 bucket, CORS configurations need to be set correctly to allow access 
for other origins. As an example, assume a browser makes a request to Bucket A to fetch an index.html file. That HTML 
file may contain references to images in Bucket B to render completely. To allow the browser to additionally fetch those 
images, Bucket B must set bucket A as an allowed origin.
```