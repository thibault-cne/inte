import axios from "axios";

let base_backend_url = `${process.env.VUE_APP_BACKEND_URI}/api`;

axios.defaults.withCredentials = true;
const getAPI = axios.create({
  baseURL: base_backend_url,
  headers: {
    "Content-Type": "application/json",
  },
});

export { getAPI };
