# Terraform Azure Web Server with Terratest

This project uses Terraform to deploy a simple Ubuntu web server on Microsoft Azure and uses Terratest to verify the deployed infrastructure.

## Project Overview

The Terraform configuration creates the following Azure resources:

- Resource Group
- Virtual Network
- Subnet
- Network Security Group
- Public IP Address
- Network Interface
- Ubuntu Linux Virtual Machine
- Apache Web Server installed with cloud-init

The infrastructure allows:

- SSH access through port 22
- HTTP access through port 80

## Infrastructure Testing

Terratest is used to automate the deployment and testing process.

The test performs the following steps:

1. Runs `terraform init`
2. Runs `terraform apply`
3. Confirms that the virtual machine exists
4. Confirms that the network interface exists
5. Confirms that the NIC is connected to the virtual machine
6. Confirms that the VM uses Ubuntu 22.04 LTS
7. Runs `terraform destroy`

## Technologies Used

- Terraform
- Terratest
- Go
- Microsoft Azure
- Azure Resource Manager (AzureRM)
- Azure CLI
- Ubuntu 22.04 LTS
- Apache HTTP Server

## Project Structure

```text
.
├── main.tf
├── variables.tf
├── outputs.tf
├── providers.tf
├── init.sh
└── test
    ├── azure_webserver_test.go
    ├── go.mod
    └── go.sum
```

## Run the Terraform Deployment Manually

### Initialize Terraform:

```bash
terraform init
```

### Review the deployment plan:

```bash
terraform plan
```

### Deploy the infrastructure:

```bash
terraform apply
```

### Destroy the infrastructure:

```bash
terraform destroy
```
## Run the Terratest Integration Test

### Go to the test folder:

```bash
cd test
```

### Run the test:

```bash
go test -v azure_webserver_test.go
```

### Terratest will deploy the infrastructure, verify the Azure resources, and destroy the infrastructure after testing.

## Test Result

The Terratest integration test successfully verified that:

- The Azure Linux virtual machine was created successfully.
- The network interface exists.
- The network interface is attached to the virtual machine.
- The virtual machine is running Ubuntu 22.04 LTS.
- All Azure resources were automatically destroyed after the test completed.