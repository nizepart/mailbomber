## API Endpoints


### POST /users

**Description:** Creates a new user.

**Request JSON example:**
```json
{ 
    "username": "<username>",
    "password": "<password>"
}
```


### POST /sessions

**Description:** Creates a new session for user login.

**Request JSON example:**
```json
{
    "username": "<username>",
    "password": "<password>"
} 
```


### GET /private/whoami

**Description:** Returns the information of the authenticated user.

**Request JSON:** No request body required.


### POST /private/email/send

**Description:** Sends an email.

**Request JSON example:**
```json
{
  "to": ["<recipient email>, ..."],
  "cc": ["<cc email>, ..."], // optional 
  "subject": "<email subject>",
  "body": "<email body>",
  "bodyType": "<email type>",
  "attach": "<email attach>" // optional 
}
```


### POST /private/email/template/schedule

**Description:** Schedules an email to be sent later.

**Request JSON example:**
```json
{ 
    "template_id": "<email template id>",
    "recipients": "<recipient email>, ...",
    "execute_after": "<datetime>", // "2024-04-19T23:40:00+03:00"
    "execution_period": "<hours>" // optional 
}
```


### POST /private/email/template/create

**Description:** Creates a new email template.

**Request JSON example:**
```json
{ 
    "subject": "<email subject>",
    "body": "<email body>",
    "bodyType": "<email type>"
}
```


### GET /private/email/template/{id}

**Description:** Retrieves the email template with the specified id.

**Request JSON:** No request body required.


### POST /private/email/template/{id}/send

**Description:** Sends an email using the specified template.

**Request JSON example:**
```json
{
  "to": ["<recipient email>, ..."],
  "cc": ["<cc email>, ..."], // optional
  "attach": "<email attach>" // optional 
}
```

## .env Configuration File

The `.env` file is used to set environment variables that your application needs to run. These variables are loaded into
the `os` package and can be accessed using `os.Getenv`.

Here is a brief description of each variable:

| Environment Variable | Fallback Value                                                                   | Description                                                                                       |
|----------------------|----------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| `BIND_ADDR`          | `:8080`                                                                          | The address and port on which your application will run.                                          |
| `LOG_LEVEL`          | `debug`                                                                          | The logging level for your application.                                                           |
| `SESSION_KEY`        | `1234567890`                                                                     | The key used for session management in your application.                                          |
| `SMTP_HOST`          | `localhost`                                                                      | The host address of your SMTP server for sending emails.                                          |
| `SMTP_PORT`          | `587`                                                                            | The port of your SMTP server.                                                                     |
| `SMTP_FROM`          | `noreply@localhost`                                                              | The default email address that will be used as the sender in the emails sent by your application. |
| `TZ`                 | `UTC`                                                                            | The timezone that will be used in your application.                                               |
| `DB_PASSWORD`        |                                                                                  | The password for your database.                                                                   |
| `DB_USER`            |                                                                                  | The username for your database.                                                                   |
| `DB_NAME`            |                                                                                  | The name of your database.                                                                        |
| `DB_HOST`            |                                                                                  | The host of your database.                                                                        |
| `DB_TEST_URL`        | `host=db user=postgres password=postgres dbname=mailbomber_test sslmode=disable` | The connection string for your test database.                                                     |
| `DB_URL`             | `host=db user=postgres password=postgres dbname=mailbomber sslmode=disable`      | The connection string for your database.                                                          |
| `GOOSE_DRIVER`       |                                                                                  | The driver used by Goose for database migrations.                                                 |
| `GOOSE_DBSTRING`     |                                                                                  | The connection string used by Goose for database migrations.                                      |
