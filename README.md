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
rancher_host_state
rancher_host_cpu_count
rancher_host_mem_free
rancher_host_mem_total
rancher_host_mountpoint_used
rancher_host_mountpoint_total
rancher_service_health_status
rancher_service_scale
rancher_service_state
rancher_stack_state
rancher_node_state
rancher_cluster_state
```
