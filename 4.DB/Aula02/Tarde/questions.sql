-- Explique o conceito de normalização e por que ele é usado.
-- Normalização em banco de dados, é um processo de padronização e validação dos dados, com o objetivo de eliminar redundancias e insconsistencias.

-- Adicione um filme à tabela de filmes.
INSERT INTO movies_db.movies
VALUES (22, null, null, 'interstellar', 10.0, 10, '2012-01-02', 200, 10);

-- Adicione um gênero à tabela de gêneros.
INSERT INTO movies_db.genres 
VALUES(13, '2012-01-02', null, 'Horror', 13, 1);

-- Associe o filme do Ex 2. ao gênero criado no Ex 3.
UPDATE movies_db.movies m
	SET m.genre_id = 13
    WHERE m.id = 22;

-- Modifique a tabela de atores para que pelo menos um ator tenha o filme adicionado no Ex.2 como favorito.
UPDATE movies_db.actors a
		SET a.favorite_movie_id = 22
		WHERE a.id = 1; 

-- Crie uma cópia de tabela temporária da tabela de filmes.
CREATE TEMPORARY TABLE IF NOT EXISTS movies_db.moviesTemp AS
    (SELECT * FROM movies_db.movies);

-- Elimine dessa tabela temporária todos os filmes que ganharam menos de 5 prêmios.
DELETE FROM movies_db.moviesTemp
		WHERE movies_db.moviesTemp.awards <5;

-- Obtenha a lista de todos os gêneros que possuem pelo menos um filme.
SELECT DISTINCT g.name, m.genre_id
        FROM movies_db.movies m
        INNER JOIN movies_db.genres g
        WHERE g.id = m.genre_id;

-- Obtenha a lista de atores cujo filme favorito ganhou mais de 3 prêmios.
SELECT a.first_name, m.awards
        FROM movies_db.movies m
        INNER JOIN movies_db.actors a
        WHERE a.favorite_movie_id = m.id
        AND m.awards > 3;

-- Use o plano de explicação para analisar as consultas em Ex.6 e 7.
EXPLAIN SELECT * FROM movies_db.filmes;

-- O que são índices? Para que servem?
-- Indices nada mais são do que mecanismos para otimizar nossas consultas SQL.

-- Crie um índice no nome na tabela de filmes.
ALTER TABLE movies_db.movies
        ADD INDEX(titulo);

-- Verifique se o índice foi criado corretamente.
SHOW INDEX 
        FROM movies_db.movies;