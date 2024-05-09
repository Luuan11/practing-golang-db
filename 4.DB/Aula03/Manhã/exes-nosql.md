### 1 - Traga 3 restaurantes que receberam pelo menos uma classificação de grau 'A' com pontuação maior que 20. A mesma classificação deve atender às duas condições simultaneamente; investigue o operador elemMatch.
Trazendo os 3 restaurantes:
db.getCollection('restaurants').find({
    "grados": {
        $elemMatch: {
            "grado": "A",
            "puntaje": { $gt: 20 }
        }
    }
}).limit(3)


### 2 - Quantos documentos estão faltando coordenadas geográficas? Em outras palavras, verifique se o tamanho de direccion.coord é 0 e contar.
os documentos:
db.getCollection('restaurants').countDocuments({
    "direccion.coord": { $size: 0 }
})

### 3 - Devolver nombre, barrio, tipo_cocina y grados para os 3 primeiros restaurantes; de cada documento apenas a última avaliação. Ver o operador slice.
trazendo:
db.getCollection('restaurants').find({}, {
    "nombre": 1,
    "barrio": 1,
    "tipo_cocina": 1,
    "grados": { $slice: -1 }
}).limit(3)
