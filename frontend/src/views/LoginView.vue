<template>
  <div class="login-comp">
    <div
      id="g_id_onload"
      data-client_id="938768924694-4b483tmfbeogob3pmn4s5u5ujngqm43b.apps.googleusercontent.com"
      data-callback="loginCallback"
    ></div>
    <div class="g_id_signin" data-type="standard"></div>
  </div>
</template>

<script>
import { handleSignIn } from "@/apis/google-api";

export default {
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
