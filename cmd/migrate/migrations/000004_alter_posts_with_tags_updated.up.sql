ALTER TABLE 
    posts
ADD 
    COLUMN tags VARCHAR(255)[];

ALTER TABLE 
    posts
ADD
    COLUMN updated_at TIMESTAMP with time zone NOT NULL DEFAULT NOW();