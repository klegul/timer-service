# timer-service

This simple project also allows a user to call an HTTP endpoint, which starts a timer that sends a POST request to the specified callback URL when the timer has expired.

## Usage

Use the Docker image `ghcr.io/klegul/timer-service:latest`.

Send a POST request to `HOST:8080/start-timer` with the following `application/json` request body:

```JSON
{
  "duration": 60,
  "calback_url": "http://other.host/callback"
}
```

The duration is given in seconds.
