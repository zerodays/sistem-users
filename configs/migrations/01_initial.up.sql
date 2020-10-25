CREATE TABLE users
(
    uid          TEXT PRIMARY KEY,

    email        TEXT                                               NOT NULL UNIQUE,
    name         TEXT                                               NOT NULL,

    password     TEXT                                               NOT NULL,

    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Refresh tokens --
CREATE TABLE authenticated_devices
(
    user_id      TEXT REFERENCES users (uid) ON DELETE CASCADE      NOT NULL,
    token        TEXT                                               NOT NULL,

    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_used    TIMESTAMP WITH TIME ZONE                           NOT NULL
);

CREATE TABLE user_permissions
(
    user_id    TEXT REFERENCES users (uid) ON DELETE CASCADE NOT NULL,
    permission TEXT                                          NOT NULL,

    UNIQUE (user_id, permission)
);
