<template>
  <div class="home">
    <h1>Welcome to Servers API</h1>
    <div class="search-form">
      <b-form-group
        class="mb-0"
        label-for="input-formatter"
        description="Enter the domain you want to consult"
      >
        <b-form-input
          id="input-formatter"
          v-model="inputValue"
          placeholder="Example: google.com"
          :formatter="formatter"
        ></b-form-input>
      </b-form-group>
    </div>
    <div>
      <b-button 
        class="home-btn" 
        variant="outline-dark"
        @click="serversClick"
      >Search</b-button>
      <b-button 
        class="home-btn" 
        variant="outline-secondary"
        @click="prevServersClick"
      >Previously consulted</b-button>
    </div>
    
    <div v-if="clicked">
      <div v-if="isLoading" class="d-flex justify-content-center mb-3 ">
        <b-spinner style="width: 3rem; height: 3rem;" label="Loading"></b-spinner>
      </div>
      <px-server-card v-if="!isLoading && !error" :title="title" :servers="servers"/>
    </div>
    <px-error-alert v-if="error" :errText="errText"/>
  </div>
</template>

<script>
// @ is an alias to /src
import PxServerCard from "@/components/PxServerCard.vue"
import PxErrorAlert from "@/components/PxErrorAlert.vue"
import api from '@/api'

export default {
  name: "Home",
  components: {PxServerCard, PxErrorAlert},

  data() {
      return {
        clicked: false,
        isLoading: false,
        error: false,
        errText: '',
        servers: [],
        title: '',
        inputValue: '',
      }
    },
    methods: {
      formatter(value) {
        return value.toLowerCase()
      },
      prevServersClick() {
        this.error = false
        this.clicked = true
        this.isLoading = true
        return api
        .getServers()
        .then(response => (this.servers = response))
        .finally(() => (
          this.servers ? {} : this.setError('There was an issue, please try again later'),
          this.isLoading = false,
          this.title = 'Domains consulted previously'
        ))
        
      },
      serversClick() {
        this.error = false
        this.clicked = true
        this.isLoading = true

        if (this.inputValue === ''){
          this.setError('Please enter a domain')
          this.isLoading = false
        } else {
          return api
          .getServerByDomain(this.inputValue)
          .then(response => (this.servers = response))
          .finally(() => (
            this.servers ? this.checkIfError() : this.setError('There was an issue, please try again later'),
            this.isLoading = false,
            this.title = this.inputValue
          ))
        } 
      },
      setError(mss) {
        this.error = true
        this.errText = mss
      },
      checkIfError() {
        if (this.servers.status === false) {
          this.setError('There was an issue, server doesnt exist or is down')
        }
      },
    }
};
</script>

<style scoped>
.search-form {
  margin-top: 5%;
  margin-bottom: 5%;
}
.form-control {
  width: 50%;
  display: inline-block;
}
.home-btn {
  margin: 0 10px 0 10px;
}
.spinner-border {
  margin-top: 40px;
}
</style>