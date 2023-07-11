<template>
  <div>
    <h2>Users</h2>
    <UsersTable />
  </div>
</template>

<script>
import UsersTable from "@/components/UsersTable";
export default {
  name: "ListUsers",
  components: {UsersTable},
}
</script>

<style scoped>

</style>