# [Azure CosmosDB](https://azure.microsoft.com/en-us/services/cosmos-db/)

_Note: This module is PREVIEW._

_This module involves the Parent-Child Model concept in OSBA, please refer to the [Parent-Child Model doc](../parent-child-model-for-multiple-layers-services.md)._

## Services & Plans

### Service: azure-cosmosdb-sql

| Plan Name | Description |
|-----------|-------------|
| `sql-api` | Database Account and Database configured to use SQL API |

#### Behaviors

##### Provision

Provisions a new CosmosDB database account that can be accessed through any of the SQL API. The new database account is named using a new UUID. Additionally
provisions an empty Database. Ready to use with existing Azure CosmosDB libraries.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `location` | `string` | The Azure region in which to provision applicable resources. | Y |  |
| `resourceGroup` | `string` | The (new or existing) resource group with which to associate new resources. | Y |  |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority. See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `multipleWriteRegionsEnabled` | `string` | Specifies if you want  the account to write in multiple regions. Valid values are [ "enabled", "disabled"]. If set to "enabled", regions in `readRegions`  will also be writable. | N | If not specified, "disabled" will be used as the default value. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `uri` | `string` | The fully-qualified address and port of the CosmosDB database account. |
| `primaryKey` | `string` | A secret key used for connecting to the CosmosDB database. |
| `primaryConnectionString` | `string` | The full connection string, which includes the URI and primary key. |
| `databaseName` | `string` | The generated database name. |
| `documentdb_database_id` | `string` | The database name provided in a legacy key for use with Azure libraries. |
| `documentdb_host_endpoint` | `string` | The fully-qualified address and port of the CosmosDB database account provided in a legacy key for use with Azure libraries. |
| `documentdb_master_key` | `string` | A secret key used for connecting to the CosmosDB database provided in a legacy key for use with Azure libraries. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database account and database.

##### Update

Idempotently update the service instance to specified state.

| Parameter Name | Type                | Description                                                  | Required | Default Value                                                |
| -------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority. See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |


### Service: azure-cosmosdb-sql-account

| Plan Name | Description |
|-----------|-------------|
| `account` | Database Account configured to use SQL API |

#### Behaviors

##### Provision

Provisions a new CosmosDB database account that can be accessed through any of the SQL API. The new database account is named using a new UUID.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `location` | `string` | The Azure region in which to provision applicable resources. | Y |  |
| `resourceGroup` | `string` | The (new or existing) resource group with which to associate new resources. | Y |  |
| `alias` | `string` | Specifies an alias that can be used by later provision actions to create databases on this DBMS. | Y | |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `consistencyPolicy` | `object` | The consistency policy for the Cosmos DB account. | N | |
| `consistencyPolicy.defaultConsistencyLevel` | `string` | The default consistency level and configuration settings of the Cosmos DB account. - Eventual, Session, BoundedStaleness, Strong, ConsistentPrefix | Y | |
| `consistencyPolicy.boundedStaleness` | `object` | Specifies the settings when using BoundedStaleness consistency. | Y - When Using `BoundedStaleness` | |
| `consistencyPolicy.maxStalenessPrefix` | `integer` | When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | Y | |
| `consistencyPolicy.maxIntervalInSeconds` | `integer` | When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | Y | |
| `readRegions` | `array ` | Replication read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority. See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `multipleWriteRegionsEnabled` | `string` | Specifies if you want  the account to write in multiple regions. Valid values are [ "enabled", "disabled"]. If set to "enabled", regions in `readRegions`  will also be writable. | N | If not specified, "disabled" will be used as the default value. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `uri` | `string` | The fully-qualified address and port of the CosmosDB database account. |
| `primaryKey` | `string` | A secret key used for connecting to the CosmosDB database. |
| `primaryConnectionString` | `string` | The full connection string, which includes the URI and primary key. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database account.

##### Update

Idempotently update the service instance to specified state.

| Parameter Name | Type                | Description                                                  | Required | Default Value                                                |
| -------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

### Service: azure-cosmosdb-sql-database

| Plan Name | Description |
|-----------|-------------|
| `database` | Database on existing CosmosDB database account configured to use SQL API |

#### Behaviors

##### Provision

