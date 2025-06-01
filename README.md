# 🎯 KubeTarget

KubeTarget is a Kubernetes extension that enables developers and test engineers to provision **virtual embedded hardware targets** on demand using Kubernetes Custom Resource Definitions (CRDs). It bridges the gap between CI/CD automation and Hardware-in-the-Loop (HiL) testing, enabling scalable and declarative GitOps workflows for embedded systems.

## 🚀 Key Features

- **Virtual Hardware**: Provision virtual embedded devices anywhere using Kubernetes resources.
- **Multiple Backends**: Support multiple virtualization backends such as QEMU, Corellium, AVD, and AWS Graviton with pluggable CRDs.
- **Unified Instance Model**: All provisioned targets are tracked using the `VirtualTargetInstance` CRD, regardless of backend.
- **Sidecar Injection**: Run test agents as sidecars directly inside the virtual target pod.
- **Autoscaling and TTL**: Ephemeral targets can be scaled up/down dynamically based on demand and automatically cleaned up after use.
- **Multi-Tenant Safe**: Namespace-scoped provider instances allow secure separation between teams or credentials.
- **Kubernetes Native**: Built on Pods, CRDs, RBAC, and native Kubernetes constructs.

## 🔁 How Does KubeTarget Compare to KubeVirt?

While both KubeTarget and KubeVirt enable virtualization within Kubernetes, they serve very different purposes:

- **KubeTarget** is designed for ephemeral, on-demand virtual hardware—ideal for embedded systems, HiL testing, and CI pipelines. It supports multiple backend providers (QEMU, Corellium, AVD, etc.) and integrates natively with tools like Jumpstarter for interface-level testing.

- **KubeVirt** focuses on long-lived VMs for enterprise applications using libvirt/KVM, with features like persistent storage, VM migration, and traditional server virtualization.

### Key Differences

- KubeTarget uses multiple lightweight, pluggable backends and CRDs, while KubeVirt relies on the more robust libvirt runtime.
- KubeTarget emphasizes sidecar support, TTL-based cleanup, and GitOps automation, making it ideal for test automation.
- KubeVirt is better suited for stateful VM workloads requiring traditional virtualization features.

### When to use KubeTarget?

Use KubeTarget when you're simulating or testing virtual devices. Use KubeVirt when you're running virtualized applications.

## 🔄 What is TargetOps?

**TargetOps** is a GitOps-native approach to embedded systems testing. Just like DevOps brought Infrastructure as Code, TargetOps brings **Hardware as Code**—version-controlled, declarative test environments that run virtual or real hardware through Kubernetes.

## 📦 Example: QEMU STM32F4 Virtual Target

```yaml
# QEMU Target Definition
apiVersion: qemu.kubetarget.dev/v1alpha1
kind: QemuVirtualTarget
metadata:
  name: stm32f4-qemu
spec:
  architecture: arm
  qemuConfig:
    machine: stm32-p103
    cpu: cortex-m3
    memory: 256M
    drives:
      - name: firmware
        type: pflash
        source:
          configMap:
            name: stm32f4-firmware
    devices:
      - name: uart0
        type: serial
      - name: can0
        type: socket
```

## 🧩 Extending KubeTarget

Want to add a new provider?

1. Implement a provisioner.
2. Register it with the KubeTarget controller.
3. Define a new `XVirtualTarget` CRD in your own API group (e.g., `corellium.kubetarget.dev/v1alpha1`).

## 🤝 Contributing

We welcome contributions to expand target support, optimize autoscaling, and improve integrations like Jumpstarter and Tekton!

> Run `make help` for available targets and development tools.

## 🪪 License

KubeTarget is licensed under the Apache 2.0 License ([LICENSE](LICENSE) or [https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)).