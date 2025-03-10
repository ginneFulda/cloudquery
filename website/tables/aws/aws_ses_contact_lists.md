# Table: aws_ses_contact_lists

This table shows data for Amazon Simple Email Service (SES) Contact Lists.

https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetContactList.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|tags|JSON|
|contact_list_name|String|
|created_timestamp|Timestamp|
|description|String|
|last_updated_timestamp|Timestamp|
|topics|JSON|