<script setup>
import { ref, onMounted } from 'vue';

// --- State ---
const dialog = ref(false);
const loading = ref(false);
const schools = ref([]);
const newSchool = ref({ name: '', location: '', contact: '' });

const headers = [
  { title: 'No.', key: 'idx' },
  { title: 'School Name', key: 'name' },
];

// --- API Logic ---
const fetchSchools = async () => {
  loading.value = true;
  try {
    const response = await fetch('http://localhost:8080/api/v1/schools');
    const result = await response.json();
    console.log(result)
    // Assuming backend returns { "data": { "schools": [...] } }
    schools.value = result.data.schools || [];
  } catch (error) {
    console.error('Error fetching schools:', error);
  } finally {
    loading.value = false;
  }
};

const saveSchool = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/v1/schools', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
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
    <h1>School List</h1>
    <v-btn color="primary" @click="dialog = true">Add New School</v-btn>
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
      <v-card-title>Add New School</v-card-title>
      <v-card-text>
        <v-text-field v-model="newSchool.name" label="School Name"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="dialog = false">Cancel</v-btn>
        <v-btn color="primary" @click="saveSchool">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
