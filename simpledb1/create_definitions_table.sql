DROP TABLE IF EXISTS Models_Branches_V_0000_0001;
CREATE TABLE Models_Branches_V_0000_0001 (
	branch uuid DEFAULT uuid_generate_v4(),
	label TEXT,
	description TEXT,
	tags TEXT[],
	created INTEGER,
	modified INTEGER,
	origin uuid,
	frozen BOOLEAN,
	PRIMARY KEY (branch)
);

