<template>
  <div style="width: 40vw; height: 40vh">
    <v-chart :option="option" autoresize :loading="false" />
  </div>
</template>

<script setup>
import { reactive } from "@vue/reactivity";

const option = reactive({
  xAxis: {
    data: ["2015", "2016", "2017", "2018", "2019", "2020"],
  },
  yAxis: {},
  series: {
    data: [400, 100, 350, 280, 170, 310, 30],
    type: "bar",
  },
});

const updateData = () => {
  option.series.data = [120, 10, 50, 380, 70, 210, 400]; //可以

  // option.series = {
  //   data: [120, 10, 50, 380, 70, 210, 230],
  // }; //也可以
};
</script>
