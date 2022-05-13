import { handleSignIn } from "@/apis/google-api";

const actions = {
  async userLogin(context) {
    await handleSignIn()
      .then((response) => {
        context.commit("updateStorage", response.data.credentials);
      })
      .catch((error) => {
        console.log(error);
      });
  },
};

export { actions };
