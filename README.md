# Ticket App

## API Endpoints

### Event Endpoints
| HTTP Method | Endpoint                          | Description          |
| ----------- | --------------------------------- | -------------------- |
| POST        | `{{base_ticket_url}}/v1/event`    | Create New Event     |
| DELETE      | `{{base_ticket_url}}/event/:event_id`     | Delete an event      |
| PUT         | `{{base_ticket_url}}/event/:event_id`     | Update event by id   |
| GET         | `{{base_ticket_url}}/event/:event_id`     | Get Event by id      |

### Seat Endpoints
| HTTP Method | Endpoint                                   | Description              |
| ----------- | ------------------------------------------ | ------------------------ |
| POST        | `{{base_ticket_url}}/seat/1`               | Create New Seat          |
| GET         | `{{base_ticket_url}}/event/:event_id/seats` | Get Seat from Event Id   |

### Public Endpoints
| HTTP Method | Endpoint                        | Description   |
| ----------- | ------------------------------- | ------------- |
| POST        | `{{base_ticket_url}}/public`    | Send Email    |

### Webhook Endpoints
| HTTP Method | Endpoint                          | Description   |
| ----------- | --------------------------------- | ------------- |
| POST        | `{{base_ticket_url}}/webhook`     | Webhook       |

## Notes

- Replace `{{base_ticket_url}}` with the actual base URL of the Ticket App service.
- For endpoints with parameters (like `:event_id`), replace them with actual values when calling the API.

## API Docs
https://crimson-meadow-438973.postman.co/workspace/PAT~5e4b20a9-a21e-48b8-8eef-baeb56a29ad7/collection/30701742-215d3d8c-31b2-4ae8-adf4-d23db163d0d6?action=share&creator=30701742&active-environment=30701742-3c17942c-d556-4a3c-b175-2402ac791441

## How to Use
1. Clone or fork this repository
```sh
https://github.com/NicholasLiem/IF4031_M1_Ticket_App
```
2. Initialize .env file using the template given (.env.example and docker.env.example)
```sh
touch .env
touch docker.env
```
3. Run docker compose and build
```sh
docker-compose up --build
```