#!/bin/sh

set -e

if [ -z "$SHIPATOR_PLACEHOLDER" ]; then
  echo "SHIPATOR_PLACEHOLDER was not provided, aborting!"
  exit 1
fi

if [ -z "$SHIPATOR_PREFIX" ]; then
  echo "SHIPATOR_PREFIX was not provided, aborting!"
  exit 1
fi

if [ -z "$SHIPATOR_TARGET" ]; then
  echo "SHIPATOR_TARGET was not provided, aborting!"
  exit 1
fi

shipator -placeholder $SHIPATOR_PLACEHOLDER -prefix $SHIPATOR_PREFIX $SHIPATOR_TARGET
exec nginx -g "daemon off;"
