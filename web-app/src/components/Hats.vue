<template>
  <b-container>
    <div class="row">
      <div class="col-6">
        <b-table :items="hats" :fields="fields" :busy="isLoading">
          <template #cell(status)="data">
            <b-badge v-if="checkIsNew(data.item.firstUse)" variant="success">New</b-badge>
            <b-badge v-else-if="checkIsInUse(data.item.partyId)" variant="secondary">In Use</b-badge>
            <b-badge v-else-if="checkIsCleaning(data.item)" variant="warning">In Cleaning</b-badge>
            <b-badge v-else variant="primary">Free</b-badge>
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
      fields: ["name", "status"],
    };
  },
  computed: {
    ...mapGetters({
      hats: "hat/getHats",
      isLoading: "hat/getIsLoading",
    }),
  },
  methods: {
    checkIsNew(firstUse) {
        return new Date(null).toISOString() == new Date(firstUse).toISOString();
    },
    checkIsInUse(partyId) {
        let newVal = partyId.split('0').join('');
        return newVal.length > 0;
    },
    checkIsCleaning(item) {
        return new Date() < new Date(item.canBeUseAfter);
    },
  },
};
</script>
