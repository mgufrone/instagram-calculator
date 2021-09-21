<template>
  <div class="hello">
    <form method="post" >
      <input v-model="username" placeholder="insert username. ex: @">
      <p>Message is: {{ username }}</p>
      <button type="button" v-on:click="getEngagement()">Calculate</button>
    </form>
    <div v-if="engagements.length > 0">
    <h3>Engagement</h3>
    <table>
      <thead>
        <tr>
          <th>Username</th>
          <th>Engagement</th>
          <th>Followers</th>
        </tr>
      </thead>
      <tr v-for="{username, followers, engagement} in engagements" :key="username">
        <td>{{username}}</td>
        <td>{{percent(engagement*100)}} %</td>
        <td>{{followers}}</td>
      </tr>
    </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data: function () {
    return {
      username: "",
      engagements: [],
    }
  },
  methods: {
    percent: (value) => Math.ceil(value * 100) / 100,
    getEngagement: async function () {
      this.engagements.push(await fetch(process.env.VUE_APP_API + "/profile", {
        method: "POST",
        body: JSON.stringify({
          username: this.username,
        })
      }).then((r) => r.json()))
      this.username = ""
    }
  },
  mounted() {
    console.log(process.env)
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
