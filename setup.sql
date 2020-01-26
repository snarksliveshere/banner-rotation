CREATE TABLE banner
(
    id        BIGSERIAL PRIMARY KEY,
    banner_id TEXT NOT NULL,
    CONSTRAINT public_banner_banner_id_uidx UNIQUE (banner_id)
);

CREATE TABLE slot
(
    id      BIGSERIAL PRIMARY KEY,
    slot_id TEXT NOT NULL,
    CONSTRAINT public_slot_slot_id_uidx UNIQUE (slot_id)
);

CREATE TABLE audience
(
    id          BIGSERIAL PRIMARY KEY,
    audience_id TEXT NOT NULL,
    CONSTRAINT public_audience_audience_id_uidx UNIQUE (audience_id)
);

CREATE TABLE audience2banner
(
    id          BIGSERIAL PRIMARY KEY,
    audience_fk BIGINT NOT NULL,
    banner_fk   BIGINT NOT NULL,
    CONSTRAINT public_audience2banner_audience_banner_uidx UNIQUE (audience_fk, banner_fk),
    CONSTRAINT public_audience2banner_audience_fk FOREIGN KEY (audience_fk) REFERENCES audience (id),
    CONSTRAINT public_audience2banner_banner_fk FOREIGN KEY (banner_fk) REFERENCES banner (id)
);

CREATE TABLE statistics_slot
(
    id          BIGSERIAL PRIMARY KEY,
    audience_fk BIGINT           NOT NULL,
    banner_fk   BIGINT           NOT NULL,
    slot_fk     BIGINT           NOT NULL,
    shows       BIGINT DEFAULT 0 NOT NULL,
    clicks      BIGINT DEFAULT 0 NOT NULL,
    CONSTRAINT public_statistics_slot_audience_banner_uidx UNIQUE (audience_fk, banner_fk, slot_fk),
    CONSTRAINT public_statistics_slot_audience_fk FOREIGN KEY (audience_fk) REFERENCES audience (id),
    CONSTRAINT public_statistics_slot_banner_fk FOREIGN KEY (banner_fk) REFERENCES banner (id),
    CONSTRAINT public_statistics_slot_slot_fk FOREIGN KEY (slot_fk) REFERENCES slot (id)
);

CREATE TABLE statistics
(
    id          BIGSERIAL PRIMARY KEY,
    audience_fk BIGINT           NOT NULL,
    banner_fk   BIGINT           NOT NULL,
    shows       BIGINT DEFAULT 0 NOT NULL,
    clicks      BIGINT DEFAULT 0 NOT NULL,
    CONSTRAINT public_statistics_audience_banner_uidx UNIQUE (audience_fk, banner_fk),
    CONSTRAINT public_statistics_audience_fk FOREIGN KEY (audience_fk) REFERENCES audience (id),
    CONSTRAINT public_statistics_banner_fk FOREIGN KEY (banner_fk) REFERENCES banner (id)
);



insert into calendar.calendar (date, description)
values ('2019-11-10', 'some desc'),
       ('2019-11-12', 'some desc2'),
       ('2019-11-15', 'some desc3'),
       ('2019-10-20', 'some desc4')
;

insert into calendar.event (date_fk, time, title, description)
values ((SELECT id FROM calendar.calendar WHERE date = '2019-11-10'), '2019-11-10 07:18:09.767953 +00:00',
        'some title event1', 'desc event1'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-11-10'), '2019-11-10 09:20:09.767953 +00:00',
        'some title event2', 'desc event2'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-11-12'), '2019-11-12 10:20:09.767953 +00:00',
        'some title event3', 'desc event3'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-10-20'), '2019-10-20 10:30:09.767953 +00:00',
        'some title event4', 'desc event4')
;
CREATE TABLE calendar.message
(
    id     BIGSERIAL PRIMARY KEY,
    status TEXT NOT NULL,
    msg    TEXT
);
WITH cte AS (
    INSERT INTO calendar.calendar (date, description)
        VALUES (NOW() :: date, 'test_msg')
        RETURNING id
)
INSERT
INTO calendar.event (date_fk, time, title, description)
VALUES ((SELECT id FROM cte),
        NOW() + interval '5 minutes',
        'test_title',
        'test_description');