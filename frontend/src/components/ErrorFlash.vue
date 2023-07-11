<template>
  <div class="alert alert-danger mt-5" role="alert" v-if="error">
    {{error}}
  </div>
</template>

<script>
export default {
  name: "ErrorFlash",
  props: ['error']
}
</script>

<style scoped>

</style>