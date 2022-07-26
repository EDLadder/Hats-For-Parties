import axios from "axios";

export const partyService = {
  getParties,
  startParty,
  stopParty,
};

function getParties() {
  return axios.get(`http://localhost:8081/party`);
}

function startParty(payload) {
    return axios.post(`http://localhost:8081/party/start`, payload);
}

function stopParty(payload) {
    return axios.patch(`http://localhost:8081/party/stop/${payload.id}`)
}
