<template>
  <h1 class="mb-5">Admin Login</h1>
  <div class="row">
    <div class="col-2"></div>
    <div class="col-8">
      <form id="LoginForm" @submit.prevent="onSubmit">
        <div class="form-group mb-3">
          <label for="username">Username</label>
          <input id="username" class="form-control" type="text" autocomplete="username" v-model="username">
        </div>
        <div class="form-group mb-3">
          <label for="password">Password</label>
          <input id="password" class="form-control" type="password" autocomplete="password" v-model="password">
        </div>
        <button class="btn btn-primary w-100">Login</button>
      </form>
      <ErrorFlash :error="this.errors"></ErrorFlash>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import ErrorFlash from "@/components/ErrorFlash";

export default {
  name: "AdminLogin",
  components: {ErrorFlash},
  data() {
    return {
      username: "",
      password: "",
      success: false,
      errors: ""
    }
  },
  methods: {
    async onSubmit(){
      let creds = {
        username: this.username,
        password: this.password
      }

      try {
        const url = "/api/v1/login"
        const response = await axios.post(url, creds)
        if (response.status === 200){
          // TODO: This is an insecure way to mock admin access. Ideally, real session tokens would be issued
          this.$router.push("users")
        }
      } catch (err) {
        if (err.response){
          this.errors = err.response.data.error;
        }
      }
    }
  }
}
</script>

<style scoped>

</style>