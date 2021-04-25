#!/bin/sh

if [ -n "$DB_USER_FILE" ]; then
  export DB_USER=$(cat $DB_USER_FILE)
fi

if [ -n "$DB_PASSWORD_FILE" ]; then
  export DB_PASSWORD=$(cat $DB_PASSWORD_FILE)
fi

revel run /accounting prod
