<template>
  <svg
    viewBox="0 0 1000 500"
    xmlns="http://www.w3.org/2000/svg"
    v-html="bubbles"
  ></svg>
</template>

<script>
const BACKEND_ADDR = "http://192.168.92.93:8888";

export default {
  data() {
    return {
      bubbles: null,
    };
  },
  created() {
    var self = this;
    setInterval(function () {
      self.getDatas();
      //console.log(self.bubbles)
    }, 300);
  },
  methods: {
    getDatas() {
      fetch(BACKEND_ADDR)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          this.jsonToSVG(data);
        });
    },
    jsonToSVG(json) {
      var self = this;
      this.bubbles = "";
      for (var [key, allBubbles] of Object.entries(json)) {
        for (var [key, bubble] of Object.entries(allBubbles)) {
          if (key == "data") {
            //console.log(bubble);
            self.bubbles += `<circle cx="${bubble.cx * 50}" cy="${
              bubble.cy * 50
            }" r="${bubble.r / 2}" fill="${bubble.color}"/>`;
          }
        }
      }
      console.log(this.bubbles);
    },
  },
};
</script>

<style scoped></style>
