
function getTime(){
  var t = new Date();
  var time =t.getHours() + ":" + t.getMinutes() + ":" + t.getSeconds();
   
   return time;
 }

var value = 0;
var layout = {
  title:'Grafica de Linea Porcentaje de Ram',
  xaxis: {
    title: "Tiempo"
  },
  yaxis: {
    title: "Porcentaje Ram Usada"
  }
};

Plotly.plot('chart',[
{
    y:[0],
    x: [0].map(getTime),
    type:'line'
}], layout);

setInterval(function(){
    $.ajax({
      //ip de ec2 que contiene todos los conteiner :5536 puerto que corre el servidor golang
//        url: 'http://18.219.48.140:5536/ram',
        url: 'http://localhost:8010/ram',
        type: 'GET',
        success: function(response) { 
            value = response;
        },
        error: function(error) {
            value = 0;
            console.log(error);
        }
    });
    Plotly.extendTraces('chart',{y:[[value]] , x: [[getTime()]] }, [0]);
},5000);
