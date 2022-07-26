import { partyService } from "@/services";

const state = {
  isLoading: false,
  parties: [],
};

const getters = {
  getIsLoading: (state) => state.isLoading,
  getParties: (state) => state.parties,
};

const actions = {
  getParties({ commit }) {
    commit("setLoading", true);
    return partyService.getParties().then(
      (response) => {
        commit("getPartiesSuccess", response);
        commit("setLoading", false);
        return Promise.resolve();
      },
      (error) => {
        commit("setLoading", false);
        return Promise.reject(error);
      }
    );
  },
  startParty({ commit }, payload) {
    commit("setLoading", true);
    return partyService.startParty(payload).then(
      () => {
        commit("setLoading", true);
        return Promise.resolve();
      },
      (error) => {
        commit("setLoading", false);
        return Promise.reject(error);
      }
    );
  },
  stopParty({ commit }, payload) {
    commit("setLoading", true);
    return partyService.stopParty(payload).then(
      () => {
        commit("setLoading", true);
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
  getPartiesSuccess(state, payload) {
    if (payload.data.msg.parties != null)
    state.parties = payload.data.msg.parties.sort((a, b) => {
        if (a.updatedAt < b.updatedAt) return 1;
        if (a.updatedAt > b.updatedAt) return -1;
        return ;
    });
  },
};

export const party = {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
