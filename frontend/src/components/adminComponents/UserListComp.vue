<template lang="">
  <table>
    <thead>
      <tr>
        <th v-for="head in headers" :key="head" @click="sortTable(head)">
          {{ head }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="user in users" :key="user.id">
        <UserRowCompVue :user="user" />
      </tr>
    </tbody>
  </table>
</template>
<script>
import _ from "lodash";
import UserRowCompVue from "@/components/adminComponents/UserRowComp.vue";

export default {
  name: "UserListView",
  components: { UserRowCompVue },
  data() {
    return {
      ascending: [false, false, false, false, false],
      objectAttributes: [
        "id",
        "name",
        "currentYear",
        "promotionYear",
        "god_father_id",
      ],
      headers: [
        "Id",
        "Name",
        "Current year",
        "Promotion year",
        "God father id",
      ],
      users: [
        {
          id: 1,
          name: "John Doe",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 20,
        },
        {
          id: 2,
          name: "Jane Doe",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 18,
        },
        {
          id: 3,
          name: "LV Capelli",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 20,
        },
        {
          id: 4,
          name: "Thibault Cheneviere",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 8,
        },
        {
          id: 5,
          name: "Louis Clériot",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 20,
        },
        {
          id: 6,
          name: "Louis Chatard",
          currentYear: 1,
          promotionYear: 2024,
          god_father_id: 18,
        },
        {
          id: 7,
          name: "Laury Thiebaux",
          currentYear: 2,
          promotionYear: 2023,
          god_father_id: 17,
        },
        {
          id: 8,
          name: "Eliott Pommier",
          currentYear: 2,
          promotionYear: 2023,
          god_father_id: 42,
        },
      ],
    };
  },
  methods: {
    sortTable: function sortTable(col) {
      let index = this.headers.indexOf(col);
      this.ascending[index] = !this.ascending[index];
      let sortBy = this.ascending[index] ? "desc" : "asc";
      let colName = this.objectAttributes[index];

      this.users = _.orderBy(this.users, colName, [sortBy]);
    },
  },
};
</script>
<style lang="scss">
table th {
  cursor: pointer;
}
</style>
