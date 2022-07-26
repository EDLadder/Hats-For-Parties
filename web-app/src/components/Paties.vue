<template>
  <b-container>
    <div class="row">
      <div class="col-3">
        <b-form @submit="onSubmit">
          <b-form-group
            id="input-group-2"
            label="Party Name:"
            label-for="input-2"
          >
            <b-form-input
              id="input-2"
              v-model="name"
              placeholder="Party name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group id="input-group-2" label="Hats:" label-for="input-2">
            <b-form-input
              id="input-2"
              v-model="hats"
              type="number"
              placeholder="0"
              required
            ></b-form-input>
          </b-form-group>
          <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>
      <div class="col-9">
        <b-table :items="parties" :fields="fields" :busy="isLoading">
          <template v-slot:cell(action)="{ item }">
            <b-button
              v-if="item.status != 'Stopped'"
              variant="danger"
              size="sm"
              @click="stopParty(item)"
              >Stop</b-button
            >
          </template>
        </b-table>
      </div>
    </div>
  </b-container>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  data() {
    return {
      name: "",
      hats: 0,
      fields: ["name", "hats", "status", "action"],
    };
  },
  computed: {
    ...mapGetters({
      parties: "party/getParties",
      isLoading: "party/getIsLoading",
    }),
  },
  methods: {
    onSubmit() {
      this.$store.dispatch("party/startParty", {"name": this.name, "hats":parseInt(this.hats)}).then(
        () => {
            this.name = "";
            this.hats = 0;
            this.getParties();
            this.getHats();
        },
        (error) => {
            alert(error.response.data.msg);
        }
      );
    },
    stopParty(payload) {
      this.$store.dispatch("party/stopParty", {id: payload.id}).then(
        () => {
            this.getParties();
            this.getHats();
        },
        (error) => {
            alert(error.response.data.msg);
        }
      );
    },
    getParties() {
      this.$store.dispatch("party/getParties");
    },
    getHats() {
      this.$store.dispatch("hat/getHats");
    },
  },
  created() {
    this.getParties();
    this.getHats();
  },
};
</script>
