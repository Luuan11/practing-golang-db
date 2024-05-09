### Quantas coleções tem o banco de dados?
    Uma
### Quantos documentos em cada coleção? Quanto ocupa cada coleção?
db.getCollection('restaurants').find({}).count()
    25359

 db.getCollection('restaurants').dataSize()
    11140976

### Quantos índices em cada coleção? Quanto espaço os índices de cada coleção ocupam?

    

### Traga um documento de exemplo de cada coleção. db.collection.find(...).pretty() nos dá um formato mais legível.

### Para cada coleção, liste os campos de nível raiz (ignore campos em documentos aninhados) e seus tipos de dados.

