# ServiceTier

Specifies the latency tier to use for processing the request. This parameter is relevant for customers subscribed to the scale tier service:

  - If set to 'auto', and the Project is Scale tier enabled, the system will utilize scale tier credits until they are exhausted.
  - If set to 'auto', and the Project is not Scale tier enabled, the request will be processed using the default service tier with a lower uptime SLA and no latency guarentee.
  - If set to 'default', the request will be processed using the default service tier with a lower uptime SLA and no latency guarentee.
  - When not set, the default behavior is 'auto'.

  When this parameter is set, the response body will include the `service_tier` utilized.



## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ServiceTierAuto`    | auto                 |
| `ServiceTierDefault` | default              |