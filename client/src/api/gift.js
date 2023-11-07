import axios from "axios";
import { API } from "./const";

export function giveGift({ login, key, pickId }) {
  return axios.post(API + "/gift" + `/${login}`, null, {
    params: { key: key, count: 1, pickId: pickId },
  });
}
