CREATE TYPE HouseKind AS ENUM  (
  'House',
  'Rooms'
);

CREATE TYPE RentalStatus AS ENUM (
  'Rented',
  'Empty'
);

CREATE TABLE "house" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "address" VARCHAR NOT NULL,
  "kind" HouseKind NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "rental_unit" (
  "id" SERIAL PRIMARY KEY,
  "house_id" INT NOT NULL,
  "price" NUMERIC(10,2) NOT NULL,
  "status" RentalStatus NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "renter" (
  "id" SERIAL PRIMARY KEY,
  "full_name" VARCHAR NOT NULL
);

CREATE TABLE "rental_agreement" (
  "id" SERIAL PRIMARY KEY,
  "renter_id" INT NOT NULL,
  "rental_id" INT NOT NULL,
  "start_date" DATE NOT NULL,
  "end_date" DATE,
  "price" NUMERIC(10,2),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "rental_unit" ADD FOREIGN KEY ("house_id") REFERENCES "house" ("id");

ALTER TABLE "rental_agreement" ADD FOREIGN KEY ("renter_id") REFERENCES "renter" ("id");

ALTER TABLE "rental_agreement" ADD CONSTRAINT "fk_house" FOREIGN KEY ("rental_id") REFERENCES "house" ("id");

ALTER TABLE "rental_agreement" ADD CONSTRAINT "fk_room" FOREIGN KEY ("rental_id") REFERENCES "rental_unit" ("id");
