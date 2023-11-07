import axios from "axios";
import { API } from "./const";

export function updateRegistration(reg, pickId) {
  return axios.patch(API + "/registration", reg, {
    params: { pickId: pickId },
  });
}
