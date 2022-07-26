import { createStore } from "vuex";
import { hat } from "./modules/hat.module";
import { party } from "./modules/party.module";

export default createStore({
    modules: {
      hat,
      party,
    },
});
