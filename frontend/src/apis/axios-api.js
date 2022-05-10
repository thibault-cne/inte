import axios from "axios";

const getAPI = axios.create({
  baseURL: "http://127.0.0.1:5454",
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
  },
});

export { getAPI };
