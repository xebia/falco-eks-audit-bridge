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
| `global.falco.enabled`                          | Dependancy chart Falco Enabling                                                                                    | `false`                                                                                                                                   |
| `image.registry`                                | The image registry to pull from                                                                                    | `docker.io`                                                                                                                               |
| `image.repository`                              | The image repository to pull from                                                                                  | `xebia/falco-eks-audit-bridge`                                                                                                            |
| `image.tag`                                     | The image tag to pull                                                                                              | `v1.0.2`                                                                                                                                  |
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
| `serviceAccount.create`                         | Create serviceAccount                                                                                              | `true`                                                                                                                                    |
| `serviceAccount.name`                           | Use this value as serviceAccountName                                                                               | `feab`                                                                                                                                    |
| `deployment.updateStrategy.type`                | The updateStrategy for updating the deployment                                                                     | `RollingUpdate`                                                                                                                           |
| `deployment.env`                                | Extra environment variables passed to deployment pods                                                              | `{}`                                                                                                                                      |
| `proxy.httpProxy`                               | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `proxy.httpsProxy`                              | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `proxy.noProxy`                                 | Set the Proxy server if is behind a firewall                                                                       | ` `                                                                                                                                       |
| `timezone`                                      | Set the deployment's timezone                                                                                      | ` `                                                                                                                                       |
| `feab.service.enabled`                          | If true, create falco-eks-audit-bridge service                                                                     | `true`                                                                                                                                    |
| `feab.service.listenPort`                       | Port on which Service will listen                                                                                  | `8080`                                                                                                                                    |
| `feab.service.svcType`                          | Type of Service                                                                                                    | `ClusterIP`                                                                                                                               |
| `feab.service.name`                             | Name of Service                                                                                                    | `falco-feab-service`                                                                                                                      |
| `feab.bucket`                                   | Name of bucket containting CloudWatch Logs                                                                         | ` `                                                                                                                                       |
| `feab.falco_ep`                                 | Falco End-Point                                                                                                    | ` `                                                                                                                                       |
| `feab.aws_region`                               | AWS Region where EKS is Installed                                                                                  | `eu-west-1`                                                                                                                               |
| `tolerations`                                   | The tolerations for scheduling                                                                                     | `node-role.kubernetes.io/master:NoSchedule`                                                                                               |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```bash
$ helm install --name my-release --set serviceAccount.create=true ./falco-eks-audit-bridge
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install --name my-release -f values.yaml ./falco-eks-audit-bridge
```

> **Tip**: You can use the default [values.yaml](values.yaml)