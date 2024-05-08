-- Mostre o título e o nome do gênero de todas as séries.
select s.title, g.name from series s INNER JOIN genres g ON (g.id = s.genre_id)

-- Mostre o título dos episódios, o nome e sobrenome dos atores que trabalham em cada um deles.
SELECT e.title, a.first_name, a.last_name from episodes e
	INNER JOIN actor_episode ae ON (ae.episode_id = e.id)
	INNER JOIN actors a ON (ae.actor_id = a.id)

-- Mostre o título de todas as séries e o número total de temporadas que cada uma delas possui.
SELECT s.title, MAX(se.number) from series s 
  INNER JOIN seasons se ON (s.id = se.serie_id) GROUP BY s.title 

-- Mostre o nome de todos os gêneros e o número total de filmes de cada um, desde que seja maior ou igual a 3.


-- Mostre apenas o nome e sobrenome dos atores que atuam em todos os filmes de Star Wars e que estes não se repitam.
