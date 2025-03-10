# Table: aws_securityhub_findings

This table shows data for AWS Security Hub Findings.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html
The `request_account_id` and `request_region` columns are added to show the account and region of where the request was made from.
This is useful when multi region and account aggregation is enabled.

The composite primary key for this table is (**request_account_id**, **request_region**, **aws_account_id**, **created_at**, **description**, **generator_id**, **id**, **product_arn**, **schema_version**, **title**, **updated_at**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|request_account_id (PK)|String|
|request_region (PK)|String|
|aws_account_id (PK)|String|
|created_at (PK)|Timestamp|
|description (PK)|String|
|generator_id (PK)|String|
|id (PK)|String|
|product_arn (PK)|String|
|resources|JSON|
|schema_version (PK)|String|
|title (PK)|String|
|updated_at (PK)|Timestamp|
|action|JSON|
|company_name|String|
|compliance|JSON|
|confidence|Int|
|criticality|Int|
|finding_provider_fields|JSON|
|first_observed_at|Timestamp|
|last_observed_at|Timestamp|
|malware|JSON|
|network|JSON|
|network_path|JSON|
|note|JSON|
|patch_summary|JSON|
|process|JSON|
|product_fields|JSON|
|product_name|String|
|record_state|String|
|region (PK)|String|
|related_findings|JSON|
|remediation|JSON|
|sample|Bool|
|severity|JSON|
|source_url|String|
|threat_intel_indicators|JSON|
|threats|JSON|
|types|StringArray|
|user_defined_fields|JSON|
|verification_state|String|
|vulnerabilities|JSON|
|workflow|JSON|
|workflow_state|String|