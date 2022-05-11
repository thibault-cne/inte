import { postRequest } from "@/requests/postRequest";

// Function to handle the user login using google
export async function handleSignIn(googleUser) {
  try {
    const user_token = googleUser.credential;

    return new Promise((resolve, reject) => {
      postRequest({ credential: user_token }, "auth-api/login", "json")
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
