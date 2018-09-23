# Self-Service Monitoring

Monitoring plays a crucial role in a microservice architecture. Restricting the
management and configuration of the monitoring stack to the operations team
results in workflow bottlenecks. Instead one could provide a self-service
monitoring platform, enabling each team to easily setup monitoring for their
applications and customize it to their needs. This gives each team the ability
to deeply introspect their application, benchmark new features and alert on
failures on their own.

This workshop will show hands-on how to achieve the above on **Kubernetes** via
the Prometheus Operator leveraging the **Prometheus** systems and service
monitoring project. By the end of this talk, one will be ready to try
self-service monitoring in their own Kubernetes infrastructure.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Self-Service Monitoring](#self-service-monitoring)
    - [Introduction to monitoring components](#introduction-to-monitoring-components)
    - [Getting started](#getting-started)
        - [Prerequisites](#prerequisites)
        - [Setup environment](#setup-environment)
    - [1. Explore default black box monitoring](#1-explore-default-black-box-monitoring)

<!-- markdown-toc end -->


## Introduction to monitoring components

- [Prometheus](https://github.com/prometheus/prometheus)

  Monitoring system and time-series database.

- [Alertmanager](https://github.com/prometheus/alertmanager)

  Deduplicate, group and route alerts send by Prometheus.

- [cAdvisor](https://github.com/google/cadvisor/)

  Analyzes resource usage and performance characteristics of running containers.

- [kube-state-metrics](https://github.com/kubernetes/kube-state-metrics/)

  Add-on agent to generate and expose cluster-level metrics.

- [Prometheus Operator](https://github.com/coreos/prometheus-operator)

  Create, configure and manage Prometheus and Alertmanager clusters on
  Kubernetes.

- [kube-prometheus](https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus)

  Provide end-to-end cluster Kubernetes cluster monitoring with the Prometheus
  Operator.


## Getting started


### Prerequisites

- [Minikube](https://github.com/kubernetes/minikube)

- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)


### Setup environment

1. Start Kubernetes cluster

    `minikube start --kubernetes-version=v1.10.0 --memory=4096 && minikube addons enable ingress`
  
2. Start *kube-prometheus* stack

    `kubectl apply -f kube-prometheus/manifests`

3. Create ingress configuration

    `kubectl create -f ingress.yaml`
  
  
## 1. Explore default black box monitoring

1. Deploy *sample-app* version 1.0.0

    `kubectl apply -f sample-app/manifests`

2. Expose Prometheus UI

    `kubectl port-forward -n monitoring prometheus-k8s-0 9090`
  
3. Go to `localhost:9090` in your browser

4. Get *sample-app* pod information

    `kube_pod_info{pod=~"sample-app.*"}`
  
5. Get *sample-app* CPU usage

    `container_cpu_usage_seconds_total{container_name="sample-app"}`

6. Get *sample-app* memory usage

    `container_memory_usage_bytes{container_name="sample-app"}`

7. Expose *sample-app* service

    `kubectl port-forward svc/sample-app 8080`

8. Query expensive `/hello-universe` endpoint

    `curl localhost:8080/hello-universe`

9. Compare CPU usage (see 4)
