# Table: aws_resourcegroups_resource_groups

This table shows data for Resourcegroups Resource Groups.

https://docs.aws.amazon.com/ARG/latest/APIReference/API_GetGroupQuery.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|group_arn|String|
|name|String|
|description|String|
|query|String|
|type|String|