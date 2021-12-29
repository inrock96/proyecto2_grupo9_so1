# Prometheus

Prometheus es un conjunto de herramientas de supervisión de sistemas y alertas de código abierto.

Prometheus recoge y almacena sus métricas como datos de series temporales, es decir, la información de las métricas se almacena con la marca de tiempo en la que se registró, junto con pares clave-valor opcionales llamados etiquetas.

### ¿Qué son las métricas?

En términos sencillos, las métricas son mediciones numéricas; las series temporales significan que se registran los cambios a lo largo del tiempo. Lo que los usuarios quieren medir difiere de una aplicación a otra. En el caso de un servidor web, pueden ser los tiempos de solicitud; en el caso de una base de datos, el número de conexiones activas o el número de consultas activas, etc.

Las métricas desempeñan un papel importante a la hora de entender por qué su aplicación funciona de una manera determinada. Supongamos que estás ejecutando una aplicación web y descubres que la aplicación es lenta. Necesitarás cierta información para saber qué está pasando con tu aplicación. Por ejemplo, la aplicación puede volverse lenta cuando el número de peticiones es alto. Si tienes la métrica del recuento de peticiones puedes detectar la razón y aumentar el número de servidores para manejar la carga. 

## Instalar Prometheus
```
wget https://github.com/prometheus/prometheus/releases/download/v2.32.1/prometheus-2.32.1.linux-amd64.tar.gz
tar xzf prometheus-2.32.1.linux-amd64.tar.gz
```
### Levantar Promethus
```
./prometheus
```

## Agregando Node Exporter
Exportador de Prometheus para las métricas de hardware y SO expuestas por los kernels *NIX, escrito en Go con colectores de métricas enchufables.

```
wget https://github.com/prometheus/node_exporter/releases/download/v1.3.1/node_exporter-1.3.1.linux-amd64.tar.gz
tar xzf node_exporter-1.3.1.linux-amd64.tar.gz
```

### Levantar Node Exporter
```
./node_exporter
```

## Agregar Node Explorar a Prometheus
Editar el archivo de configuracion de Prometheus, agregando node_exporter como un "job"
```
nano prometheus.yml
# Agregar el job
  - job_name: "node_exporter"
    static_configs:
      - targets: ["localhost:9100"]
```

# Grafana
Grafana es una aplicación web de análisis y visualización interactiva de código abierto multiplataforma. Proporciona tablas, gráficos y alertas para la web cuando se conecta a fuentes de datos compatibles.
## Instalar Grafana
```
#FROM: https://grafana.com/grafana/download?pg=get&plcmt=selfmanaged-box1-cta1
wget https://dl.grafana.com/enterprise/release/grafana-enterprise-8.3.3.linux-amd64.tar.gz
tar -zxvf grafana-enterprise-8.3.3.linux-amd64.tar.gz
```

### Levantar Grafana
```
cd bin
./grafana-server
```

# Agregar Dashboard de Node Explorer a Grafana
Dashboard: https://grafana.com/grafana/dashboards/1860
