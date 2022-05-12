<template>
  <div class="login-comp">
    <div
      id="g_id_onload"
      :data-client_id="google_client_id"
      data-callback="loginCallback"
    ></div>
    <div class="g_id_signin" data-type="standard"></div>
  </div>
</template>

<script>
import { handleSignIn } from "@/apis/google-api";

export default {
  data() {
    return {
      google_client_id: process.env.VUE_APP_GOOGLE_CLIENT_ID,
    };
  },
  name: "LoginView",
  mounted() {
    window.loginCallback = function (googleUser) {
      handleSignIn(googleUser);
    };
    let recaptchaScript = document.createElement("script");
    recaptchaScript.setAttribute(
      "src",
      "https://accounts.google.com/gsi/client"
    );
    document.head.appendChild(recaptchaScript);
  },
};
</script>

<style scoped lang="scss"></style>
