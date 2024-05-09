### 1 Devolver restaurante_id, nombre, barrio e tipo_cocina mas excluindo _id para um documento (o primeiro).
buscando: 
db.getCollection('restaurants').find({}, {restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}).limit(1)

### 2 Devolver restaurante_id, nombre, barrio e tipo_cocina para os primeiros 3 restaurantes que contenham 'Bake' em alguma parte do seu nome.
buscando:
db.getCollection('restaurants').find({nombre: /Bake/}, {restaurante_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1, _id: 0}).limit(3)

### 3 Contar os restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) do bairro (bairro) Bronx. Consultar or versus in.
contando:  
db.getCollection('restaurants').find({$or: [{tipo_cocina: 'Chinese'}, {tipo_cocina: 'Thai'}], barrio: 'Bronx'}).count()
