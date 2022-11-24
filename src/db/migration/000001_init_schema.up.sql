CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "balance" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payment_by" (
  "id" bigserial PRIMARY KEY,
  "payment_name" varchar NOT NULL
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "payment_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "desc" varchar NOT NULL,
  "amount" bigint NOT NULL
);

CREATE INDEX ON "accounts" ("user_name");

CREATE INDEX ON "payment_by" ("payment_name");

CREATE INDEX ON "transactions" ("account_id");

CREATE INDEX ON "transactions" ("payment_id");

CREATE INDEX ON "transactions" ("account_id", "payment_id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("payment_id") REFERENCES "payment_by" ("id");
