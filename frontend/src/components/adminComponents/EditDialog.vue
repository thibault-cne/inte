<template>
  <div class="text-center">
    <v-dialog v-model="t">
      <v-card class="dialog">
        <v-card-title>
          <span class="text-h5">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-text-field
                label="Id*"
                required
                v-model="this.u.id"
                disabled
                class="mr"
              ></v-text-field>
              <v-text-field
                label="Username*"
                required
                v-model="this.u.name"
                class="mr"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-text-field
                label="Current year*"
                required
                v-model="this.u.currentYear"
                class="mr"
                type="number"
              ></v-text-field>
              <v-text-field
                label="Promotion year*"
                required
                v-model="this.u.promotionYear"
                class="mr"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-text-field
                label="God fathe id*"
                required
                v-model="this.u.god_father_id"
                class="mr"
                type="number"
              ></v-text-field>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue-darken-1" text @click="this.saveUser()">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
<script>
export default {
  name: "EditDialog",
  props: {
    user: undefined,
    toggle: Boolean,
  },
  watch: {
    toggle: function () {
      this.t = this.toggle;
    },
  },
  data() {
    return {
      u: this.user,
      t: this.toggle,
    };
  },
  methods: {
    saveUser() {
      this.$emit("saveUser", this.user);
    },
  },
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standards/mixin";

.mr {
  margin-right: 8px;
}

.dialog {
  min-width: 100vw;

  @include mixin.lg {
    min-width: 45vw;
  }
}

.buttons {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
}
</style>
