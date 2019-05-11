# Minio S3

Upload or Download files from Minio Simple Storage Service (S3)


## Installation

```bash
flogo install github.com/adesaeher/flogo-components/activity/minios3
```
Link for flogo web:
```
https://github.com/adesaegher/flogo-components/activity/minios3
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "action",
            "type": "string",
            "allowed": [
                "download",
                "upload",
                "delete",
                "copy"
            ],
            "required": true
        },
        {
            "name": "s3Endpoint",
            "type": "string",
            "required": true
        },
        {
            "name": "awsAccessKeyID",
            "type": "string",
            "required": true
        },
        {
            "name": "awsSecretAccessKey",
            "type": "string",
            "required": true
        },
        {
            "name": "useSsl",
            "type": "bool",
            "required": true
        },
        {
            "name": "awsRegion",
            "type": "string",
            "required": true
        },
        {
            "name": "s3BucketName",
            "type": "string",
            "required": true
        },
        {
            "name": "s3Location",
            "type": "string",
            "required": true
        },
        {
            "name": "localLocation",
            "type": "string"
        },
        {
            "name": "s3NewLocation",
            "type": "string"
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "string"
        }
    ]
}
```
## Inputs
| Input              | Description    |
|:-------------------|:---------------|
| action             | The action you want to take, either `download`, `upload`, `delete`, or `copy` |
| s3Endpoint         | Your Minio S3 Endpoint |
| awsAccessKeyID     | Your AWS Access Key (only needed if you don't give your Lambda function rights to invoke S3 reads and puts) |
| awsSecretAccessKey | Your AWS Secret Key (only needed if you don't give your Lambda function rights to invoke S3 reads and puts) |
| useSsl             | Use ssl |
| awsRegion          | The AWS region your S3 bucket is in |
| s3BucketName       | The name of your S3 bucket |
| s3Location         | The file location on S3, this should be a full path (like `/bla/temp.txt`) |
| localLocation      | The `localLocation` is the full path to a file (like `/bla/temp.txt`) when uploading a file or the full path to a directory (like `./tmp`) when downloading a file |
| s3NewLocation      | The new file location on S3 of you want to copy a file, this should be a full path (like `/bla/temp.txt`) |

## Ouputs
| Output    | Description    |
|:----------|:---------------|
| result    | The result will contain OK if the action was carried out successfully or will contain an error message |
