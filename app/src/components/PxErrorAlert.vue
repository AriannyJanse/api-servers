<template>
  <div>
    <b-alert
      :show="dismissCountDown"
      dismissible
      variant="warning"
      @dismissed="dismissCountDown=0"
      @dismiss-count-down="countDownChanged"
    >
    {{ errText }}
    </b-alert>
    
  </div>
</template>

<script>
  export default {
    name: "PxServerCard",
    props: {
      errText: String
    },
    data() {
      return {
        dismissSecs: 5,
        dismissCountDown: 0
      }
    },
    created() {
        this.showAlert()
    },
    methods: {
      countDownChanged(dismissCountDown) {
        this.dismissCountDown = dismissCountDown
      },
      showAlert() {
        this.dismissCountDown = this.dismissSecs
      }
    }
  }
</script>

<style scoped>
.alert {
  width: 50%;
  display: inline-block;
  margin-top: 5%;
}
</style>