const { io } = require('./ServerExpress');
const { MongoClient } = require('./ServerExpress');
const { redis } = require('./ServerExpress');

//constantes
const MONGO_HOST = '34.123.181.161'
const REDIS_HOST = '34.123.181.161'

const NOMBRE = "nombre"
const DEPARTAMENTO = "departamento"
const EDAD = "edad"
const FORMA = "forma de contagio"
const ESTADO = "estado"

//conexion
io.on('connection', async (client) => {
    try {
        console.log('Conexion Establecida');
        let redis = await getRedis();
        let mongo = await getMongo();
        client.emit('data', {
            redis: redis,
            mongo: mongo
        });

        client.on('disconnect', () => {
            console.log('Conexion Terminada');
        });

        //Escuchar el cliente
        client.on(' sendData', async (data, callback) => {
            //console.log(data);
            let redis = await getRedis();
            let mongo = await getMongo();
            client.emit(' sendData', {
                redis: redis,
                mongo: mongo
            });
        });
    } catch (error) {
        client.emit('data', {
            error: error.toString()
        });
    }
});

async function getMongo() {
    let client;
    try {
        client = await MongoClient.connect(`mongodb://AdminSopes1:Sopes1Grupo9@${MONGO_HOST}:27017/?authSource=admin`, { useNewUrlParser: true, useUnifiedTopology: true });
        const db = client.db('CORONAVIRUS');
        const personCollection = db.collection('PACIENTES');

        let piegraph = await personCollection.aggregate([
            { $match: {} }, {
                $group: { _id: '$location', total: { $sum: 1 } }
            }
        ]).sort({ total: -1 }).toArray();

        let top3 = await personCollection.aggregate([
            { $match: {} }, {
                $group: { _id: '$location', total: { $sum: 1 } }
            }
        ]).sort({ total: -1 }).limit(3).toArray();

        //LOCALHOST
        //let alldata = await personCollection.find({}).limit(10).sort({ nombre: 1 }).toArray();
        //SERVER
        let alldata = await personCollection.find({}).sort({ nombre: 1 }).toArray();
        return {
            piegraph: piegraph,
            top3: top3,
            alldata: alldata
        };
    } catch (e) {
        console.error(e);
    } finally {
        client.close();
    }
}

async function getRedis() {
    let vector = [];
    let vector_ultimos = [];
    return new Promise(function (resolve, reject) {
        const clt = redis.createClient({
            host: REDIS_HOST,
            port: 6379,
            password  : 'Sopes1Grupo9'
        });
        clt.on("error", function (error) {
            console.error(error);
        });
        clt.get("CONTADOR", function (err, req) {
            if (err) {
                throw new Error(err);
            }
            console.log("CONTADOR = " + req);
            if (req != null) {
                let valor = req.toString();

                for (let x = 1; x <= 5; x++) {
                    valor = valor - 1;

                    let caso ={NOMBRE: "",
                        EDAD : "",
                        DEPARTAMENTO : "",
                        FORMA : "",
                        ESTADO : ""}


                    let key = "nombre[" + valor.toString() + "]";
                    console.log(key);
                    clt.hget("PACIENTES", key, function (err, result) {
                        if (err) throw err;
                        caso.NOMBRE = result;
                        console.log("NOMBRE: "+caso.NOMBRE);
                    });

                    key = "Estado[".toLowerCase() + valor.toString() + "]";
                    clt.hget("PACIENTES", key, function (err, result) {
                        if (err) throw err;
                        caso.ESTADO = result;
                    });

                    key = "departamento[" + valor.toString() + "]";
                    clt.hget("PACIENTES", key, function (err, result) {
                        if (err) throw err;
                        caso.DEPARTAMENTO = result;
                    });

                    key = "forma de Contagio[".toLowerCase() + valor.toString() + "]";
                    clt.hget("PACIENTES", key, function (err, result) {
                        if (err) throw err;
                        caso.FORMA = result;
                    });

                    key = "edad[" + valor.toString() + "]";
                    clt.hget("PACIENTES", key, function (err, result) {
                        if (err) throw err;
                        caso.EDAD = result;
                    });
                    vector_ultimos.push(caso);
                }
            }
        });


        clt.hgetall("PACIENTES", function (err, value) {
            clt.get("CONTADOR", function (err, req) {
                if (err) {
                    throw new Error(err);
                }

                let valor = req.toString();
                valor = valor - 1;
                for (let x = 0; x <= valor; x++) {
                    let edad = Number(value[EDAD + "[" + x.toString() + "]"]);
                    edad = edad / 10;
                    let intvalue = Math.floor(edad);
                    let contador = 1;
                    if (vector[intvalue] !== undefined) {
                        contador = vector[intvalue] + 1;
                    }
                    vector[intvalue] = contador;
                }

                let redisQuery = {
                    last_case: vector_ultimos,
                    barras: vector
                }
                resolve(redisQuery);

            });
        });
    });
}