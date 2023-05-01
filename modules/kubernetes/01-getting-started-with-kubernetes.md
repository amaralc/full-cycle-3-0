# Getting Started with Kubernetes

[<- Index](../../README.md)

</br>

## 01 Introduction to Kubernetes

```
"Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications."
```

Source: https://kubernetes.io/

---

- It is a containers management solution;
- you can ask Kubernetes to run your containers for you;
- Created by Google, as an evolution of Borg:

  - Borg;
  - Omega;
  - Kubernetes;

- It is made available through a set of APIs;
- We normally access that API using a CLI: **kubctl**;
- Everything is based in state. You configure the state of every object;
- It works as a cluster (a set of machines);
- One of the cluster nodes is the Kubernetes Master, which controls all the coordination process;
- Kubernetes Master has some services running:
  - Kube-apiserver;
  - Kube-controller-manager;
  - Kube-scheduler;
- Other nodes have:
  - Kubelet;
  - Kubeproxy;

These other nodes communicate with Kubernetes master to understand how to communicate with other nodes through the proxies;

## Cluster

- Cluster: set of virtual machines (nodes)
- Each machine has a portion of vCPU and Memory;

### Pods

- Pods: units that have provisioned containers;
- A pod represent running processes in a cluster;
- Generally you have one container per pod, but you can also have multiple containers running in a single pod;

### Deployment

- The deployment is another Kubernetes object;
- Its goal is to provision pods;
- In order to do that, the deployment must know how many replicas each pod will have, which is defined by the number of **ReplicaSets**;
- We can define the **ReplicaSets** in the deployment object;

Example:

- Lets suppose we need 3 replicas for the back-end and 2 for the front-end:
  - Node 1:
    - Back-end: 3 replicas
    - Front-end: 2 replicas
- Then we need to add another replica to the front-end. We than change the number of ReplicaSets in one file and kubernetes create another front-end replica for us:
  - Node 1:
    - Back-end: 3 replicas
    - Front-end: 3 replicas
- Then for the high demand, we need another back-end replica, but the available memory and CPU for that node is not enough. Then nothing will happen until we create another node. Once we do, kubernetes manage to set a new pod in that node:
  - Node 1:
    - Back-end: 3 replicas
    - Front-end: 3 replicas
  - Node 2:
    - Back-end: 1 replica
- Then the first node shuts off for some reason. In that case, Kubernetes manage to distribute the remaining pods in the existing node, up until the memory and cpu capacities of that node;
  - Node 2:
    - Back-end: 4 replicas
    - Front-end: 2 replicas

### Services

- The aggregation dynamics of a set of pods to implement visibility policies;
