-- Drop tables in reverse order of creation (to handle foreign key dependencies)
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS states;
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS user_addresses;