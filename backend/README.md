# Bug Bug Tracker's Backend Design

## Endpoints

### User

- `/users`:
    - `GET`: get all user's public data

- `/users/me`:
    - `GET`: get all data (public and private) about the requester user

- `/users/:username`:
    - `GET`: get a specific user's public data

### History

- `/user/:username/histories`:
    - `GET`: get History that is attached to specific User

- `/issue/:issueId/histories`:
    - `GET`: get History that is attached to specific Issue

## Concepts

### History

- Must be retrieved according to either Issue or User that is linked to it.
