import axios from "axios";
import { API } from "./const";

export function getEnv(key) {
  return axios.get(API + "/env", { params: { key: key } });
}

export function updateSheetTitle() {
  return axios.patch(API + "/env/title");
}
