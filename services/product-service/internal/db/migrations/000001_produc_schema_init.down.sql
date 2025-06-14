DROP TABLE IF EXISTS products;

DO $$
BEGIN
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_status') THEN
    DROP TYPE product_status;
  END IF;

  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'availability_status') THEN
    DROP TYPE availability_status;
  END IF;
END
$$;
