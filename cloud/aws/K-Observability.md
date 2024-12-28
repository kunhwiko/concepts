### CloudWatch
---
##### CloudWatch Metrics
```
a) CloudWatch provides metrics for every service in AWS and allows for creating custom metrics. Metrics can be separated
   into different namespaces to provide isolation per service.
b) CloudWatch metrics can be streamed to a different destination, such as Kinesis Data Firehose or Splunk.
c) Alarms can be set to trigger notifications on a metric. Alarms can also be composite based on the state of multiple
   other metrics. 
```

##### CloudWatch Logs
```
a) CloudWatch logs are organized into log groups and log streams. Log streams are a sequence of log events that share the 
   same source, while log groups are a collection of log streams that are typically divided per application.
b) CloudWatch logs can be sent to a different destination such as Kinesis Data Firehose, S3, Lambda, or Splunk through 
   batch processing or streaming. Subscription filters can be set to filter logs before they are sent to the destination.
c) CloudWatch logs from different accounts can be shipped and aggregated to another account for centralized logging.
```

##### CloudWatch Agent
```
Agents need to be installed on EC2 instances to collect logs and granular system-level metrics (e.g. I/O, RAM, netstat, 
processes) beyond out-of-the-box basic metrics (e.g. CPU, high level network usage). The instance must also have an IAM 
role to allow the agent to push logs to CloudWatch.
```

##### CloudWatch Insights
```
a) Container Insights can be used to monitor containerized applications on ECS, EKS, and Fargate. CloudWatch agents need
   to be deployed as containers on the instances to collect and aggregate metrics and logs.
b) Lambda Insights can be used to monitor Lambda functions and identify performance issues (e.g. cold starts).
c) Contributor Insights can be used to identify the heaviest network users through VPC flow logs and source IPs.
```

### EventBridge
---
##### EventBridge Basics
```
EventBridge can be used to schedule cron jobs to trigger scripts (e.g. Lambda functions) periodically or react to event 
rules (e.g. send SNS notification or trigger Lambda function on instance termination). Events sent to EventBridge can be 
filtered (e.g. certain S3 buckets).
```

##### Event Bus
```
a) Event buses are routers that receieve events and deliver them to different targets based on event rules. The default 
   event bus is used for events from AWS services, partner event buses can be created to integrate with third-party 
   services (e.g. Datadog), and custom event buses can be created to route events from custom applications.
b) Event buses can be accessed by other accounts for centralization given the necessary resource policy.
```