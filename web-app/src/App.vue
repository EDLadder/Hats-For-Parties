<template>
  <div>
  <b-tabs content-class="mt-3">
    <b-tab title="Parties" active>
      <Paties />
    </b-tab>
    <b-tab title="Hats">
      <Hats />
    </b-tab>
  </b-tabs>
</div>
</template>

<script>
import Paties from './components/Paties.vue';
import Hats from './components/Hats.vue';

export default {
    name: "App",
    components: { Paties, Hats }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 10px;
}
</style>
