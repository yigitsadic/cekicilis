# Cekicilis

ENV Variables

```
DB_CONNECTION_STRING
PORT
```

## Tables

Events:

```sql
CREATE TABLE events
(
    id uuid DEFAULT uuid_generate_v4()  PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    expiresAt bigint NOT NULL,
    status int DEFAULT 0 NOT NULL
);
CREATE UNIQUE INDEX events_id_uindex ON events (id);
```

Participants:
```sql
CREATE TABLE public.participants
(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
    reference varchar(255),
    userReference varchar(255) NOT NULL,
    eventId uuid NOT NULL,
    CONSTRAINT fk_event
    FOREIGN KEY(eventId) REFERENCES events(id)
);
CREATE UNIQUE INDEX participants_id_uindex ON public.participants (id);
```
