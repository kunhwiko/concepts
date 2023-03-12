### GCP Characteristics
---
##### GCP Differentiators
```
a) GCP utilizes Google's massive network and security infrastructure.
b) GCP is developed with global availiability in mind.
   Services such as load balancing are globally available rather than regionally available.
c) GCP VM instances are priced per second with a minimum runtime of 1 minute.
d) GCP data centers run on renewable energy.
```

### Locations
---
##### Region
```
a) Region (e.g. us-west1) is a geographical area that consists of multiple zones (e.g. us-west1-a).
b) Not all services are available in each region or zone.
c) Network latencies within regions should be under 1 ms in 95% of cases.
```

##### Zone
```
a) Zone is a deployment area within a region that could potentially consist of multiple clusters across different data centers. 
b) Single zone deployments are considered to be a single point of failure.
```

##### Network Edge Location
```
Connections to GCP services located in a particular metropolitan area.
```

##### Resource Scope
```
Global Resources
  a) Resources that are globally available within the same project.
  b) Examples include VPC networks, images, firewalls, routes, persistent disk snapshots.

Regional Resources
  a) Resources that are available only within the same region.
  b) Examples include subnets, regional persistent disks, regional managed instance groups.

Zonal Resources
  a) Resources that are available only within the same zone.
  b) Examples include VM instances, zonal persistent disks, zonal managed instance groups.
```

### Resource Management
---
##### Hierarchy
```
a) GCP organizes resources through an organization --> folder --> project --> resources hierarchy.
b) GCP resources are managed through Google APIs, which Google Cloud Console and command line tools use under the hood. 
```

##### Organization
```
a) Organization is the highest logical group that is only available for Google Workspace and Cloud Identity customers.
   Google Workspace or Cloud Identity accounts can have exactly one organization resource.
b) Organization provides centralized visibility and control to every resource that belongs to an organization.
   All projects must belong to some organization if an organization exists.
c) Initial IAM role for a newly created organization grants "Project Creator" and "Billing Account Creator" to all users in the domain.
```

##### Folders
```
a) Folders are logical groups that group projects or other sub-folders.
b) Folders can be used to model different departments or teams within a company.
c) Folders are optional and requires an organization as a prerequisite to use.
   This means folders are available only to Google Workspace and Cloud Identity customers.
```

##### Projects
```
a) Projects represent the smallest logical grouping in GCP and every GCP resource must belong to exactly one project.
b) Projects are identified through either a project ID or project number.
     * Project ID     : human readable identifiers that are frequently used alongside GCP APIs.
     * Project number : used for Google Cloud internal use (e.g. error logging, cloud support).  
c) Quotas limit the number of projects per account.
d) Initial IAM role for a newly created project grants Owner to the creator of the project.
```

##### IAM Inheritance
```
a) IAM policies can be set at an organization, folder, project, or resource level and are inherited throughout the hierarchy.
b) If a policy is set on a folder that grants Project Editor role to Alice, then Alice will have editor role on all projects under the folder.
c) If a policy is set on a folder that grants Project Editor role to Alice but the role is removed at a project level, Alice will still inherit and have editor role for the project.
d) If a policy is set on a folder that grants Owner to Alice but Viewer at a project level, Alice will still be granted Owner to the project.
e) If a project is moved to a new folder, it will remove previously inherited policies and inherit from its new parent resource. 
```

##### Best Practices
```
Google Workspace and Cloud Identity accounts come with super administrator roles that have high privileges.
It is best to limit the direct use of this role on GCP and assign Organization Administrator IAM roles to designated users. 
```

### Billing
---
##### Billing Accounts
```
Smallest entity that is billed is a project and a project can only be associated with one billing account.
There can still be multiple billing accounts that a project can choose from.
```

##### Billing Exports
```
a) GCP allows for exporting billing information to a BigQuery dataset or a file on Cloud Storage.
   After data is exported, users can perform queries to generate useful insights and reports.
b) Budgets and alerts can be set for each billing account or project, where the default alert thresholds are 50%, 90%, 100%.
   If an alert is attached to a billing account used by multiple projects, the alert counts towards the total cost of all projects. 
```

