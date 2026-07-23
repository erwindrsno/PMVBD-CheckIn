<script setup>
import { ref, onMounted } from 'vue';

// --- State ---
const dialog = ref(false);
const loading = ref(false);
const schools = ref([]);
const newSchool = ref({ name: '', location: '', contact: '' });

const headers = [
  { title: '#', key: 'idx' },
  { title: 'Nama Sekolah', key: 'name' },
];

// --- API Logic ---
const fetchSchools = async () => {
  loading.value = true;
  try {
    // 1. Retrieve the token from localStorage
    const token = sessionStorage.getItem('token');

    const response = await fetch('http://localhost:8080/api/v1/schools', {
      method: 'GET', // (GET is default, but good to be explicit)
      headers: {
        'Content-Type': 'application/json',
        // 2. Attach the Bearer token
        'Authorization': `Bearer ${token}`
      }
    });

    // Optional: Handle unauthorized cases (e.g., if token expired)
    if (response.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login'; // Or use router.push('/login')
      return;
    }

    const result = await response.json();
    schools.value = result.data.schools || [];
  } catch (error) {
    console.error('Error fetching schools:', error);
  } finally {
    loading.value = false;
  }
};

const saveSchool = async () => {
  try {
    const token = sessionStorage.getItem('token');
    const response = await fetch('http://localhost:8080/api/v1/schools', {
      method: 'POST',
      headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(newSchool.value),
    });

    if (response.ok) {
      await fetchSchools(); // Refresh the list
      dialog.value = false;
      newSchool.value = { name: '', location: '', contact: '' };
    }
  } catch (error) {
    console.error('Error saving school:', error);
  }
};

// --- Lifecycle ---
onMounted(() => {
  fetchSchools();
});
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Daftar sekolah</h1>
    <v-btn color="primary" @click="dialog = true">Tambah sekolah baru</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="schools"
    :loading="loading"
    :items-per-page="5"
    class="elevation-1"
  >
    <!-- Optional: Add a loading overlay -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>
    <template v-slot:loading>
      <v-skeleton-loader type="table-row@5"></v-skeleton-loader>
    </template>
  </v-data-table>

  <!-- Add School Dialog -->
  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title>Tambah sekolah baru</v-card-title>
      <v-card-text>
        <v-text-field v-model="newSchool.name" label="School Name"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="dialog = false">Batal</v-btn>
        <v-btn color="primary" @click="saveSchool">Simpan</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
