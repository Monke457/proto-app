CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	firstname VARCHAR(50),
	lastname VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(80)
);
