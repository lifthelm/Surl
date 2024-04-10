#!/bin/bash

eval "psql -U postgres -d postgres -a -f /scripts/db_init.sql"
eval "psql -U postgres -d surl -a -f /scripts/surl_table_creation.sql"
eval "psql -U postgres -d surl -a -f /scripts/surl_constraints.sql"
eval "psql -U postgres -d surl -a -f /scripts/surl_data_load.sql"
eval "psql -U postgres -d surl -a -f /scripts/surl_sequence_update.sql"

