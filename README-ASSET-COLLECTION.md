# Example Plugin

- [Overview](#overview)
- [Examples](#examples)
- [Configuration](#configuration)
  - [Asset Definitions](#asset-definitions)
  - [Check Definitions](#check-definitions)
- [Metrics](#metrics)
  - [Service A Metrics](#service-a-metrics)
  - [Service B Metrics](#service-b-metrics)
- [Service Checks](#service-checks)
  - [Service A Service Checks](#service-a-service-checks)

## Overview 

Example Plugin is a Sensu Go plugin for monitoring [Example Service][example].
The more high-level description we can provide here, the better!  

This plugin provides the following features: 

- [Metrics collection](#metrics) 
- [Service checks](#service-checks)

[example]: http://link/to/example/website 

## Examples 

Something something usage examples... 

## Configuration 

### Asset Definitions 

```
{
  "type": "Asset",
  "spec": {
    "namespace": "default",
    "name": "example-plugin",
    "url": "http://bonsai.sensu.io/calebhailey/example-plugin",
    "sha512": "cd9f3b756e110d5d68af9b9235eb9e38efaf7ea118ce0d6ce2ff1a4049d205e417495c962585d160a534088eb2c1390672116a75c7a667d49cd5fbcf70d102c9",
    "filters": [
      "System.Platform=='debian'",
      "System.OS=='linux'", 
      "System.Arch=='amd64'"
    ]
  }
}
```

### Check Definitions

```
{
  "type": "Check",
  "spec": {
    "organization": "default",
    "environment": "default",
    "name": "example-check",
    "command": "example-plugin.sh",
    "runtime_assets": ["example-plugin"],
    "publish": true,
    "interval": 10,
    "subscriptions": ["example"]
  }
}
```

## Metrics 

### Service A Metrics 

- **name**: `example.response_time`  
  **description**: Request response time  
  **unit**: milliseconds  

- **name**: `example.response_size`  
  **description**: The size of the response payload  
  **unit**: bytes  

### Service B Metrics 

- **name**: `example.error_rate`  
  **description**: The percentage of error responses, calculated as "count of 
  error responses / total count of responses".   
  **unit**: percentile  

- **name**: `metric_name`  
  **description**: This is an example metric description. The description 
  should clearly explain what data is being measured, and provide some 
  information regarding what it should mean to the user.    
  **unit**: unit specification (e.g. if time, is the unit "minutes", "seconds",
  or "milliseconds"?).  

## Service Checks 

### Service A Service Checks

- **name**: `service_metrics`  
  **description**: The service check which collects "Service A Metrics" (see above for details).  
  **output**: metrics  

- **name**: `service_health`  
  **description**: Monitors the service `/healthz` endpoint to confirm the service is healthy.  
  **output**: status information  
