import { API } from "./const";
import axios from "axios";

export async function getEvents() {
  const res = await fetch(API + "/event", {
    method: "GET",
  });

  return res.json();
}

export async function getEventLists(query) {
  const uuid = query.queryKey[1];
  const res = await fetch(API + `/event/${uuid}` + "/lists");

  return res.json();
}

export async function postEvent(data) {
  return axios.post(API + "/event", data);
}

export async function deleteEvent(uuid) {
  return axios.delete(API + "/event/" + uuid);
}

export async function updateEvent(event) {
  return axios.patch(API + "/event", event);
}
