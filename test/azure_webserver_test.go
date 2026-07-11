package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "50e7699c-94b9-43a7-a21c-34d7a59e2715"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Override the default Terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "duan0029",
		},
	}

	// Destroy the resources after the test finishes
	defer terraform.Destroy(t, terraformOptions)

	// Run terraform init and terraform apply
	terraform.InitAndApply(t, terraformOptions)

	// Read Terraform output values
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	// Test 1: Confirm the VM exists
	vmExists := azure.VirtualMachineExists(
		t,
		vmName,
		resourceGroupName,
		subscriptionID,
	)

	assert.True(t, vmExists, "The virtual machine should exist")

	// Test 2: Confirm the NIC exists
	nicExists := azure.NetworkInterfaceExists(
		t,
		nicName,
		resourceGroupName,
		subscriptionID,
	)

	assert.True(t, nicExists, "The network interface should exist")

	// Test 3: Confirm the NIC is connected to the VM
	vmNICs := azure.GetVirtualMachineNics(
		t,
		vmName,
		resourceGroupName,
		subscriptionID,
	)

	assert.Contains(
		t,
		vmNICs,
		nicName,
		"The NIC should be connected to the virtual machine",
	)

	// Test 4: Confirm the VM uses the expected Ubuntu image
	vmImage := azure.GetVirtualMachineImage(
		t,
		vmName,
		resourceGroupName,
		subscriptionID,
	)

	assert.Equal(t, "Canonical", vmImage.Publisher)
	assert.Equal(t, "0001-com-ubuntu-server-jammy", vmImage.Offer)
	assert.Equal(t, "22_04-lts-gen2", vmImage.SKU)
}