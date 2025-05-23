# 🎯 KubeTarget

KubeTarget is a Kubernetes extension that enables developers and test engineers to provision virtual embedded hardware targets on demand using Kubernetes Custom Resource Definitions (CRDs). It bridges the gap between CI/CD automation and Hardware-in-the-Loop (HiL) testing, enabling scalable and declarative GitOps workflows for hardware engineers.

## 🚀 Key Features

- **Virtual Hardware CRDs**: Provision virtual hardware anywhere using familiar Kubernetes resources and GitOps.
- **Pluggable Provider Model**: Easily add multiple local and remote virtualization providers including QEMU, Corellium, and Android Virtual Devices (AVDs).
- **Secure, Multi-Tenant Deployment**: Namespace-scoped providers support different teams or credentials (e.g., different AWS accounts).
- **Jumpstarter Integration**: Run Hardware-in-the-Loop (HiL) tests directly against provisioned devices using structured test APIs.
- **Kubernetes-Native**: Built on Pods, CRDs, RBAC, and standard Kubernetes APIs.

## 🔄 What is TargetOps?

**TargetOps** is a GitOps-based methodology to manage virtual hardware targets, enabling teams to version control, automate, and declaratively manage their virtual hardware infrastructure through Kubernetes.

## 🧱 Architecture

- **`VirtualTargetDevice`**: CRD that defines a virtual device under test.
- **`VirtualTargetProvider`**: CRD that specifies a container image and runtime config for a gRPC-based provisioning provider.

## 📦 Example

Provision an STM32F4 device in QEMU and run CAN bus loopback tests with Jumpstarter:

```yaml
apiVersion: kubetarget.dev/v1alpha1
kind: VirtualTargetDevice
metadata:
  name: stm32f4
spec:
  providerRef: qemu-arm-provider
  hardwareType: stm32f4
  configuration:
    peripherals:
      can:
        count: 2
      uart:
        count: 1
    firmware:
      configMap: stm32f4-firmware
```

## Getting Started

### Prerequisites
- go version v1.23.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=ghcr.io/kubetarget/controller:latest
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don't work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=ghcr.io/kubetarget/controller:latest
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following the options to release and provide this solution to the users.

### By providing a bundle with all YAML files

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/controller:tag
```

**NOTE:** The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without its
dependencies.

2. Using the installer

Users can just run 'kubectl apply -f <URL for YAML BUNDLE>' to install
the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/controller/<tag or branch>/dist/install.yaml
```

### By providing a Helm Chart

1. Build the chart using the optional helm plugin

```sh
kubebuilder edit --plugins=helm/v1-alpha
```

2. See that a chart was generated under 'dist/chart', and users
can obtain this solution from there.

**NOTE:** If you change the project, you need to update the Helm Chart
using the same command above to sync the latest changes. Furthermore,
if you create webhooks, you need to use the above command with
the '--force' flag and manually ensure that any custom configuration
previously added to 'dist/chart/values.yaml' or 'dist/chart/manager/manager.yaml'
is manually re-applied afterwards.

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

KubeTarget is licensed under the Apache 2.0 License ([LICENSE](LICENSE) or
[https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)).