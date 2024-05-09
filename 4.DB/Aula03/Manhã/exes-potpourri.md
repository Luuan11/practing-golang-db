### Quais são os 3 principais tipos de culinária (cuisine) que podemos encontrar nos dados? Googlear "mongodb group by field, count it and sort it". Ver a etapa limit do pipeline de agregação.

esses são os 3 principais:
db.getCollection('restaurants').aggregate([
    { $group: { _id: "$tipo_cocina", count: { $sum: 1 } } },
    { $sort: { count: -1 } },
    { $limit: 3 }
])

### Quais são os bairros mais desenvolvidos gastronomicamente? Calcular a média ($avg) a pontuação (grades.score) por bairro; considerando restaurantes com mais de três avaliações; oClassifique os bairros com a melhor pontuação. Ajuda:
db.getCollection('restaurants').aggregate([
    { $unwind: "$grados" }, 
    { $group: { 
        _id: "$barrio", // Agrupa por bairro
        pontuacaoMedia: { $avg: "$grados.puntaje" }, 
        count: { $sum: 1 } 
    }},
    { $match: { count: { $gt: 3 } } }, 
    { $sort: { pontuacaoMedia: -1 } } 
])

### match é um estágio que filtra documentos com base em uma condição, similar a db.orders.find(<condição>).
### Parece necessário desconstruir as listas grades para produzir um documento para cada ponto usando o palco unwind.
db.getCollection('restaurants').aggregate([
    { $unwind: "$grados" }, // Desdobra o array de grados para que cada documento tenha uma única avaliação
    { $group: { 
        _id: "$barrio", // Agrupa por bairro
        pontuacaoMedia: { $avg: "$grados.puntaje" }, // Calcula a média da pontuação para cada bairro
        count: { $sum: 1 } // Conta o número de avaliações para cada bairro
    }},
    { $match: { count: { $gt: 3 } } }, // Filtra bairros com mais de três avaliações
    { $sort: { pontuacaoMedia: -1 } } // Ordena os resultados pela média de pontuação em ordem decrescente
])

### Uma pessoa querendo comer está na longitude -73.93414657 e na latitude 40.82302903, Quais opções você tem em um raio de 500 metros? Consultar geospatial tutorial.
db.restaurants.createIndex({ "direccion.coord": "2dsphere" })

### Para buscar os 500 metros
db.restaurants.find({
    "direccion.coord": {
        $nearSphere: {
            $geometry: {
                type: "Point",
                coordinates: [-73.93414657, 40.82302903]
            },
            $maxDistance: 500
        }
    }
})
