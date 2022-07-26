import { hatService } from "@/services";

const state = {
  isLoading: false,
  hats: [],
};

const getters = {
  getIsLoading: (state) => state.isLoading,
  getHats: (state) => state.hats,
};

const actions = {
  getHats({ commit }) {
    commit("setLoading", true);
    return hatService.getHats().then(
      (response) => {
        commit("getHatsSuccess", response);
        commit("setLoading", false);
        return Promise.resolve();
      },
      (error) => {
        commit("setLoading", false);
        return Promise.reject(error);
      }
    );
  },
};

const mutations = {
  setLoading(state, payload) {
    state.isLoading = payload;
  },
  getHatsSuccess(state, payload) {
    state.hats = payload.data.msg.hats;
  },
};

export const hat = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
