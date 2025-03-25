Este repositorio contiene la solucion para el desafio de desarrollo que consiste en:

- Conectar con una API para descargar informacion.
- Almacenar la informacion obtenida en CockroachDB.
- Crear una API y una interfaz de usuario para visualizar los datos.
- Desarrollar un algoritmo que recomiende las mejores acciones para invertir*.

## Descripcion

El objetivo del proyecto es construir un sistema que se conecte a una API REST para obtener datos sobre acciones, los almacene en una base de datos CockroachDB y los muestre en una interfaz de usuario amigable. Ademas, se debe implementar una logica que analice estos datos y recomiende la mejor opcion de inversion.

---

## Tecnologias Utilizadas

- **Backend:** Golang  
- **Frontend:** Vue 3 con TypeScript, Pinia y Tailwind CSS  
- **Base de datos:** CockroachDB  

---

## Estructura del Proyecto

- **/backend:** Codigo fuente en Golang para la conexion con la API, procesamiento de datos y exposicion de endpoints.
- **/frontend:** Aplicacion Vue 3 para visualizar la informacion de stocks.

---

## Instalacion y Configuracion

1. **Clonar el repositorio:**

   ```bash
   git clone git@github.com:devAngelOrtiz/DashboardAcciones.git
   ```

2. **Configurar el entorno:**

   - Configura las variables de entorno necesarias:
   ```
   configuracion para dockers
   ./.env
   
   PORT=3000
   FRONT_PORT=8080
   ```
   ```
   configuracion para el api de backedn
   ./backend/.env
   DATABASE_URL=postgresql://root@db:26257/defaultdb?sslmode=disable
   INFO_URL=
   INFO_JWT=
```
   - Asegurate de tener instalados Docker.

3. **Ejecutar la aplicacion:**

   - Levantar la base de datos, api y frontend:
   ```bash
   docker-compose --profile all up
   ```

   si prefieres iniciarlos individualmente usa los siguieentes comandos
   - Levantar base de datos:
   ```bash
   docker-compose --profile db up
   ```
   - Levantar api:
   ```bash
   docker-compose --profile api up
   ```
   - Levantar frontend:
   ```bash
   docker-compose --profile front up
   ```
---

## Uso y Ejecucion

- **Visualizacion:** Accede a la interfaz web para ver el listado de acciones, su analisis y las recomendaciones de inversion*.
- **API:** Utiliza los endpoints provistos por el backend para interactuar con la informacion almacenada.
- **Analisis:** El algoritmo implementado procesa los datos de la base de datos y muestra la recomenda