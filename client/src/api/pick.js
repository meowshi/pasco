import axios from "axios";
import { API } from "./const";

export function pick(key) {
  return axios.get(API + "/yandexoid", { params: { key: key } });
}

export function picks() {
  return axios.get(API + "/pick");
}
