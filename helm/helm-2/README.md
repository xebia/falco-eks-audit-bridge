Falco-EKS-Audit-Bridge

[Falco-EKS-Audit-Bridge](https://github.com/xebia/falco-eks-audit-bridge) is designed to send EKS CloudWatch logs to an S3 bucket and then transfer the audit events to [Falco](https://falco.org/) for compliance checking activity of the applications.

To know more about Falco-EKS-Audit-Bridge have a look at:

- [Monitoring-AWS-EKS-Audit-Logs-with-Falco](https://xebia.com/blog/monitoring-aws-eks-audit-logs-with-falco)

## Introduction

This chart adds Falco-EKS-Audit-Bridge pods to the nodes in your cluster using a Deployment.

A Daemonset can also be used  in place of Deployment as [Falco](https://falco.org/) is deployed as Daemonset.

## Installing the Chart

To install the chart with the release name `my-release` run:

```bash
$ helm install --name my-release ./falco-eks-audit-bridge
```

After a few seconds, falco-eks-audit-bridge should be up and running.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```bash
$ helm delete my-release
```
> **Tip**: Use helm delete --purge my-release to completely remove the release from Helm internal storage

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the Falco-EKS-Audit-Bridge chart and their default values.

| Parameter                                       | Description                                                                                                        | Default                                                                                                                                   |
| ---                                             | ---                                                                                                                | ---                                                                                                                                       |
| `image.registry`                                | The image registry to pull from                                                                                    | `docker.io`                                                                                                                               |
| `image.repository`                              | The image repository to pull from                                                                                  | `falcosecurity/falco`                                                                                                                     |
| `image.tag`                                     | The image tag to pull                                                                                              | `0.17.1`                                                                                                                                  |
| `image.pullPolicy`                              | The image pull policy                                                                                              | `IfNotPresent`                                                                                                                            |
| `image.replicas`                                | How many replicas of falco-eks-audit-bridge to run                                                                  | `3`                                                                                                                                      |
| `image.namespace`                               | The namespace to be use for workloads                                                                               | `falco`                                                                                                                                  |
| `resources.requests.cpu`                        | CPU requested for being run in a node                                                                              | `100m`                                                                                                                                    |
| `resources.requests.memory`                     | Memory requested for being run in a node                                                                           | `256Mi`                                                                                                                                   |
| `resources.limits.cpu`                          | CPU limit                                                                                                          | `200m`                                                                                                                                    |
| `resources.limits.memory`                       | Memory limit                                                                                                       | `512Mi`                                                                                                                                   |
| `extraArgs`                                     | Specify additional container args                                                                                   | `[]`                                                                                                                                     |
| `podSecurityPolicy.create`                      | If true, create & use podSecurityPolicy                                                                             | `false`                                                                                                                                  |
| `rbac.create`                                   | If true, create & use RBAC resources                                                                                | `true`                                                                                                                                   |
| `rbac.clusterRole.name`                         | If rbac.create=true, clusterRole Name                                                                              | `cr`                                                                                                                                      |
| `rbac.clusterRoleBinding.name`                  | If rbac.create=true, clusterRoleBinding Name                                                                       | `crb`                                                                                                                                     |
| `serviceAccount.create`                         | Create serviceAccount                                                                                              | `true`                                                                                                                                    |
| `serviceAccount.name`                           | Use this value as serviceAccountName                                                                               | `feab`                                                                                                                                    |
| `deployment.updateStrategy.type`                | The updateStrategy for updating the deployment                                                                     | `RollingUpdate`                                                                                                                           |
| `deployment.env`                                | Extra environment variables passed to deployment pods                                                              | `{}`                                                                                                                                      |
| `proxy.httpProxy`                               | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `proxy.httpsProxy`                              | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `proxy.noProxy`                                 | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `timezone`                                      | Set the daemonset's timezone                                                                                       | ` `                                                                                                                                       |
| `feab.service.enabled`                          | If true, create falco-eks-audit-bridge service                                                                     | `true`                                                                                                                                    |
| `feab.service.listenPort`                       | Port on which Service will listen                                                                                  | `8080`                                                                                                                                    |
| `feab.service.svcType`                          | Type of Service                                                                                                    | `ClusterIP`                                                                                                                               |
| `feab.service.name`                             | Name of Service                                                                                                    | `falco-feab-service`                                                                                                                      |
| `feab.bucket`                                   | Name of bucket containting CloudWatch Logs                                                                         | ` `                                                                                                                                       |
| `feab.falco_ep`                                 | Falco End-Point                                                                                                    | ` `                                                                                                                                       |
| `feab.aws_region`                               | AWS Region where EKS is Installed                                                                                  | `eu-west-1`                                                                                                                               |





| `tolerations`                                   | The tolerations for scheduling                                                                                     | `node-role.kubernetes.io/master:NoSchedule`                                                                                               |



| `fakeEventGenerator.enabled`                    | Run falco-event-generator for sample events                                                                        | `false`                                                                                                                                   |





| `priorityClassName`                             | Set the daemonset's priorityClassName                                                                              | ` `                                                                                                                                       |
| `ebpf.enabled`                                  | Enable eBPF support for Falco instead of `falco-probe` kernel module                                               | `false`                                                                                                                                   |
| `ebpf.settings.hostNetwork`                     | Needed to enable eBPF JIT at runtime for performance reasons                                                       | `true`                                                                                                                                    |
| `ebpf.settings.mountEtcVolume`                  | Needed to detect which kernel version are running in Google COS                                                    | `true`                                                                                                                                    |
| `falco.rulesFile`                               | The location of the rules files                                                                                    | `[/etc/falco/falco_rules.yaml, /etc/falco/falco_rules.local.yaml, /etc/falco/rules.available/application_rules.yaml, /etc/falco/rules.d]` |
| `falco.timeFormatISO8601`                       | Display times using ISO 8601 instead of local time zone                                                            | `false`                                                                                                                                   |
| `falco.jsonOutput`                              | Output events in json or text                                                                                      | `false`                                                                                                                                   |
| `falco.jsonIncludeOutputProperty`               | Include output property in json output                                                                             | `true`                                                                                                                                    |
| `falco.logStderr`                               | Send Falco debugging information logs to stderr                                                                    | `true`                                                                                                                                    |
| `falco.logSyslog`                               | Send Falco debugging information logs to syslog                                                                    | `true`                                                                                                                                    |
| `falco.logLevel`                                | The minimum level of Falco debugging information to include in logs                                                | `info`                                                                                                                                    |
| `falco.priority`                                | The minimum rule priority level to load and run                                                                    | `debug`                                                                                                                                   |
| `falco.bufferedOutputs`                         | Use buffered outputs to channels                                                                                   | `false`                                                                                                                                   |
| `falco.syscallEventDrops.actions`               | Actions to be taken when system calls were dropped from the circular buffer                                        | `[log, alert]`                                                                                                                            |
| `falco.syscallEventDrops.rate`                  | Rate at which log/alert messages are emitted                                                                       | `.03333`                                                                                                                                  |
| `falco.syscallEventDrops.maxBurst`              | Max burst of messages emitted                                                                                      | `10`                                                                                                                                      |
| `falco.outputs.rate`                            | Number of tokens gained per second                                                                                 | `1`                                                                                                                                       |
| `falco.outputs.maxBurst`                        | Maximum number of tokens outstanding                                                                               | `1000`                                                                                                                                    |
| `falco.syslogOutput.enabled`                    | Enable syslog output for security notifications                                                                    | `true`                                                                                                                                    |
| `falco.fileOutput.enabled`                      | Enable file output for security notifications                                                                      | `false`                                                                                                                                   |
| `falco.fileOutput.keepAlive`                    | Open file once or every time a new notification arrives                                                            | `false`                                                                                                                                   |
| `falco.fileOutput.filename`                     | The filename for logging notifications                                                                             | `./events.txt`                                                                                                                            |
| `falco.stdoutOutput.enabled`                    | Enable stdout output for security notifications                                                                    | `true`                                                                                                                                    |
| `falco.webserver.enabled`                       | Enable Falco embedded webserver to accept K8s audit events                                                         | `false`                                                                                                                                   |
| `falco.webserver.listenPort`                    | Port where Falco embedded webserver listen to connections                                                          | `8765`                                                                                                                                    |
| `falco.webserver.k8sAuditEndpoint`              | Endpoint where Falco embedded webserver accepts K8s audit events                                                   | `/k8s-audit`                                                                                                                              |
| `falco.webserver.clusterIP`                     | ClusterIP address where Falco will listen to K8s audit events. If you enable the webserver, this field is required | ` `                                                                                                                                       |
| `falco.programOutput.enabled`                   | Enable program output for security notifications                                                                   | `false`                                                                                                                                   |
| `falco.programOutput.keepAlive`                 | Start the program once or re-spawn when a notification arrives                                                     | `false`                                                                                                                                   |
| `falco.programOutput.program`                   | Command to execute for program output                                                                              | `mail -s "Falco Notification" someone@example.com`                                                                                        |
| `falco.httpOutput.enabled`                      | Enable http output for security notifications                                                                      | `false`                                                                                                                                   |
| `falco.httpOutput.url`                          | Url to notify using the http output when a notification arrives                                                    | `http://some.url`                                                                                                                         |
| `customRules`                                   | Third party rules enabled for Falco                                                                                | `{}`                                                                                                                                      |
| `integrations.gcscc.enabled`                    | Enable Google Cloud Security Command Center integration                                                            | `false`                                                                                                                                   |
| `integrations.gcscc.webhookUrl`                 | The URL where sysdig-gcscc-connector webhook is listening                                                          | `http://sysdig-gcscc-connector.default.svc.cluster.local:8080/events`                                                                     |
| `integrations.gcscc.webhookAuthenticationToken` | Token used for authentication and webhook                                                                          | `b27511f86e911f20b9e0f9c8104b4ec4`                                                                                                        |
| `integrations.natsOutput.enabled`               | Enable NATS Output integration                                                                                     | `false`                                                                                                                                   |
| `integrations.natsOutput.natsUrl`               | The NATS' URL where Falco is going to publish security alerts                                                      | `nats://nats.nats-io.svc.cluster.local:4222`                                                                                              |
| `integrations.pubsubOutput.credentialsData`     | Contents retrieved from `cat $HOME/.config/gcloud/legacy_credentials/<email>/adc.json                              | jq -c .`                                                                                                                                  | ` ` |
| `integrations.pubsubOutput.enabled`             | Enable GCloud PubSub Output Integration                                                                            | `false`                                                                                                                                   |
| `integrations.pubsubOutput.projectID`           | GCloud Project ID where the Pub/Sub will be created                                                                | ` `                                                                                                                                       |
| `integrations.snsOutput.enabled`                | Enable Amazon SNS Output integration                                                                               | `false`                                                                                                                                   |
| `integrations.snsOutput.topic`                  | The SNS topic where Falco is going to publish security alerts                                                      | ` `                                                                                                                                       |
| `integrations.snsOutput.aws_access_key_id`      | The AWS Access Key Id credentials for access to SNS n                                                              | ` `                                                                                                                                       |
| `integrations.snsOutput.aws_secret_access_key`  | The AWS Secret Access Key credential to access to SNS                                                              | ` `                                                                                                                                       |
| `integrations.snsOutput.aws_default_region`     | The AWS region where SNS is deployed                                                                               | ` `                                                                                                                                       |
| `nodeSelector`                                  | The node selection constraint                                                                                      | ` `                                                                                                                                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```bash
$ helm install --name my-release --set falco.jsonOutput=true stable/falco
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install --name my-release -f values.yaml ./falco-eks-audit-bridge
```

> **Tip**: You can use the default [values.yaml](values.yaml)