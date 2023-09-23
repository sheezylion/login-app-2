<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="formData.username" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input
          type="password"
          id="password"
          v-model="formData.password"
          required
        />
      </div>
      <button type="submit">Login</button>
    </form>
    <div v-if="message" class="message">{{ message }}</div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      formData: {
        username: "",
        password: "",
      },
      message: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post(
          `${process.env.VUE_APP_API_ENDPOINT}/api/login`,
          this.formData
        );
        if (response.status === 200) {
          this.message = "Login successful";
        }
      } catch (error) {
        this.message = "Invalid credentials";
      }
    },
  },
};
</script>

<style scoped>
form {
  max-width: 300px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f9f9f9;
}

label {
  font-weight: bold;
}

input {
  width: 100%;
  padding: 8px;
  margin-bottom: 10px;
}

button {
  background-color: #007bff;
  color: #fff;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
}

.message {
  margin-top: 10px;
  color: #ff0000;
}
</style>
