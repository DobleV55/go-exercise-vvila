# Bitcoin Last Traded Price API

Este proyecto proporciona una API para recuperar el último precio negociado de Bitcoin en varios pares de divisas, así como el precio promedio a partir de múltiples fuentes de datos.

## Características

- Recupera el último precio negociado de Bitcoin para los pares de divisas BTC/USD, BTC/CHF y BTC/EUR.
- Calcula el precio promedio de Bitcoin utilizando datos de múltiples servicios (Kraken y Blockchain.com).
- Datos actualizados con precisión de hasta el último minuto.
- Implementación de caché con Redis para mejorar el rendimiento y reducir la carga en los servicios externos.

## Endpoints

- `/api/v1/ltp`: Recupera el último precio negociado de Bitcoin para los pares de divisas especificados.
- `/api/v1/average`: Recupera el precio promedio de Bitcoin a partir de múltiples servicios.

## Requisitos

- Docker
- Docker Compose

## Configuración y Ejecución

1. Clona este repositorio en tu máquina local:

   ```bash
   git clone https://github.com/DobleV55/go-exercise-vvila.git
   cd go-exercise-vvila
   ```

2. Construye la imagen Docker del servidor:

    ```bash
    docker-compose build
    ```

3. Inicia los contenedores Docker para ejecutar el servidor y Redis:

        ```bash
        docker-compose up
        ```
4. Accede a la API desde tu navegador o cliente HTTP:

    LTP (Last Traded Price): http://localhost:8080/api/v1/ltp
    Precio Promedio: http://localhost:8080/api/v1/average
