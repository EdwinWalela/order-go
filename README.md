# Go Order API

Repository for the Go Order API

## Tech Stack

- Golang (API)
- GCP (Firestore & CloudRun)
- Docker (local development)
- Terraform (IAC)

## Run locally 

The project runs locally using `docker-compose`. 

### Environment variables

Create an .env file using the sample file

```bash
cp .sample.en .env
```

### Firestore Authentication

The project uses Google Firestore as a database and requires authentication via a service account json file.

The file should be names `firestore-sa.json` and located in `./orders` directory


### Start the service

```
docker-compose up
```

## CI/CD

The project uses Terraform as IAC. PRs opened targeting main trigger the build and push workflow to Google Container Registry

## API Documentation

[Documentation](https://documenter.getpostman.com/view/4654837/2s93zB5Mm9)