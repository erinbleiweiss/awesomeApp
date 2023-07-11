<template>
  <NavBar />
  <div class="container">
    <router-view />
  </div>
</template>

<script>
import NavBar from "@/layouts/NavBar";
export default {
  name: "MainLayout",
  components: {NavBar}
}
</script>

<style scoped>

</style>