Provisions a new CosmosDB database onto an existing database account that can be accessed through any of the SQL API. The new database is named using a new UUID.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `parentAlias` | `string` | Specifies the alias of the CosmosDB database account upon which the database should be provisioned. | Y | |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `uri` | `string` | The fully-qualified address and port of the CosmosDB database account. |
| `primaryKey` | `string` | A secret key used for connecting to the CosmosDB database. |
| `primaryConnectionString` | `string` | The full connection string, which includes the URI and primary key. |
| `databaseName` | `string` | The generated database name. |
| `documentdb_database_id` | `string` | The database name provided in a legacy key for use with Azure libraries. |
| `documentdb_host_endpoint` | `string` | The fully-qualified address and port of the CosmosDB database account provided in a legacy key for use with Azure libraries. |
| `documentdb_master_key` | `string` | A secret key used for connecting to the CosmosDB database provided in a legacy key for use with Azure libraries. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database. The existing database account is not deleted.

### Service: azure-cosmosdb-mongo-account

| Plan Name | Description |
|-----------|-------------|
| `account` | MongoDB on Azure provided by CosmosDB |

#### Behaviors

##### Provision

Provisions a new CosmosDB database account that can be accessed through the MongoDB API. The new database account is named using a new UUID.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `location` | `string` | The Azure region in which to provision applicable resources. | Y |  |
| `resourceGroup` | `string` | The (new or existing) resource group with which to associate new resources. | Y |  |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `consistencyPolicy` | `object` | The consistency policy for the Cosmos DB account. | N | |
| `consistencyPolicy.defaultConsistencyLevel` | `string` | The default consistency level and configuration settings of the Cosmos DB account. - Eventual, Session, BoundedStaleness, Strong, ConsistentPrefix | Y | |
| `consistencyPolicy.maxStalenessPrefix` | `integer` | When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `consistencyPolicy.maxIntervalInSeconds` | `integer` | When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `readRegions` | `array ` | Replication read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `multipleWriteRegionsEnabled` | `string` | Specifies if you want  the account to write in multiple regions. Valid values are [ "enabled", "disabled"]. If set to "enabled", regions in `readRegions`  will also be writable. | N | If not specified, "disabled" will be used as the default value. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `host` | `string` | The fully-qualified address of the CosmosDB database account. |
| `port` | `int` | The port number to connect to on the CosmosDB database account. |
| `username` | `string` | The name of the database user. |
| `password` | `string` | The password for the database user. |
| `connectionstring` | `string` | The full connection string, which includes the host, port, username, and password. |
| `uri` | `string` | URI encoded string that represents the connection information |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database account.

##### Update

Idempotently update the service instance to specified state.

| Parameter Name | Type                | Description                                                  | Required | Default Value                                                |
| -------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

### Service: azure-cosmosdb-graph-account

| Plan Name | Description |
|-----------|-------------|
| `account` | Database Account configured to use Graph (Gremlin) API |

#### Behaviors

##### Provision

Provisions a new CosmosDB database account that can be accessed through any of the Graph (Gremlin) API. The new database account is named using a new UUID.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `location` | `string` | The Azure region in which to provision applicable resources. | Y |  |
| `resourceGroup` | `string` | The (new or existing) resource group with which to associate new resources. | Y |  |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `consistencyPolicy` | `object` | The consistency policy for the Cosmos DB account. | N | |
| `consistencyPolicy.defaultConsistencyLevel` | `string` | The default consistency level and configuration settings of the Cosmos DB account. - Eventual, Session, BoundedStaleness, Strong, ConsistentPrefix | Y | |
| `consistencyPolicy.boundedStaleness` | object | Settings for to determine staleness when used with `BoundedStaleness` consistency | Yes - If using `BoundedStaleness` consistency | |
| `consistencyPolicy.boundedStaleness.maxStalenessPrefix` | `integer` | When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `consistencyPolicy.boundedStaleness.maxIntervalInSeconds` | `integer` | When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `readRegions` | `array ` | Replication read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `multipleWriteRegionsEnabled` | `string` | Specifies if you want  the account to write in multiple regions. Valid values are [ "enabled", "disabled"]. If set to "enabled", regions in `readRegions`  will also be writable. | N | If not specified, "disabled" will be used as the default value. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `uri` | `string` | The fully-qualified address and port of the CosmosDB database account. |
| `primaryKey` | `string` | A secret key used for connecting to the CosmosDB database account. |
| `primaryConnectionString` | `string` | The full connection string, which includes the URI and primary key. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database account.

##### Update

Idempotently update the service instance to specified state.

