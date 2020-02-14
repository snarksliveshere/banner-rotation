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
    CONSTRAINT public_audience2banner_audience_banner_uidx UNIQUE (audience_fk, banner_fk)
);

CREATE TABLE banner2slot
(
    id        BIGSERIAL PRIMARY KEY,
    slot_fk   BIGINT NOT NULL,
    banner_fk BIGINT NOT NULL,
    CONSTRAINT public_banner2slot_slot_banner_uidx UNIQUE (slot_fk, banner_fk)
);


INSERT INTO banner (banner_id)
VALUES ('some_male1_adult_app_id'),
       ('some_male2_adult_app_id'),
       ('some_male3_adult_app_id'),
       ('some_male1_kid_app_id'),
       ('some_male2_kid_app_id'),
       ('some_male3_kid_app_id'),
       ('some_female1_adult_app_id'),
       ('some_female2_adult_app_id'),
       ('some_female3_adult_app_id'),
       ('some_female1_kid_app_id'),
       ('some_female2_kid_app_id'),
       ('some_female3_kid_app_id')
;
INSERT INTO audience (audience_id)
VALUES ('male_kid'),
       ('female_kid'),
       ('male_adult'),
       ('female_adult')
;
INSERT INTO slot (slot_id)
VALUES ('top_slot_id'),
       ('bottom_slot_id')
;

INSERT INTO banner2slot (slot_fk, banner_fk)
VALUES ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male2_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male3_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'top_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_kid_app_id')),

       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male3_adult_app_id')),

       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_adult_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_adult_app_id')),

       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_kid_app_id')),

       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male3_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male2_kid_app_id')),
       ((SELECT id FROM slot WHERE slot_id = 'bottom_slot_id'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_kid_app_id'))
;

INSERT INTO audience2banner (audience_fk, banner_fk)
VALUES ((SELECT id FROM audience WHERE audience_id = 'male_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_male2_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_male2_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_male3_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_adult_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_male1_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_male2_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'male_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_male3_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_female1_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_adult'),
        (SELECT id FROM banner WHERE banner_id = 'some_female2_kid_app_id')),
       ((SELECT id FROM audience WHERE audience_id = 'female_kid'),
        (SELECT id FROM banner WHERE banner_id = 'some_female3_kid_app_id'))
;


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