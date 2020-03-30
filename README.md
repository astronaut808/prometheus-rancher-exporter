# prometheus-rancher-exporter

Exposes the health of Stacks / Services, Hosts, information about resources used by hosts from the Rancher API, to a Prometheus compatible endpoint.

Forked from `infinityworks/prometheus-rancher-exporter` with metrics about resources used by hosts (cpu count, memory information, mountpoint information) and labels (environment id, environment name).


## Description

The application requires at a minimum, the URL of the Rancher API. If you have authentication enabled on your Rancher server, the application will require a `RANCHER_ACCESS_KEY` and a `RANCHER_SECRET_KEY` providing.


**Required**
* `CATTLE_URL` // Either provisioned through labels, or set by the user. Should be in a format similar to `http://<YOUR_IP>:8080/v2-beta`.

**Optional**
* `CATTLE_ACCESS_KEY`   // Rancher API access Key, if supplied this will be used when authentication is enabled.
* `CATTLE_SECRET_KEY`   // Rancher API secret Key, if supplied this will be used when authentication is enabled.
* `METRICS_PATH`        // Path under which to expose metrics.
* `LISTEN_ADDRESS`      // Port on which to expose metrics.
* `HIDE_SYS`            // If set to `true` then this hides any of Ranchers internal system services from being shown. *If used, ensure `false` is encapsulated with quotes e.g. `HIDE_SYS="false"`.
* `LABELS_FILTER`       // Optional regular expression for filtering service and host labels, defaults to `^io.prometheus`.
* `LOG_LEVEL`           // Optional - Set the logging level, defaults to Info.
* `API_LIMIT`           // Optional - Rancher API resource limit (default: 100)

## Metrics

```
rancher_environment_state{id="test-id",name="test-cattle",state="active"} 1
rancher_host_agent_state{labels="",name="test-agent",state="active"} 1
rancher_host_state{labels="",name="test-agent",state="active"} 1
rancher_host_cpu_count{environment_id="test-id",environment_name="test-cattle",labels="",name="test-app-1"} 4
rancher_host_mem_free{environment_id="test-id",environment_name="test-cattle",labels="",name="test-app-1"} 1604
rancher_host_mem_total{environment_id="test-id",environment_name="test-cattle",labels="",name="test-app-1"} 7983
rancher_host_mountpoint_used{environment_id="test-id",environment_name="test-cattle",labels="",mountpoint="/dev/sda1",name="test-app-1"} 8640
rancher_host_mountpoint_total{environment_id="test-id",environment_name="test-cattle",labels="",mountpoint="/dev/sda1",name="test-app-1"} 18538
rancher_service_health_status{health_state="healthy",labels="",name="test-service-1",stack_name="test-stack-1"} 0
rancher_service_scale{labels="",name="test-service-1",stack_name="test-stack-1"} 1
rancher_service_state{labels="",name="test-service-1",stack_name="test-stack-1",state="active"} 1
rancher_stack_state{name="test-stack-1",state="active",system="false"} 1
```
