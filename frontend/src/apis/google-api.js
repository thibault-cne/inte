import { postRequest } from "@/requests/postRequest";
import { authStore } from "@/store/authStore";

// Function to handle the user login using google
export async function handleSignIn(googleUser) {
  try {
    const user_token = googleUser.credential;

    return new Promise((resolve, reject) => {
      postRequest({ credential: user_token }, "/login/g-oauth", "data")
        .then((response) => {
          let credentials = {
            access_token: response.data.access_token,
            refresh_token: response.data.refresh_token,
          };
          authStore.commit("updateStorage", credentials);
          resolve();
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
