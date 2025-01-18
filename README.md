# Microservices-Based Software Defined Network (SDN) Controller

## Overview

This project implements a microservices-based SDN controller using **Go (Golang)** and **Kubernetes**, aimed at overcoming the scalability and performance limitations of traditional monolithic SDN controllers like RYU. By leveraging modern container orchestration tools and modular design principles, the controller achieves high performance, scalability, and fault tolerance.

## Key Features

- **Microservices Architecture**: 
  - **Connection Manager**: Manages the initial connection and communication with OpenFlow switches.
  - **PacketHandler**: Processes incoming packets and identifies key properties.
  - **FlowAdd**: Handles flow addition operations and forwards instructions back to the Connection Manager.

- **Technologies Used**:
  - **Golang**: High-performance and efficient concurrent programming.
  - **Kubernetes**: Ensures reliability, scalability, and automatic resource management.
  - **Docker**: Facilitates containerized deployments.
  - **Mininet**: Simulates SDN environments for controlled testing.
  - **gRPC & Protocol Buffers**: Enables efficient communication between microservices.

- **Performance**:
  - Low latency (0.072 ms average for 100 ICMP packets).
  - Minimal memory footprint (13 MB vs. 49 MB for RYU).
  - Seamless scaling and fault recovery.

## Architecture

The system is composed of the following microservices:

1. **Connection Manager**: Establishes and maintains communication with the OpenFlow switch.
2. **PacketHandler**: Interprets and processes incoming packets.
3. **FlowAdd**: Adds flow rules based on packet properties.

Each microservice runs as a containerized application managed by Kubernetes, ensuring independent scalability and fault isolation.

## Setup and Installation

### Prerequisites

- **Docker**: Ensure Docker is installed and running.
- **Kubernetes**: Set up a Kubernetes cluster (e.g., using Minikube).
- **Helm**: For managing Kubernetes deployments.
- **Mininet**: For network simulation.

### Steps

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/your-repo/microservices-sdn-controller.git
    cd microservices-sdn-controller
    ```

2. **Build Docker Images**:
    ```bash
    docker build -t connection-manager ./connection-manager
    docker build -t packet-handler ./packet-handler
    docker build -t flow-add ./flow-add
    ```

3. **Deploy on Kubernetes**:
    ```bash
    helm install sdn-controller ./helm-chart
    ```

4. **Simulate Network Topology** (using Mininet):
    ```bash
    sudo mn --topo single,3 --mac --controller=remote,ip=<controller-ip>,port=6633 --switch ovsk
    ```

## Usage

- Monitor logs using:
  ```bash
  kubectl logs -l app=connection-manager
  kubectl logs -l app=packet-handler
  kubectl logs -l app=flow-add
  ```

- Access Kubernetes Dashboard for detailed resource utilization and scaling insights.

## Results

- **Latency**: Achieved an average of **0.072 ms** for 100 ICMP packets.
- **Memory Usage**: Only **13 MB**, compared to 49 MB for RYU.
- **Fault Tolerance**: Seamless recovery during Kubernetes replica crashes.

## Future Work

- Implementing flow deletion capabilities.
- Enhancing Quality of Service (QoS) metrics.
- Adding security features to prevent potential threats.
- Exploring machine learning-based optimizations for traffic analysis.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributors

- **Alper Sağnak**

## References

1. Arzo et al., "MSN: A Playground Framework for Design and Evaluation of MicroServices-Based SDN Controller."
2. Hölscher et al., "Evaluation of an SDN-based Microservice Architecture."
