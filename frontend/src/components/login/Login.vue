<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router' // 1. Import it

const router = useRouter()

const username = ref('')
const password = ref('')
const isLoading = ref(false)

const snackbar = ref({
  show: false,
  text: '',
  color: ''
});

const showSnackbar = (text, color = 'success') => {
  snackbar.value = { show: true, text, color };
};

const handleLogin = async () => {
  isLoading.value = true

  try {
    const response = await fetch('http://localhost:8080/api/v1/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    })
    if (response.status === 401){
      showSnackbar("Username/Password salah", "error")
    } else if(response.status === 200){
      const result = await response.json();
      sessionStorage.setItem("token", result.data.token)
      router.push('/attendees')
    }
  } catch (error) {
    console.error('Error logging in:', error);
  } finally {
    isLoading.value = false;
  }
}
</script>

<template>
  <div class="login-container">
    <form @submit.prevent="handleLogin" class="login-form">
      <h2>PMVBD Hadir</h2>

      <div class="form-group">
        <label for="username">Username</label>
        <input
          id="username"
          type="text"
          v-model="username"
          required
          placeholder="Input username anda"
        />
      </div>

      <div class="form-group">
        <label for="password">Password</label>
        <input
          id="password"
          type="password"
          v-model="password"
          required
          placeholder="Enter your password"
        />
      </div>

      <button type="submit" :disabled="isLoading">
        {{ isLoading ? 'Logging in...' : 'Login' }}
      </button>
    </form>
  </div>


  <v-snackbar
    v-model="snackbar.show"
    :color="snackbar.color"
    location="top"
    timeout="5000"
  >
      {{ snackbar.text }}
  </v-snackbar>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f8fafc;
}

.login-form {
  background: white;
  padding: 2.5rem;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  width: 100%;
  max-width: 400px;
}

.login-form h2 {
  margin-bottom: 1.5rem;
  color: #1e293b;
  text-align: center;
  font-size: 1.5rem;
}

.form-group {
  margin-bottom: 1.25rem;
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: #475569;
}

.form-group input {
  padding: 0.75rem;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.form-group input:focus {
  outline: none;
  border-color: #2563eb;
}

.error-alert {
  background-color: #fef2f2;
  color: #dc2626;
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1.25rem;
  font-size: 0.875rem;
  text-align: center;
  border: 1px solid #fca5a5;
}

button {
  width: 100%;
  padding: 0.75rem;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  font-weight: 500;
  margin-top: 0.5rem;
  transition: background-color 0.2s;
}

button:hover {
  background-color: #1d4ed8;
}

button:disabled {
  background-color: #93c5fd;
  cursor: not-allowed;
}
</style>
