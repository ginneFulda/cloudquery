# Table: digitalocean_databases

This table shows data for DigitalOcean Databases.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_databases:
  - [digitalocean_database_backups](digitalocean_database_backups)
  - [digitalocean_database_firewall_rules](digitalocean_database_firewall_rules)
  - [digitalocean_database_replicas](digitalocean_database_replicas)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|engine|String|
|version|String|
|connection|JSON|
|private_connection|JSON|
|users|JSON|
|num_nodes|Int|
|size|String|
|db_names|StringArray|
|region|String|
|status|String|
|maintenance_window|JSON|
|created_at|Timestamp|
|private_network_uuid|String|
|tags|StringArray|
|project_id|String|