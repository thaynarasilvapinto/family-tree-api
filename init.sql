CREATE TABLE family_tree (
    id INTEGER PRIMARY KEY,
    name TEXT,
    lft INTEGER,
    rgt INTEGER,
    parent_id INTEGER REFERENCES family_tree(id)
);