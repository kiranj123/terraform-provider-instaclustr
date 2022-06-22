---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "instaclustr_cassandra_cluster_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# instaclustr_cassandra_cluster_v2 (Resource)



## Example Usage

```terraform
resource "instaclustr_cassandra_cluster_v2" "test_cluster" {
  name                    = "MyTestCluster"
  cassandra_version       = "4.0.1"
  lucene_enabled          = true
  password_and_user_auth  = true
  pci_compliance_mode     = false
  private_network_cluster = false
  sla_tier                = "PRODUCTION"

  spark {
    version = "3.0.1"
  }

  data_centre {
    name                               = "MyTestDataCentre"
    client_to_cluster_encryption       = true
    cloud_provider                     = "AWS_VPC"
    continuous_backup                  = true
    node_size                          = "m5l-250-v2"
    number_of_nodes                    = 2
    replication_factor                 = 2
    private_ip_broadcast_for_discovery = true
    network                            = "10.0.0.0/16"
    region                             = "US_EAST_1"
  }

  two_factor_delete {
    confirmation_email = "test@email.com"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cassandra_version** (String) Version of Cassandra to run on the cluster.
- **data_centre** (Block List, Min: 1) List of data centre settings. (see [below for nested schema](#nestedblock--data_centre))
- **lucene_enabled** (Boolean) Adds Apache Lucene to the Cassandra cluster.
- **name** (String) Name of the cluster.
- **password_and_user_auth** (Boolean) Enables Password Authentication and User Authorization.
- **pci_compliance_mode** (Boolean) Creates a PCI compliant cluster, see [PCI Compliance](https://www.instaclustr.com/support/documentation/useful-information/pci-compliance/).
- **private_network_cluster** (Boolean) Creates the cluster with private network only, see [Private Network Clusters](https://www.instaclustr.com/support/documentation/useful-information/private-network-clusters/).
- **sla_tier** (String) SLA Tier of the cluster. Non-production clusters may receive lower priority support and reduced SLAs. Production tier is not available when using Developer class nodes. See [SLA Tier](https://www.instaclustr.com/support/documentation/useful-information/sla-tier/) for more information.

### Optional

- **current_cluster_operation_status** (String) Indicates if the cluster is currently performing any restructuring operation such as being created or resized
- **default_user_password** (String, Sensitive) Password of the default user created for the Cassandra cluster.
- **id** (String) The ID of this resource.
- **spark** (Block List) Adds the specified version of Apache Spark to the Cassandra cluster. (see [below for nested schema](#nestedblock--spark))
- **status** (String) Status of the cluster.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **two_factor_delete** (Block List) (see [below for nested schema](#nestedblock--two_factor_delete))

<a id="nestedblock--data_centre"></a>
### Nested Schema for `data_centre`

Required:

- **client_to_cluster_encryption** (Boolean) Enables Client ⇄ Node Encryption.
- **cloud_provider** (String) Name of the cloud provider service in which the Data Centre will be provisioned.
- **continuous_backup** (Boolean) Enables commitlog backups and increases the frequency of the default snapshot backups.
- **name** (String) A logical name for the data centre within a cluster. These names must be unique in the cluster.
- **network** (String) The private network address block for the Data Centre specified using CIDR address notation. The network must have a prefix length between `/12` and `/22` and must be part of a private address space.
- **node_size** (String) Size of the nodes provisioned in the Data Centre.
- **number_of_nodes** (Number) Total number of nodes in the Data Centre. Must be a multiple of `replicationFactor`.
- **private_ip_broadcast_for_discovery** (Boolean) Enables broadcast of private IPs for auto-discovery.
- **region** (String) Region of the Data Centre.
- **replication_factor** (Number) Number of racks to use when allocating nodes.

Optional:

- **aws_settings** (Block List) AWS specific settings for the Data Centre. Cannot be provided with `gcpSettings` and `azureSettings`. (see [below for nested schema](#nestedblock--data_centre--aws_settings))
- **azure_settings** (Block List) Azure specific settings for the Data Centre. Cannot be provided with `gcpSettings` and `awsSettings`. (see [below for nested schema](#nestedblock--data_centre--azure_settings))
- **gcp_settings** (Block List) GCP specific settings for the Data Centre. Cannot be provided with `awsSettings` and `azureSettings`. (see [below for nested schema](#nestedblock--data_centre--gcp_settings))
- **id** (String) ID of the Cluster Data Centre.
- **nodes** (Block List) (see [below for nested schema](#nestedblock--data_centre--nodes))
- **provider_account_name** (String) For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.
- **status** (String) Status of the Data Centre.
- **tag** (Block List) List of tags to apply to the Data Centre. Tags are metadata labels which  allow you to identify, categorize and filter clusters. This can be useful for grouping together clusters into applications, environments, or any category that you require. (see [below for nested schema](#nestedblock--data_centre--tag))

<a id="nestedblock--data_centre--aws_settings"></a>
### Nested Schema for `data_centre.aws_settings`

Optional:

- **custom_virtual_network_id** (String) VPC ID into which the Data Centre will be provisioned. The Data Centre's network allocation must match the IPv4 CIDR block of the specified VPC.
- **ebs_encryption_key** (String) ID of a KMS encryption key to encrypt data on nodes. KMS encryption key must be set in Cluster Resources through the Instaclustr Console before provisioning an encrypted Data Centre.


<a id="nestedblock--data_centre--azure_settings"></a>
### Nested Schema for `data_centre.azure_settings`

Optional:

- **resource_group** (String) The name of the Azure Resource Group into which the Data Centre will be provisioned.


<a id="nestedblock--data_centre--gcp_settings"></a>
### Nested Schema for `data_centre.gcp_settings`

Optional:

- **custom_virtual_network_id** (String) Network name or a relative Network or Subnetwork URI e.g. projects/my-project/regions/us-central1/subnetworks/my-subnet. The Data Centre's network allocation must match the IPv4 CIDR block of the specified subnet.


<a id="nestedblock--data_centre--nodes"></a>
### Nested Schema for `data_centre.nodes`

Optional:

- **id** (String) ID of the node.
- **node_roles** (List of String) The roles or purposes of the node. Useful for filtering for nodes that have a specific role.
- **node_size** (String) Size of the node.
- **private_address** (String) Private IP address of the node.
- **public_address** (String) Public IP address of the node.
- **rack** (String) Rack name in which the node is located.
- **status** (String) Provisioning status of the node.


<a id="nestedblock--data_centre--tag"></a>
### Nested Schema for `data_centre.tag`

Required:

- **key** (String) Key of the tag to be added to the Data Centre.
- **value** (String) Value of the tag to be added to the Data Centre.



<a id="nestedblock--spark"></a>
### Nested Schema for `spark`

Required:

- **version** (String) Adds the specified version of Apache Spark to the Cassandra cluster.


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)
- **default** (String)
- **update** (String)


<a id="nestedblock--two_factor_delete"></a>
### Nested Schema for `two_factor_delete`

Required:

- **confirmation_email** (String) The email address which will be contacted when the cluster is requested to be deleted.

Optional:

- **confirmation_phone_number** (String) The phone number which will be contacted when the cluster is requested to be delete.

