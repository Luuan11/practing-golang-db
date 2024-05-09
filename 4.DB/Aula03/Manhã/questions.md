### Quantas coleções tem o banco de dados?
    Uma
### Quantos documentos em cada coleção? Quanto ocupa cada coleção?
db.getCollection('restaurants').find({}).count()
    25359

db.getCollection('restaurants').dataSize()
    11140976

### Quantos índices em cada coleção? Quanto espaço os índices de cada coleção ocupam?
Para ver os indices, deve ser rodar:
db.getCollection('restaurants').stats

para ver quanto espaço os indices ocupam é:
db.getCollection('restaurants').totalIndexSize()

### Traga um documento de exemplo de cada coleção. db.collection.find(...).pretty() nos dá um formato mais legível.
rodando:
db.getCollection('restaurants').find({}).limit(2).pretty()

obtenhe:
 _id: ObjectId('5eb3d668b31de5d588f4294f'),
  direccion: {
    edificio: '625',
    coord: [
      -73.990494,
      40.7569545
    ],
    calle: '8 Avenue',
    codigo_postal: '10018'
  },
  barrio: 'Manhattan',
  tipo_cocina: 'American',
  grados: [
    {
      date: 2014-06-09T00:00:00.000Z,
      grado: 'A',
      puntaje: 12
    },
    {
      date: 2014-01-10T00:00:00.000Z,
      grado: 'A',
      puntaje: 9
    },
    {
      date: 2012-12-07T00:00:00.000Z,
      grado: 'A',
      puntaje: 4
    },
    {
      date: 2011-12-13T00:00:00.000Z,
      grado: 'A',
      puntaje: 9
    },
    {
      date: 2011-09-09T00:00:00.000Z,
      grado: 'A',
      puntaje: 13
    }
  ],
  nombre: 'Cafe Metro',
  restaurante_id: '40363298'
}
{
  _id: ObjectId('5eb3d668b31de5d588f42930'),
  direccion: {
    edificio: '8825',
    coord: [
      -73.8803827,
      40.7643124
    ],
    calle: 'Astoria Boulevard',
    codigo_postal: '11369'
  },
  barrio: 'Queens',
  tipo_cocina: 'American',
  grados: [
    {
      date: 2014-11-15T00:00:00.000Z,
      grado: 'Z',
      puntaje: 38
    },
    {
      date: 2014-05-02T00:00:00.000Z,
      grado: 'A',
      puntaje: 10
    },
    {
      date: 2013-03-02T00:00:00.000Z,
      grado: 'A',
      puntaje: 7
    },
    {
      date: 2012-02-10T00:00:00.000Z,
      grado: 'A',
      puntaje: 13
    }
  ],
  nombre: 'Brunos On The Boulevard',
  restaurante_id: '40356151'
}{
  _id: ObjectId('5eb3d668b31de5d588f42955'),
  direccion: {
    edificio: '464',
    coord: [
      -73.9791458,
      40.744328
    ],
    calle: '3 Avenue',
    codigo_postal: '10016'
  },
  barrio: 'Manhattan',
  tipo_cocina: 'Pizza',
  grados: [
    {
      date: 2014-08-05T00:00:00.000Z,
      grado: 'A',
      puntaje: 3
    },
    {
      date: 2014-03-06T00:00:00.000Z,
      grado: 'A',
      puntaje: 11
    },
    {
      date: 2013-07-09T00:00:00.000Z,
      grado: 'A',
      puntaje: 12
    },
    {
      date: 2013-01-30T00:00:00.000Z,
      grado: 'A',
      puntaje: 4
    },
    {
      date: 2012-01-05T00:00:00.000Z,
      grado: 'A',
      puntaje: 2
    },
    {
      date: 2011-09-26T00:00:00.000Z,
      grado: 'A',
      puntaje: 0
    }
  ],
  nombre: "Domino'S Pizza",
  restaurante_id: '40363644'
} 

### Para cada coleção, liste os campos de nível raiz (ignore campos em documentos aninhados) e seus tipos de dados.
listando os campos e cada tipo dos seus dados:

deve rodar: 
db.getCollection('restaurants').find({}).limit(1).forEach(function(doc) {
        for (var field in doc) {
            print(field + ': ' + typeof doc[field]);
        }
    });

_id: object
direccion: object
barrio: string
tipo_cocina: string
grados: object
nombre: string
restaurante_id: string
