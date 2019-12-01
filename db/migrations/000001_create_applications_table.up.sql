CREATE TABLE applications (
    name String unique,
    api_key String,
    enabled Boolean,
    created_at DateTime,
    updated_at DateTime,
    deleted_at DateTime
) engine=Memory;
