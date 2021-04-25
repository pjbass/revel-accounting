#!/bin/sh

if [ -n "$DB_USER_FILE" ]; then
  export DB_USER=$(cat $DB_USER_FILE)
fi

if [ -n "$DB_PASSWORD_FILE" ]; then
  export DB_PASSWORD=$(cat $DB_PASSWORD_FILE)
fi

# This one to accomodate Heroku Postgres (or to provide a single secret value
# vs multiple)
if [ -n "$DATABASE_URL" ]; then
  export DB_DRIVER=$(echo "$DATABASE_URL" | cut -d ':' -f 1)
  export DB_USER=$(echo "$DATABASE_URL" | cut -d '/' -f 3 | cut -d ':' -f 1)
  export DB_PASSWORD=$(echo "$DATABASE_URL" | cut -d ':' -f 3 | cut -d '@' -f 1)
  export DB_HOST=$(echo "$DATABASE_URL" | cut -d '@' -f 2 | cut -d '/' -f 1)
  export DB_NAME=$(echo "$DATABASE_URL" | cut -d '/' -f 4)
fi

/accounting/run.sh
