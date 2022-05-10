import { postRequest } from "@/requests/postRequest";

// Function to handle the user login using google
export async function handleSignIn() {
  try {
    const googleUser = await this.$gAuth.signIn();
    if (!googleUser) {
      return null;
    }
    const user_token = googleUser.getAuthResponse().id_token;

    return new Promise((resolve, reject) => {
      postRequest({ token: user_token }, "auth-api/login", "json")
        .then((response) => {
          resolve(response);
        })
        .catch((error) => {
          reject(error);
        });
    });
  } catch (error) {
    console.log(error);
    return null;
  }
}

// Function to handle the user logout
export async function handleSignOut() {
  try {
    await this.$gAuth.signOut();
    // console.log(this.$gAuth.signOut);
    this.user = "";
  } catch (error) {
    console.log(error);
  }
}
