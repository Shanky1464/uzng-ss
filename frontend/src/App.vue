<template>
  <div id="app">
    <nav>
      <img src="https://www.uniphore.com/wp-content/uploads/2024/08/uniphore-default-featured-card.jpg" alt="Uniphore Logo" class="logo" />
      <ul class="nav-links">
        <li><a href="#">Home</a></li>
        <li><a href="#">About</a></li>
        <li><a href="#">Contact</a></li>
      </ul>
    </nav>

    <h1>Service Availability Status</h1>
    <table>
      <thead>
        <tr>
          <th>Service</th>
          <th>Startup</th>
          <th>Readiness</th>
          <th>Liveness</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(service, index) in services" :key="index">
          <td>{{ service.serviceName }}</td>
          <td>
            <span v-if="service.startup.isAvailable" class="status available">&#10003;</span>
            <span v-else class="status not-available">&#10007;</span>
          </td>
          <td>
            <span v-if="service.readiness.isAvailable" class="status available">&#10003;</span>
            <span v-else class="status not-available">&#10007;</span>
          </td>
          <td>
            <span v-if="service.liveness.isAvailable" class="status available">&#10003;</span>
            <span v-else class="status not-available">&#10007;</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      services: [],
    };
  },
  mounted() {
    this.fetchServiceStatuses();
  },
  methods: {
    async fetchServiceStatuses() {
      try {
        const response = await axios.get("http://localhost:8080/check-api");
        const statusCodes = response.data;

        // Services list with all three APIs for each service
        this.services = [
          {
            serviceName: "uzng-gateway",
            startup: { statusCode: statusCodes.StatusCode11 },
            readiness: { statusCode: statusCodes.StatusCode12 },
            liveness: { statusCode: statusCodes.StatusCode13 },
          },
          {
            serviceName: "uzng-flow-manager",
            startup: { statusCode: statusCodes.StatusCode21 },
            readiness: { statusCode: statusCodes.StatusCode22 },
            liveness: { statusCode: statusCodes.StatusCode23 },
          },
          {
            serviceName: "score-card-studio-backend",
            startup: { statusCode: statusCodes.StatusCode31 },
            readiness: { statusCode: statusCodes.StatusCode32 },
            liveness: { statusCode: statusCodes.StatusCode33 },
          },
          {
            serviceName: "score-card-studio-runtime",
            startup: { statusCode: statusCodes.StatusCode41 },
            readiness: { statusCode: statusCodes.StatusCode42 },
            liveness: { statusCode: statusCodes.StatusCode43 },
          },
          {
            serviceName: "uzng-audio-processor",
            startup: { statusCode: statusCodes.StatusCode51 },
            readiness: { statusCode: statusCodes.StatusCode52 },
            liveness: { statusCode: statusCodes.StatusCode53 },
          },
          {
            serviceName: "uzng-conversation-facts-backend",
            startup: { statusCode: statusCodes.StatusCode71 },
            readiness: { statusCode: statusCodes.StatusCode72 },
            liveness: { statusCode: statusCodes.StatusCode73 },
          },
          {
            serviceName: "uzng-data-replicator",
            startup: { statusCode: statusCodes.StatusCode81 },
            readiness: { statusCode: statusCodes.StatusCode82 },
            liveness: { statusCode: statusCodes.StatusCode83 },
          },
          {
            serviceName: "uzng-facts-resolver",
            startup: { statusCode: statusCodes.StatusCode91 },
            readiness: { statusCode: statusCodes.StatusCode92 },
            liveness: { statusCode: statusCodes.StatusCode93 },
          },
          {
            serviceName: "uzng-orchestrator",
            startup: { statusCode: statusCodes.StatusCode101 },
            readiness: { statusCode: statusCodes.StatusCode102 },
            liveness: { statusCode: statusCodes.StatusCode103 },
          },
          {
            serviceName: "uzng-orchestrator-sweeper",
            startup: { statusCode: statusCodes.StatusCode111 },
            readiness: { statusCode: statusCodes.StatusCode112 },
            liveness: { statusCode: statusCodes.StatusCode113 },
          },
          {
            serviceName: "uzng-sentiments",
            startup: { statusCode: statusCodes.StatusCode121 },
            readiness: { statusCode: statusCodes.StatusCode122 },
            liveness: { statusCode: statusCodes.StatusCode123 },
          },
        ];

        // Add availability flags based on statusCode === 200
        this.services = this.services.map((service) => ({
          ...service,
          startup: {
            ...service.startup,
            isAvailable: service.startup.statusCode === 200, // Check if the status code is 200
          },
          readiness: {
            ...service.readiness,
            isAvailable: service.readiness.statusCode === 200, // Check if the status code is 200
          },
          liveness: {
            ...service.liveness,
            isAvailable: service.liveness.statusCode === 200, // Check if the status code is 200
          },
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
  },
};
</script>

<style scoped>
#app {
  font-family: Arial, sans-serif;
  margin: 10px;
}

nav {
  display: flex;
  justify-content: space-between;
  background-color: #6d6d6db8;
  padding: 10px;
  color: rgba(0, 0, 0);
  align-items: center;
  margin:0px;
}

nav .logo {
  height: 100px;
  width: 200px;
}

nav .nav-links {
  list-style: none;
  display: flex;
  gap: 20px;
}

nav .nav-links li {
  display: inline;
}

nav .nav-links a {
  text-decoration: none;
  color: rgb(0, 0, 0);
  font-weight: bold;
}

h1 {
  margin-top: 40px;
  text-align: center;
}

table {
  width: 100%;
  margin-top: 20px;
}

table, th, td {
  border: 1px solid #ddd;
}

th, td {
  padding: 12px;
  text-align: left;
}

th {
  background-color: #f4f4f4;
}

.status {
  font-size: 18px;
}

.status.available {
  color: green;
}

.status.not-available {
  color: red;
}
</style>
