-- Mostre todos os registros da tabela de filmes.
SELECT * FROM movies

-- Mostre o nome, sobrenome e classificação de todos os atores.
SELECT first_name, last_name, rating FROM actors

-- Mostre o título de todas as séries e use aliases para que o nome da tabela e o campo estejam em espanhol


-- Mostre o nome e sobrenome dos atores cuja classificação é superior a 7,5.
SELECT first_name, last_name FROM actors WHERE rating > 7.5

-- Mostre o título dos filmes, a classificação e os prêmios dos filmes com classificação superior a 7,5 e com mais de dois prêmios.
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2

-- Mostre o título dos filmes e a classificação ordenados por classificação em ordem crescente.
SELECT title, rating FROM movies ORDER BY rating

-- Mostre os títulos dos três primeiros filmes no banco de dados.
SELECT title FROM movies LIMIT 0,03

-- Mostre os 5 melhores filmes com a classificação mais alta.


-- Mostre o top 5 a 10 dos filmes com a classificação mais alta.


-- Liste os 10 primeiros atores (seria a página 1),


-- Em seguida, use o deslocamento para buscar a página 3


-- Faça o mesmo para a página 5


-- Mostre o título e a classificação de todos os filmes cujo título é Toy Story.


-- Mostre todos os atores cujos nomes começam com Sam.


-- Mostre o título dos filmes que saíram entre 2004 e 2008.


-- Traga o título dos filmes com classificação superior a 3, com mais de 1 prêmio e com data de lançamento entre 1988 e 2009. Classifique os resultados por classificação.


-- Busque os 3 primeiros do registro 10 da consulta anterior.