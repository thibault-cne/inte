<template>
  <div class="hello">
    <h1>{{ msg }}</h1>

    <h1>Is Initialized: {{ Vue3GoogleOauth.isInit }}</h1>
    <h1>Is Authorized: {{ Vue3GoogleOauth.isAuthorized }}</h1>
    <h2 v-if="user">Logged in user: {{ user }}</h2>

    <button
      @click="handleSignIn"
      :disabled="!Vue3GoogleOauth.isInit || Vue3GoogleOauth.isAuthorized"
    >
      Sign In
    </button>
    <button @click="handleSignOut" :disabled="!Vue3GoogleOauth.isAuthorized">
      Sign Out
    </button>
  </div>
</template>

<script>
import { inject } from "vue";
import { postRequest } from "@/requests/postRequest";

export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
  data() {
    return {
      user: "",
    };
  },
  mounted: function () {
    const data = {
      token:
        "eyJhbGciOiJSUzI1NiIsImtpZCI6ImZjYmQ3ZjQ4MWE4MjVkMTEzZTBkMDNkZDk0ZTYwYjY5ZmYxNjY1YTIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiOTM4NzY4OTI0Njk0LTRiNDgzdG1mYmVvZ29iM3BtbjRzNXU1dWpuZ3FtNDNiLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiOTM4NzY4OTI0Njk0LTRiNDgzdG1mYmVvZ29iM3BtbjRzNXU1dWpuZ3FtNDNiLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTExMDY1OTc4ODUyNzUwMzgzNjI3IiwiaGQiOiJ0ZWxlY29tbmFuY3kubmV0IiwiZW1haWwiOiJ0aGliYXVsdC5jaGVuZXZpZXJlQHRlbGVjb21uYW5jeS5uZXQiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6Ik5TMElsVGE3aWlNT3oxZ2QwZjlRdGciLCJuYW1lIjoiVGhpYmF1bHQgQ0hFTkVWScOIUkUiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUFUWEFKd19OU21WWVRuTmFFcC1kZkhlbk5ZcmlzOUhWam1DSVB4U21RZzc9czk2LWMiLCJnaXZlbl9uYW1lIjoiVGhpYmF1bHQiLCJmYW1pbHlfbmFtZSI6IkNIRU5FVknDiFJFIiwibG9jYWxlIjoiZnIiLCJpYXQiOjE2NTE5MTY3OTMsImV4cCI6MTY1MTkyMDM5MywianRpIjoiZjg5ODEzOTMyZTgzMWE3ZTVlY2M2YWFhZmU3OWJlZDNjNTIyMjg0MCJ9.RSnprIxKPQH3D43fdqPGvdljaCQAhrAbCq7_cTDZlaF9bZ9-eaZ2KoaummvU3LTyxa9Gen69zUFA6lwLnQeDJV7byfD9HzUcMss8JnIXsBHHjigL67Zc1OAYfcTVVV-KFerXWlxaNqNNhugKCKROjkeH_8QFgst-mfMkGGD4xnwTPn4nnql_bb_du8UV-GJxmlNi5qGd1Ti03Z-_SNkMOjjbdVtEWJzTsPJ-GcuRtkOTtA9pcods9x8nn0Jyrec_x-pxyJv5Yw7_YOSUeUlPNXdXQD2cOuwGaaYhzgnFdpe2vsjlHBuJJ6ZJlA_NO60Q7gM3gy5H5O67PSi3W2RsIA",
    };
    postRequest(data, "/auth-api/login").then((response) => {
      console.log(response);
    });
    console.log("mounted");
  },
  methods: {
    async handleSignIn() {
      try {
        const googleUser = await this.$gAuth.signIn();
        // console.log(this.$gAuth.signIn);
        if (!googleUser) {
          return null;
        }
        this.user = googleUser.getAuthResponse().id_token;
      } catch (error) {
        console.log(error);
        return null;
      }
    },
    async handleSignOut() {
      try {
        await this.$gAuth.signOut();
        // console.log(this.$gAuth.signOut);
        this.user = "";
      } catch (error) {
        console.log(error);
      }
    },
  },
  setup() {
    const Vue3GoogleOauth = inject("Vue3GoogleOauth");
    return {
      Vue3GoogleOauth,
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
