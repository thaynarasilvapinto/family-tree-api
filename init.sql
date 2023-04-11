CREATE TABLE family_tree (
    id SERIAL PRIMARY KEY,
    name TEXT,
    parent1_id INTEGER REFERENCES family_tree(id),
    parent2_id INTEGER REFERENCES family_tree(id)
);


-- AVÓS
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('John', null, null);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Mary', null, null);

-- PAIS
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Peter', 1, 2);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Maria', null, null);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Lucy', 1, 2);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Marcos', null, null);

-- FILHOS
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Tom', 3, 4);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Kate', 3, 4);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Jack', 5, 6);
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('Sophie', 5, 6);

-- NETOS
INSERT INTO family_tree (name, parent1_id, parent2_id) VALUES ('João', 10, null);