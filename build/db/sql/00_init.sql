CREATE TABLE IF NOT EXISTS todo (
  id SERIAL NOT NULL,
  title VARCHAR(40) NOT NULL,
  content VARCHAR(100) NOT NULL,
  create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO todo (title, content) VALUES ('test', 'test');
INSERT INTO todo (title, content) VALUES ('sample', 'sample');
