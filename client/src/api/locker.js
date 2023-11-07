import axios from "axios";
import { API } from "./const";

export async function getLockerEvents() {
  return axios.get(API + "/locker");
}

export async function getPRinters() {
  return axios.get(API + "/locker/printers");
}

export async function printBracelet(req, pickId) {
  return axios.post(API + "/locker/print", req, { params: { pickId: pickId } });
}
