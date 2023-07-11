<template>
  <h1 class="mb-5">Sign Up</h1>
  <div class="row">
    <div class="col-2"></div>
    <div class="col-8">
      <form id="RegistrationForm" @submit.prevent="onSubmit">
        <div class="row form-row mb-3">
          <div class="col-md-6 form-group">
            <label for="firstName">First Name</label>
            <input id="firstName" class="form-control" type="text" autocomplete="given-name" v-model="first_name">
          </div>
          <div class="col-md-6 form-group">
            <label for="lastName">Last Name</label>
            <input id="lastName" class="form-control" type="text" autocomplete="family-name" v-model="last_name">
          </div>
        </div>
        <div class="form-group mb-3">
          <label for="username">Username</label>
          <input id="username" class="form-control" type="text" autocomplete="username" v-model="username">
        </div>
        <div class="form-group mb-3">
          <label for="email">Email</label>
          <input id="email" class="form-control" type="email" autocomplete="email" v-model="email">
        </div>
        <div class="form-group mb-3">
          <label for="password">Password</label>
          <input id="password" class="form-control" type="password" autocomplete="new-password" v-model="password">
        </div>
        <div class="form-group mb-5">
          <label for="repeatPassword">Repeat Password</label>
          <input id="repeatPassword" class="form-control" type="password" autocomplete="new-password" v-model="repeat_password">
        </div>
        <button class="btn btn-primary w-100">Sign up</button>
      </form>
      <ErrorFlash :error="this.errors"></ErrorFlash>
      <h3 v-if="this.success" class="mt-5">Thank you for signing up!</h3>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import ErrorFlash from "@/components/ErrorFlash";

function initialState(){
  return {
    first_name: "",
    last_name: "",
    email: "",
    username: "",
    password: "",
    repeat_password: "",
    errors: "",
    success: false
  }
}

export default {
  name: "RegistrationForm",
  components: {ErrorFlash},
  data: function(){
    return initialState()
  },
  methods: {
    async onSubmit() {
      this.errors = "";
      this.success = false;

      if (this.password !== this.repeat_password) {
        this.errors = "Passwords must match.";
        return
      }

      // TODO: Validate email format and enforce password requirements client-side

      let newUser = {
        first_name: this.first_name,
        last_name: this.last_name,
        email: this.email,
        username: this.username,
        password: this.password,
      }

      try {
        const url = "/api/v1/register"
        const response = await axios.post(url, newUser);
        if (response.status === 200){
          Object.assign(this.$data, initialState());
          this.success = true;
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
  h3 {
    text-align: center;
  }
</style>