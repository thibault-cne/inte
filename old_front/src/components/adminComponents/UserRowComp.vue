<template>
  <td>{{ user.id }}</td>
  <td>{{ user.name }}</td>
  <td>{{ user.currentYear }}</td>
  <td>{{ user.promotionYear }}</td>
  <td>{{ user.god_father_id }}</td>
  <td>
    <v-icon class="mr icon" @click="this.toggle()">mdi-pencil </v-icon>
  </td>
  <EditDialog :toggle="t" :user="user" @save-user="(u) => this.saveUser(u)" />
</template>
<script>
import EditDialog from "@/components/adminComponents/EditDialog.vue";

export default {
  name: "UserRowComp",
  components: { EditDialog },
  props: {
    user: Object,
  },
  data() {
    return {
      t: false,
    };
  },
  methods: {
    toggle() {
      this.t = !this.t;
    },
    saveUser(u) {
      console.log(u);
      this.toggle();
    },
  },
};
</script>
<style lang=""></style>
