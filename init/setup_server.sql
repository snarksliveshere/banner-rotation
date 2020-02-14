CREATE TABLE statistics
(
    id          BIGSERIAL PRIMARY KEY NOT NULL,
    audience_id TEXT                  NOT NULL,
    banner_id   TEXT                  NOT NULL,
    slot_id     TEXT                  NOT NULL,
    shows       BIGINT DEFAULT 0      NOT NULL,
    clicks      BIGINT DEFAULT 0      NOT NULL,
    CONSTRAINT public_statistics__uidx UNIQUE (audience_id, banner_id, slot_id)
);
