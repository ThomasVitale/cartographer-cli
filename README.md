# Cartographer CLI

[![The Apache 2.0 license badge](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

The Cartographer CLI offers a convenient way to manage a [Cartographer](https://cartographer.sh) installation and related workflows. It's built on top of the [Apps Plugin for the Tanzu CLI](https://github.com/vmware-tanzu/apps-cli-plugin), which provides the ability to create, view, update, and delete application `Workloads` as well as list and view `ClusterSupplyChains`.

## 🚀&nbsp; Getting Started

### Prerequisites

* Kubernetes 1.24+
* Carvel [`kctrl`](https://carvel.dev/kapp-controller/docs/latest/install/#installing-kapp-controller-cli-kctrl) CLI.
* [Cert Manager](https://cert-manager.io) deployed in your Kubernetes cluster. You can install it with Carvel [`kapp`](https://carvel.dev/kapp/docs/latest/install) (recommended choice) or `kubectl`.

  ```shell
  kapp deploy -a cert-manager -y \
    -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
  ```
* [Cartographer](https://github.com/vmware-tanzu/cartographer) deployed in your Kubernetes cluster. You can install it with Carvel [`kapp`](https://carvel.dev/kapp/docs/latest/install) (recommended choice) or `kubectl`.

  ```shell
  kapp deploy -a cert-manager -y \
    -f https://github.com/vmware-tanzu/cartographer/releases/download/v0.7.0/cartographer.yaml
  ```

### Installation

Check out this code repository, then build the Cartographer CLI:

  ```shell
  go build -o carto
  ```

### Usage

Get the help information for the `carto apps` subcommand.

  ```shell
  ./carto apps
  ```

## 📙&nbsp; Documentation

Documentation, tutorials and examples for this package are available in the [docs](docs) folder.
For documentation specific to the Apps Plugin for the Tanzu CLI, check out [github.com/vmware-tanzu/apps-cli-plugin](https://github.com/vmware-tanzu/apps-cli-plugin).

## 🛡️&nbsp; Security

The security process for reporting vulnerabilities is described in [SECURITY.md](SECURITY.md).

## 🖊️&nbsp; License

This project is licensed under the **Apache License 2.0**. See [LICENSE](LICENSE) for more information.

## 🙏&nbsp; Acknowledgments

The Cartographer CLI embeds the [Apps Plugin for the Tanzu CLI](https://github.com/vmware-tanzu/apps-cli-plugin) as subcommands providing the ability to create, view, update, and delete application `Workloads` as well as list and view `ClusterSupplyChains`.
