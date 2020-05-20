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
          placeholder="Example: google.com"
          :formatter="formatter"
        ></b-form-input>
      </b-form-group>
    </div>
    <div>
      <b-button 
        class="home-btn" 
        variant="outline-dark"
        @click="isClicked"
      >Search</b-button>
      <b-button 
        class="home-btn" 
        variant="outline-secondary"
        @click="isClicked"
      >Previously consulted</b-button>
    </div>
    <div v-show="clicked">
      <card />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Card from "@/components/Card.vue";

export default {
  name: "Home",
  components: {Card},

  data() {
      return {
        clicked: false,
      }
    },
    methods: {
      formatter(value) {
        return value.toLowerCase()
      },
      isClicked() {
        this.clicked = true;
      }
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

</style>