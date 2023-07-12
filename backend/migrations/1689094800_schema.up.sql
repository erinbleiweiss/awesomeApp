CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL
);

INSERT INTO user ( first_name, last_name, email, username, password, role)
VALUES (
        'Jack',
        'Box',
        'jack@jackboxgames.com',
        'admin',
        '$2a$08$nkNm7OG2UrCbRK4yJElsKe4al9kSR8afMVoaPwV8G9klvuDZ3692i',
        'admin'
);

INSERT INTO user ( first_name, last_name, email, username, password, role)
VALUES (
           'Ron',
           'Swanson',
           'rswanson@pawnee.gov',
           'rswanson',
           '$2a$08$PnddSYjF.C98gBX5NeW1tOJHRMLcOO0jZBEY0HCwSJ7g5gmvNUJ.e',
           'user'
       );

INSERT INTO user ( first_name, last_name, email, username, password, role)
VALUES (
           'Leslie',
           'Knope',
           'lknope@pawnee.gov',
           'lknope',
           '$2a$08$RfwujkXmh28BeR2JicQvpuU9ZCt9VgaMTVoWWnz4ek2dufHf//fei',
           'user'
       );

INSERT INTO user ( first_name, last_name, email, username, password, role)
VALUES (
           'Ben',
           'Wyatt',
           'bwyatt@pawnee.gov',
           'bwyatt',
           '$2a$08$uyG.nsxf/ErmkUhRYSUvGOP3zAnR.bkVcTMdD3eKFZjYPvsi2RKve',
           'user'
       );