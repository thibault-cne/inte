<template>
  <div class="login-comp">
    <div
      id="g_id_onload"
      :data-client_id="google_client_id"
      data-auto_prompt="false"
      data-callback="loginCallback"
    ></div>
    <div
      class="g_id_signin"
      data-type="standard"
      data-size="large"
      data-theme="outline"
      data-text="sign_in_with"
      data-shape="rectangular"
      data-logo_alignment="left"
    ></div>
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

<style scoped lang="scss">
.g_id_signin {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100px;
  height: 40px;
  margin: 0 auto;
}
</style>
