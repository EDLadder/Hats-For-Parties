import axios from "axios";

export const hatService = {
  getHats,
};

function getHats() {
  return axios.get(`http://localhost:8081/hat`);
}