##### Billing Roles
```
Billing Account Creator
  a) Role for initial billing setup and to allow creation of additional billing accounts.
  b) This is an initial IAM role when an organization is created and the number of users with this role should be limited.

Billing Account Administrator
  a) Owner role of a billing account to manage payments, other user roles on the account, billing exports, and link or unlink projects to the account.
  b) Role for managing a given billing account but cannot create new billing accounts.

Billing Account Costs Manager
  a) Role for creating, editing, and deleting budgets, viewing costs and transactions, and managing billing exports.
  b) Role does not give right to export pricing data, view custom pricing, link projects, and manage properties of the billing account. 

Billing Account Viewer
  a) Role for allowing access to view billing information but nothing more.

Billing Account User
  a) Role that has very restricted permissions and can be granted broadly.
  b) If user has "Project Creator" role, this role allows for the creation of new projects linked to the given billing account.
  c) If user has "Project Billing Manager" role, this role allows for linking or unlinking those projects to the given billing account.
     This does not give the rights to resources in the project.
```

### Computing and Hosting Services
---
##### Google Compute Engine (GCE)
```
GCE is an infrastructure-as-a-service offering and gives full control over instance hardware, operating system, region/zone, networking, and autoscaling.
```

##### Google Kubernetes Engine (GKE)
```
a) GKE is Google's version of Kubernetes and container-as-a-service offering.
b) GKE leverages GCE for hosting cluster nodes and integrates GCP software defined networks.
```

##### Anthos on VMWare
```
Anthos on VMWare (a.k.a. GKE on-prem) is a GKE service that can be installed on on-premise data centers.
```

##### Google App Engine (GAE)
```
a) GAE is a platform-as-a-service offering that allows users to focus on writing code while abstracting infrastructure.
b) GAE supports two types of environments:
     * Standard: Supports a set of common languages. 
     * Flexible: Supports more languages and custom runtimes but lose some out of the box integrations.
```

##### Cloud Function
```
a) Cloud Function is a function-as-a-service offering that allows users to focus on writing functions in one of the supported languages.
b) Cloud Function is serverless and can be executed through an event trigger or HTTP endpoint.
```

##### Cloud Run
```
Cloud Run is a function-as-a-service offering that allows users to define on-demand containers that will listen for HTTP requests.
```

### Storage Services
---
##### Google Cloud Storage (GCS)
```
GCS is a managed blob storage that can store data in the form of "buckets" and can be infinitely scaled.
```

##### Filestore
```
a) Filestore is a managed file storage service (not block storage service).
b) Filestore is a network attached storage that can be used alongside GCE or GKE.
c) Filestore comes with two tiers, "premium" and "standard", which differs for IOPS and throughput.
```

##### Cloud SQL
```
a) Cloud SQL is a managed relational database for MySQL or PostgreSQL and allows users to easily migrate existing MySQL or PostgreSQL databases.
b) Cloud SQL offers data replication, backup, and monitoring out of the box.
```

##### Cloud Spanner
```
a) Cloud Spanner is a managed relational database that is globally available, supports global strong consistency, and has 99.999% SLA.
b) Compared to Cloud SQL, Spanner provides higher SLA, better concurrent request handling, more scalability, and more storage at the cost of higher prices.  
```

##### Cloud Firestore
```
a) Cloud Firestore is a managed NoSQL document database and is an upgraded version of Cloud Datastore.
b) Firestore supports "Native" mode as well as the legacy "Datastore" mode without some of the previous limitations of Datastore. 
c) Firestore supports multi-region replication, strong consistency, secondary indices, and is billed per operation. 
```

##### Bigtable
```
a) Bigtable is a managed NoSQL wide column database based on Apache HBase and is in use by Gmail and Google Maps.
b) Bigtable indexes on row keys, is typically used as a mass scale database for IoT/time series, and is billed per provisioned node.
   Additional nodes can be provisioned to increase queries per second.
c) Compared to Firestore, Bigtable supports faster reads, range scans, and higher write throughput through eventual consistency.
```

### Data Services
---
##### BigQuery
```
a) BigQuery is a managed serverless data warehouse that helps analyze mass scale data with built-in features like ML and geospatial analysis.
   BigQuery abstracts away the computing and storage infrastructure so users can focus on data analysis.
b) BigQuery allows for data to be ingested in batches or in streams which can then be queried through SQL.
c) BigQuery provides the flexibility to analyze data not only from BigQuery storages but for other storage options where data might reside.
d) BigQuery comes with two forms of payment models, one that involves paying for storage and queries, and another involves flat monthly rates.  
```

##### Cloud Pub/Sub
```
a) Cloud Pub/Sub is a managed asynchronous messaging service and is serverless with global availability.
b) Cloud Pub/Sub allows messages to be pushed through webhooks.
```

##### Dataproc
```
Dataproc is a managed service based off of Apache Spark and Hadoop to bring up clusters for data processing.
These clusters can be brought up on-demand and are billed per second.
```

##### Dataflow
```
Dataflow is a managed service based off of Apache Beam to process data in batches or streams and is completely serverless.
```

##### Datalab
```
Datalab is a service based off of Jupyter to explore, analyze, and transform data.
```

