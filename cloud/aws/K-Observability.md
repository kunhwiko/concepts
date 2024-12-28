### CloudWatch
---
##### CloudWatch Metrics
```
a) CloudWatch provides metrics for every service in AWS and allows for creating custom metrics. Metrics can be separated
   into different namespaces to provide isolation per service.
b) CloudWatch metrics can be streamed to a different destination, such as Kinesis Data Firehose or Splunk.
```

##### CloudWatch Logs
```
a) CloudWatch logs are organized into log groups and log streams. Log streams are a sequence of log events that share the 
   same source, while log groups are a collection of log streams that are typically divided per application.
b) CloudWatch logs can be sent to a different destination such as Kinesis Data Firehose, S3, Lambda, or Splunk through 
   batch processing or streaming. Subscription filters can be set to filter logs before they are sent to the destination.
c) CloudWatch logs from different account can be shipped and aggregated to another account for centralized logging.
```

##### CloudWatch Agent
```
Agents need to be installed on EC2 instances to collect logs and granular system-level metrics (e.g. I/O, RAM, netstat, 
processes) beyond out-of-the-box basic metrics (e.g. CPU, high level network usage). The instance must also have an IAM 
role to allow the agent to push logs to CloudWatch.
```