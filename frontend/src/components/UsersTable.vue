<template>
  <div>
    <table class="table">
      <thead>
        <th scope="col">ID</th>
        <th scope="col">First Name</th>
        <th scope="col">Last Name</th>
        <th scope="col">Email</th>
        <th scope="col">Username</th>
        <th scope="col">Role</th>
      </thead>
      <tbody>
        <tr v-for="user of users" :key="user.id">
          <td>{{user.id}}</td>
          <td>{{user.firstName}}</td>
          <td>{{user.lastName}}</td>
          <td>{{user.email}}</td>
          <td>{{user.username}}</td>
          <td>{{user.role}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "UsersTable",
  data(){
    return {
      users: []
    }
  },
  methods: {
    async getUsers(){
      try {
        const url = "/api/v1/users"
        const response = await axios.get(url);
        const users = response.data;
        this.users = users.map(user => ({
          id: user.id,
          firstName: user.first_name,
          lastName: user.last_name,
          email: user.email,
          username: user.username,
          role: user.role,
        }))
      } catch (err) {
        if (err.error){
          console.log(`Error getting users: ${err.error}`)
        }
      }
    }
  },
  mounted() {
    this.getUsers()
  }
}
</script>

<style scoped>

</style>