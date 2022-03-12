-- CREATE TABLE total_balance (
--   id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
--   total bigint NOT NULL
-- );

CREATE TABLE monthly_balance (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY key,
  -- total_id bigint,
  Amount bigint,
  year_month timestamp
);

CREATE TABLE entries (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY key,
  title varchar not null,
  entry_type varchar(40) not null,
  balance_id bigint,
  amount bigint NOT NULL,
  -- installments int NOT NULL,
  -- is_planned boolean NOT NULL,
  entry_date date not null,
  created_at timestamptz NOT NULL DEFAULT (now())
);

-- ALTER TABLE monthly_balance ADD FOREIGN KEY (total_id) REFERENCES total_balance (id);

-- ALTER TABLE entries ADD FOREIGN KEY (balance_id) REFERENCES monthly_balance (id);

COMMENT ON COLUMN "entries"."amount" IS 'can be positive or negative';
