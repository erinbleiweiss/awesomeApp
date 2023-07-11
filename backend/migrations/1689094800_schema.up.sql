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
