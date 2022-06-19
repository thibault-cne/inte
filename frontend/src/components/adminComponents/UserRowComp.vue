<template lang="">
  <td>{{ user.id }}</td>
  <td>{{ user.name }}</td>
  <td>{{ user.currentYear }}</td>
  <td>{{ user.promotionYear }}</td>
  <td>{{ user.god_father_id }}</td>
</template>
<script>
export default {
  name: "UserRowComp",
  props: {
    user: Object,
  },
};
</script>
<style lang=""></style>
