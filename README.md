# Khaikhan-back

## Installation

- Install [gin](https://github.com/gin-gonic/gin) package /includes doc/.
- Install [viper](https://github.com/spf13/viper) package /includes doc/.

## Usage

- Change config files in `./config` directory (`/config/dev.json /config/prod.json`)
- Rename database in `./Db/query.sql`

- if you can't change your user as admin(your user) then add table as postgres(owner) after run that give privilage to admin(your user) by following code:

```
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA MY_SCHEMA TO USER_NAME;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA MY_SCHEMA TO USER_NAME;
```

## Template includes

- Todu login / login /
- Viper / config loader /
- Userman / user management base model /
- Gin / Web Framework /
