DROP TABLE  IF EXISTS"councils";

DROP TABLE  IF EXISTS"schedule";


DROP TABLE  IF EXISTS"time_slots";

DROP TABLE  IF EXISTS"thesis";

ALTER TABLE "thesis_commitees"
DROP COLUMN "time_slots_id";

ALTER TABLE "thesis_commitees"
DROP COLUMN "time";