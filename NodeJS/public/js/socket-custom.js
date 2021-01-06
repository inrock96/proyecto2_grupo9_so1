let socket = io();

//Escuchar sucesos
socket.on('connect', function () {
    console.log('Conectado al servidor');
});

socket.on('disconnect', function () {
    console.log('Perdimos conexion con el servidor');
});

//intervalo
setInterval(function () {
    socket.emit(' sendData', {
        usuario: 'Checha',
        mensaje: 'Hola mundo'
    }, function (resp) {
        //console.log('Respuesta server', resp);
    });
}, 2000);

//Escuchar informacion
socket.on(' sendData', function (mensaje) {
    let top3, piegraph, last_case, barras, alldata;
    top3 = mensaje.mongo.top3;
    piegraph = mensaje.mongo.piegraph;
    alldata = mensaje.mongo.alldata;
    //last_case = mensaje.redis.last_case.shift();
    barras = mensaje.redis.barras;
    console.log('Servidor: ', mensaje);

    /*** ALL DATA */
    $('#alldata').html("")
    alldata.forEach(e => {
        let valor_actual = $('#alldata').html()
        $('#alldata').html(valor_actual + " <tr><td>" + e.name + "</td><td>" + e.age + "</td><td>" + e.location + "</td><td>" + e.infectedtype + "</td><td>" + e.state + "</td></tr>")
    })

    /*   ULTIMO CASO */
    for (let x = 1; x <= 5; x++) {
        last_case = mensaje.redis.last_case.shift();

        $("#nombre" + x).text(last_case.NOMBRE);
        $("#edad" + x).text(last_case.EDAD);
        $("#depto" + x).text(last_case.DEPARTAMENTO);
        $("#forma" + x).text(last_case.FORMA);
        $("#estado" + x).text(last_case.ESTADO);
    }


    /** TOP 3 */
    let cont = 0;
    top3.forEach(element => {
        if (element._id !== null) {
            switch (cont) {
                case 0:
                    $("#depto1_n").text(element._id);
                    $("#depto_val1").text(element.total);
                    break;
                case 1:
                    $("#depto2_n").text(element._id);
                    $("#depto_val2").text(element.total);
                    break;
                case 2:
                    $("#depto3_n").text(element._id);
                    $("#depto_val3").text(element.total);
                    break;
            }
            cont++;
        }
    });

    /***  PIE GRAPH  */
    let valores = [];
    let labels = [];
    piegraph.forEach(element => {
        if (element._id !== null)
            valores.push(element.total);
    });
    piegraph.forEach(element => {
        if (element._id !== null)
            labels.push(element._id);
    });
    let data_pie = [{
        values: valores,
        labels: labels,
        type: 'pie'
    }];


    let layout = {
        height: 500,
        width: 700
    };

    /** BARRAS */
    Plotly.newPlot('chart', data_pie, layout);
    let labeledades = [];
    let valoresedades = [];
    cont = 0;
    barras.forEach(element => {
        if (element !== null) {
            if (cont === 0) {
                labeledades.push(0 + " - " + (cont + 1) * 10);
                valoresedades.push(element);
            } else {
                labeledades.push((cont * 10) + " - " + (cont + 1) * 10);
                valoresedades.push(element);
            }
        }
        cont++;
    });

    let bar_data = [{
        x: labeledades,
        y: valoresedades,
        type: 'bar'
    }];
    Plotly.newPlot('edades', bar_data);
});