| Parameter Name | Type                | Description                                                  | Required | Default Value                                                |
| -------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

### Service: azure-cosmosdb-table-account

| Plan Name | Description |
|-----------|-------------|
| `account` | Database Account configured to use Table API |

#### Behaviors

##### Provision

Provisions a new CosmosDB database account that can be accessed through any of the Azure Table API. The new database account is named using a new UUID.

###### Provisioning Parameters

| Parameter Name | Type | Description | Required | Default Value |
|----------------|------|-------------|----------|---------------|
| `location` | `string` | The Azure region in which to provision applicable resources. | Y |  |
| `resourceGroup` | `string` | The (new or existing) resource group with which to associate new resources. | Y |  |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `consistencyPolicy` | `object` | The consistency policy for the Cosmos DB account. | N | |
| `consistencyPolicy.defaultConsistencyLevel` | `string` | The default consistency level and configuration settings of the Cosmos DB account. - Eventual, Session, BoundedStaleness, Strong, ConsistentPrefix | Y | |
| `consistencyPolicy.maxStalenessPrefix` | `integer` | When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `consistencyPolicy.maxIntervalInSeconds` | `integer` | When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'. | N | |
| `readRegions` | `array ` | Replication read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority.  See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `multipleWriteRegionsEnabled` | `string` | Specifies if you want  the account to write in multiple regions. Valid values are [ "enabled", "disabled"]. If set to "enabled", regions in `readRegions`  will also be writable. | N | If not specified, "disabled" will be used as the default value. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name | Type | Description |
|------------|------|-------------|
| `uri` | `string` | The fully-qualified address and port of the CosmosDB database account. |
| `primaryKey` | `string` | A secret key used for connecting to the CosmosDB database account. |
| `primaryConnectionString` | `string` | The full connection string, which includes the URI and primary key. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the CosmosDB database account.

##### Update

Idempotently update the service instance to specified state.

| Parameter Name | Type                | Description                                                  | Required | Default Value                                                |
| -------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `tags` | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |
| `ipFilters` | `object` | IP Range Filter to be applied to new CosmosDB account | N | A default filter is created that allows only Azure service access |
| `ipFilters.allowAccessFromAzure` | `string` | Specifies if Azure Services should be able to access the CosmosDB account. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowAccessFromPortal` | `string` | Specifies if the Azure Portal should be able to access the CosmosDB account. If `allowAccessFromAzure` is set to enabled, this value is ignored. Valid valued are `""` (unspecified), `enabled`, or `disabled`. | N | If left unspecified, defaults to enabled. |
| `ipFilters.allowedIPRanges` | `array` | Values to include in IP Filter. Can be IP Address or CIDR range. | N | If not specified, no additional values will be included in filters. |
| `readRegions` | `array ` | Read regions to be created, your data will be synchronized across these regions, providing high availability and disaster recovery ability. Region's order in the array will be treated as failover priority. **Note**: you can't update `readRegions` and other properties at the same time; if you want to update both `readRegions` and other properties, please update them in separate update operations. See [here](#About Read Regions) for points to pay attention to. | N | If not specified, no replication region will be created. |
| `autoFailoverEnabled` | `string ` | Specifies if you want Cosmos DB to perform automatic failover of the write region to one of the read regions in the rare event of a data center outage. Valid values are [ "enabled", "disabled"]. **Note**: If `multipleWriteRegionsEnabled` is set to `enabled`, all regions will be writable, and this attribute will not work. | N | If not specified, default "disabled". |

## Supplementary explanation

### About Read Regions

**Caution**: This feature has several constraint in ` Strong ` and ` Bounded Staleness ` consistency level, we recommend you use this feature in `Session`, ` Consistent Prefix ` and ` Eventual `  consistency level.

**Caution**: Do NOT fill provision parameter `location` in the `readRegions` array. For example, if the `location` of your account is `eastus` and you want to add a read region `westus`, you should use `{"readRegions": ["westus"]}` instead of `{"readRegions": ["eastus", "westus"]}`.

**Caution**: Allowed elements in `readRegions` array:  `"westus2", "westus", "southcentralus", "centralus", "northcentralus", "canadacentral", "eastus2", "canadaeast", "northeurope", "ukwest", "uksouth", "francecentral", "westeurope", "westindia", "centralindia", "southindia", "southeastasia", "eastasia", "koreacentral", "koreasouth", "japaneast", "japanwest", "australiasoutheast", "australiaeast"`.
