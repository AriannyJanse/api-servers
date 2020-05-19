<template>
  <div>
    <h1>{{ msg }}</h1>
    <b-card border-variant="dark" header="Dark" align="center">
        <b-card-text>Lorem ipsum dolor sit amet, consectetur adipiscing elit.<br>
            aaaaa<br>
            aaaaa<br>
            asdasds<br>
            displaydsad<br>
            ssdsds<br>
            asdasds<br>
            asdsdsd<br>
            dssadsd<br>
            adsa
        </b-card-text>
    </b-card>   
  </div>
</template>

<script>
export default {
  name: "Card",
  props: {
    msg: String
  }
};
</script>

<style scoped>
.card {
    display: inline-block;
    width: 40%;
    margin: 3% 0 5% 0;
}
</style>