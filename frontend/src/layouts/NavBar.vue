<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark ps-4 mb-4">
    <a class="navbar-brand" href="/">Awesome App</a>
    <button class="navbar-toggler" type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            @click="visible=!visible"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav me-auto">
      </ul>
      <ul class="navbar-nav navbar-right me-3">
        <li class="nav-item">
          <router-link to="/admin" class="nav-item nav-link">Admin Login</router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavBar"
}
</script>

<style scoped>

</style>