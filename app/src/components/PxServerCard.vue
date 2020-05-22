<template>
  <div>
    <b-card-group deck>
      <b-card v-if="title == 'Domains consulted previously'" border-variant="dark" :header="title" align="center">
        <b-card-text
          v-for="server in servers"
          :key="server"
        >
        {{ server }}
      </b-card-text>
    </b-card> 
    <b-card v-else border-variant="dark" :header="title" align="center">
        <b-card-text
        >
          General info: { <br>
            Logo: {{ servers.data.logo }} <br>
            Title: {{ servers.data.title }} <br>
            Is Down: {{ servers.data.is_down }}
          } <br> <br>
          <span
            v-for="server in servers.data.Servers"
            :key="server"
          >
            Server detail: { <br>
              Address: {{ server.address }} <br>
              SSL grade: {{ server.ssl_grade }} <br>
              Country: {{ server.country }} <br>
              Owner: {{ server.owner }}
            } <br><br>
          </span>
          <span v-if="servers.data.ServerChange">
            Server Changes: { <br>
              Servers changed: {{ servers.data.ServerChange.servers_changed }} <br>
              SSL grade: {{ servers.data.ServerChange.ssl_grade }} <br>
              Previous SSL grade: {{ servers.data.ServerChange.previous_ssl_grade }} <br>
            } <br><br>
          </span>

      </b-card-text>
    </b-card>
    </b-card-group>  
  </div>
</template>

<script>
export default {
  name: "PxServerCard",
  props: {
      servers: {
        type: Array,
        default: () => []
      },
      title: String
    },
};
</script>

<style scoped>
.card-deck {
    display: inline-block;
    width: 60% ;
    margin: 3% 0 5% 0;
}
</style>