##### Dataprep
```
Dataprep is a managed service to perform data visualization and exploration without the need of prior coding experience.
```

##### Data Studio
```
Data Studio is a service to visualize data as reports and dashboards.
```

##### Cloud Composer
```
Cloud Composer is a managed service based off of Apache Airflow that helps to create and orchestrate big data pipelines.
```

### Networking Services
---
##### Virtual Private Cloud (VPC)
```
a) Google VPC is the foundation of GCP networking and each project by default comes with a default VPC network.
b) Google VPC can contain one or more regional subnets and allows for VMs of the same VPC to communicate with one another.
```

##### Load Balancer
```
Google Load Balancer is Google's load balancing solution for GCE, GKE, and GAE.
Google Load Balancer allows users to choose from HTTP, SSL proxy, TCP proxy, network, and internal TCP/UDP load balancer.
```

##### Cloud Router
```
a) Cloud Router is a means to dynamically exchange route and next hop information between GCP VPC and on-prem network using BGP.
b) Cloud Router is used with Cloud VPNs or Interconnect to achieve dynamic route propagation to GCP resources.

More information here: https://www.youtube.com/watch?v=K_xb_j46YOk
```

##### Cloud VPN
```
a) Cloud VPN is a means for establishing a connection between GCP VPC and on-prem network through an IPsec tunnel.
b) Cloud VPN entails traffic will traverse public networks, but the service is cheap, easy to setup, and packets are encrypted.
```

##### Cloud Interconnect
```
a) Cloud Interconnect is a means for establishing low latency connectivity to Google's network without traversing the Internet.
   Cloud Interconnect achieves this by creating a cross connect between an on-prem router and Google network at a co-location facility.
   Cloud Router can then be used to dynamically advertise efficient routes to resources in the GCP VPC network.
b) Cloud Interconnect supports two different modes:
     * Dedicated Interconnect :  Connects with Google directly for the highest guaranteed performance.
     * Partner Interconnect   :  Connects with partners of Google if Dedicated Interconnect is not possible. 
c) Cloud Interconnect gives the on-prem network direct access to internal IP spaces in the VPC without need for NATs or VPN tunneling.
```

##### Cloud DNS
```
Cloud DNS is Google's managed DNS service with 100% SLA that can also host private zones accessible only to a user's GCP network.
```

##### Cloud CDN
```
Cloud CDN is Google's managed CDN service that enables caching of HTTP load balanced content (e.g. GCS bucket objects) and mitigation of DDoS.
```

##### Cloud NAT
```
Cloud NAT is a regional service that allows VMs without external IPs to communicate to the outside network. 
```

##### GCP Firewall
```
Firewall rules are created per VPC and rules can be applied based on IP addresses, tags, and service accounts.
```

##### Identity Aware Proxy (IAP)
```
a) IAP inforces access control policies to control which users or groups have access to certain applications and GCP resources.
b) IAP grants user permissions through IAM roles and performs authentication and authorization when a user tries to access an IAP secured resource.
   This ensures access control even without the use of a VPN.
```

##### Cloud Armor
```
Cloud Armor is a service designed to protect against DDoS by integrating with HTTP load balancers and blocking traffic based on IP address ranges. 
```

### Machine Learning Services
---
##### Vertex AI
```
Vertex AI is a managed ML platform that provides a unified way to maintain all the steps in an ML workflow, including:
  * Ingest, label, prepare, and transform data into trainable datasets.
  * Start Jupyter instances with specified hardware for data exploration.
  * Train models from datasets through AutoML or containers containing custom training logic.
  * Assess and evaluate models through services like Explainable AI.
  * Deploy trained models to an endpoint with the necessary infrastructure and hardware required.
```

##### AutoML
```
AutoML is a service to allow developers to train ML models without extensive knowledge about data science by simply passing in labeled data.
```

##### Pretrained APIs
```
GCP offers several different pretrained models for users to leverage.
These include text-to-speech, video intelligence, vision, natural language, and translation.
```

##### Dialogflow
```
Dialogflow is a service to build conversational applications that can respond through text or speech by training from customer text and audio inputs.
```

### Identity Services
---
##### Identity and Access Management (IAM)
```
IAM allows GCP administrators to grant granular permissions for GCP resources to users or groups of users.
```

##### Cloud Identity
```
a) Cloud Identity is an identity-as-a-service offering that sits outside of GCP but can be easily integrated to GCP.
b) Cloud Identity allows for a centralized way to create, assign IAM, and manage user and group accounts.
c) Cloud Identity enables multi-factor authentication, SSO, and various login methods (e.g. SAML, OIDC) for applications. 
```